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

func usePotion(joueur *Personnage, index int) {
	const potionHeal = 50

	if joueur.PvActuels == joueur.PvMax {
		fmt.Println("Vous avez déjà plein de vie !")
		return
	}

	healAmount := joueur.PvActuels + potionHeal
	if healAmount > joueur.PvMax {
		healAmount = joueur.PvMax
	}

	healed := healAmount - joueur.PvActuels
	fmt.Printf("Vous avez utilisé une Redbull et avez été soigné de %d points de vie. Vous avez maintenant %d points de vie actuels sur %d.\n", healed, healAmount, joueur.PvMax)

	joueur.PvActuels = healAmount

	// Supprimer la potion de l'inventaire
	joueur.Inventaire = append(joueur.Inventaire[:index], joueur.Inventaire[index+1:]...)
}

func usePoisonPotion(joueur *Personnage, index int) {
	const poisonDamage = 10
	const poisonDuration = 3
	clearScreen()
	fmt.Println("=======================================")
	fmt.Println("Vous avez utilisé une Potion de poison.")
	fmt.Println("Vous subissez des dégâts de poison pendant 3 secondes.")
	fmt.Println("=======================================")

	for i := 0; i < poisonDuration; i++ {
		joueur.PvActuels -= poisonDamage
		if joueur.PvActuels < 0 {
			joueur.PvActuels = 0
		}

		fmt.Printf("Points de vie actuels: %d / %d\n", joueur.PvActuels, joueur.PvMax)
		time.Sleep(1 * time.Second)
	}

	// Supprimer la potion de l'inventaire
	joueur.Inventaire = append(joueur.Inventaire[:index], joueur.Inventaire[index+1:]...)
	fmt.Println("=======================================")
	fmt.Println("Les effets du poison ont disparu.")

}

func AddToInventory(joueur *Personnage, item Item) {
	if len(joueur.Inventaire) >= 10 {
		fmt.Println("=======================================")
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas ajouter plus d'articles.")
		return
	}
	joueur.Inventaire = append(joueur.Inventaire, item)
}

