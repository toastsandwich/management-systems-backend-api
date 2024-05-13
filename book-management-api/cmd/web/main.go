package web

import (
	"net/http"
)

func Start() {
	app := LoadApplication()
	http.ListenAndServe(":4000", app.Mux)
}
