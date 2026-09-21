package projet_red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func camp(perso *character) {
	reader := bufio.NewReader(os.Stdin)
	bo := true
	for {
		fmt.Println("\n========================================")
		fmt.Println("               CAMPEMENT                ")
		fmt.Println("========================================")
		fmt.Printf(" Votre vie : %d Vie\n", (*perso).Vie_actuel)
		fmt.Println("Que voulez vous faire ?")
		fmt.Println("1. Se reposer (restore votre vie à 100%)")
		fmt.Println("2. Bricolage ")
		fmt.Println("3. Quitter la camp")
		fmt.Print(" Que voulez-vous faire ? (1-3) : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			fmt.Println("Un bon dodo !")
			(*perso).Vie_actuel += 20
			if (*perso).Vie_actuel > (*perso).Vie_max {
				(*perso).Vie_actuel = (*perso).Vie_max
			}
			fmt.Println(" Plein d'énergie !")
			return

		case "2":

			for bo {
				fmt.Println("=============================")
				fmt.Println("Que voulez vous fabriquer ?")
				fmt.Println("=============================")
				fmt.Println("1. Petite potion de soin")
				fmt.Println("2. Potion de poison ")
				fmt.Println("3. Retour au menu du camp ")
				fmt.Println(" Que voulez-vous faire ? (1-3) : ")
				fmt.Println("=============================")

				input2, _ := reader.ReadString('\n')
				input2 = strings.TrimSpace(input2)

				switch input2 {

				case "1":

					elementASupprimer := "Herbe"
					index := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer {
							index = i
							break
						}
					}

					if index != -1 {
						(*perso).Inventaire = append((*perso).Inventaire[:index], (*perso).Inventaire[index+1:]...)
						(*perso).Inventaire = append((*perso).Inventaire, "Potion de soin")
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Potion de soin !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas d'herbe, pas de potion ...")
					}

				case "2":

					elementASupprimer := "Champignon"
					index := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer {
							index = i
							break
						}
					}

					if index != -1 {
						(*perso).Inventaire = append((*perso).Inventaire[:index], (*perso).Inventaire[index+1:]...)
						(*perso).Inventaire = append((*perso).Inventaire, "Potion de poison")
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Potion de poison !")

					} else {
						fmt.Println("=============================")
						fmt.Println("Pas de champignon, pas de poison ...")
					}

				case "3":

					bo  = false

				default:
					fmt.Println("=============================")
					fmt.Println("Choix invalide, veuillez choisir entre 1 et 3 !")
				}
			}

		case "3":
			fmt.Println("=============================")
			fmt.Println("Il est temps de se remmetre en route!'")
			return

		default:
			fmt.Println("=============================")
			fmt.Println("Choix invalide, veuillez choisir entre 1 et 3 !")
		}
	}
}
