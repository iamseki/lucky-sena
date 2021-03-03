package main

import (
	"log"
	irisapp "lucky-sena/main/http/iris-app"
	"net/http"
	"os"
	"time"
)

func main() {
	app := irisapp.New()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	addr := `localhost:` + port
	srv := &http.Server{
		Addr:         addr,
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println(`listen on`, addr)
	log.Fatalln(srv.ListenAndServe())
}
