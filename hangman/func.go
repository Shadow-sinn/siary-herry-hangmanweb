package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func (p *Player) Init() {
	p.Mot = WriteWord(p.Level)
	p.Result = false
	p.Test = Count(p.Mot)
	p.Lettre = ""
	p.Score = 0
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteWord(s string) string {
	f, err := ReadLines("mot/" + s + ".txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f))
	return f[ale]
}

func Count(m string) string {
	var guess string

	for n := 0; n < len(m); n++ {
		guess += "_ "
	}
	return guess
}

func IsInWord(word, s string) bool {
	for _, l := range word {
		if string(l) == s {
			return true
		}
	}
	return false
}

func TransformString(s string) []string {
	slice := []string{}
	for _, c := range s {
		slice = append(slice, string(c))
	}
	return slice
}

func TransformSlice(s []string) string {
	var str string
	for _, c := range s {
		str += c
	}
	return str
}

func (p *Player) IsUnderscore() bool {
	for _, c := range p.Mot {
		if string(c) == "_" {
			return false
		}
	}
	return true
}

func (p *Player) TestWord(guess, mot string) bool {
	if mot == guess {
		fmt.Println("Vous avez trouvé le mot")
		p.Result = true
	} else {
		fmt.Println("Ce n'est pas le bon mot")
	}
	return p.Result
}

func (p *Player) GuessLetter() {
	for i, t := range p.Mot {
		if string(t) == p.Lettre {
			slc := TransformString(p.Test)
			slc[i*2] = p.Lettre
			p.Test = TransformSlice(slc)
		}
	}
}
