package projet_red

import "fmt"

func Cheat(perso *character) {
	for {
		fmt.Println("=====================")
		fmt.Println("CHEAT")
		fmt.Println("1 : +100 XP")
		fmt.Println("2 : +1000 pièces")
		fmt.Println("3 : remetre le Res")
		fmt.Println("4 : Combattre un monstre")
		fmt.Println("5 : Aller vers un événement")
		fmt.Println("6 : Quitter le cheat")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			(*perso).Xp += 100
			fmt.Println("Vous gagnez 100 XP.")
		case 2:
			(*perso).Piece += 1000
			fmt.Println("Vous gagnez 1000 pièces.")
		case 3:
			(*perso).Res = true
			fmt.Println("perso.Res est maintenant true.")
		case 4:
			fmt.Println("Choisissez le monstre :")
			fmt.Println("1 : Zombie")
			fmt.Println("2 : Claqueur")
			fmt.Println("3 : Solar")
			fmt.Println("4 : Goliath")
			fmt.Println("5 : M.A.X.I.M.E")
			var monstreChoisi int
			fmt.Scan(&monstreChoisi)

			var nom string
			switch monstreChoisi {
			case 1:
				nom = "Zombie"
			case 2:
				nom = "Claqueur"
			case 3:
				nom = "Solar"
			case 4:
				nom = "Goliath"
			case 5:
				nom = "M.A.X.I.M.E"
			default:
				fmt.Println("Choix invalide.")
				continue
			}

			(*perso).Combat = true
			m := InitMonster(nom)
			Combat(perso, &m)
		case 5:
			fmt.Println("Choisissez l'événement :")
			fmt.Println("1 : Combat")
			fmt.Println("2 : Campement")
			fmt.Println("3 : Marchand")
			fmt.Println("4 : Trouver 10 pièces")
			fmt.Println("5 : Trouver un objet aléatoire")
			var eventChoisi int
			fmt.Scan(&eventChoisi)

			switch eventChoisi {
			case 1:
				fmt.Println("Choisissez le monstre :")
				fmt.Println("1 : Zombie")
				fmt.Println("2 : Claqueur")
				fmt.Println("3 : Solar")
				fmt.Println("4 : Goliath")
				fmt.Println("5 : M.A.X.I.M.E")
				var monstreChoisi int
				fmt.Scan(&monstreChoisi)

				var nom string
				switch monstreChoisi {
				case 1:
					nom = "Zombie"
				case 2:
					nom = "Claqueur"
				case 3:
					nom = "Solar"
				case 4:
					nom = "Goliath"
				case 5:
					nom = "M.A.X.I.M.E"
				default:
					fmt.Println("Choix invalide.")
					continue
				}

				(*perso).Combat = true
				m := InitMonster(nom)
				Combat(perso, &m)
			case 2:
				camp(perso)
			case 3:
				openMarchand(perso)
			case 4:
				(*perso).Piece += 10
				fmt.Println("Vous trouvez 10 pièces.")
			case 5:
				res := randomCraftRessource()
				(*perso).Inventaire = append((*perso).Inventaire, res)
				fmt.Println("Vous trouvez une ressource de craft aléatoire :", res)
			default:
				fmt.Println("Événement invalide.")
			}
		case 6:
			fmt.Println("Fin du cheat.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
