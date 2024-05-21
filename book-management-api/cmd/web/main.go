package web

import (
	"fmt"
	"net/http"
)

func Start() {
	app := LoadApplication()
	fmt.Println("server on port 4000")
	http.ListenAndServe(":4000", app.Mux)
}
