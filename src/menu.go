package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// handleMenu gère l'affichage et l'exécution du menu principal
func handleMenu(joueur *Personnage) {
	afficherMenu()
	choix := getChoice()
	executeChoice(joueur, choix)
}

// afficherMenu affiche les différentes options disponibles dans le menu
func afficherMenu() {
	clearScreen()
	fmt.Println("Menu:")
	fmt.Println("=======================================")
	// Options du menu
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder à l'inventaire du personnage")
	// ...
	fmt.Println("=======================================")
	fmt.Print("Choisissez une option (1, 2, 3, 4, 5, 6 ou 7): ")
	fmt.Println('\n')
	fmt.Println("=======================================")
}

// getChoice récupère le choix de l'utilisateur dans le menu
func getChoice() int {
	var choix int
	_, err := fmt.Scan(&choix)
	if err != nil {
		// Gestion des erreurs de saisie
		fmt.Println("Erreur de saisie. Veuillez entrer 1, 2, 3, 4, 5 ou 6.")
		return getChoice()
	}
	return choix
}

// executeChoice exécute l'action correspondant au choix de l'utilisateur
func executeChoice(joueur *Personnage, choix int) {
	switch choix {
	case 1:
		// Affiche les informations du personnage
		clearScreen()
		joueur.displayInfo()
		time.Sleep(5 * time.Second)
	case 2:
		// Gère l'inventaire du personnage
		clearScreen()
		menuInventaire(joueur)
	case 3:
		// Affiche le menu du marchand
		clearScreen()
		marchand(joueur)
	case 4:
		// Affiche le menu du forgeron
		clearScreen()
		Forgeron(joueur)
	case 5:
		// Gère un combat d'entraînement
		clearScreen()
		combatEntrainement(joueur)
	case 6:
		// Bonus : Affiche les noms d'artistes cachés
		clearScreen()
		fmt.Println("Les artistes cachés sont : Banksy et Da Vinci, A toi de jouer !")
		time.Sleep(5 * time.Second)
	case 7:
		// Quitte le jeu
		clearScreen()
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		// Gestion d'un choix invalide
		fmt.Println("Option invalide. Veuillez entrer 1, 2, 3, 4, 5, 6 ou 7.")
		time.Sleep(1 * time.Second)
		clearScreen()
	}
}

// charTurn gère le tour du personnage pendant un combat
func charTurn(joueur *Personnage, monstre *Monstre, Inventaire []Item, tour int) {
	var choice string // Variable pour stocker le choix de l'utilisateur

	// Affichage des options disponibles pour le joueur pendant son tour
	fmt.Println("Menu")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Scanln(&choice) // Lecture du choix de l'utilisateur

	// Exécution de l'action correspondant au choix de l'utilisateur
	switch choice {
	case "1":
		// Génère des dégâts aléatoires basés sur les PointsDAttaque du joueur
		damage := rand.Intn(joueur.PointsDAttaque)
		// Applique les dégâts au monstre
		monstre.PvActuels -= damage
		// Affichage des informations de l'attaque
		fmt.Printf("%s attaque et inflige %d de dégâts ! Monstre Pv Restant %d/%d\n", joueur.Nom, damage, monstre.PvActuels, monstre.PvMax)
		time.Sleep(3 * time.Second)
		fmt.Println("============================================================")

		// Vérifie si le monstre est vaincu
		if monstre.PvActuels <= 0 {
			fmt.Println("Le monstre est vaincu!")
			break
		}
	case "2":
		// Gère l'utilisation d'items depuis l'inventaire
		if len(Inventaire) == 0 {
			fmt.Println("Votre inventaire est vide.") // Si l'inventaire est vide
		} else {
			// Liste les items disponibles dans l'inventaire
			fmt.Println("Inventaire:")
			for i, item := range Inventaire {
				fmt.Printf("%d. %s\n", i+1, item.Name)
			}

			// Demande à l'utilisateur quel item utiliser
			fmt.Println("Choisissez un objet à utiliser:")
			var itemChoice int
			fmt.Scanln(&itemChoice)

			// Utilise l'item si le choix est valide
			if itemChoice <= len(Inventaire) && itemChoice > 0 {
				item := Inventaire[itemChoice-1]
				joueur.utiliserItem(item.Name)
			} else {
				fmt.Println("Choix invalide.") // Si le choix est invalide
			}
		}
	default:
		fmt.Println("Choix invalide.") // Si l'utilisateur ne choisit ni '1' ni '2'
	}
}
