package projet_red

import (

	"fmt"
	"math/rand"
)

func first(a int, b int) string {
	if a > b {
		return "1"
	} else if a == b {
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
	AccessInventoryC(perso, monstre)
	fmt.Println("===================")
	fmt.Println((*perso).Nom ," : ",(*perso).Vie_actuel,"/",(*perso).Vie_max)
	
	TMonster(perso, monstre, useM)
}
	

func Pfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	TPersonnage(perso, monstre, useP)
	
	if (*monstre).Vie_actuel != 0 {
		TMonster(perso, monstre,useM)
	}
}

func Mfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	TMonster(perso, monstre, useM)
	
	if (*perso).Vie_actuel != 0 {
		TPersonnage(perso, monstre, useM)
	}
}

func Combat(perso *character, monstre *monster) {
	fmt.Println("===================")
	fmt.Println("Vous tombez sur : ",(*monstre).Nom)
	fmt.Println("COMBATTEZ !!!!")
	//roll := rand.Intn(100) + 1
	tour := 1
	rien := InitAttack("Rien")
	for IsDead(perso) && (*monstre).Vie_actuel >= 0 {
		fmt.Println("===================")

		fmt.Println("TOUR ",tour)
		Poison(perso)
		Pois(monstre)
		Feu(perso)
		FeuM(monstre)

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

		if a >0 && 6>= a {
			if a != 6 {
				useP = (*perso).Action[a-1]
			} else {
				invet = true
			}
		}

		if tour > len((*monstre).turn) {
			if tour%len((*monstre).turn) != 0 {
				useM = (*monstre).turn[ tour%len((*monstre).turn) - 1 ]
			}
			useM = (*monstre).turn[len((*monstre).turn)-1]
		} else {
			useM = (*monstre).turn[tour-1]
		}

		if (*perso).Stun {
			fmt.Println("Vous êtes étourdi et ne pouvez pas agir")
			useP = &rien
			(*perso).Stun = false
		}
		if (*monstre).Stun {
			fmt.Println("Le monstre est étourdi et ne peut pas agir")
			useM = &rien
			(*monstre).Stun = false
		}
        v1 := int((*perso).BuffV * float64(useP.Vitesse))
		v2 := int((*monstre).BuffV * float64(useM.Vitesse))
		if invet {
			invet = false
			SoloMonster(perso, monstre, useM)
		} else if first(v1, v2) == "1" {
			Pfirst(perso, monstre, useP, useM)
		} else {
			Mfirst(perso, monstre, useP, useM)
		}
	}
	if IsDead(perso) {
		fmt.Println((*monstre).Nom, " est vaincu !!!!!")
		perso.Piece += 10
		fmt.Println("Vous avez gagner 10 pièces")
		roll := rand.Intn(len((*monstre).Loot)) 
		(*perso).Inventaire = append((*perso).Inventaire, (*monstre).Loot[roll])
		fmt.Println("Vous avez récuperé" (*monstre).Loot[roll] )
	} else {
		fmt.Println("Vous avez failli à votre mission")
	}
}

func TPersonnage(perso *character, monstre *monster, useP *attack, ) {
	var a int
	fmt.Println( (*perso).Nom ," utilise ",(*useP).Nom)

	if (*useP).Degat {
		a = int( (float64(Critique(useP)) * (*perso).BuffA ) * (*monstre).Buff )
		(*monstre).Vie_actuel -= a
		fmt.Println( (*perso).Nom ," inflige -",(a))
	}

	if (*useP).Buff {
		UseBuff(perso, useP)
	}

	if (*useP).Debuff {
		UseDebuff(perso, monstre, useP)
	}
	if (*useP).Essence {
		(*perso).Essence_actuel -= (*useP).EssenceCost
	}

	TB(monstre, (*monstre).Buff)
	TurnBuff(perso, (*perso).BuffA)
	fmt.Println((*perso).Nom ," : ",(*perso).Vie_actuel,"/",(*perso).Vie_max)
}

func TMonster(perso *character, monstre *monster, useM *attack) {
	var a int
	fmt.Println( (*monstre).Nom ," utilise ",(*useM).Nom)

	if (*useM).Degat {
		a = int( (float64(Critique(useM)) * (*monstre).BuffA ) * (*perso).Buff )
		(*monstre).Vie_actuel -= a
		fmt.Println( (*monstre).Nom ," inflige -",(a))
		
	}

	if (*useM).Buff {
		UseBuff(perso, useM)
	}

	if (*useM).Debuff {
		UseDebuff(perso, monstre, useM)
	}

	TB(monstre, (*monstre).BuffA)
	TurnBuff(perso, (*perso).Buff)
	fmt.Println((*monstre).Nom ," : ",(*monstre).Vie_actuel,"/",(*monstre).Vie_max)
}