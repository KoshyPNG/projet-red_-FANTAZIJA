package projet_red

import "fmt"

type character struct{
	nom string
	vie_actuel int
	vie_max int
	xp float64
	piece int
	inventaire []string
	action []string
}

func InitCharacter() character {
	var perso character
	var i string
	fmt.Println("Quellle est votre nom?")
	fmt.Scan(&i)
	perso.nom = i
	perso.vie_actuel = 100
	perso.vie_max = 100
	perso.xp = 1.0
	for i := 0; i <3; i++{
		perso.inventaire = append(perso.inventaire , "potion")
	}
	perso.piece = 20
	perso.action = append(perso.action , "Coup d'épée")
	perso.action = append(perso.action , "Cri de guerre")
	perso.action = append(perso.action , "Bloquer")
	return perso
}
