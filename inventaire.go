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
		time.Sleep(2 * time.Second)
		return
	}

	healAmount := joueur.PvActuels + potionHeal
	if healAmount > joueur.PvMax {
		healAmount = joueur.PvMax
	}

	healed := healAmount - joueur.PvActuels
	fmt.Printf("Vous avez utilisé une Redbull et avez été soigné de %d points de vie. Vous avez maintenant %d points de vie actuels sur %d.\n", healed, healAmount, joueur.PvMax)
	time.Sleep(2 * time.Second)
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

func (joueur *Personnage) AddToInventory(item Item, cost int) bool {
	if len(joueur.Inventaire) >= 10 {
		fmt.Println("=======================================")
		fmt.Println("Votre inventaire est plein. Vous ne pouvez pas ajouter plus d'articles.")
		fmt.Println("=======================================")
		time.Sleep(1 * time.Second)
		clearScreen()
		return false
	}
	if joueur.Argent < cost {
		fmt.Println("Vous n'avez pas assez d'argent.")
		return false
	}
	joueur.Inventaire = append(joueur.Inventaire, item)
	fmt.Printf("Item %s ajouté à l'inventaire.\n", item.Name)
	time.Sleep(1 * time.Second)
	joueur.Argent -= cost
	return true
}

func marchand(joueur *Personnage) {
	for {
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
		fmt.Println("8: Augmentation d'inventaire - 30 pièces d'or")
		fmt.Println("9. Quitter")
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
			if joueur.AddToInventory(Item{Name: "RedBull", Value: 3}, 3) {
				clearScreen()
			}
		case 2:
			if joueur.AddToInventory(Item{Name: "Potion de poison", Value: 6}, 6) {
				clearScreen()
			}
		case 3:
			if joueur.AddToInventory(Item{Name: "Livre de Sort : Boule de feu", Value: 25}, 25) {
				clearScreen()
			}
		case 4:
			if joueur.AddToInventory(Item{Name: "Fourrure de Loup", Value: 4}, 4) {
				clearScreen()
			}
		case 5:
			if joueur.AddToInventory(Item{Name: "Peau de Troll", Value: 7}, 7) {
				clearScreen()
			}
		case 6:
			if joueur.AddToInventory(Item{Name: "Cuir de Sanglier", Value: 3}, 3) {
				clearScreen()
			}
		case 7:
			if joueur.AddToInventory(Item{Name: "Plume de Corbeau", Value: 1}, 1) {
				clearScreen()
			}
		case 8:
			if joueur.Argent >= 30 {
				joueur.upgradeInventorySlot()
			} else {
				fmt.Println("Vous n'avez pas assez d'argent.")
			}
		case 9:
			return
		default:
			fmt.Println("Option invalide. Veuillez entrer un nombre entre 1 et 8.")
		}

		if choix >= 1 && choix <= 7 && joueur.Argent < choix {
			fmt.Println("Vous n'avez pas assez d'argent.")
			time.Sleep(2 * time.Second)
			clearScreen()
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
func Forgeron(joueur *Personnage) {
	inventaire := make(map[string]int)
	or := 20

	// Initialisation des ressources de l'inventaire
	inventaire["Plume de Corbeau"] = 2
	inventaire["Cuir de Sanglier"] = 2
	inventaire["Fourrure de loup"] = 2
	inventaire["Peau de Troll"] = 1

	for {
		for {
			fmt.Println("Menu du Forgeron:")
			fmt.Println("=======================================")
			fmt.Println("1. Chapeau de l'aventurier")
			fmt.Println("2. Tunique de l'aventurier")
			fmt.Println("3. Bottes de l'aventurier")
			fmt.Println("4. Retour")
			fmt.Println("=======================================")

			var choix int

			if choix == 4 {
				return
			}

			var sousChoix int
			fmt.Scanln(&sousChoix)
			clearScreen()

			if sousChoix == 4 {
				return
			}

			if or < 5 {
				fmt.Println("Vous n'avez pas assez d'or pour fabriquer un équipement.")
				continue
			}

			switch sousChoix {
			case 1:
				if inventaire["Plume de Corbeau"] < 1 {
					fmt.Println("Il vous manque 1 Plume de Corbeau.")
					time.Sleep(2 * time.Second)
					clearScreen()
					continue
				}
				if inventaire["Cuir de Sanglier"] < 1 {
					fmt.Println("Il vous manque 1 Cuir de Sanglier.")
					time.Sleep(2 * time.Second)
					clearScreen()
					continue
					}
				or -= 5
				inventaire["Plume de Corbeau"]--
				inventaire["Cuir de Sanglier"]--
				joueur.Equipement["Chapeau de l'aventurier"]++
				fmt.Println("Vous avez fabriqué un chapeau de l'aventurier!")
				time.Sleep(2 * time.Second)
				clearScreen()

			case 2:

				if inventaire["Fourrure de loup"] < 2 {
				fmt.Println("Il vous manque de la Fourrure de loup.")
				time.Sleep(2 * time.Second)
				clearScreen()
				continue
			}
			if inventaire["Peau de Troll"] < 1 {
				fmt.Println("Il vous manque 1 Peau de Troll.")
				time.Sleep(2 * time.Second)
				clearScreen()
				continue
				}
				or -= 5
				inventaire["Fourrure de loup"] -= 2
				inventaire["Peau de Troll"]--
				joueur.Equipement["Tunique de l'aventurier"]++
				fmt.Println("Vous avez fabriqué une tunique de l'aventurier!")
				time.Sleep(2 * time.Second)
				clearScreen()
			case 3:

				if inventaire["Fourrure de loup"] < 1 {
					fmt.Println("Il vous manque 1 Fourrure de loup.")
					time.Sleep(2 * time.Second)
					clearScreen()
					continue
				}
				if inventaire["Cuir de Sanglier"] < 1 {
					fmt.Println("Il vous manque 1 Cuir de Sanglier.")
					time.Sleep(2 * time.Second)
					clearScreen()
					continue
				}
				or -= 5
				inventaire["Fourrure de loup"]--
				inventaire["Cuir de Sanglier"]--
				joueur.Equipement["Bottes de l'aventurier"]++
				fmt.Println("Vous avez fabriqué des bottes de l'aventurier!")
				time.Sleep(2 * time.Second)
				clearScreen()
			default:
				fmt.Println("Choix non valide.")
				clearScreen()
			}
		}
	}
}
func ForgeronInventory() {
	p := &Personnage{
		Equipement: make(map[string]int),
	}

	Forgeron(p)
}

func (joueur *Personnage) upgradeInventorySlot() {
	clearScreen()
    if joueur.AmeliorationInventaire >= 3 {
        fmt.Println("Vous avez atteint la limite maximale d'augmentations d'inventaire.")
        return
    }
    if joueur.Argent < 30 {
        fmt.Println("Vous n'avez pas assez d'argent pour cette amélioration.")
        return
    }
    
    joueur.MaxSlotsInventaire += 10
    joueur.AmeliorationInventaire++
    joueur.Argent -= 30
    fmt.Println("La capacité de votre inventaire a été augmentée de 10 slots.")
}
