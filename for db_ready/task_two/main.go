package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

// TODO: Структуры для студентов, курсов, зачислений, платежей
type Student struct {
	Id    int
	Name  string
	Group string
	Year  int
}

type Course struct {
	Id      int
	Title   string
	Credits int
}
type Enrollment struct {
	Id         int
	Student_id int
	Course_id  int
	Grade      *int
	Paid       bool
}

type Payments struct {
	Id            int
	Enrollment_id int
	Amount        float64
}

// TODO: Функция подключения к БД (через pgxpool, с переменными окружения)
func connectDB() (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")
	return pgxpool.New(context.Background(), dsn)
}

// TODO: Функция создания таблиц (чтение schema.sql)
func createTables(db *pgxpool.Pool) error {
	sqlBytes, err := os.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(context.Background(), string(sqlBytes))
	return err
}

// TODO: Функция зачисления студента на курс с созданием платежа (транзакция)
func enrollWithPaymentTx(db *pgxpool.Pool, studentID, courseID int, amount float64) (enrollmentID, paymentID int, err error) {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return 0, 0, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(),
		"INSERT INTO enrollments(student_id, course_id) VALUES($1, $2) RETURNING id;",
		studentID, courseID).Scan(&enrollmentID); err != nil {
		return 0, 0, err
	}

	if err := tx.QueryRow(context.Background(),
		"INSERT INTO payments (enrollment_id, amount) VALUES($1, $2) RETURNING id;",
		enrollmentID, amount).Scan(&paymentID); err != nil {
		return 0, 0, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return 0, 0, err
	}
	return enrollmentID, paymentID, nil

}

// TODO: Функция обновления оценки и оплаты (транзакция)
func updateGradeAndMarkPaidTx(db *pgxpool.Pool, enrollmentID, grade int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	if _, err := tx.Exec(context.Background(), `UPDATE enrollments SET grade=$1 WHERE id=$2`, grade, enrollmentID); err != nil {
		return err
	}
	if _, err := tx.Exec(context.Background(),
		`UPDATE enrollments SET paid=TRUE WHERE id=$1`, enrollmentID); err != nil {
		return err
	}
	return tx.Commit(context.Background())
}

// TODO: Функция удаления студента и связанных записей (транзакция)
func deleteStudent(db *pgxpool.Pool, studentID int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	if _, err := tx.Exec(context.Background(),
		"DELETE FROM students WHERE id=$1", studentID); err != nil {
		return err
	}
	return tx.Commit(context.Background())
}

// TODO: Функции для выполнения сложных SQL-запросов
// Список студентов которые учатся последний курс
func studentsWithAnyPayment(db *pgxpool.Pool) ([]Student, error) {
	rows, err := db.Query(context.Background(), `SELECT DISTINCT s.id, s.name, s."group",COALESCE(s.birth_year, 0)
FROM students AS s 
    JOIN enrollments AS e ON e.student_id = s.id
    JOIN payments AS p ON p.enrollment_id = e.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var st Student
		if err := rows.Scan(&st.Id, &st.Name, &st.Group, &st.Year); err != nil {
			return nil, err
		}
		students = append(students, st)
	}
	return students, rows.Err()
}

// Курсы на которые никто не записан
func coursesWithNoEnrollments(db *pgxpool.Pool) ([]Course, error) {
	rows, err := db.Query(context.Background(),
		`SELECT c.id, c.title, coalesce(c.credits, 0)
FROM courses AS c
LEFT JOIN enrollments AS e ON e.course_id = c.id
WHERE e.id is NULL
ORDER BY c.title ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.Id, &c.Title, &c.Credits); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, rows.Err()
}

// Общая сумам для каждого курса
func totalPayments(db *pgxpool.Pool) (map[int]float64, error) {
	rows, err := db.Query(context.Background(), `SELECT e.course_id, SUM(p.amount)::float8
		   FROM payments p
		   JOIN enrollments e ON e.id = p.enrollment_id
		  GROUP BY e.course_id
		  ORDER BY e.course_id;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make(map[int]float64)
	for rows.Next() {
		var courseID int
		var sum float64
		if err := rows.Scan(&courseID, &sum); err != nil {
			return nil, err
		}
		res[courseID] = sum
	}
	return res, rows.Err()
}

type StudentRank struct {
	StudentID int
	Name      string
	Total     float64
	Rank      int
}

func rankStudentsByPayments(db *pgxpool.Pool, limit int) ([]StudentRank, error) {
	rows, err := db.Query(context.Background(),
		`WITH totals AS (
		   SELECT s.id AS student_id, s.name, COALESCE(SUM(p.amount), 0)::float8 AS total
		     FROM students s
		     LEFT JOIN enrollments e ON e.student_id = s.id
		     LEFT JOIN payments p    ON p.enrollment_id = e.id
		    GROUP BY s.id, s.name
		)
		SELECT student_id, name, total,
		       RANK() OVER (ORDER BY total DESC) AS rnk
		  FROM totals
		  ORDER BY total DESC, name
		  LIMIT $1;`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []StudentRank
	for rows.Next() {
		var st StudentRank
		if err := rows.Scan(&st.StudentID, &st.Name, &st.Total, &st.Rank); err != nil {
			return nil, err
		}
		out = append(out, st)
	}
	return out, rows.Err()
}

func main() {
	// TODO: Подключение, создание таблиц, тестовые данные, вызовы функций
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		panic(err)
	}
	var studentID, courseID int
	_ = db.QueryRow(context.Background(),
		`INSERT INTO students (name, "group", birth_year) VALUES ('Иван', 'A-01', 2003) RETURNING id;`,
	).Scan(&studentID)
	_ = db.QueryRow(context.Background(),
		`INSERT INTO courses (title, credits) VALUES ('Базы данных', 4) RETURNING id;`,
	).Scan(&courseID)

	enrollID, payID, err := enrollWithPaymentTx(db, studentID, courseID, 199.99)
	if err != nil {
		panic(err)
	}
	println("enrolled:", enrollID, "payment:", payID)

	if err := updateGradeAndMarkPaidTx(db, enrollID, 5); err != nil {
		panic(err)
	}
	println("grade updated & paid marked")

	students, _ := studentsWithAnyPayment(db)
	println("students with any payment:", len(students))

	courses, _ := coursesWithNoEnrollments(db)
	println("courses with no enrollments:", len(courses))

	total, _ := totalPayments(db)
	println("total payments:", len(total))

	ranks, _ := rankStudentsByPayments(db, 10)
	println("ranked students:", len(ranks))
}
