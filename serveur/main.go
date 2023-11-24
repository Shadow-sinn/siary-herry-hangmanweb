package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	hangman_classic "hangman-web/hangman_classic_for_students"
)

type game struct {
	State          string
	Randomwordhide string
	Randomword     string
	Essaie         int
	Usedletter     []string
}

var Game game

func NewGame() {

	Game = game{}
	// On récupère les arguments
	args := os.Args[1:]
	// S'il n'y a pas qu'un seul argument, on arrête le programme
	if len(args) != 1 {
		os.Exit(0)
	}
	// Sinon on ouvre le fichier correspondant au nom passé en argument
	f, _ := os.OpenFile(args[0], os.O_RDWR, 0644)

	scanner := bufio.NewScanner(f)

	// On déclare notre liste de mots
	wordlist := []string{}

	// On ajoute tous les mots à notre liste de mots
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	// On initialise le random
	rand.Seed(time.Now().UnixNano())
	// On génère un nombre aléatoire qui va récupérer un mot grâce à son index dans notre liste de mots
	Game.Randomword = wordlist[rand.Intn(len(wordlist))]
	// On crée notre mot avec quelques lettres d'afficher
	Game.Randomwordhide = hangman_classic.CreateWord(Game.Randomword)
	// On déclare nos variables qui vont servir à l'avancée de notre jeu
	Game.State = ""
	Game.Essaie = 10
	Game.Usedletter = []string{}

}

func main() {

	NewGame()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		t := template.Must(template.ParseFiles("index.html"))
		t.Execute(res, Game)
		//http.ServeFile(res, req, "../index.html")

	})

	http.HandleFunc("/hangman", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			http.Redirect(res, req, "/", 301)
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(res, "ParseForm() err: %v", err)
				return
			}
			fmt.Printf("POST from website r.PostFrom = %v\n", req.PostForm)
			LettreReq := (req.Form.Get("lettre"))
			// On vérifie si ce qu'il a marqué est valide
			Game.Randomwordhide, Game.State = hangman_classic.IsInputOk(LettreReq, Game.Randomword, Game.Randomwordhide, &Game.Usedletter)
			// On clear le terminal
			hangman_classic.Clear()
			// Si le joueur a déjà fait une erreur
			if Game.Essaie != 10 {
				// On affiche l'état du pendu
				//fmt.Println(hang[9-Game.Essaie])
			}
			// Si ce que le joueur a marqué n'est pas valide
			if Game.State == "fail" {
				// On diminue le nombre d'essai restant du joueur
				Game.Essaie--
				// On affiche l'état de la partie
				//fmt.Print(hang[9-essaie])
				fmt.Printf("La lettre %v n'est pas comprise dans le mot, il ne te reste plus que : %v essais\n", LettreReq, Game.Essaie)

				// Si la lettre a déjà été utilisé
			} else if Game.State == "usedletter" {
				// On affiche le message correspondant
				fmt.Printf("Lettre déjà utiliser\n")
				// Si la lettre est valide
			} else if Game.State == "good" {
				// On affiche le message correspondant
				fmt.Printf("La lettre %v est bien comprise dans le mot\n", LettreReq)
				// Si le mot rentré n'est pas de la bonne taille
			} else if Game.State == "wordinvalid" {
				// On affiche le message correspondant
				fmt.Printf("Le format n'est pas valide, veuillez rentrer une lettre ou un mot de bonne taille\n")
				// Si le mot est le bon
			} else if Game.State == "wordgood" {
				// On affiche le message correspondant
				fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", Game.Essaie, Game.Randomword)
				// On arrête le programme
				Game.State = "win"
				// Si l'input n'est pas une lettre
			} else if Game.State == "error" {
				// On affiche le message correspondant
				fmt.Println("La lettre est invalide, veuillez recommencer")
				// Si la mot n'est pas le bon
			} else if Game.State == "wordwrong" {
				// On retire 2 essais au lieu de 1
				Game.Essaie -= 2
				// On affiche le message correspondant
				fmt.Printf("Le mot proposé n'est pas le bon, il te reste %v essais\n", Game.Essaie)
			}
			// Si le joueur n'a plus d'essais
			if Game.Essaie <= 0 {
				// On clear le terminal
				hangman_classic.Clear()
				// On affiche l'état de la partie
				//fmt.Print(hang[9])
				fmt.Printf("Tu as perdu, le mot était : %v", Game.Randomword)
				// On arrête le programme
				Game.State = "lose"
			}
			fmt.Println(Game.Randomwordhide)
			// Si le mot a été totalement découvert
			if Game.Randomwordhide == Game.Randomword {
				// On clear le terminal
				hangman_classic.Clear()
				// On affiche le message correspondant
				fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", Game.Essaie, Game.Randomword)
				// On arrête le programme
				// os.Exit(0)
				Game.State = "win"
			}
		}

		http.Redirect(res, req, "/", 301)

	})
	http.HandleFunc("/restart", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			http.Redirect(res, req, "/", 301)
		case "POST":
			NewGame()
			http.Redirect(res, req, "/", 301)
		}
	})

	fs := http.FileServer(http.Dir("style/"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fmt.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
