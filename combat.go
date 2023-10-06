package main

import (
	"fmt"
	"math/rand"
	"time"
)

// combatEntrainement gère le combat entre un personnage joueur et un monstre
func combatEntrainement(joueur *Personnage) {

	tour := 1                // Initialise le compteur de tours
	monstre := InitMonstre() // Initialise le monstre

	// Affiche les points de vie du joueur avant le combat
	fmt.Printf("Points de vie du personnage avant le combat: %d\n", joueur.PvActuels)
	time.Sleep(2 * time.Second)
	fmt.Println("============================================================")
	// Affiche les points de vie du monstre avant le combat
	fmt.Printf("Points de vie du monstre avant le combat: %d\n", monstre.PvActuels)
	time.Sleep(2 * time.Second)
	fmt.Println("============================================================")

	// Boucle principale du combat
	for {

		fmt.Printf("===== Tour %d =====\n", tour)

		// Exécute le pattern d'attaque du monstre
		MonstrePattern(tour, joueur, &monstre)
		time.Sleep(3 * time.Second)

		// Vérifie si le joueur est mort
		if joueur.PvActuels <= 0 {
			fmt.Println("Vous avez été vaincu!")
			joueur.dead()
			break
		}

		// Exécute le tour du joueur
		charTurn(joueur, &monstre, joueur.Inventaire, tour)

		// Incrémente le compteur de tours
		tour++
	}
}

// MonstrePattern exécute le pattern d'attaque du monstre
func MonstrePattern(tour int, joueur *Personnage, monstre *Monstre) {
	var damage int                                  // Initialise les dégâts
	baseDamage := rand.Intn(monstre.PointsDAttaque) // Calcule les dégâts de base

	// Si le tour est un multiple de 3, double les dégâts
	if tour%3 == 0 {
		damage = baseDamage * 2
		fmt.Println("Coup critique du monstre! Dégats doublés!")
	} else {
		damage = baseDamage
	}

	// Applique les dégâts au joueur
	joueur.PvActuels -= damage
	fmt.Printf("%s attaque et inflige %d de dégâts ! %s Pv Restant %d/%d\n", monstre.Nom, damage, joueur.Nom, joueur.PvActuels, joueur.PvMax)
}
