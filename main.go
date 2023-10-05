package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leandro-portugal/MyTask.git/handlers"
	"github.com/leandro-portugal/MyTask.git/settings"
)

func main() {

	err := settings.Load()

	if err != nil {
		panic(err)
	}

	route := chi.NewRouter()
	route.Post("/create", handlers.Insert)
	route.Put("/update/{id}", handlers.Update)
	route.Delete("/delete/{id}", handlers.Delete)
	route.Get("/", handlers.List)
	route.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", settings.GetServerPort()), route)
}
