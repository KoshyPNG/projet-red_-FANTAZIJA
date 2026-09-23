package projet_red

type equip struct {
	Nom         string
	description string
	buff 	 string
	val         float64
	pos         int
}

func EquipR() equip {
	var rien equip
	rien.Nom = "Rien"
	rien.description = "Ne fait rien"
	rien.val = 0
	rien.buff = "aucun"
	return rien
}

func EquipChapeau() equip {
	var e equip
	e.Nom = "Casque Renforcer"
	e.description = "Vous vous sentez plus protégé"
	e.pos = 0
	e.val = 10
	e.buff = "Vie_max"
	return e
}

func EquipTunique() equip {
	var e equip
	e.Nom = "Gilet par balles"
	e.description = "Vous vous sentez plus protégé"
	e.pos = 1
	e.val = 30
	e.buff = "Vie_max"
	return e
}

func EquipBotte() equip {
	var e equip
	e.Nom = "Bottes Renforcé"
	e.description = "Vous vous sentez plus protégé"
	e.pos = 2
	e.val = 10
	e.buff = "Vie_max"
	return e
}

func Equiper(a *equip, perso *character) {
	len := (*a).pos
	if (*a).buff != "aucun" {
		switch (*a).buff {
		case "Vie_max":
			DesEquiper(perso.Equipement[len], perso)
			perso.Equipement[len] = a
			(*perso).Vie_max += int(perso.Equipement[len].val)
			(*perso).Vie_actuel += int(perso.Equipement[len].val)
		case "Essence_max":
			DesEquiper(perso.Equipement[len], perso)
			perso.Equipement[len] = a
			(*perso).Essence_max += int(perso.Equipement[len].val)
			(*perso).Essence_actuel += int(perso.Equipement[len].val)
		case "Buff":
			DesEquiper(perso.Equipement[len], perso)
			perso.Equipement[len] = a
			(*perso).Buffmin += float64(perso.Equipement[len].val)
			(*perso).Buff = (*perso).Buffmin
		case "BuffA":
			DesEquiper(perso.Equipement[len], perso)
			perso.Equipement[len] = a
			(*perso).BuffAmin += float64(perso.Equipement[len].val)
			(*perso).BuffA = (*perso).BuffAmin
		case "BuffV":
			DesEquiper(perso.Equipement[len], perso)
			perso.Equipement[len] = a
			(*perso).BuffVmin += float64(perso.Equipement[len].val)
			(*perso).BuffV = (*perso).BuffVmin
		}
	}
}

func DesEquiper(a *equip, perso *character) {
	if (*a).buff != "aucun" {
		switch (*a).buff {
		case "Vie_max":
			(*perso).Vie_max -= int(a.val)
			if (*perso).Vie_max < (*perso).Vie_actuel {
				(*perso).Vie_actuel = (*perso).Vie_max
			}
		case "Essence_max":
			(*perso).Essence_max -= int(a.val)
			if (*perso).Essence_max < (*perso).Essence_actuel {
				(*perso).Essence_actuel = (*perso).Essence_max
			}
		case "Buff":
			(*perso).Buffmin -= float64(a.val)
		case "BuffA":
			(*perso).BuffAmin -= float64(a.val)
		case "BuffV":
			(*perso).BuffVmin -= float64(a.val)
		}
	}
}

func InitEquip(nom string) equip {
	switch nom {
	case "Casque Renforcer":
		return EquipChapeau()
	case "Gilet par balles":
		return EquipTunique()
	case "Bottes Renforcé":
		return EquipBotte()
	}
	return EquipR()
}