package main

import (
	"fmt"
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
	fmt.Println("4. Quitter")
	fmt.Println("=======================================")
	fmt.Print("Choisissez une option (1, 2, 3 ou 4): ")
	fmt.Println('\n')
	fmt.Println("=======================================")
}

func getChoice() int {
	var choix int
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de saisie. Veuillez entrer 1, 2, 3 ou 4.")
		return getChoice()
	}
	return choix
}
func executeChoice(joueur *Personnage, choix int) {
	switch choix {
	case 1:
		clearScreen()
		joueur.displayInfo()
	case 2:
		clearScreen()
		menuInventaire(joueur)
	case 3:
		clearScreen()
		marchand(joueur)
	case 4:
		clearScreen()
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide. Veuillez entrer 1, 2, 3 ou 4.")
		time.Sleep(1 * time.Second)
		clearScreen()
	}
}
