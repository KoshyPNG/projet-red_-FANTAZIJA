package projet_red

import "fmt"

func AccessInventory(perso *character) {
	fmt.Println("==========================")
	fmt.Println("         INVENTAIRE       ")
	fmt.Println("==========================")
	for i, a := range (*perso).Inventaire {
		fmt.Print(i+1, ": ")
		fmt.Println(a)
	}
	var a int
	fmt.Println("==========================")
	fmt.Println("Inventaire MAX : ",(*perso).Inventaire_max)
	fmt.Println("==========================")
	fmt.Println("Ecrire le numero de l'objet pour utilisé ou 0 pour rien faire.")
	fmt.Scan(&a)
	if a > 0 && a <= len((*perso).Inventaire) {
		Use_object(a, perso)
	}
}

func AccessInventoryC(perso *character, monstre *monster) {
	fmt.Println("==========================")
	fmt.Println("         INVENTAIRE       ")
	fmt.Println("==========================")
	for i, a := range (*perso).Inventaire {
		fmt.Print(i+1, ": ")
		fmt.Println(a)
	}
	var a int
	fmt.Println("==========================")
	fmt.Println("Inventaire MAX : ",(*perso).Inventaire_max)
	fmt.Println("==========================")
	fmt.Println("Ecrire le numero de l'objet pour utilisé ou 0 pour rien faire.")
	fmt.Scan(&a)
	if a > 0 && a <= len((*perso).Inventaire) {
		Use_object_C(a, perso, monstre)
	}
}