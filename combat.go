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
			return ((*at).ValDegat) * 2
		}
	}
	return (*at).ValDegat
}

func SoloMonster(perso *character, monstre *monster, useM *attack) {
	AccessInventoryC(perso, monstre)
	fmt.Println("===================")
	fmt.Println((*perso).Nom, " : ", (*perso).Vie_actuel, "/", (*perso).Vie_max)

	TMonster(perso, monstre, useM)
}

func Pfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	TPersonnage(perso, monstre, useP)

	if (*monstre).Vie_actuel > 0 {
		TMonster(perso, monstre, useM)
	}
}

func Mfirst(perso *character, monstre *monster, useP *attack, useM *attack) {
	fmt.Println("===================")
	TMonster(perso, monstre, useM)

	if (*perso).Vie_actuel > 0 {
		TPersonnage(perso, monstre, useP)
	}
}

func Combat(perso *character, monstre *monster) {
	fmt.Println("===================")
	fmt.Println("Vous tombez sur : ", (*monstre).Nom)
	fmt.Println("COMBATTEZ !!!!")
	//roll := rand.Intn(100) + 1
	tour := 1
	rien := InitAttack("Rien")
	Maxime := false
	for IsDead(perso) && (*monstre).Vie_actuel > 0 {
		fmt.Println("===================")

		fmt.Println("TOUR ", tour)

		fmt.Print((*monstre).Nom, "  :  ")
		fmt.Print((*monstre).Vie_actuel, "/", (*monstre).Vie_max)
		fmt.Println("  ")
		fmt.Println((*perso).Nom, " : ", (*perso).Vie_actuel, "/", (*perso).Vie_max, " et Essence : ", (*perso).Essence_actuel, "/", (*perso).Essence_max)
		for i, v := range (*perso).Action {
			fmt.Println(i+1, " : ", v.Nom)
		}
		fmt.Println(" 6: inventaire ")

		if (*monstre).Nom == "M.A.X.I.M.E (Modèle Anatomique X-pert Interractif Mesurable Ergonomique)" {
			fmt.Println("7 : quittez M.A.X.I.M.E ")
			Maxime = true
		}

		invet := false
		useP := &rien
		var useM *attack
		var a int
		fmt.Scan(&a)

		if a > 0 && a <= 6 {
			if a < 6 {
				useP = (*perso).Action[a-1]
			} else {
				invet = true
			}
		} else if a == 7 && Maxime {
			(*monstre).Vie_actuel = 0
		}

		if tour > len((*monstre).turn) {
			if tour%len((*monstre).turn) != 0 {
				useM = (*monstre).turn[tour%len((*monstre).turn)-1]
			}
			useM = (*monstre).turn[len((*monstre).turn)%tour-1]
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

		if invet {
			invet = false
			SoloMonster(perso, monstre, useM)
		} else if first(int((*perso).BuffV*float64(useP.Vitesse)), int((*monstre).BuffV*float64(useM.Vitesse))) == "1" {
			Pfirst(perso, monstre, useP, useM)
		} else {

			Mfirst(perso, monstre, useP, useM)
		}
		Poison(perso)
		Pois(monstre)
		Feu(perso)
		FeuM(monstre)
		tour += 1
	}
	(*perso).Stun = false
	if IsDead(perso) {
		fmt.Println("===============")
		fmt.Println((*monstre).Nom, " est vaincu !!!!!")
		if !Maxime {
			(*perso).Piece += (*monstre).Or
			fmt.Println("Vous avez gagner ", (*monstre).Or, " piece.")
			(*perso).Xp += (*monstre).Xp
			fmt.Println("Vous avez gagner ", (*monstre).Xp, " Xp.")
		}

		roll := rand.Intn(len((*monstre).Loot))
		(*perso).Inventaire = append((*perso).Inventaire, (*monstre).Loot[roll])
		fmt.Println("Vous avez récuperé", (*monstre).Loot[roll])
	} else {
		fmt.Println("Vous avez failli à votre mission")
		Res(perso)
	}
}

func TPersonnage(perso *character, monstre *monster, useP *attack) {
	var a int
	fmt.Println((*perso).Nom, " utilise ", (*useP).Nom)

	if (*useP).Degat {
		if (*useP).Essence {
			if Useessence(perso, useP) {
				a = int((float64(Critique(useP)) * (*perso).BuffA) * (*monstre).Buff)
				(*monstre).Vie_actuel -= a
				fmt.Println((*perso).Nom, " reçois -", (a))
			}
		} else {
			a = int((float64(Critique(useP)) * (*perso).BuffA) * (*monstre).Buff)
			(*monstre).Vie_actuel -= a
			fmt.Println((*perso).Nom, " reçois -", (a))
		}
	}

	if (*useP).Buff {
		UseBuff(perso, useP)
	}

	if (*useP).Debuff {
		UseDebuff(perso, monstre, useP)
	}

	TB(monstre, (*monstre).Buff)
	TurnBuff(perso, (*perso).BuffA)
}

func TMonster(perso *character, monstre *monster, useM *attack) {
	var a int
	fmt.Println((*monstre).Nom, " utilise ", (*useM).Nom)

	if (*useM).Degat {
		a = int((float64(Critique(useM)) * (*monstre).BuffA) * (*perso).Buff)
		(*perso).Vie_actuel -= a
		fmt.Println((*perso).Nom, " -", (a))

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

func Useessence(perso *character, a *attack) bool {
	b := true
	if (*a).EssenceCost > 0 {
		fmt.Println("Vous dépensez ", (*a).EssenceCost, " d'essence avec ", (*a).Nom)
	}
	if (*perso).Essence_actuel >= (*a).EssenceCost {
		(*perso).Essence_actuel -= (*a).EssenceCost
	} else {
		fmt.Println("Vous n'avez pas assez d'essence")
		b = false
	}
	fmt.Println("Essence : ", (*perso).Essence_actuel, "/", (*perso).Essence_max)
	return b
}
