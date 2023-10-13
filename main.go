package main

import (
	"fmt"
	h "hangman/jeu"
)

func main() {
	h.Recup()
	h.Convert()
	h.Convertstr()
	h.Motrandom()
	fmt.Println(h.MotR)
	//fmt.Println(h.TabData[0])

}
