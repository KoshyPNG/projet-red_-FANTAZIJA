package main

import (
	"fmt"
	"projet_red"
)

func main() {
	perso := projet_red.InitCharacter()
	v := true
	for v {
		var i int
		fmt.Println("clivk")
		fmt.Scan(&i)
		if i == 1 {
			projet_red.DisplayInfo(perso)
			v = false
		} else if i == 2 {
			projet_red.AccessInventory(perso)
			v = false
		}
	}
}
