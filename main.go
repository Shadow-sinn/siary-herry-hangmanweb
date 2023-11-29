package main

import (
	routeur "hangman-web/routeur"
	initTemplate "hangman-web/temp"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServ()
}
