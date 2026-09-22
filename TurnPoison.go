package projet_red

import "fmt"

func GetPoison(perso *character) {
	(*perso).Poison = true
	if (*perso).Tpois != 0 {
		(*perso).Tpois++
	} else {
		(*perso).Tpois = 3
	}
}

func GP(monstre *monster) {
	(*monstre).Poison = true
	if (*monstre).Tpois != 0 {
		(*monstre).Tpois++
	} else {
		(*monstre).Tpois = 3
	}
}

func Poison(perso *character) {
	if (*perso).Poison {
			fmt.Println("Vous vous sentez malade")
			(*perso).Tpois--
			if (*perso).Tpois == 0 {
				(*perso).Poison = false
			}
			(*perso).Vie_actuel -= 10
		}
}

func Pois(monstre *monster) {
	if (*monstre).Poison {
			fmt.Println("Le monstre se sent malade")
			(*monstre).Tpois--
			if (*monstre).Tpois == 0 {
				(*monstre).Poison = false
			}
			(*monstre).Vie_actuel -= 10
		}
}
