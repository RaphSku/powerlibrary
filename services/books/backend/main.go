package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/handlers"
	"github.com/RaphSku/powerlibrary/tree/main/services/books/validation"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	logger := log.New(os.Stdout, "books", log.LstdFlags)

	servermux := mux.NewRouter()
	getRouter := servermux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/books/", handlers.GetBooks)
	getRouter.HandleFunc("/api/v1/books/{id:[0-9]+}", handlers.GetBookById)

	postRouter := servermux.Methods(http.MethodPost).Subrouter()
	postRouter.Use(validation.ValidateRequestBodyMw)
	postRouter.HandleFunc("/api/v1/book/", handlers.PostBook)
	postRouter.HandleFunc("/api/v1/books/", handlers.PostBooks)

	putRouter := servermux.Methods(http.MethodPut).Subrouter()
	putRouter.Use(validation.ValidateRequestBodyMw)
	putRouter.HandleFunc("/api/v1/books/{id:[0-9]+}", handlers.PutBook)

	deleteRouter := servermux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/api/v1/books/{id:[0-9]+}", handlers.DeleteBook)

	handler := cors.Default().Handler(servermux)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      handler,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-signalChannel
	logger.Println("Graceful shutdown", sig)

	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	server.Shutdown(timeoutContext)
}
