package projet_red

import "fmt"

func DisplayInfo(perso character) {
	fmt.Println("==========================")
	fmt.Print("===========")
	fmt.Print(perso.Nom)
	fmt.Print("===========\n")
	fmt.Print("Vie : ",perso.Vie_actuel)
	fmt.Print("/",perso.Vie_max)
}
