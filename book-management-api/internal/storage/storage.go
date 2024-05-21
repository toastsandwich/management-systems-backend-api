package storage

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/toastsanwich/management-systems-api/book-management-api/internal/models"
)

type Storage interface {
	CreateBook(*models.Book) error
	DeleteBook(int) error
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(int) (*models.Book, error)
	RecentAdds() ([]*models.Book, error)
}

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore() (*MySQLStore, error) {
	cfg := mysql.Config{
		User:                 "gomon",
		Passwd:               "smpmsmim",
		Net:                  "tcp",
		DBName:               "bookManagementApi",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
	}
	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping fail")
	}
	return &MySQLStore{
		db: db,
	}, nil
}

func (m *MySQLStore) CreateBook(book *models.Book) error {
	qry := `INSERT INTO books VALUES(title, genre, description, isbn, page_count, cost) (
		?,
		?,
		?,
		?,
		?,
		?
	)`
	_, err := m.db.Exec(qry,
		book.Title,
		book.Genre,
		book.Description,
		book.ISBN,
		book.PageCount,
		book.Cost,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQLStore) DeleteBook(id int) error {
	qry := `DELETE FROM books WHERE id = ?`
	_, err := m.db.Exec(qry, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQLStore) GetAllBooks() ([]*models.Book, error) {
	qry := `SELECT * FROM books`

	rows, err := m.db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []*models.Book{}
	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Genre, &book.Description, &book.ISBN, &book.PageCount, &book.Cost)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return books, nil
}

func (m *MySQLStore) GetBookByID(id int) (*models.Book, error) {
	qry := `SELECT * FROM books WHERE id = ?`
	row := m.db.QueryRow(qry, id)

	book := &models.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Genre, &book.Description, &book.ISBN, &book.PageCount, &book.Cost)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (m *MySQLStore) RecentAdds() ([]*models.Book, error) {
	qry := `SELECT * FROM books ORDER BY id DEC LIMIT 10`
	books := []*models.Book{}
	rows, err := m.db.Query(qry)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &models.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Genre, &book.Description, &book.ISBN, &book.PageCount, &book.Cost)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
