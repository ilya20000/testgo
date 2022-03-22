package dom

import "net/http"

// Обработчик главной странице.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox home"))
}
