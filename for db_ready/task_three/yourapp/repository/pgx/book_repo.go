package pgxrepo

//реализация интерфейса через pgxpool
import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
	"yourapp/domain"
)

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) Create(ctx context.Context, b domain.Book) (int, error) {
	var id int
	err := r.db.QueryRow(ctx,
		`INSERT INTO books (title, author_id, publication_year)  VALUES($1, $2, $3) RETURNING id`,
		b.Title, b.AuthorId, b.PublicationYear).Scan(&id)
	return id, err
}

func (r *BookRepo) GetByID(ctx context.Context, id int) (*domain.Book, error) {
	var book domain.Book
	err := r.db.QueryRow(ctx,
		`SELECT id, title, author_id, publication_year FROM books WHERE id=$1;`, id,
	).Scan(&book.Id, &book.Title, &book.AuthorId, &book.PublicationYear)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepo) List(ctx context.Context, f domain.BookFilter) ([]domain.Book, error) {
	var args []any
	var where []string
	if f.Title != nil && *f.Title != "" {
		args = append(args, "%"+*f.Title+"%")
		where = append(where, fmt.Sprintf("b.title ILIKE $%d", len(args)))
	}
	if f.Author != nil && *f.Author != "" {
		args = append(args, "%"+*f.Author+"%")
		where = append(where, fmt.Sprintf("a.name ILIKE $%d", len(args)))
	}
	q := `
    SELECT  b.id, b.title, b.author_id, b.publication_year
	FROM books b 
	LEFT JOIN authors a ON a.id=b.author_id`
	if len(where) > 0 {
		q += " WHERE " + strings.Join(where, " AND ") + "\n"
	}
	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.PublicationYear); err != nil {
			return nil, err
		}
		out = append(out, book)
	}
	return out, rows.Err()
}
func (r *BookRepo) Update(ctx context.Context, b domain.Book) error {
	_, err := r.db.Exec(ctx, `UPDATE books
SET title=$1, author_id=$2, publication_year=$3
WHERE id=$4;`,
		b.Title, b.AuthorId, b.PublicationYear, b.Id)
	return err
}

func (r *BookRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM books WHERE id=$1;`, id)
	return err
}
