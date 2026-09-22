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
func Critique(at *attack) int {
	if (*at).Crit {
		roll := rand.Intn(100) + 1
		if roll <= (*at).Valcrit {
			return ((*at).ValDegat)*2
		}
	}
	return (*at).ValDegat
}

func SoloMonster(perso *character, monstre *monster, useM *attack) {
	AccessInventory(perso)
	fmt.Println("===================")
	fmt.Println( (*monstre).Nom ," utilise ",(*useM).Nom)
	(*perso).Vie_actuel -= int(float64(Critique(useM))*(*perso).Buff)
	fmt.Println("===================")
	fmt.Println((*perso).Nom ," : ",(*perso).Vie_actuel,"/",(*perso).Vie_max)
}

func Pfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	fmt.Println( (*perso).Nom ," utilise ",(*useP).Nom)
	if (*useP).Degat {
		(*monstre).Vie_actuel -= int(float64(Critique(useP))*(*perso).BuffA)
		if (*monstre).Vie_actuel != 0 {
			fmt.Println( (*monstre).Nom ," utilise ",(*useM).Nom)
			(*perso).Vie_actuel -= Critique(useM)
		}
	} else if (*useP).Buff {
		UseBuff(perso, monstre, useP)
		fmt.Println( (*monstre).Nom ," utilise ",(*useM).Nom)
		if (*useM).Nom != "RIEN" {
			(*perso).Vie_actuel -= int(float64(Critique(useM))*(*perso).Buff)
		}
	} else if (*useP).Debuff {
		UseDebuff(perso, monstre)
		if (*useM).Nom != "RIEN" {
			(*perso).Vie_actuel -= int(float64(Critique(useM))*(*perso).Buff)
		}
	}
}

func Mfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	fmt.Println( (*monstre).Nom ," utilise ",(*useM).Nom)
	(*perso).Vie_actuel -= int(float64(Critique(useM))*(*perso).Buff)
	if (*perso).Vie_actuel != 0 {
		fmt.Println( (*perso).Nom ," utilise ",(*useP).Nom)
		if (*useP).Degat {
			(*monstre).Vie_actuel -= int(float64(Critique(useP))*(*perso).BuffA)
		} else if (*useP).Buff {
			UseBuff(perso, monstre, useP)
		} else if (*useP).Debuff {
			UseDebuff(perso, monstre)
		}
	}
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
		fmt.Println((*perso).Nom ," : ",(*perso).Vie_actuel,"/",(*perso).Vie_max)
		for i,v := range (*perso).Action {
			fmt.Println(i+1," : ",v.Nom )
		}
		invet := false
		var useP *attack
		var useM *attack
		var a int
		fmt.Scan(&a)
		switch a {
		case 1 :
			useP = (*perso).Action[0]
		case 2 :
			useP = (*perso).Action[1]
		case 3 :
			useP = (*perso).Action[3]
		case 4 :
			invet = true
		}
		if tour > len((*monstre).turn) {
			if tour%len((*monstre).turn) != 0 {
				useM = (*monstre).turn[tour%len((*monstre).turn)-1]
			}
			useM = (*monstre).turn[len((*monstre).turn)-1]
		}
		if invet {
			invet = false
			SoloMonster(perso, monstre, useM)
		} else if first(useP, useM) == "1" {
			Pfirst(perso, monstre, useP, useM)
		} else {
			Mfirst(perso, monstre, useP, useM)
		}
	}
}