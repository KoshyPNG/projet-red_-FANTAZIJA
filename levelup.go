package projet_red

import "fmt"

func LevelUp(perso *character) {
	if (*perso).Xp/100 > 0 {
		fmt.Println("Vous gagner un Niveau !!!!")
		for (*perso).Xp > 99 {
			(*perso).Xp -= 100
			(*perso).Level += 1
			NewAttack(perso)
		}
	}
}

func NewAttack(perso *character) {
	fmt.Println("Quelle attack vous voulez ?")
	attaques := []string{
		"Coup de matraque",
		"Cri de guerre",
		"Bloquer",
		"Glock 26",
		"Flèchette de poison",
		"Dodge",
		"Lance flamme",
		"Rechargement",
		"Coup de crosse",
		"Coup de griffe",
		"Morsure",
		"Charge",
		"Lancer de RedBull",
		"Morsure de loup",
		"Traque empoisonné",
		"Poing titanesque",
	}

	for {
		for i, val := range attaques {
			fmt.Println(i+1, " : ", val)
		}
		fmt.Println("0 : Annuler")

		var choix int
		fmt.Scan(&choix)
		if choix == 0 {
			return
		}
		if choix < 1 || choix > len(attaques) {
			fmt.Println("Choix invalide.")
			continue
		}

		fmt.Println("Où placer cette attaque dans perso.Action ?")
		for i, val := range (*perso).Action {
			if val != nil {
				fmt.Println(i+1, " : ", val.Nom)
			}
		}
		fmt.Println("0 : Ajouter à la fin")

		var place int
		fmt.Scan(&place)
		if place < 0 || place > len((*perso).Action) {
			fmt.Println("Emplacement invalide.")
			continue
		}

		c := InitAttack(attaques[choix-1])
		if place == 0 {
			return
		}
		(*perso).Action[place-1] = &c
		fmt.Println("Attaque ajoutée :", c.Nom, "à la position", place)
		return
	}
}
