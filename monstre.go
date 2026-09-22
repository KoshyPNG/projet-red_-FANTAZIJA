package projet_red

type monster struct {
	Nom        string
	Vie_actuel int
	Vie_max    int
    turn       [](*attack)
	Attaque    [](*attack)
	Loot       []string
}

func monsterZombie() monster {
	monster := monster{
		Nom:        "un Zombie",
		Vie_actuel: 40,
		Vie_max:    40,
		Loot:       []string{"Pièce d'or", "Herbe", "Champignon"},
	}
	for _, nom := range []string{"Coup de griffe", "Morsure"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
    attack1 := monster.Attaque[0]
    attack2 := monster.Attaque[1]
    r := InitAttack("rien")
    monster.turn = [](*attack){&r,attack1,&r,attack2}
	return monster
}

func monsterClaqueur() monster {
	monster := monster{
		Nom:        "un Claqueur",
		Vie_actuel: 65,
		Vie_max:    65,
		Loot:       []string{"Pièce d'or", "Tissus", "Corde"},
	}
	for _, nom := range []string{"Charge", "Morsure"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
    attack1 := monster.Attaque[0]
    attack2 := monster.Attaque[1]
    r := InitAttack("rien")
    monster.turn = [](*attack){attack1,&r,attack2}
	return monster
}

func monsterMaxime() monster {
	monster := monster{
		Nom:        "Solar",
		Vie_actuel: 200,
		Vie_max:    200,
		Loot:       []string{"PP7 silencieux", "RedBull"},
	}
	for _, nom := range []string{"Lancer de RedBull", "Morsure de loup", "Traque empoisonné"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
    attack1 := monster.Attaque[0]
    attack2 := monster.Attaque[2]
    attack3 := monster.Attaque[1]
    monster.turn = [](*attack){attack1,attack3,attack2}
	return monster
}

func InitMonster(nom string) monster {
	switch nom {
	case "Zombie":
		return monsterZombie()
	case "Claqueur":
		return monsterClaqueur()
	case "Solar":
		return monsterMaxime()
	}
	return monster{}
}
