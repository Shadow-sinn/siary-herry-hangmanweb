package controller

import (
	code "hangman-web/hangman"
	initTemplates "hangman-web/temp"
	"net/http"
)

func Jeu(w http.ResponseWriter, r *http.Request) {
	if code.Joueur.Lettre == "" {
		initTemplates.Temp.ExecuteTemplate(w, "jeu", nil)
	}
	if code.IsInWord(code.Joueur.Mot, code.Joueur.Lettre) {
		code.Joueur.GuessLetter()
	} else {
		code.Joueur.Score++
		if code.Joueur.Score >= 11 {
			code.Joueur.Result = false
			http.Redirect(w, r, "/result", http.StatusMovedPermanently)
		}
	}
	if !code.Joueur.IsUnderscore() {
		code.Joueur.Result = true
		http.Redirect(w, r, "/result", http.StatusMovedPermanently)
	}
	initTemplates.Temp.ExecuteTemplate(w, "jeu", code.Joueur)
}

func InitJeu(w http.ResponseWriter, r *http.Request) {
	code.Joueur.Lettre = r.FormValue("Lettre")
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}
