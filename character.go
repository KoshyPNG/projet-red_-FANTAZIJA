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
	Poison bool
	Combat bool
	Tpois int
	Equipement [3](*equip)
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

func creation(perso *character) {
	var i int
	fmt.Println("===========================")
	fmt.Println("Quellle est votre classe ?")
	fmt.Println("1 : CHEVALIER ;  2 : ARCHER ; 3 : MAGICIEN(pas fini)")
	fmt.Scan(&i)
	switch i {
	case 1 :
		perso.Vie_actuel = 100
		perso.Vie_max = 100
		perso.Action = append(perso.Action , "Coup d'épée")
		perso.Action = append(perso.Action , "Cri de guerre")
		perso.Action = append(perso.Action , "Bloquer")

	case 2 :
		perso.Vie_actuel = 80
		perso.Vie_max = 80
		perso.Action = append(perso.Action , "Flèche de fer")
		perso.Action = append(perso.Action , "Flèche de poison")
		perso.Action = append(perso.Action , "Dodge")

	case 3 :
		perso.Vie_actuel = 60
		perso.Vie_max = 60
		perso.Action = append(perso.Action , "Aiguille de mana")
		perso.Action = append(perso.Action , "Boule de feu")
		perso.Action = append(perso.Action , "Bouclier magique")

	}
}

func InitCharacter() character {
	var perso character
	var i string
	p := &perso
	fmt.Println("Quellle est votre nom?")
	fmt.Scan(&i)
	perso.Nom = corrected(i)
	creation(p)
	perso.Level = 1.0
	for i := 0; i < 3; i++ {
		perso.Inventaire = append(perso.Inventaire, "Potion de soin")
	}
	tab := []string{"Rien","Rien","Rien"}
	var a equip
	for i, val := range tab {
		a = InitEquip(val)
		perso.Equipement[i] = &a

	}
	perso.Inventaire_max = 10
	perso.Piece = 20
	perso.Res = true
	perso.Poison = false
	perso.Combat = false
	perso.Tpois = 0
	return perso
}
