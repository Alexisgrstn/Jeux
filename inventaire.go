package main

import (
	"fmt"
	"time"
)

type Marchand struct {
	ItemsEnVente []string
}
func menuInventaire(joueur *Personnage) {
	for {
		clearScreen()
		fmt.Println("=======================================")
		fmt.Println("Inventaire du personnage:")
		fmt.Println("=======================================")
		for i, item := range joueur.Inventaire {
			fmt.Printf("%d. %s (%d)\n", i+1, item.Name, item.Value)
		}
		fmt.Println("=======================================")
		fmt.Println("0. Retour au menu précédent")
		inventoryChoice := AskForInventoryChoice(joueur)
		if inventoryChoice == 0 {
			break
		}
		useInventoryItem(joueur, inventoryChoice)
	}
}

