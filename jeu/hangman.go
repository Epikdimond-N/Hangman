package hangman

import (
	"fmt"
	"os"
)

func Jeu(mot string) {
	var c int
	stock := make([]string, len(mot))
	Displaystock(mot, stock)
	for i := 1; c <= 10; i++ {
		Choose(mot, stock, c)
		if c >= 10 {
			fmt.Println("Vous avez perdu !")
			return
		}
	}
}

func Recup() {
	data, err := os.ReadFile("dico_test.txt")
	if err != nil {
		fmt.Println("Fichier vide !")
	}
	os.Stdout.Write(data)
}

func Displaystock(mot string, stock []string) int {
	var lenmo int
	for _, i := range stock {
		if i == "" {
			fmt.Print("_ ")
		} else {
			fmt.Print(i, " ")
			lenmo++
		}
	}
	return lenmo
}

func Inputletter(mot string, stock []string, c int) {
	fmt.Println("\n\nEntrez une lettre : ")
	var letter string
	var estpresent bool
	fmt.Scan(&letter)
	for i := 0; i <= len(mot)-1; i++ {
		if letter == string(mot[i]) {
			stock[i] = letter
			estpresent = true
		}
	}
	if !estpresent {
		c++
	}
	estpresent = false
}

func Inputword(mot string, c int) {
	fmt.Println("Entrez le mot que vous pensez bon")
	var guess string
	fmt.Scan(&guess)
	if guess == mot {
		fmt.Println("C'est le bon mot !")
	} else {
		fmt.Println("Ce n'est pas le bon mot...")
		c++
	}
}

func Choose(mot string, stock []string, c int) {

	var choix int
	var a int
	fmt.Println("\nChoisissez une option : \n1. Emettre une hypothèse sur une lettre présente dans le mot\n2. Entrer directement le mot")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		Inputletter(mot, stock, c)
		a = Displaystock(mot, stock)
		if a == len(mot) {
			fmt.Println("\nBien joué ! Vous avez trouvé le mot !")
			fmt.Println(a)
			return
		} else {
			Choose(mot, stock, c)
		}
	case 2:
		Inputword(mot, c)
		Choose(mot, stock, c)
	default:
		fmt.Println("Choix invalide, Veuillez choisir une option valide")
		Choose(mot, stock, c)
	}
}
