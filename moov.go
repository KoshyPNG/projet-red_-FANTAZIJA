package projet_red

import (
	"fmt"
	"math/rand"
)

func moov() {
	fmt.Println("\n Vous vous deplacez vers une nouvelle zone...")

	roll := rand.Intn(100) + 1

	if roll <= 20 {
		fmt.Println(" Un gobelin surgit des buissons et vous attaque !")
	} else if roll <= 40 {
		fmt.Println(" Vous tombez sur un campement. C'est l'ideal pour vous reposer ou crafter.")
	} else if roll <= 60 {
		fmt.Println(" Vous croisez un marchand ambulant qui propose de bonnes affaires.")
	} else {
		fmt.Println(" La route est calme et sure. Vous avancez tranquillement.")
	}
}
