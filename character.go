package projet_red

import "fmt"

type character struct {
	Nom            string
	Vie_actuel     int
	Vie_max        int
	Essence_actuel int
	Essence_max    int
	Res            bool
	Level          float64
	Piece          int

	Inventaire     []string
	Action         [](*attack)
	Inventaire_max int
	NbAchatSacoche int

	Combat bool

	Equipement [3](*equip)
	
	Tpois int
	Poison bool
	BuffA float64
	TbuffA int
	Buff float64
	Tbuff int
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
	case 1:
		perso.Vie_actuel = 100
		perso.Vie_max = 100
		Essence_actuel = 50
		Essence_max = 50
		tab := []string{"Coup d'épée","Cri de guerre","Bloquer"}
		for _,val := range tab {
			c := InitAttack(val)
			perso.Action = append(perso.Action , &c)
		}
	
	case 2 :
		perso.Vie_actuel = 80
		perso.Vie_max = 80
		Essence_actuel = 65
		Essence_max = 65
		tab := []string{"Flèche de fer","Flèche de poison","Dodge"}
		for _,val := range tab {
			c := InitAttack(val)
			perso.Action = append(perso.Action , &c)
		}

	case 3 :
		perso.Vie_actuel = 60
		perso.Vie_max = 60
		Essence_actuel = 100
		Essence_max = 100
		tab := []string{ "Aiguille de mana", "Boule de feu","Bouclier magique"}
		for _,val := range tab {
			c := InitAttack(val)
			perso.Action = append(perso.Action , &c)
		}

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
	perso.Piece = 50
	perso.Res = true
	perso.Poison = false
	perso.Combat = false
	perso.Tpois = 0
	perso.NbAchatSacoche  = 0
	perso.BuffA = 1
	perso.Buff = 1
	perso.TbuffA = 0
	perso.Tbuff = 0
	return perso
}
