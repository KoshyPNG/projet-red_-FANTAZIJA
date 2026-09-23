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
	
	if (*monstre).Vie_actuel <= 0 {
		TMonster(perso, monstre,useM)
	}
}

func Mfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	TMonster(perso, monstre, useM)
	
	if (*perso).Vie_actuel <= 0 {
		TPersonnage(perso, monstre, useP)
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
		fmt.Println("4 : inventaire ")

		invet := false
		var useP *attack
		var useM *attack
		var a int
		fmt.Scan(&a)

		if a >0 && 4>= a {
			if a < 4 {
				useP = (*perso).Action[a-1]
			} else {
				invet = true
			}
		}

		if tour > len((*monstre).turn) {
			if tour%len((*monstre).turn) != 0 {
				useM = (*monstre).turn[ tour%len((*monstre).turn) - 1 ]
			}
			useM = (*monstre).turn[len((*monstre).turn)%tour -1]
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
		if invet && !(*perso).Stun {
			invet = false
			SoloMonster(perso, monstre, useM)
		} else if first(int((*perso).BuffV * float64(useP.Vitesse)), int((*monstre).BuffV * float64(useM.Vitesse))) == "1" {
			Pfirst(perso, monstre, useP, useM)
		} else {
			Mfirst(perso, monstre, useP, useM)
		}
		tour += 1
	}
	(*perso).Stun = false
	if IsDead(perso) {
		fmt.Println("===============")
		fmt.Println((*monstre).Nom, " est vaincu !!!!!")
		perso.Piece += 10
		fmt.Println("Vous avez gagner 10 pièces")
		roll := rand.Intn(len((*monstre).Loot)) 
		(*perso).Inventaire = append((*perso).Inventaire, (*monstre).Loot[roll])
		fmt.Println("Vous avez récuperé", (*monstre).Loot[roll] )
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
		fmt.Println( (*perso).Nom ," infligez -",(a))
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
}

func TMonster(perso *character, monstre *monster, useM *attack) {
	var a int
	fmt.Println( (*monstre).Nom ," utilise ",(*useM).Nom)

	if (*useM).Degat {
		a = int( (float64(Critique(useM)) * (*monstre).BuffA ) * (*perso).Buff )
		(*perso).Vie_actuel -= a
		fmt.Println( (*perso).Nom ," infligez -",(a))
		
	}

	if (*useM).Buff {
		UB(monstre, useM)
	}

	if (*useM).Debuff {
		UD(perso, monstre, useM)
	}

	TB(monstre, (*monstre).BuffA)
	TurnBuff(perso, (*perso).Buff)
}