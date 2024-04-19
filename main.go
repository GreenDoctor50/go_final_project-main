package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/greendoctor50/go_final_project-main/packages/handler/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// здесь регистрируйте ваши обработчики
	r.Get("/", getIndexHTML)
	if err := http.ListenAndServe(":7540", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}