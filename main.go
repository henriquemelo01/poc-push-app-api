package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"poc-push-app-api/di"
)

func main() {

	serveMux := chi.NewRouter()

	// Setup MiddleWares
	serveMux.Use(middleware.Logger)

	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {

		writer.WriteHeader(http.StatusOK)

		writer.Header().Set("Content-Type", "text/plain")

		_, _ = writer.Write([]byte("Pong 🧙‍♂️"))
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
