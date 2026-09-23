package projet_red

import (
	"math/rand"
)

func UseBuff(perso *character, useP *attack) {
	switch (*useP).TypeBuff {
	case "attack":
		if (*perso).BuffA == (*perso).BuffAmin {
			(*perso).BuffA = (*useP).ValBuff
			(*perso).TbuffA = (*useP).Tbuff
		} else {
			(*perso).BuffA += (*useP).ValBuff - 1
			(*perso).TbuffA += 1
		}
	case "vie_actuel":
		if (*perso).Buff == (*perso).Buffmin {
			(*perso).Buff = (*useP).ValBuff
			(*perso).Tbuff = (*useP).Tbuff
		} else {
			(*perso).Buff += (*useP).ValBuff /2
			(*perso).Tbuff += 1
		}
	case "esquive":
		roll := rand.Intn(100) + 1
		if roll <= 66 {
			(*perso).Buff = (*useP).ValBuff
			(*perso).Tbuff = (*useP).Tbuff
		} else {
			(*perso).Buff = 1.0
			(*perso).Tbuff = 0
		}
	case "Vitesse":
		if (*perso).BuffV == (*perso).BuffVmin {
			(*perso).BuffV = (*useP).ValBuff
			(*perso).TbuffV = (*useP).Tbuff
		} else {
			(*perso).BuffV += (*useP).ValBuff - 1
			(*perso).TbuffV += 1
		}
	}
}

func UB(monstre *monster, useM *attack) {
	switch (*useM).TypeBuff {
        case "attack":
		if (*monstre).BuffA == 1.0 {
			(*monstre).BuffA = (*useM).ValBuff
			(*monstre).TbuffA = (*useM).Tbuff
		} else {
			(*monstre).BuffA += (*useM).ValBuff - 1
			(*monstre).TbuffA += 1
		}
	case "vie_actuel":
		if (*monstre).Buff == 1.0 {
			(*monstre).Buff = (*useM).ValBuff
			(*monstre).Tbuff = (*useM).Tbuff
		} else {
			(*monstre).Buff += (*useM).ValBuff - 1
			(*monstre).Tbuff += 1
		}
	case "esquive":
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
			(*perso).BuffA = (*perso).BuffAmin
		}
	} else if (*perso).Buff == a && (*perso).Tbuff > 0 {
		(*perso).Tbuff -= 1
		if (*perso).Tbuff == 0 {
			(*perso).Buff = (*perso).Buffmin
		}
	} else if (*perso).BuffV == a && (*perso).TbuffV > 0 {
		(*perso).TbuffV -= 1
		if (*perso).TbuffV == 0 {
			(*perso).BuffV = (*perso).BuffVmin
		}
	}
}

func TB(monstre *monster, a float64) {
	
	if (*monstre).BuffA == a && (*monstre).TbuffA > 0 {
		(*monstre).TbuffA -= 1
		if (*monstre).TbuffA == 0 {
			(*monstre).BuffA = (*monstre).BuffAmin
		}
	} else if (*monstre).Buff == a && (*monstre).Tbuff > 0 {
		(*monstre).Tbuff -= 1
		if (*monstre).Tbuff == 0 {
			(*monstre).Buff = (*monstre).Buffmin
		}
	} else if (*monstre).BuffV == a && (*monstre).TbuffV > 0 {
		(*monstre).TbuffV -= 1
		if (*monstre).TbuffV == 0 {
			(*monstre).BuffV = (*monstre).BuffVmin
		}
	}
}