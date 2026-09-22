package projet_red

type attack struct {
	Nom string
	Degat bool
	Crit bool

	Buff bool
	Debuff bool

	ValDegat int
	Valcrit int
	
	TypeBuff string
	Tbuff int
	ValBuff float64
}

func attackCoupEpee() attack {
	return attack{
		Nom:      "Coup d'épée",
		Degat:    true,
		Crit: true,
		Buff:     false,
		Debuff:   false,
		ValDegat: 20,
		Valcrit: 10,
	}
}

func attckCriGuerre() attack {
	return attack{
		Nom:      "Cri de guerre",
		Degat:    false,
		Buff:     true,
		Debuff:   false,
		ValBuff:  1.3,
		Tbuff:  2,
	}
}

func attackBloquer() attack {
	return attack{
		Nom:      "Bloquer",
		Degat:    false,
		Buff:     true,
		Debuff:   false,
		TypeBuff: "vie_actuel",
		ValBuff:  0.3,
	}
}

func attackFlecheFer() attack {
	return attack{
		Nom:      "Flèche de fer",
		Degat:    true,
		Crit: true,
		Buff: false,
		Debuff:   false,
		ValDegat: 15,
		Valcrit: 33,
	}
}

func attackFlechePoison() attack {
	return attack{
		Nom:      "Flèche de poison",
		Degat:    true,
		Crit: true,
		Debuff:   true,
		Buff: false,
		ValDegat: 10,
		TypeBuff: "poison",
		Tbuff:    3,
		Valcrit: 33,
	}
}

func attackDodge() attack {
	return attack{
		Nom:      "Dodge",
		Buff:     true,
		TypeBuff: "esquive",
		Tbuff:    1,
		ValBuff:  0.66,
	}
}

func attackAiguilleMana() attack {
	return attack{
		Nom:      "Aiguille de mana",
		Degat:    true,
		Crit: false,
		ValDegat: 25,
	}
}

func attackBouleFeu() attack {
	return attack{
		Nom:      "Boule de feu",
		Degat:    true,
		Crit: false,
		ValDegat: 35,
	}
}

func attackBouclierMagique() attack {
	return attack{
		Nom:      "Bouclier magique",
		Buff:     true,
		TypeBuff: "vie_max",
		Tbuff:    2,
		ValBuff:  0.3,
	}
}

func InitAttack(nom string) attack {
	switch nom {
	case "Coup d'épée":
		return attackCoupEpee()
	case "Cri de guerre":
		return attckCriGuerre()
	case "Bloquer":
		return attackBloquer()
	case "Flèche de fer":
		return attackFlecheFer()
	case "Flèche de poison":
		return attackFlechePoison()
	case "Dodge":
		return attackDodge()
	case "Aiguille de mana":
		return attackAiguilleMana()
	case "Boule de feu":
		return attackBouleFeu()
	case "Bouclier magique":
		return attackBouclierMagique()
	}
	return attack{}
}
