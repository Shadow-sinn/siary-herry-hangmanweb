package routeur

import (
	"fmt"
	routeur "hangman-web/controller"
	"net/http"
	"os"
)

func InitServ() {
	http.HandleFunc("/", routeur.Id)
	http.HandleFunc("/treatment/id", routeur.InitId)
	http.HandleFunc("/level", routeur.Level)
	http.HandleFunc("/treatment/level", routeur.InitLevel)
	http.HandleFunc("/jeu", routeur.Jeu)
	http.HandleFunc("/treatment/jeu", routeur.InitJeu)
	http.HandleFunc("/result", routeur.Result)

	rootDoc, _ := os.Getwd()
	FileServ := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServ))

	fmt.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
