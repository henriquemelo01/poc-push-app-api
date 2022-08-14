package main

import (
	"awesomeProject/di"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"log"
	"net/http"
)

func main() {

	serveMux := chi.NewRouter()

	// Setup MiddleWares
	serveMux.Use(middleware.Logger)

	// Setup Hello World Handler
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		writer.WriteHeader(http.StatusOK)

		writer.Header().Set("Content-Type", "text/plain")

		_, err := io.WriteString(writer, "Hello World")

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	// Setup Controllers
	reportsController := di.CreateReportsController()
	reportsController.SetupRoute(serveMux)

	// Setup Server
	httpServer := http.Server{
		Addr:    "localhost:8080",
		Handler: serveMux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("Houve um erro por aqui: ", err.Error())
	}
}
