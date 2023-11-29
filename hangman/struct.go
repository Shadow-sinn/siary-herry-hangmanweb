package hangman

type Player struct {
	Pseudo string
	Level  string
	Result bool
	Mot    string
	Test   string
	Lettre string
	Score  int
}

var Joueur Player
