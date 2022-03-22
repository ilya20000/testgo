package login

import (
	"encoding/json"
	"net/http"
)

func Reg(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1/reg" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	books := "sssss"
	json.NewEncoder(w).Encode(books)

}
