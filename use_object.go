package projet_red

import "fmt"

func remove(a int, perso *character) {
	var new []string
	for i, val := range perso.Inventaire {
		if i == a{
			continue
		} else {
			new = append(new,val)
		}
	}
	(*perso).Inventaire = new
}

func PotH(perso *character) {
	perso.Vie_actuel += 30
	if perso.Vie_actuel > perso.Vie_max {
		perso.Vie_actuel = perso.Vie_max
	}
	fmt.Print("Vous avez maintenant ",perso.Vie_actuel)
	fmt.Println(" PV")
}

func PotGH(perso *character) {
	perso.Vie_actuel += 50
	if perso.Vie_actuel > perso.Vie_max {
		perso.Vie_actuel = perso.Vie_max
	}
	fmt.Print("Vous avez maintenant ",perso.Vie_actuel)
	fmt.Println(" PV")
}

func PotP(perso *character) {
	(*perso).poison = true
	if (*perso).tpois != 0 {
		(*perso).tpois++
	} else {
		(*perso).tpois = 3
	}
}

func Use_object(a int, perso *character) {
	a--
	objet := (*perso).Inventaire[a]
	switch objet {
	case "Potion de soin" :
		PotH(perso)

	case "Grande Potion de soin" :
		PotGH(perso)
	}
	remove(a, perso)
}
