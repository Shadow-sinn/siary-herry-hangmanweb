package controller

import (
	code "hangman-web/hangman"
	initTemplate "hangman-web/temp"
	"net/http"
)

func Result(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "result", code.Joueur)
}
