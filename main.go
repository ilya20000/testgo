package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"./dom"
	"./login"
)

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Метод запрещен!", 405)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", dom.Home)
	mux.HandleFunc("/v1/snippet", showSnippet)
	mux.HandleFunc("/v1/snippet/create", createSnippet)
	mux.HandleFunc("/v1/reg", login.Reg)

	log.Println("Запуск сервера на :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
