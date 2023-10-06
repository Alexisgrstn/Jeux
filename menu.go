package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func handleMenu(joueur *Personnage) {
	afficherMenu()
	choix := getChoice()
	executeChoice(joueur, choix)
}

func afficherMenu() {
	clearScreen()
	fmt.Println("Menu:")
	fmt.Println("=======================================")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder à l'inventaire du personnage")
	fmt.Println("3. Marchand")
	fmt.Println("4. Forgeron")
	fmt.Println("5. Combat d'entraînement")
	fmt.Println("6. Qui sont ils ?(Bonus)")
	fmt.Println("7. Quitter")
	fmt.Println("=======================================")
	fmt.Print("Choisissez une option (1, 2, 3, 4, 5, 6 ou 7): ")
	fmt.Println('\n')
	fmt.Println("=======================================")
}

func getChoice() int {
	var choix int
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de saisie. Veuillez entrer 1, 2, 3, 4, 5 ou 6.")
		return getChoice()
	}
	return choix
}
func executeChoice(joueur *Personnage, choix int) {
	switch choix {
	case 1:
		clearScreen()
		joueur.displayInfo()
		time.Sleep(5 * time.Second)
	case 2:
		clearScreen()
		menuInventaire(joueur)
	case 3:
		clearScreen()
		marchand(joueur)
	case 4:
		clearScreen()
		Forgeron(joueur)
	case 5:
		clearScreen()
		combatEntrainement(joueur)
	case 6:
		clearScreen()
		fmt.Println("Les artistes cachés sont : Banksy et Da Vinci")
		time.Sleep(5 * time.Second)
	case 7:
		clearScreen()
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide. Veuillez entrer 1, 2, 3, 4, 5, 6 ou 7.")
		time.Sleep(1 * time.Second)
		clearScreen()
	}
}

func charTurn(joueur *Personnage, monstre *Monstre, Inventaire []Item, tour int) {
	var choice string

	fmt.Println("Menu")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		damage := rand.Intn(joueur.PointsDAttaque)
		monstre.PvActuels -= damage
		fmt.Printf("%s attaque et inflige %d de dégâts ! Monstre Pv Restant %d/%d\n", joueur.Nom, damage, monstre.PvActuels, monstre.PvMax)
		time.Sleep(3 * time.Second)
		fmt.Println("============================================================")

		if monstre.PvActuels <= 0 {
			fmt.Println("Le monstre est vaincu!")
			break
		}
	case "2":
		if len(Inventaire) == 0 {
			fmt.Println("Votre inventaire est vide.")
		} else {
			fmt.Println("Inventaire:")
			for i, item := range Inventaire {
				fmt.Printf("%d. %s\n", i+1, item.Name)
			}

			fmt.Println("Choisissez un objet à utiliser:")
			var itemChoice int
			fmt.Scanln(&itemChoice)

			if itemChoice <= len(Inventaire) && itemChoice > 0 {
				item := Inventaire[itemChoice-1]
				joueur.utiliserItem(item.Name)
			} else {
				fmt.Println("Choix invalide.")
			}
		}
	default:
		fmt.Println("Choix invalide.")
	}
}
