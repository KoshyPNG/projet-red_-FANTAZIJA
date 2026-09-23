package projet_red

import "fmt"

type character struct {
	Nom            string
	Classe         string
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

	Buffmin float64
	BuffAmin float64
	BuffVmin float64
	
	Stun bool
	Tpois int
	Poison bool
	stun bool

	BuffA float64
	TbuffA int
	Buff float64
	Tbuff int
	BuffV float64
	TbuffV int
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
	fmt.Println("1 : CRS S.T.A.R.S  " )
	fmt.Println("2 : Unité Tactique S.T.A.R.S")
	fmt.Println("3 : Unité d'éxtermination S.T.A.R.S")
	fmt.Scan(&i)
	switch i {
	case 1:
		perso.Classe = "CRS S.T.A.R.S"
		perso.Vie_actuel = 100
		perso.Vie_max = 100
		perso.Essence_actuel = 50
		perso.Essence_max = 50
		tab := []string{"Coup de matraque","Cri de guerre","Bloquer"}
		for _,val := range tab {
			c := InitAttack(val)
			perso.Action = append(perso.Action , &c)
		}
		perso.Buffmin = 1.2
		perso.BuffAmin = 1.0
		perso.BuffVmin = 1.0
	
	case 2 :
		perso.Classe = "Unité Tactique S.T.A.R.S"
		perso.Vie_actuel = 80
		perso.Vie_max = 80
		perso.Essence_actuel = 65
		perso.Essence_max = 65
		tab := []string{"Glock 26","Flèchette de poison","Dodge"}
		for _,val := range tab {
			c := InitAttack(val)
			perso.Action = append(perso.Action , &c)
		}
		perso.Buffmin = 1.0
		perso.BuffAmin = 1.0
		perso.BuffVmin = 1.1

	case 3 :
		perso.Classe = "Unité d'éxtermination S.T.A.R.S"
		perso.Vie_actuel = 60
		perso.Vie_max = 60
		perso.Essence_actuel = 100
		perso.Essence_max = 100
		tab := []string{ "Lance flamme", "Rechargement","Coup de crosse"}
		for _,val := range tab {
			c := InitAttack(val)
			perso.Action = append(perso.Action , &c)
		}
		perso.Buffmin = 1.0
		perso.BuffAmin = 1.1
		perso.BuffVmin = 1.0

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

	perso.BuffA = perso.BuffAmin
	perso.Buff = perso.Buffmin
	perso.BuffV = perso.BuffVmin

	perso.TbuffV = 0
	perso.TbuffA = 0
	perso.Tbuff = 0
	perso.Stun = false

	return perso
}
