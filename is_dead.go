package projet_red

import "fmt"

func IsDead(perso character) bool {
	val := true
	if perso.Vie_actuel <= 0 {
		if perso.Res {
			fmt.Println("Vous avez survécu !!! Quelle miracle !!!")
			perso.Vie_actuel = 30
			perso.Res = false
		} else {
			val = false
		}
	}
	return val
}
