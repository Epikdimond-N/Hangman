package main

import (
	h "hangman/jeu"
)

func main() {
	var mot = h.WriteWord("dico_test.txt")
	h.Jeu(mot)
}
