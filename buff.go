package projet_red

import (
	"math/rand"
)

func UseBuff(perso *character, useP *attack) {
	if (*useP).TypeBuff == "attack" {
		if (*perso).BuffA == 1.0 {
			(*perso).BuffA = (*useP).ValBuff
			(*perso).TbuffA = (*useP).Tbuff
		} else {
			(*perso).BuffA += (*useP).ValBuff - 1
			(*perso).TbuffA += 1
		}
	} else if (*useP).TypeBuff == "vie_actuel" {
		if (*perso).Buff == 1.0 {
			(*perso).Buff = (*useP).ValBuff
			(*perso).Tbuff = (*useP).Tbuff
		} else {
			(*perso).Buff += (*useP).ValBuff - 1
			(*perso).Tbuff += 1
		}
	} else if (*useP).TypeBuff == "esquive" {
		roll := rand.Intn(100) + 1
		if roll <= 66 {
			(*perso).Buff = (*useP).ValBuff
			(*perso).Tbuff = (*useP).Tbuff
		} else {
			(*perso).Buff = 1.0
			(*perso).Tbuff = 0
		}
	}
}

func UB(monstre *monster, useM *attack) {
	if (*useM).TypeBuff == "attack" {
		if (*monstre).BuffA == 1.0 {
			(*monstre).BuffA = (*useM).ValBuff
			(*monstre).TbuffA = (*useM).Tbuff
		} else {
			(*monstre).BuffA += (*useM).ValBuff - 1
			(*monstre).TbuffA += 1
		}
	} else if (*useM).TypeBuff == "vie_actuel" {
		if (*monstre).Buff == 1.0 {
			(*monstre).Buff = (*useM).ValBuff
			(*monstre).Tbuff = (*useM).Tbuff
		} else {
			(*monstre).Buff += (*useM).ValBuff - 1
			(*monstre).Tbuff += 1
		}
	} else if (*useM).TypeBuff == "esquive" {
		roll := rand.Intn(100) + 1
		if roll <= 66 {
			(*monstre).Buff = (*useM).ValBuff
			(*monstre).Tbuff = (*useM).Tbuff
		} else {
			(*monstre).Buff = 1.0
			(*monstre).Tbuff = 0
		}
	}
}

func TurnBuff(perso *character , a float64) {
	
	if (*perso).BuffA == a && (*perso).TbuffA > 0 {
		(*perso).TbuffA -= 1
		if (*perso).TbuffA == 0 {
			(*perso).BuffA = 1.0
		}
	} else if (*perso).Buff == a && (*perso).Tbuff > 0 {
		(*perso).Tbuff -= 1
		if (*perso).Tbuff == 0 {
			(*perso).Buff = 1.0
		}
	}
}

func TB(monstre *monster, a float64) {
	
	if (*monstre).BuffA == a && (*monstre).TbuffA > 0 {
		(*monstre).TbuffA -= 1
		if (*monstre).TbuffA == 0 {
			(*monstre).BuffA = 1.0
		}
	} else if (*monstre).Buff == a {
		(*monstre).Tbuff -= 1
		if (*monstre).Tbuff == 0 {
			(*monstre).Buff = 1.0
		}
	}
}