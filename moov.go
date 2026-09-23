package projet_red

import (
	"fmt"
	"math/rand"
)

func Moov(perso *character) {
	fmt.Println("=====================")
	fmt.Println("\nVous vous deplacez vers une nouvelle zone...")
	Poison(perso)
	roll := rand.Intn(100) + 1
	if (*perso).count == 3 {
		(*perso).count = 0
		roll = rand.Intn(3) + 1
		if roll == 1 {
			roll = 1
		} else if roll == 2 {
			roll  = 70
		} else {
			roll = 76
		}
	}
	if roll <= 20 {
		fmt.Println("Un ennemi surgit de derrière le mur et vous fonce dessus !")
		perso.Combat = true
		m := InitMonster("Zombie")
		Combat(perso , &m)
	} else if roll <= 40 {
		fmt.Println("Vous tombez sur un campement. C'est l'ideal pour vous reposer ou crafter.")
		camp(perso)
		(*perso).count++
	} else if roll <= 60 {
		fmt.Println("Vous croisez un marchand ambulant qui propose de bonnes affaires.")
		openMarchand(perso)
		(*perso).count++
	} else if roll <= 75 {
		fmt.Println("Un ennemi surgit de derrière le mur et vous fonce dessus !")
		perso.Combat = true
		m := InitMonster("Claqueur")
		Combat(perso , &m)
	} else if roll <= 79 {
		fmt.Println("Un ennemi surgit de derrière le mur et vous fonce dessus !")
		perso.Combat = true
		m := InitMonster("Solar")
		Combat(perso , &m)
	} else {
		fmt.Println("Vous trouvez 10 pièce !!")
		(*perso).count++
	}
}
