package web

import (
	"log"
	"net/http"
	"os"

	"github.com/toastsanwich/management-systems-api/book-management-api/internal/storage"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Storage  storage.Storage
	Mux      *http.ServeMux
}

func LoadApplication() *Application {
	return &Application{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Storage: func() *storage.MySQLStore {
			store, err := storage.NewMySQLStore()
			if err != nil {
				panic(err)
			}
			return store
		}(),
		Mux: http.NewServeMux(),
	}
}

func (app *Application) Routes() {
	app.Mux.HandleFunc("/", app.Home)
	app.Mux.HandleFunc("/book/create", app.CreateBook)
	app.Mux.HandleFunc("/book/get", app.GetBookByID)
	app.Mux.HandleFunc("/book/delete", app.DeleteBook)
	app.Mux.HandleFunc("/book/all", app.GetAllBooks)
	app.Mux.HandleFunc("/book/recent", app.RecentAdds)
}
