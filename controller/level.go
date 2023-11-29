package controller

import (
	code "hangman-web/hangman"
	initTemplates "hangman-web/temp"
	"net/http"
)

func Level(w http.ResponseWriter, r *http.Request) {
	initTemplates.Temp.ExecuteTemplate(w, "level", nil)
}

func InitLevel(w http.ResponseWriter, r *http.Request) {
	code.Joueur.Level = r.FormValue("level")
	code.Joueur.Init()
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}
