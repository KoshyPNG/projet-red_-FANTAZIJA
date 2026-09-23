package projet_red

import "fmt"

func Feu(perso *character) {
	if (*perso).Feu {
		a := ((*perso).Tfeu) * 5
		fmt.Println("Vous brulez de -",a)
		(*perso).Vie_actuel -= a
	}
}

func FeuM(perso *monster) {
	if (*perso).Feu {
		a := ((*perso).Tfeu) * 5
		fmt.Println("Vous brulez de -",a)
		(*perso).Vie_actuel -= a
	}
}