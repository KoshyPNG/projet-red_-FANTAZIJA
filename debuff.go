package projet_red

import "math/rand"

func UseDebuff(perso *character,monstre *monster, useP *attack) {
	switch (*useP).TypeBuff {
	case "poison":
		GP(monstre)
	case "stun":
		(*monstre).Stun = true
	}
}

func UD(perso *character,monstre *monster, useM *attack) {
	switch (*useM).TypeBuff {
	case "poison":
		GetPoison(perso)
	case "stun":
		roll := rand.Intn(100) + 1
		if roll <= 66 {
			(*perso).Stun = true
		}
	}
}

