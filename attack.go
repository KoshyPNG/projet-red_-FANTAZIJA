package projet_red

type attack struct {
	Nom   string
	Vitesse int
	Degat bool
	Crit  bool

	Buff   bool
	Debuff bool

	ValDegat int
	Valcrit  int

	TypeBuff string
	Tbuff    int
	ValBuff  float64
}

func attackCoupMatraque() attack {
	return attack{
		Nom:      "Coup de matraque",
		Vitesse: 50    ,
		Degat:    true,
		Crit:     true,
		Buff:     false,
		Debuff:   false,
		ValDegat: 20,
		Valcrit:  10,
	}
}

func attckCriGuerre() attack {
	return attack{
		Nom:     "Cri de guerre",
		Vitesse: 70,
		Degat:   false,
		Buff:    true,
		Debuff:  false,
		TypeBuff: "attack",
		ValBuff: 1.3,
		Tbuff:   2,
	}
}

func attackBloquer() attack {
	return attack{
		Nom:      "Bloquer",
		Vitesse: 100,
		Degat:    false,
		Buff:     true,
		Debuff:   false,
		TypeBuff: "vie_actuel",
		ValBuff:  0.3,
	}
}

func attackFlecheFer() attack {
	return attack{
		Nom:      "Glock 26",
		Vitesse: 80,
		Degat:    true,
		Crit:     true,
		Buff:     false,
		Debuff:   false,
		ValDegat: 15,
		Valcrit:  33,
	}
}

func attackFlechePoison() attack {
	return attack{
		Nom:      "Flèchette de poison",
		Vitesse: 50,
		Degat:    true,
		Crit:     true,
		Debuff:   true,
		Buff:     false,
		ValDegat: 10,
		TypeBuff: "poison",
		Tbuff:    3,
		Valcrit:  33,
	}
}

func attackDodge() attack {
	return attack{
		Nom:      "Dodge",
		Buff:     true,
		Vitesse: 100,
		TypeBuff: "esquive",
		Tbuff:    1,
		ValBuff:  0.0 ,
	}
}

func attackLanceFlamme() attack {
	return attack{
		Nom:      "Lance flamme",
		Vitesse: 60,
		Degat:    true,
		Crit:     false,
		ValDegat: 25,
	}
}

func attackRechargement() attack {
	return attack{
		Nom:      "Rechargement",
		Vitesse: 30,
		Degat:    true,
		Crit:     false,
		ValDegat: 35,
	}
}

func attackCoupCrosse() attack {
	return attack{
		Nom:      "Coup de crosse",
		Vitesse: 80,
		Buff:     true,
		TypeBuff: "vie_max",
		Tbuff:    2,
		ValBuff:  0.3,
	}
}

func attackCoupGriffe() attack {
	return attack{
		Nom:      "Coup de griffe",
		Vitesse: 50,
		Degat:    true,
		ValDegat: 5,
		Valcrit:  10,
	}
}

func attackMorsure() attack {
	return attack{
		Nom:      "Morsure",
		Degat:    true,
		Vitesse: 35,
		ValDegat: 20,
		Valcrit:  10,
	}
}

func attackCharge() attack {
	return attack{
		Nom:      "Charge",
		Degat:    true,
		Vitesse: 90,
		ValDegat: 10,
		Valcrit:  10,
	}
}

func attackLancerRedBull() attack {
	return attack{
		Nom:      "Lancer de RedBull",
		Degat:    true,
		Vitesse: 80,
		ValDegat: 30,
		Valcrit:  10,
	}
}

func attackMorsureLoup() attack {
	return attack{
		Nom:      "Morsure de loup",
		Degat:    true,
		Vitesse: 101,
		ValDegat: 50,
		Valcrit:  10,
	}
}

func attackTraqueEmpoisonne() attack {
	return attack{
		Nom:      "Traque empoisonné",
		Degat:    true,
		Debuff:   true,
		Vitesse: 70,
		ValDegat: 30,
		TypeBuff: "poison",
		Tbuff:    3,
		Valcrit:  10,
	}
}

func attackR() attack {
	var a attack
	a.Nom = "RIEN"
	a.Vitesse = 101
	a.ValDegat = 0

	return a 
}

func InitAttack(nom string) attack {
	switch nom {
	case "Coup de matraque":
		return attackCoupMatraque()
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
	case "Lance flamme":
		return attackLanceFlamme()
	case "Boule de feu":
		return attackRechargement()
	case "Bouclier magique":
		return attackCoupCrosse()
	case "Coup de griffe":
		return attackCoupGriffe()
	case "Morsure":
		return attackMorsure()
	case "Charge":
		return attackCharge()
	case "Lancer de RedBull":
		return attackLancerRedBull()
	case "Morsure de loup":
		return attackMorsureLoup()
	case "Traque empoisonné":
		return attackTraqueEmpoisonne()
	}
	return attackR()
}
