package projet_red

import "fmt"

func remove(a int, perso *character) {
	var new []string
	for i, val := range perso.Inventaire {
		if i == a {
			continue
		} else {
			new = append(new, val)
		}
	}
	(*perso).Inventaire = new
}

func Baril(perso *character) {
	perso.Essence_actuel += 50
	if perso.Essence_actuel > perso.Essence_max {
		perso.Essence_actuel= perso.Essence_max
	}
	fmt.Print("Vous avez maintenant ", perso.Essence_actuel)
	fmt.Println(" PV")
}

func RedB(perso *character) {
	if (*perso).BuffV == (*perso).BuffVmin {
		(*perso).BuffV = 1.5
		(*perso).TbuffV = 2
	} else {
		(*perso).BuffV += 0.1
		(*perso).TbuffV += 1
	}
}

func PotH(perso *character) {
	perso.Vie_actuel += 30
	if perso.Vie_actuel > perso.Vie_max {
		perso.Vie_actuel = perso.Vie_max
	}
	fmt.Print("Vous avez maintenant ", perso.Vie_actuel)
	fmt.Println(" PV")
}

func PotGH(perso *character) {
	perso.Vie_actuel += 50
	if perso.Vie_actuel > perso.Vie_max {
		perso.Vie_actuel = perso.Vie_max
	}
	fmt.Print("Vous avez maintenant ", perso.Vie_actuel)
	fmt.Println(" PV")
}


func PotP(perso *character) {
	GetPoison(perso)
}

func PotP_C(perso *character, monstre *monster) {
	GP(monstre)
}

func Use_object(a int, perso *character) {
	a--
	objet := (*perso).Inventaire[a]
	switch objet {
	case "Potion de soin":
		PotH(perso)
		remove(a, perso)
	case "Grande Potion de soin":
		PotGH(perso)
		remove(a, perso)
	case "Potion de poison" :
		PotP(perso)
		remove(a, perso)
	case "Casquette Gucci Fraise" :
		equip := InitEquip(objet)
		Equiper(&equip, perso)
		remove(a, perso)
	case "Gilet par balles" :
		equip := InitEquip(objet)
		Equiper(&equip, perso)
		remove(a, perso)
	case "Timberland" :
		equip := InitEquip(objet)
		Equiper(&equip, perso)
		remove(a, perso)
	case "RedBull" :
		RedB(perso)
		remove(a, perso)
	case "Baril d'essence" :
		Baril(perso)
		remove(a, perso)
	default:
		fmt.Println("Vous ne pouvez pas utiliser cet objet.")
	}
}

func Use_object_C(a int, perso *character, monstre *monster) {
	a--
	objet := (*perso).Inventaire[a]
	switch objet {
	case "Potion de soin":
		PotH(perso)
		remove(a, perso)
	case "Grande Potion de soin":
		PotGH(perso)
		remove(a, perso)
	case "Potion de poison" :
		PotP_C(perso, monstre)
		remove(a, perso)
	case "RedBull" :
		RedB(perso)
		remove(a, perso)
	case "Baril d'essence" :
		Baril(perso)
		remove(a,perso)
	default:
		fmt.Println("Vous ne pouvez pas utiliser cet objet.")
	}
}