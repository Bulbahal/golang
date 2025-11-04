package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), dsn)
	must(err)
	return db
}

func applySchema(db *pgxpool.Pool) {
	b, err := os.ReadFile("schema.sql")
	must(err)
	_, err = db.Exec(context.Background(), string(b))
	must(err)
}

func printAvg(db *pgxpool.Pool, courseID int, label string) {
	var avg float64
	err := db.QueryRow(context.Background(),
		`SELECT avg_grade::float8 FROM course_avg_grade WHERE course_id=$1`, courseID,
	).Scan(&avg)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Printf("%s course %d avg = (no row)\n", label, courseID)
			return
		}
		must(err)
	}
	fmt.Printf("%s course %d avg = %.2f\n", label, courseID, avg)
}

func main() {
	db := connect()
	defer db.Close()

	applySchema(db)

	var courseID int
	must(db.QueryRow(context.Background(),
		`INSERT INTO courses(title) VALUES($1) RETURNING id`, "DB 101",
	).Scan(&courseID))

	_, err := db.Exec(context.Background(),
		`INSERT INTO enrollments(student_id, course_id, grade) VALUES (1,$1,5),(2,$1,3)`,
		courseID)
	must(err)
	printAvg(db, courseID, "after insert:")

	_, err = db.Exec(context.Background(),
		`UPDATE enrollments SET grade=4 WHERE student_id=2 AND course_id=$1`,
		courseID)
	must(err)
	printAvg(db, courseID, "after update:")

	_, err = db.Exec(context.Background(),
		`DELETE FROM enrollments WHERE student_id=1 AND course_id=$1`,
		courseID)
	must(err)
	printAvg(db, courseID, "after delete:")

	_, err = db.Exec(context.Background(),
		`DELETE FROM enrollments WHERE course_id=$1`,
		courseID)
	must(err)
	printAvg(db, courseID, "after delete all:")
}
