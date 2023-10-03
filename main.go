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
	personnage.Init(nom, classe, 1, pvMax, pvActuels, "Coup de Poing", 100, make(map[string]int))
	return personnage
}

func NormalizeName(nom string) string {
	if len(nom) == 0 {
		return ""
	}
	return strings.ToUpper(string(nom[0])) + strings.ToLower(nom[1:])
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ForgeronInventory() {
	p := &Personnage{
		Equipement: make(map[string]int),
	}

	Forgeron(p)
}
