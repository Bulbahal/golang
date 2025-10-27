package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Author struct {
	ID   int
	Name string
}
type Book struct {
	ID       int
	Title    string
	AuthorID int
	Year     int
}

// TODO: Добавьте структуры для работы с данными

// TODO: Реализуйте функцию подключения к БД
func connectDB() (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")
	return pgxpool.New(context.Background(), dsn)
}

// TODO: Реализуйте функцию создания таблиц
func createTables(db *pgxpool.Pool) error {
	sqlBytes, err := os.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(context.Background(), string(sqlBytes))
	return err
}

// TODO: Реализуйте функцию добавления автора
func addAuthor(db *pgxpool.Pool, author Author) (int, error) {
	var id int
	err := db.QueryRow(context.Background(), "INSERT INTO authors (name) VALUES ($1) RETURNING id;", author.Name).Scan(&id)
	return id, err
}

// TODO: Реализуйте функцию добавления книги
func addBook(db *pgxpool.Pool, book Book) (int, error) {
	// TODO: Вставьте книгу и верните её ID
	var id int
	err := db.QueryRow(context.Background(),
		"INSERT INTO books (title, author_id, publication_year) VALUES ($1, $2, $3) RETURNING id;",
		book.Title, book.AuthorID, book.Year,
	).Scan(&id)
	return id, err
}

// TODO: Реализуйте функцию поиска книг по автору
func findBookByAuthor(db *pgxpool.Pool, authorName string) ([]Book, error) {
	rows, err := db.Query(context.Background(),
		`SELECT b.id, b.title, b.author_id, b.publication_year
		 FROM books b
		 JOIN authors a ON b.author_id = a.id
		 WHERE a.name = $1;`, authorName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// TODO: Реализуйте функцию подсчета книг по годам
func countBooksByYear(db *pgxpool.Pool) (map[int]int, error) {
	rows, err := db.Query(context.Background(),
		"SELECT publication_year, COUNT(*) FROM books WHERE publication_year IS NOT NULL GROUP BY publication_year;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[int]int)
	for rows.Next() {
		var year, count int
		if err := rows.Scan(&year, &count); err != nil {
			return nil, err
		}
		result[year] = count
	}
	return result, nil
}

// TODO: Реализуйте функцию обновления информации о книге
func updateBook(db *pgxpool.Pool, book Book) error {
	_, err := db.Exec(context.Background(),
		"UPDATE books SET title=$1, author_id=$2, publication_year=$3 WHERE id=$4", book.Title, book.AuthorID, book.Year, book.ID)
	return err
}

// TODO: Реализуйте функцию удаления книги
func deleteBook(db *pgxpool.Pool, id int) error {
	_, err := db.Exec(context.Background(),
		"DELETE FROM books where id=$1", id)
	return err
}
func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//err = createTables(db)
	//if err != nil {
	//	panic(err)
	//}
	authorID, _ := addAuthor(db, Author{Name: "Толстой"})
	fmt.Println("Added author with ID:", authorID)
	bookID, _ := addBook(db, Book{Title: "Война и мир",
		AuthorID: authorID,
		Year:     1869})
	fmt.Println("Added book with ID:", bookID)
	books, _ := findBookByAuthor(db, "Толстой")
	fmt.Println("Found:", len(books))

	byYear, _ := countBooksByYear(db)
	fmt.Println("Stats:", byYear)
}

//exec
