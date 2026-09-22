package projet_red

import "fmt"

type character struct {
	Nom            string
	Vie_actuel     int
	Vie_max        int
	Level          float64
	Piece          int
	Inventaire     []string
	Action         []string
	Res            bool
	Inventaire_max int
	poison         bool
	combat         bool
	tpois          int
}

func corrected(i string) string {
	var new string
	for a, val := range i {
		if a == 0 {
			if val >= 'a' && val <= 'z' {
				new = new + string(val-32)
			} else {
				new = new + string(val)
			}
		} else if val >= 'A' && val <= 'Z' {
			new = new + string(val+32)
		} else {
			new = new + string(val)
		}
	}
	return new
}

func InitCharacter() character {
	var perso character
	var i string
	fmt.Println("Quellle est votre nom?")
	fmt.Scan(&i)
	perso.Nom = corrected(i)
	perso.Vie_actuel = 100
	perso.Vie_max = 100
	perso.Level = 1.0
	for i := 0; i < 3; i++ {
		perso.Inventaire = append(perso.Inventaire, "Potion de soin")
	}
	perso.Inventaire_max = 10
	perso.Piece = 20
	perso.Action = append(perso.Action, "Coup d'épée")
	perso.Action = append(perso.Action, "Cri de guerre")
	perso.Action = append(perso.Action, "Bloquer")
	perso.Res = true
	perso.poison = false
	perso.combat = false
	perso.tpois = 0
	return perso
}
