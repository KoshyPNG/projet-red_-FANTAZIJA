package projet_red

type Monster struct {
	Nom        string
	Vie_actuel int
	Vie_max    int
	Attaque    [](*attack)
	Loot       []string
}

func monsterZombie() Monster {
	monster := Monster{
		Nom:        "Zombie",
		Vie_actuel: 40,
		Vie_max:    40,
		Loot:       []string{"Pièce d'or", "Herbe", "Champignon"},
	}
	for _, nom := range []string{"Coup de griffe", "Morsure"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
	return monster
}

func monsterClaqueur() Monster {
	monster := Monster{
		Nom:        "Claqueur",
		Vie_actuel: 65,
		Vie_max:    65,
		Loot:       []string{"Pièce d'or", "Tissus", "Corde"},
	}
	for _, nom := range []string{"Charge", "Morsure"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
	return monster
}

func monsterMaxime() Monster {
	monster := Monster{
		Nom:        "Solar",
		Vie_actuel: 200,
		Vie_max:    200,
		Loot:       []string{"PP7 silencieux"},
	}
	for _, nom := range []string{"Lancer de RedBull", "Morsure de loup", "Traque empoisonné"} {
		attack := InitAttack(nom)
		monster.Attaque = append(monster.Attaque, &attack)
	}
	return monster
}

func InitMonster(nom string) Monster {
	switch nom {
	case "Zombie":
		return monsterZombie()
	case "Claqueur":
		return monsterClaqueur()
	case "Solar":
		return monsterMaxime()
	}
	return Monster{}
}
