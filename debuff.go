package projet_red

import (
	"math/rand"
)

func UseDebuff(perso *character,monstre *monster, useP *attack) {
	switch (*useP).TypeBuff {
	case "poison":
		GP(monstre)
	case "stun":
		(*monstre).Stun = true
	case "feu" :
		if (*monstre).Feu {
			(*monstre).Tfeu += 1
		} else {
            (*monstre).Feu = true
		    (*monstre).Tfeu = 2
		}
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
	case "feu" :
		if (*perso).Feu {
			(*perso).Tfeu += 1
		} else {
            (*perso).Feu = true
		    (*perso).Tfeu = 2
		}
	}
}

