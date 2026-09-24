package projet_red

import (
	"fmt"
	"math/rand"
)

func openMarchand(perso *character) {

	for {
		fmt.Println("\n========================================")
		fmt.Println("        BOUTIQUE DU MARCHAND          ")
		fmt.Println("========================================")
		fmt.Printf(" Votre or : %d pièces\n", (*perso).Piece)
		fmt.Println("1. Potion de soin (+30 PV) - 15 Or")
		fmt.Println("2. Grande Potion (+50 PV) - 30 Or")
		fmt.Println("3. Baril d'essence (+30 d'essence) - 15 Or")
		fmt.Println("4. Potion de poison (-10 par action) - 20 Or")
		fmt.Println("5. Sacoche à la flèche (+ 10 d'emplacement) - 50 Or")
		fmt.Println("6. Ressource de craft aléatoire - 10 Or")
		fmt.Println("7. Quitter la boutique")
		fmt.Print(" Que voulez-vous faire ? (1-7) : ")

		var input int
		fmt.Scan(&input)

		if len((*perso).Inventaire) < (*perso).Inventaire_max {
			switch input {
			case 1:
				if (*perso).Piece >= 15 {
					(*perso).Piece = (*perso).Piece - 15
					(*perso).Inventaire = append((*perso).Inventaire, "Potion de soin")
					fmt.Println(" Vous avez acheté une Potion de soin !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case 2:
				if (*perso).Piece >= 30 {
					(*perso).Piece -= 30
					(*perso).Inventaire = append((*perso).Inventaire, "Grande Potion de soin")
					fmt.Println(" Vous avez acheté une Grande Potion de soin !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case 3:
				if (*perso).Piece >= 25 {
					(*perso).Piece -= 25
					(*perso).Inventaire = append((*perso).Inventaire, "Baril d'essence")
					fmt.Println(" Vous avez acheté une Baril d'essence !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case 4:
				if (*perso).Piece >= 20 {
					(*perso).Piece -= 20
					(*perso).Inventaire = append((*perso).Inventaire, "Potion de poison")
					fmt.Println(" Vous avez acheté une Potion de poison !")
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case 5:
				if (*perso).NbAchatSacoche < 3 {
					if (*perso).Piece >= 50 {
						(*perso).Piece -= 50
						(*perso).Inventaire_max += 10
						fmt.Println(" Les choses simples... une magnifique sacoche !")
					} else {
						fmt.Println(" Vous n'avez pas assez d'or !")
					}
				} else {
					fmt.Println(" La sacoche est déjà trop grande, pas possible de faire plus (3/3) !")
				}

			case 6:
				if (*perso).Piece >= 10 {
					ressources := []string{"Herbe", "Champignon", "Peau fermenté", "Tissus", "Corde", "Plaque en fer", "Kevlar", "Caoutchouc"}
					ressource := ressources[rand.Intn(len(ressources))]
					(*perso).Piece -= 10
					(*perso).Inventaire = append((*perso).Inventaire, ressource)
					fmt.Println(" Vous avez acheté la ressource :", ressource)
				} else {
					fmt.Println(" Vous n'avez pas assez d'or !")
				}

			case 7:
				fmt.Println("Le marchand vous salue : 'Revenez quand vous voulez l'ami !'")
				return

			default:
				fmt.Println("Choix invalide, veuillez choisir entre 1 et 7 !")
			}
		} else {
			fmt.Println("Poche pleine, pas possible reviens quand tu te seras vidé !")
			return
		}
	}
}
