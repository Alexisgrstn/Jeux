package main

import (
	"fmt"
	"os"
)

func handleMenu(joueur *Personnage) {
	afficherMenu()
	choix := getChoice()
	executeChoice(joueur, choix)
}

func afficherMenu() {
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
	}
	return choix
}

func executeChoice(joueur *Personnage, choix int) {
	switch choix {
	case 1:
		joueur.displayInfo()
	case 2:
		menuInventaire(joueur)
	case 3:
		marchand(joueur)
	case 4:
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide. Veuillez entrer 1, 2, 3 ou 4.")
	}
}
