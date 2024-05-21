package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/toastsanwich/management-systems-api/book-management-api/internal/models"
)

func (a *Application) Home(w http.ResponseWriter, r *http.Request) {
	a.InfoLog.Println("called home")
	fmt.Fprintf(w, "This is Home Page")
}

func (a *Application) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		a.ErrorLog.Fatal("invalid method")
		a.MethodError(w)
		return
	}
	a.InfoLog.Println("called createbook")

	newBook := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(newBook)
	if err != nil {
		a.ServerError(w, err, http.StatusInternalServerError)
	}

	err = a.Storage.CreateBook(newBook)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
	a.InfoLog.Println("book created successfully")
}

func (a *Application) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		a.ErrorLog.Fatal("invalid method")
		a.MethodError(w)
		return
	}
	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
	err = a.Storage.DeleteBook(id)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
}

func (a *Application) GetBookByID(w http.ResponseWriter, r *http.Request) {
	a.InfoLog.Println("called getbookbyid handler")

	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
	book, err := a.Storage.GetBookByID(id)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%+v", book)
}

func (a *Application) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	a.InfoLog.Println("called handler getallbooks")

	books, err := a.Storage.GetAllBooks()
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
}

func (a *Application) RecentAdds(w http.ResponseWriter, r *http.Request) {
	a.InfoLog.Println("called handler recentadds")

	recent, err := a.Storage.RecentAdds()
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(recent)
	if err != nil {
		a.ErrorLog.Fatal(err.Error())
		a.ServerError(w, err, http.StatusInternalServerError)
	}
}
