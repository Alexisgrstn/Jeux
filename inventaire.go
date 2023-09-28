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

func marchand(joueur *Personnage) {
	for {
		clearScreen()
		fmt.Println("=======================================")
		fmt.Println("Bienvenue au magasin!")
		fmt.Println("=======================================")
		fmt.Println("1. RedBull - 3 pièces d'or")
		fmt.Println("2. Potion de poison - 6 pièces d'or")
		fmt.Println("3. Livre de Sort : Boule de feu - 25 pièces d'or")
		fmt.Println("4. Fourrure de Loup - 4 pièces d'or")
		fmt.Println("5. Peau de Troll - 7 pièces d'or")
		fmt.Println("6. Cuir de Sanglier - 3 pièces d'or")
		fmt.Println("7. Plume de Corbeau - 1 pièce d'or")
		fmt.Println("8. Quitter")
		fmt.Println("=======================================")
		fmt.Printf("Argent disponible: %d\n", joueur.Argent)
		fmt.Print("Faites votre choix: ")

		var choix int
		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Erreur de saisie. Veuillez entrer un nombre entre 1 et 8.")
			continue
		}
		switch choix {
		case 1:
			if joueur.Argent >= 3 {
				AddToInventory(joueur, Item{Name: "RedBull", Value: 3})
				joueur.Argent -= 3
			}
			clearScreen()
		case 2:
			if joueur.Argent >= 6 {
				AddToInventory(joueur,Item{Name: "Potion de poison", Value: 6})
				joueur.Argent -= 6
			}
			clearScreen()
		case 3:
			if joueur.Argent >= 25 {
				AddToInventory(joueur,Item{Name: "Livre de Sort : Boule de feu", Value: 25})
				joueur.Argent -= 25
			}
			clearScreen()
		case 4:
			if joueur.Argent >= 4 {
				AddToInventory(joueur,Item{Name: "Fourrure de Loup", Value: 4})
				joueur.Argent -= 4
			}
			clearScreen()
		case 5:
			if joueur.Argent >= 7 {
				AddToInventory(joueur,Item{Name: "Peau de Troll", Value: 7})
				joueur.Argent -= 7
			}
			clearScreen()
		case 6:
			if joueur.Argent >= 3 {
				AddToInventory(joueur,Item{Name: "Cuir de Sanglier", Value: 3})
				joueur.Argent -= 3
			}
			clearScreen()
		case 7:
			if joueur.Argent >= 1 {
				AddToInventory(joueur,Item{Name: "Plume de Corbeau", Value: 1})
				joueur.Argent -= 1
			}
			clearScreen()
		case 8:
			return
		default:
			fmt.Println("Option invalide. Veuillez entrer un nombre entre 1 et 8.")
		}

		if choix >= 1 && choix <= 7 && joueur.Argent < choix {
			fmt.Println("Vous n'avez pas assez d'argent.")
		}
	}
}

func askForMarchandChoice() int {
	var choix int
	fmt.Print("Choisissez une option (0, 1 ou 2): ")
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de saisie. Veuillez entrer 0, 1 ou 2.")
		return -1
	}
	return choix
}

func NewMarchand() *Marchand {
	return &Marchand{
		ItemsEnVente: []string{"Livre de Sort : Boule de Feu"},
	}
}

func (marchand *Marchand) vendreItem(joueur *Personnage, item string) {
	for _, v := range marchand.ItemsEnVente {
		if v == item {
			joueur.Inventaire = append(joueur.Inventaire, Item{})
			return
		}
	}
}

func (joueur *Personnage) utiliserItem(item string) {
	for index, v := range joueur.Inventaire {
		if v.Name == item {
			switch item {
			case "Livre de Sort : Boule de Feu":
				joueur.SpellBook()
				joueur.Inventaire = append(joueur.Inventaire[:index], joueur.Inventaire[index+1:]...)
				return
			}
		}
	}
	fmt.Println("L'item n'est pas dans l'inventaire.")
}
