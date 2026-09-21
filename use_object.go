package projet_red

import "fmt"

func remove(a int, perso character) {
	var new []string
	for i, val := range perso.Inventaire {
		if i == a{
			continue
		} else {
			new = append(new,val)
		}
	}
	perso.Inventaire = new
}

func Use_object(a int, perso character) {
	a--
	objet := perso.Inventaire[a]
	if objet == "potion" {
		perso.Vie_actuel += 50
		if perso.Vie_actuel > perso.Vie_max {
			perso.Vie_actuel = perso.Vie_max
		}
		fmt.Print("Vous avez maintenant ",perso.Vie_actuel)
		fmt.Println(" PV")
	}
	remove(a, perso)
}
