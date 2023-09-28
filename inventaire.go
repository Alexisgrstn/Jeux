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

func AskForInventoryChoice(joueur *Personnage) int {
	var inventoryChoice int
	fmt.Println("=======================================")
	fmt.Print("Choisissez une option (0 pour retourner au menu principal): ")
	_, err := fmt.Scan(&inventoryChoice)
	if err != nil {
		fmt.Println("Erreur de saisie. Veuillez entrer un nombre valide.")
		return -1
	}
	if inventoryChoice < 0 || inventoryChoice > len(joueur.Inventaire) {
		fmt.Println("Option invalide. Veuillez entrer un nombre valide ou 0 pour retourner au menu principal.")
		return -1
	}
	return inventoryChoice
}

func useInventoryItem(joueur *Personnage, inventoryChoice int) {
	index := inventoryChoice - 1
	selectedItem := joueur.Inventaire[index]
	if selectedItem.Name == "RedBull" {
		usePotion(joueur, index)
	} else if selectedItem.Name == "Potion de poison" {
		usePoisonPotion(joueur, index)
	} else if selectedItem.Name == "Livre de Sort : Boule de Feu" {
		joueur.SpellBook()
		// ...
	} else {
		fmt.Println("Vous ne pouvez utiliser que des RedBull, des potions de poison ou des livres de sort.")
	}
}

