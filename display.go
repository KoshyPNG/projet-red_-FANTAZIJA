package projet_red

import "fmt"

func DisplayInfo(perso *character) {
	fmt.Println("==========================")
	fmt.Println("         CHARACTER        ")
	fmt.Print("===========")
	fmt.Print(perso.Nom)
	fmt.Print("===========\n")
	fmt.Println("Classe :  ", perso.Classe)
	fmt.Print("Vie : ",perso.Vie_actuel)
	fmt.Println("/",perso.Vie_max)
	fmt.Print("Essence : ",perso.Essence_actuel)
	fmt.Println("/",perso.Essence_max)
	if (*perso).Poison {
		fmt.Println("Vous ètes empoisonné")
	}
	tab := []string{"TÊTE : " ,"TORSE : " ,"JAMBE : "}
	for i,v := range tab {
		println(v, perso.Equipement[i].Nom)
	}
	fmt.Println()
	fmt.Println("==========================")
	fmt.Println(perso.Piece,"pièce d'or")
	fmt.Println("==========================")
	fmt.Println("Niveau", perso.Level)
	fmt.Println(perso.Xp, "% du niveau")
	fmt.Println("==========================")
}
