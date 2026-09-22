package projet_red

import (
	"fmt"
	"math/rand"
)

func Moov(perso *character) {
	fmt.Println("=====================")
	fmt.Println("\nVous vous deplacez vers une nouvelle zone...")
	if (*perso).Poison {
		fmt.Println("Vous vous sentez malade")
		(*perso).Tpois--
		if (*perso).Tpois == 0 {
			(*perso).Poison = false
		}
		(*perso).Vie_actuel -= 10
	}

	roll := rand.Intn(100) + 1

	if roll <= 20 {
		fmt.Println("Un gobelin surgit des buissons et vous attaque !")
	} else if roll <= 40 {
		fmt.Println("Vous tombez sur un campement. C'est l'ideal pour vous reposer ou crafter.")
		camp(perso)
	} else if roll <= 60 {
		fmt.Println("Vous croisez un marchand ambulant qui propose de bonnes affaires.")
		openMarchand(perso)
	} else {
		fmt.Println("La route est calme et sure. Vous avancez tranquillement.")
		m := InitMonster("Zombie")
		Combat(perso , &m)
	}
}
