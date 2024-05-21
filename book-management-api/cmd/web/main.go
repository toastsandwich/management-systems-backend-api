package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start() {
	app := LoadApplication()
	app.Routes()
	
	server := http.Server{
		Addr:           "localhost:8080",
		Handler:        app.Mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	fmt.Println("server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
