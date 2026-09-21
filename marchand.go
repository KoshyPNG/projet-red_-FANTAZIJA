package projet_red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func openMarchand(perso *character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========================================")
		fmt.Println("        BOUTIQUE DU MARCHAND          ")
		fmt.Println("========================================")
		fmt.Printf(" Votre or : %d pièces\n", (*perso).Piece)
		fmt.Println("1. Potion de soin (+50 PV) - 15 Or")
		fmt.Println("2. Grande Potion (+100 PV) - 30 Or")
		fmt.Println("3. Potion de poison (-5 PV / s) - 20 Or")
		fmt.Println("4. Quitter la boutique")
		fmt.Print(" Que voulez-vous faire ? (1-4) : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len((*perso).Inventaire) <= 10 {
			switch input {
			case "1":
				if (*perso).Piece >= 15 {
					(*perso).Piece = (*perso).Piece - 15
					(*perso).Inventaire = append((*perso).Inventaire, "Potion de soin")
					fmt.Println(" Vous avez acheté une Potion de soin !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case "2":
				if(*perso).Piece >= 30 {
					(*perso).Piece -= 30
					(*perso).Inventaire = append((*perso).Inventaire, "Grande Potion de soin")
					fmt.Println(" Vous avez acheté une Grande Potion de soin !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case "3":
				if (*perso).Piece >= 20 {
					(*perso).Piece -= 20
					(*perso).Inventaire = append((*perso).Inventaire, "Potion de poison")
					fmt.Println(" Vous avez acheté une Potion de poison !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case "4":
				fmt.Println("Le marchand vous salue : 'Revenez quand vous voulez l'ami !'")
				return

			default:
				fmt.Println("Choix invalide, veuillez choisir entre 1 et 3 !")
			}
		} else {
			fmt.Println("Poche pleine, pas possible reviens quand tu te seras vidé !")
		}
	}
}
