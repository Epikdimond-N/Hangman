package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func Jeu(mot string) {
	var c int
	stock := make([]string, len(mot)-2)
	Displaystock(mot, stock)
	for i := 1; c <= 10; i++ {
		Choose(mot, stock, c)
		if c >= 10 {
			fmt.Println("Vous avez perdu !")
			return
		}
	}
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

func WriteWord(path string) string {
	f, err := ReadLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f))
	fmt.Println(f[ale])
	return f[ale]
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
	estpresent = false
	fmt.Scan(&letter)
	fmt.Scan()
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
	fmt.Println("Le mot c'est :", mot)
	var guess string
	fmt.Scan(&guess)
	fmt.Println(len(guess))
	fmt.Println(len(mot))
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
	fmt.Scan(&choix)
	fmt.Scan()
	fmt.Println(choix)
	switch choix {
	case 1:
		Inputletter(mot, stock, c)
		a = Displaystock(mot, stock)
		if a == len(mot) {
			fmt.Println("\nBien joué ! Vous avez trouvé le mot !")
			fmt.Println(a)
			return
		}
		fmt.Println("\033[H\033[2J")
		Displaystock(mot, stock)
		fmt.Println(c)
	case 2:
		Inputword(mot, c)
		Choose(mot, stock, c)
		fmt.Println("\033[H\033[2J")
		Displaystock(mot, stock)
		fmt.Println(c)
	default:
		fmt.Println("Choix invalide, Veuillez choisir une option valide")
		Choose(mot, stock, c)
		fmt.Println("\033[H\033[2J")
		Displaystock(mot, stock)
	}
}
