package projet_red

func remoov(index int , perso *character) {
	if index != -1 {
		(*perso).Inventaire = append((*perso).Inventaire[:index], (*perso).Inventaire[index+1:]...)
	}
}
