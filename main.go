package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	joueur := charCreation()
	for {
		joueur.dead()
		handleMenu(joueur)
	}
}

func charCreation() *Personnage {
	var nom, classe string
	var pvMax, pvActuels int

	fmt.Print("Entrez le nom de votre personnage (uniquement des lettres): ")
	fmt.Scanln(&nom)
	nom = NormalizeName(nom)
	clearScreen()
	fmt.Println("Choisissez la classe de votre personnage (Humain, Elfe, Nain):")
	fmt.Scanln(&classe)

	switch classe {
	case "Humain":
		pvMax = 100
	case "Elfe":
		pvMax = 80
	case "Nain":
		pvMax = 120
	default:
		fmt.Println("Classe non reconnue, définie en Humain par défaut.")
		classe = "Humain"
		pvMax = 100
	}
	clearScreen()

	pvActuels = pvMax / 2

	personnage := NewPersonnage()
	personnage.Init(nom, classe, 1, pvMax, pvActuels, "Coup de Poing", 100)
	return personnage
}
