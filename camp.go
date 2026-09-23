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
			(*perso).Essence_actuel += 20
			if (*perso).Essence_actuel > (*perso).Essence_max {
				(*perso).Essence_actuel = (*perso).Essence_max
			}
			fmt.Println(" Plein d'énergie !")
			return

		case "2":

			for bo {
				fmt.Println("=============================")
				fmt.Println("Que voulez vous fabriquer ?")
				fmt.Println("=============================")
				fmt.Println("1. Potion de soin")
				fmt.Println("2. Potion de mana")
				fmt.Println("3. Potion de poison ")
				fmt.Println("4. Chapeau de l'aventurier ")
				fmt.Println("5. Tunique de l'aventurie ")
				fmt.Println("6. Bottes de l'aventurier ")
				fmt.Println("7. Retour au menu du camp ")

				fmt.Println(" Que voulez-vous faire ? (1-7) : ")
				fmt.Println("=============================")

				input2, _ := reader.ReadString('\n')
				input2 = strings.TrimSpace(input2)

				switch input2 {

				case "1":

					objet := "Potion de soin"
					elementASupprimer := "Herbe"
					index := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer {
							index = i
							break
						}
					}

					if index != -1 {
						remoov(index, perso)
						(*perso).Inventaire = append((*perso).Inventaire, objet)
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Potion de soin !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas d'herbe, pas de potion ...")
					}

				case "2":

					objet := "Baril d'essence"
					elementASupprimer := "Peau fermenté"
					index := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer {
							index = i
							break
						}
					}

					if index != -1 {
						remoov(index, perso)
						(*perso).Inventaire = append((*perso).Inventaire, objet)
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée un Baril d'essence !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas de peau, vous en aurez peut être la prochaine fois ...")
					}

				case "3":

					objet := "Potion de poison"
					elementASupprimer := "Champignon"
					index := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer {
							index = i
							break
						}
					}

					if index != -1 {
						remoov(index, perso)
						(*perso).Inventaire = append((*perso).Inventaire, objet)
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Potion de poison !")

					} else {
						fmt.Println("=============================")
						fmt.Println("Pas de champignon, pas de poison ...")
					}

				case "4":

					objet := "Chapeau de l'aventurier"
					elementASupprimer := "Tissus"
					elementASupprimer2 := "Corde"
					index1 := -1
					index2 := -1

					for i, v := range (*perso).Inventaire {
						if v ==elementASupprimer && index1 == -1 {
							index1 = i
						} else if v ==elementASupprimer2 && index2 == -1 {
							index2 = i
						}
					}
					if index1 != -1 && index2 != -1 {
						if index1 > index2 {
							remoov(index1, perso)
							remoov(index2, perso)
						} else {
							remoov(index2, perso)
							remoov(index1, perso)
						}

						(*perso).Inventaire = append((*perso).Inventaire, objet)

						fmt.Println("=============================")
						fmt.Println(" Vous avez crée un Chapeau de l'aventurier !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas assez de matières ...")
					}

				case "5":

					objet := "Tunique de l'aventurie"
					elementASupprimer := "Fourrure de loup"
					elementASupprimer2 := "Corde"
					index1 := -1
					index2 := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer && index1 == -1 {
							index1 = i
						} else if v == elementASupprimer2 && index2 == -1 {
							index2 = i
						}
					}
					if index1 != -1 && index2 != -1 {
						if index1 > index2 {
							remoov(index1, perso)
							remoov(index2, perso)
						} else {
							remoov(index2, perso)
							remoov(index1, perso)
						}

						(*perso).Inventaire = append((*perso).Inventaire, objet)

						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Tunique de l'aventurie !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas assez de matières ...")
					}

				case "6":

					objet := "Bottes de l'aventurier"
					elementASupprimer := "Tissus"
					elementASupprimer2 := "Caoutchouc"
					index1 := -1
					index2 := -1

					for i, v := range (*perso).Inventaire {
						if v == elementASupprimer && index1 == -1 {
							index1 = i
						} else if v == elementASupprimer2 && index2 == -1 {
							index2 = i
						}
					}
					if index1 != -1 && index2 != -1 {
						if index1 > index2 {
							remoov(index1, perso)
							remoov(index2, perso)
						} else {
							remoov(index2, perso)
							remoov(index1,perso)
						}

						(*perso).Inventaire = append((*perso).Inventaire, objet)

						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une paire de Bottes de l'aventurier !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas assez de matières ...")
					}

				case "7":

					bo = false

				default:
					fmt.Println("=============================")
					fmt.Println("Choix invalide, veuillez choisir entre 1 et 6 !")
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
