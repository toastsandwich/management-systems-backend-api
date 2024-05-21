package web

import (
	"net/http"
)

func (app *Application) ServerError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

func (app *Application) MethodError(w http.ResponseWriter) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
