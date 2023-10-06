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

func charTurn(joueur *Personnage, monstre *Personnage, inventory []string) {
	fmt.Println("Choisissez une action :")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		monstre.PvActuels -= joueur.PointsDAttaque
		fmt.Printf("%s inflige %d dégâts à %s\n", joueur.Nom, joueur.PointsDAttaque, monstre.Nom)
	case 2:
		fmt.Println("Objets dans l'inventaire :")
		for i, item := range inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		// Gérez le choix de l'objet ici
	case 3:
		fmt.Println("Choisissez un sort :")
		fmt.Println("1. Coup de poing (8 dégâts)")
		fmt.Println("2. Boule de feu (18 dégâts)")

		var spellChoice int
		fmt.Scanln(&spellChoice)

		switch spellChoice {
		case 1:
			if joueur.Mana >= 5 {
				damage := 8
				monstre.PvActuels -= damage
				joueur.Mana -= 5
				fmt.Printf("Coup de poing inflige %d dégâts à %s\n", damage, monstre.Nom)
			} else {
				fmt.Println("Mana insuffisant!")
			}
		case 2:
			if joueur.Mana >= 10 {
				damage := 18
				monstre.PvActuels -= damage
				joueur.Mana -= 10
				fmt.Printf("Boule de feu inflige %d dégâts à %s\n", damage, monstre.Nom)
			} else {
				fmt.Println("Mana insuffisant!")
			}
		default:
			fmt.Println("Choix non valide.")
		}
	default:
		fmt.Println("Choix non valide.")
	}
	fmt.Printf("%s PV : %d / %d, Mana : %d / %d\n", joueur.Nom, joueur.PvActuels, joueur.PvMax, joueur.Mana, joueur.MaxMana)
}
