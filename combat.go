package projet_red

import (

	"fmt"
	"math/rand"
)

func first(a *attack, b *attack) string {
	if (*a).Vitesse > (*b).Vitesse {
		return "1"
	} else if (*a).Vitesse == (*b).Vitesse {
		roll := rand.Intn(2) + 1
		if roll == 1 {
			return "1"
		}
	}
	return "2"
}

func Combat(perso *character, monstre *monster) {
	fmt.Println("===================")
	fmt.Println("Vous tombez sur : ",(*monstre).Nom)
	fmt.Println("COMBATTEZ !!!!")
	//roll := rand.Intn(100) + 1
	tour := 1
	perso.Combat = true
	for IsDead(perso) && (*monstre).Vie_actuel != 0 {
		fmt.Println("===================")
		fmt.Println("TOUR ",tour)
		fmt.Print((*monstre).Nom,"  :  ")
		fmt.Print((*monstre).Vie_actuel,"/",(*monstre).Vie_max)
		fmt.Println("  ")
		for i,v := range (*perso).Action {
			fmt.Println(i+1," : ",v.Nom )
		}
		invet := false
		var a int
		fmt.Scan(&a)
		switch a {
		case 1 :
			useP := (*perso).Action[0]
		case 2 :
			useP := (*perso).Action[1]
		case 3 :
			useP := (*perso).Action[3]
		case 4 :
			AccessInventory(perso)
			invet = true
		}
		if tour > len((*monstre).turn) {
			if tour%len((*monstre).turn) != 0 {
				useM := (*monstre).turn[tour%len((*monstre).turn)-1]
			}
			useM := (*monstre).turn[len((*monstre).turn)-1]
		}
		if invet {
			invet = false
			SoloMonster(perso, monstre)
		} else if first(useP, useM) == "1" {
			Pfirst(perso, monstre, useP, useM)
		} else {
			Mfirst(perso, monstre, useP, useM)
		}
	}
}