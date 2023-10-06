package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Fonction principale où le programme commence
func main() {
	// Création d'un nouveau personnage
	joueur := charCreation()
	for {
		// Vérifie si le joueur est mort
		joueur.dead()
		// Gère le menu principal du jeu
		handleMenu(joueur)
	}
}

// Fonction pour créer un nouveau personnage
func charCreation() *Personnage {
	var nom, classe string
	var pvMax, pvActuels int

	// Obtient le nom du personnage de l'utilisateur
	fmt.Print("Entrez le nom de votre personnage (uniquement des lettres): ")
	fmt.Scanln(&nom)

	// Normalise le nom (première lettre en majuscule, reste en minuscules)
	nom = NormalizeName(nom)
	clearScreen()

	// L'utilisateur choisit une classe pour le personnage
	fmt.Println("Choisissez la classe de votre personnage (Humain, Elfe, Nain):")
	fmt.Scanln(&classe)

	// Définit les points de vie maximaux en fonction de la classe choisie
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

	// Initialise les points de vie actuels à la moitié des points de vie maximaux
	pvActuels = pvMax / 2

	// Crée et initialise un nouveau personnage
	personnage := NewPersonnage()
	personnage.Init(nom, classe, 1, pvMax, pvActuels, "Coup de Poing", 100, make(map[string]int), 10, 100, 1, 0, 80, 100)
	return personnage
}

// Fonction pour normaliser le nom du personnage
func NormalizeName(nom string) string {
	if len(nom) == 0 {
		return charCreation().Nom
	}
	return strings.ToUpper(string(nom[0])) + strings.ToLower(nom[1:])
}

// Fonction pour effacer l'écran (Windows)
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return
}

// Fonction pour l'inventaire du forgeron (non utilisée dans main)
func ForgeronInventory() {
	p := &Personnage{
		Equipement: make(map[string]int),
	}
	Forgeron(p)
}
