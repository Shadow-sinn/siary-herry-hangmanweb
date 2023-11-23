package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	wins, losses, abortions int
	commonWords             = []string{"leur", "aurait", "à propos", "là", "penser", "lequel", "gens", "pourrait", "autre", "ces"} // Liste de mots simplifiée
	gameInProcess           bool
	answer, maskedAnswer   string
	wrongGuesses            int
)

func main() {
	gameInProcess = false
	answer = newRandomWord()

	// Commence un nouveau jeu
	newGame()
}

func newGame() {
	if gameInProcess {
		aborted()
	}
	gameInProcess = true
	fmt.Println("Jeu de pendu")
	maskedAnswer = strings.Repeat("_", len(answer))
	wrongGuesses = 0
	resetKeypad()
	updateDisplayWord()
}

func newRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return commonWords[rand.Intn(len(commonWords))]
}

func verifyGuess(guessedLetter rune) {
	if strings.ContainsRune(answer, guessedLetter) {
		// Logique pour traiter la lettre correcte devinée
	} else {
		// Logique pour traiter la lettre incorrecte devinée
	}
}

func updateDisplayWord() {
	display := ""
	for _, char := range maskedAnswer {
		display += string(char) + " "
	}
	display = display[:len(display)-1] // Supprimer le dernier espace
	fmt.Println(display)
}

func aborted() {
	abortions++
	fmt.Println("Nombre d'abandons:", abortions)
}

func resetKeypad() {
	fmt.Println("Réinitialiser le clavier")
	// Logique pour réinitialiser le clavier (pas de manipulation réelle du DOM)
	// Vous pouvez créer une structure de clavier pour stocker les lettres
}