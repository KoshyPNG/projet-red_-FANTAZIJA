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
				fmt.Println("3. Casquette Gucci fraise ")
				fmt.Println("4. Fourure Canada Goose ")
				fmt.Println("5. Dior B30 ")
				fmt.Println("6. Retour au menu du camp ")

				fmt.Println(" Que voulez-vous faire ? (1-3) : ")
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
						remoov(index)
						(*perso).Inventaire = append((*perso).Inventaire, objet)
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Potion de soin !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas d'herbe, pas de potion ...")
					}

				case "2":

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
						remoov(index)
						(*perso).Inventaire = append((*perso).Inventaire, objet)
						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Potion de poison !")

					} else {
						fmt.Println("=============================")
						fmt.Println("Pas de champignon, pas de poison ...")
					}

				case "3":

					objet := "Casquette Gucci fraise"
					elementASupprimer := "Tissus"
					elementASupprimer2 := "Corde"
					index1 := -1
					index2 := -1

					for i, v := range (*perso).Inventaire {
						if v == element1 && index1 == -1 {
							index1 = i
						} else if v == element2 && index2 == -1 {
							index2 = i
						}
					}
					if index1 != -1 && index2 != -1 {
						if index1 > index2 {
							remoov(index1)
							remoov(index2)
						} else {
							remoov(index2)
							remoov(index1)
						}

						(*perso).Inventaire = append((*perso).Inventaire, objet)

						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une Casquette Gucci fraise !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas assez de matières ...")
					}

				case "4":

					objet := "Fourure Canada Goose"
					elementASupprimer := "Fourrure de loup"
					elementASupprimer2 := "Corde"
					index1 := -1
					index2 := -1

					for i, v := range (*perso).Inventaire {
						if v == element1 && index1 == -1 {
							index1 = i
						} else if v == element2 && index2 == -1 {
							index2 = i
						}
					}
					if index1 != -1 && index2 != -1 {
						if index1 > index2 {
							remoov(index1)
							remoov(index2)
						} else {
							remoov(index2)
							remoov(index1)
						}

						(*perso).Inventaire = append((*perso).Inventaire, objet)

						fmt.Println("=============================")
						fmt.Println(" Vous avez crée un Fourure Canada Goose !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas assez de matières ...")
					}

				case "5":

					objet := "Dior B30"
					elementASupprimer := "Tissus"
					elementASupprimer2 := "Caoutchouc"
					index1 := -1
					index2 := -1

					for i, v := range (*perso).Inventaire {
						if v == element1 && index1 == -1 {
							index1 = i
						} else if v == element2 && index2 == -1 {
							index2 = i
						}
					}
					if index1 != -1 && index2 != -1 {
						if index1 > index2 {
							remoov(index1)
							remoov(index2)
						} else {
							remoov(index2)
							remoov(index1)
						}

						(*perso).Inventaire = append((*perso).Inventaire, objet)

						fmt.Println("=============================")
						fmt.Println(" Vous avez crée une paire de Dior B30 !")
					} else {
						fmt.Println("=============================")
						fmt.Println("Pas assez de matières ...")
					}

				case "6":

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
