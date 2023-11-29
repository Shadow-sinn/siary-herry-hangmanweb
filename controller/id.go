package controller

import (
	code "hangman-web/hangman"
	initTemplate "hangman-web/temp"
	"net/http"
)

func Id(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "id", nil)
}

func InitId(w http.ResponseWriter, r *http.Request) {
	code.Joueur.Pseudo = r.FormValue("Pseudo")
	http.Redirect(w, r, "/level", http.StatusMovedPermanently)
}
