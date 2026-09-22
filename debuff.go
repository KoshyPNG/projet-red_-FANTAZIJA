package projet_red

func UseDebuff(perso *character,monstre *monster, useP *attack) {
	if (*useP).TypeBuff == "poison" {
		GP(monstre)
	}
}

func UD(perso *character,monstre *monster, useM *attack) {
	if (*useM).TypeBuff == "poison" {
		GetPoison(perso)
	}
}

