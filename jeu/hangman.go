package hangman

import (
	"fmt"
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

func Recup() {
	data, err := os.ReadFile("dico_test.txt")
	if err != nil {
		fmt.Println("Fichier vide !")
	}
	os.Stdout.Write(data)
}
