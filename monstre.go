package projet_red

type Monster struct {
    Nom          string
    Vie_actuel   int
    Vie_max      int
    Attaque      []string  
    Degats       []int
    Loot         []string
}


Zombie := Monster{
    Nom:        "Zombie",
    Vie_actuel: 40,
    Vie_max:    40,
    Attaques:   []string{"Coup de griffe", "Morsure"},
    Degats:     []int{5, 10},
    Loot:       []string{"Pièce d'or", "Herbe" , "Champignon"}
}


Claqueur := Monster{
    Nom:        "Claqueur",
    Vie_actuel: 65,
    Vie_max:    65,
    Attaques:   []string{"Charge", "Morsure"},
    Degats:     []int{10, 20},
    Loot:       []string{"Pièce d'or", "Tissus", "Corde"}
}


Maxime := Monster{
    Nom:        "Solar",
    Vie_actuel: 200,
    Vie_max:    200,
    Attaques:   []string{"Lancer de RedBull", "Morsure de loup", "Traque empoisonné"},
    Degats:     []int{30, 50},
    Loot:       []string{"PP7 silencieux"}
}


