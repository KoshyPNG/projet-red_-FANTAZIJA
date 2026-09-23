package projet_red

type equip struct {
	Nom         string
	description string
	val         int
	pos         int
}

func EquipR() equip {
	var rien equip
	rien.Nom = "Rien"
	rien.description = "Ne fait rien"
	rien.val = 0
	return rien
}

func EquipChapeau() equip {
	var e equip
	e.Nom = "Casquette Gucci Fraise"
	e.description = "Vous vous sentez plus protégé"
	e.pos = 0
	e.val = 10
	return e
}

func EquipTunique() equip {
	var e equip
	e.Nom = "Gilet par balle"
	e.description = "Vous vous sentez plus protégé"
	e.pos = 1
	e.val = 30
	return e
}

func EquipBotte() equip {
	var e equip
	e.Nom = "Timberland"
	e.description = "Vous vous sentez plus protégé"
	e.pos = 2
	e.val = 10
	return e
}

func Equiper(a *equip, perso *character) {
	len := (*a).pos
	(*perso).Vie_max -= perso.Equipement[len].val
	if (*perso).Vie_max < (*perso).Vie_actuel {
		(*perso).Vie_actuel = (*perso).Vie_max
	}
	perso.Equipement[len] = a
	(*perso).Vie_max += perso.Equipement[len].val
	(*perso).Vie_actuel += perso.Equipement[len].val
}

func InitEquip(nom string) equip {
	switch nom {
	case "Chapeau de l'aventurier":
		return EquipChapeau()
	case "Tunique de l'aventurie":
		return EquipTunique()
	case "Bottes de l'aventurier":
		return EquipBotte()
	}
	return EquipR()
}