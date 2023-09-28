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

