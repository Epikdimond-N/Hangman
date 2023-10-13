package hangman

import (
	"fmt"
	"math/rand"
	"os"
)

/*func jeu(mot string) {
	var c = 0
	fmt.Println("Le jeu du pendu :")
	for i := 0; c <= 10; i++ {
		var lettre string
		fmt.Println("Entrez une lettre")
		fmt.Scanln(&lettre)
		//parcour d'indice mot
		// si la letter = une lettre du mot
		for i := 0; i <= len(mot); i++ {
			if lettre == mot[i] {

			}
		}

	}

}*/

var Data []byte
var err error

func Recup() {
	Data, err = os.ReadFile("dico_test.txt")
	if err != nil {
		fmt.Println("Fichier vide !")
	}
	os.Stdout.Write(Data)
}

var Sdata string

func Convert() {
	Sdata = (string(Data))
	/* 	for _, b := range Data {
	   		Sdata = append(Sdata, string(b))
	   	}
	   	return Sdata */
}

var TabData []string

func Convertstr() []string {
	fmt.Println(Sdata)
	var Tampon string = ""
	for _, b := range Sdata {
		i := 0
		Tampon += string(b)
		// fmt.Print("Tampon : " + Tampon)
		if b == '\n' {
			i += 1
			TabData = append(TabData, Tampon)
			Tampon = ""
		}
	}
	//fmt.Println(TabData)
	return TabData
}

var MotR string

func Motrandom() {
	d := len(TabData)
	fmt.Println(d)
	m := rand.Intn(d)
	MotR = TabData[m]
}
