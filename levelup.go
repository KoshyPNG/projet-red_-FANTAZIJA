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
	a := []string{
		"Coup de matraque",
		"Cri de guerre",
		"Bloquer",
		"Glock 26",
		"Flèchette de poison",
		"Dodge",
		"Lance flamme",
		"Rechargement",
		"Coup de crosse",
		"Charge",
		"Lancer de RedBull",
		"Morsure de loup",
		"Traque empoisonné",
		"Poing titanesque",
	}


	
	for {
		for i, val := range a {
			fmt.Println(i+1 ," : ",val)
		}
		fmt.Println("0 : rien")
		var choix int
		fmt.Scan(&choix)
		if choix > 0 && choix < 4 {
			var place int
			fmt.Scan(&place)
			for i, val := range (*perso).Action {
				fmt.Println(i+1 ," : ",val)
			}
			if place > 0 && place < 6 {
				c := InitAttack(a[choix])
				(*perso).Action[place] = &c
				return
			}
		} else {
			return
		}
	}
}
