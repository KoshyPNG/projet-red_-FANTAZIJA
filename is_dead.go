package projet_red

import "fmt"

func IsDead(perso *character) bool {
	val := true
	if (*perso).Vie_actuel <= 0 {
		val = false
	}
	return val
}

func Res(perso *character) {
	if (*perso).Res {
			fmt.Println("Vous avez survécu !!! Quelle miracle !!!")
			(*perso).Vie_actuel = (*perso).Vie_max/2
			(*perso).Res = false
	}
}