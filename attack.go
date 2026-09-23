package projet_red

type attack struct {
	Nom   string
	Vitesse int
	Degat bool
	Crit  bool
	Essence bool

	Buff   bool
	Debuff bool

	ValDegat int
	Valcrit  int
	EssenceCost int

	TypeBuff string
	Tbuff    int
	ValBuff  float64
}

func attackCoupMatraque() attack {
	return attack{
		Nom:      "Coup de matraque",
		Vitesse: 50    ,
		Essence: false,
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
		Essence: false,
		Degat:   false,
		Buff:    true,
		Debuff:  false,
		TypeBuff: "attack",
		ValBuff: 1.3,
		Tbuff:   3,
	}
}

func attackBloquer() attack {
	return attack{
		Nom:      "Bloquer",
		Vitesse: 100,
		Essence: false,
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
		Essence: false,
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
		Essence: false,
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
		Essence: false,
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
		Essence: true,
		Debuff: true,
		EssenceCost: 30,
		TypeBuff: "feu",
		Degat:    true,
		Crit:     false,

		ValDegat: 30,
	}
}

func attackRechargement() attack {
	return attack{
		Nom:      "Rechargement",
		Vitesse: 30,
		Degat:    false,
		TypeBuff  : "essence",
		ValBuff: 45,
		Essence: true,
	}
}

func attackCoupCrosse() attack {
	return attack{
		Nom:      "Coup de crosse",
		Vitesse: 80,
		Degat:    false,
		Buff:    false,
		Debuff:     true,
		TypeBuff: "stun",
		Tbuff:    1,
		Essence: true,
	}
}

func attackCoupGriffe() attack {
	return attack{
		Nom:      "Coup de griffe",
		Vitesse: 50,
		Degat:    true,
		Buff:    false,
		Debuff:     false,
		ValDegat: 5,
		Valcrit:  10,
		Essence: true,
	}
}

func attackMorsure() attack {
	return attack{
		Nom:      "Morsure",
		Degat:    true,
		Essence:  false,
		Buff:    false,
		Debuff:     false,
		Vitesse: 35,
		ValDegat: 20,
		Valcrit:  10,
	}
}

func attackCharge() attack {
	return attack{
		Nom:      "Charge",
		Degat:    true,
		Essence:  false,
		Buff:    true,
		Debuff:     false,
		Vitesse: 90,
		ValDegat: 10,
		Valcrit:  10,
		TypeBuff: "Vitesse",
		Tbuff:    2,
		ValBuff:  1.2,
		
	}
}

func attackLancerRedBull() attack {
	return attack{
		Nom:      "Lancer de RedBull",
		Degat:    true,
		Essence:  false,
		Buff:    false,
		Debuff:     true,
		Vitesse: 80,
		ValDegat: 30,
		Valcrit:  10,
		TypeBuff: "stun",
		Tbuff:    1,
	}
}

func attackMorsureLoup() attack {
	return attack{
		Nom:      "Morsure de loup",
		Degat:    true,
		Essence:  false,
		Buff:    false,
		Debuff:     false,
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
		Essence:  false,
		Buff:    false,
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
	case "Glock 26":
		return attackFlecheFer()
	case "Flèchette de poison":
		return attackFlechePoison()
	case "Dodge":
		return attackDodge()
	case "Lance flamme":
		return attackLanceFlamme()
	case "Rechargement":
		return attackRechargement()
	case "Coup de crosse":
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
