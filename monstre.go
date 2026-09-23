package projet_red

type monster struct {
	Nom        string
	Vie_actuel int
	Vie_max    int
	turn       [](*attack)
	Attaque    [](*attack)
	Loot       []string

	Buffmin float64
	BuffAmin float64
	BuffVmin float64

    Feu bool
	Tfeu int
	Stun bool
	Tpois  int
	Poison bool

	BuffA  float64
	TbuffA int
	Buff   float64
	Tbuff  int
	BuffV  float64
	TbuffV int
}

func monsterZombie() monster {
	monster := monster{
		Nom:        "Zombie",
		Vie_actuel: 40,
		Vie_max:    40,
		Loot:       []string{"Herbe", "Champignon", "Kevlar", "Peau fermenté"},
		Tpois:      0,
		Poison:     false,
		BuffA:      1.0,
		TbuffA:     0,
		Buff:       1.0,
		Tbuff:      0,
		BuffV:       1.0,
		TbuffV:      0,
	}
	for _, nom := range []string{"Coup de griffe", "Morsure"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
	attack1 := monster.Attaque[0]
	attack2 := monster.Attaque[1]
	r := InitAttack("rien")
	monster.turn = [](*attack){&r, attack1, &r, attack2}
	monster.Buffmin = 1.0
	monster.BuffAmin = 1.0
	monster.BuffVmin = 1.0

	return monster
}

func monsterClaqueur() monster {
	monster := monster{
		Nom:        "Claqueur",
		Vie_actuel: 65,
		Vie_max:    65,
		Loot:       []string{"Tissus", "Corde", "Plaque en fer", "Caoutchouc"},
		Tpois:      0,
		Poison:     false,
		BuffA:      1.0,
		TbuffA:     0,
		Buff:       1.0,
		Tbuff:      0,
		BuffV:      1.0,
		TbuffV:     0,
	}
	for _, nom := range []string{"Charge", "Morsure"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
	attack1 := monster.Attaque[0]
	attack2 := monster.Attaque[1]
	r := InitAttack("rien")
	monster.turn = [](*attack){attack1, &r, attack2}
	monster.Buffmin = 1.0
	monster.BuffAmin = 1.0
	monster.BuffVmin = 1.0
	return monster
}

func monsterMaxime() monster {
	monster := monster{
		Nom:        "Solar",
		Vie_actuel: 200,
		Vie_max:    200,
		Loot:       []string{"PP7 silencieux", "RedBull"},
		Tpois:      0,
		Poison:     false,
		BuffA:      1.0,
		TbuffA:     0,
		Buff:       1.0,
		Tbuff:      0,
		BuffV:      1.0,
		TbuffV:     0,
	}
	for _, nom := range []string{"Lancer de RedBull", "Morsure de loup", "Traque empoisonné"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
	attack1 := monster.Attaque[0]
	attack2 := monster.Attaque[2]
	attack3 := monster.Attaque[1]
	monster.turn = [](*attack){attack1, attack3, attack2}
	monster.Buffmin = 1.2
	monster.BuffAmin = 1.2
	monster.BuffVmin = 1.2
	return monster
}

func monsterMannequin() monster {
	monster := monster{
		Nom:        "M.A.X.I.M.E (Modèle Anatomique X-pert Interractif Mesurable Ergonomique)",
		Vie_actuel: 56100,
		Vie_max:    56100,
		Loot:       []string{"Herbe", "Champignon"},
		Tpois:      0,
		Poison:     false,
		BuffA:      1.0,
		TbuffA:     0,
		Buff:       1.0,
		Tbuff:      0,
		BuffV: 1.0,
	}
	r := InitAttack("rien")
	for i := 0; i< 3;i++{
		monster.turn = append(monster.turn, &r)
	}
	return monster
}

func InitMonster(nom string) monster {
	var m monster
	switch nom {
	case "Zombie":
		m = monsterZombie()
	case "Claqueur":
		m = monsterClaqueur()
	case "Solar":
		m = monsterMaxime()
	case "M.A.X.I.M.E" :
		m = monsterMannequin()
	}
	m.Stun = false
	m.Feu = false
	m.Tfeu = 0
	return m
}
