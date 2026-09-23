package main

import (
	"fmt"

	"projet_red"
)

func main() {
	me := projet_red.InitCharacter()
	perso := &me
	fmt.Println("=====================")
	fmt.Println("Voici une expliquation des commandes :")
	fmt.Println("=====================")
	fmt.Println("écrire le numero de l'action pour la réaliser")
	fmt.Println("Quand vous avancez il se passent un event, il peut y avoir un combat, un marchand ou un camp")
	fmt.Println("Combattez pour votre vie en combat !!!!!!")
	fmt.Println("Acheté des objets au marchand")
	fmt.Println("Au camp vous pouvez Craft OU vous reposez")
	for projet_red.IsDead(perso) {
		fmt.Println("=====================")
		fmt.Println("Que vous voulez faire ?")
		fmt.Println("1 : AVANCEZ ; 2 : PERSONNAGE ; 3 : INVENTAIRE ; 4 : ENTRAINEMENT ; 5 : QUITTER")
		var a int
		fmt.Scan(&a)
		switch a {
		case 1:
			projet_red.Moov(perso)
		case 2:
			projet_red.DisplayInfo(perso)
		case 3:
			projet_red.AccessInventory(perso)
		case 4:
			perso.Combat = true
			m := InitMonster("M.A.X.I.M.E")
			Combat(perso , &m)
		case 5:
			var b string
			fmt.Println("OUI pour quittez ou NON pour annulé")
			fmt.Scan(&b)
			if b == "OUI" {
				(*perso).Vie_actuel = 0
				(*perso).Res = false
			}
		}
	}
	fmt.Println("=====================")
	fmt.Println("Votre aventure se termine")
	fmt.Println("=====================")
}
