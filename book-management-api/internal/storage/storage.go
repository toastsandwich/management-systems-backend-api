package storage

import (
	"database/sql"

	"github.com/toastsanwich/management-systems-api/book-management-api/internal/models"
)

type Storage interface {
	CreateBook(*models.Book) error
	DeleteBook(int) error
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(int) (*models.Book, error)
	UpdateBook(*models.Book) error
	RecentAdds() ([]*models.Book, error)
}

type MySQLStore struct {
	db *sql.DB
}

func (m *MySQLStore) OpenDB() {

}

func SQLStore() *MySQLStore {
	return nil
}

func (m *MySQLStore) CreateBook(book *models.Book) error {
	return nil
}

func (m *MySQLStore) DeleteBook(id int) error {
	return nil
}
func (m *MySQLStore) GetAllBooks() ([]*models.Book, error) {
	return nil, nil
}

func (m *MySQLStore) GetBookByID(id int) (*models.Book, error) {
	return nil, nil
}

func (m *MySQLStore) RecentAdds() ([]*models.Book, error) {
	return nil, nil
}

func (m *MySQLStore) UpdateBook(book *models.Book) error {
	return nil
}
