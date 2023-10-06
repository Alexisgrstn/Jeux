package main

import (
	"fmt"
	"math/rand"
	"time"
)

func combatEntrainement(joueur *Personnage) {

	tour := 1
	monstre := InitMonstre()

	fmt.Printf("Points de vie du personnage avant le combat: %d\n", joueur.PvActuels)
	time.Sleep(2 * time.Second)
	fmt.Println("============================================================")
	fmt.Printf("Points de vie du monstre avant le combat: %d\n", monstre.PvActuels)
	time.Sleep(2 * time.Second)
	fmt.Println("============================================================")

	for {

		fmt.Printf("===== Tour %d =====\n", tour)

		MonstrePattern(tour, joueur, &monstre)
		time.Sleep(3 * time.Second)

		if joueur.PvActuels <= 0 {
			fmt.Println("Vous avez été vaincu!")
			joueur.dead()
			break
		}

		charTurn(joueur, &monstre, joueur.Inventaire, tour)

		tour++
	}
}

func MonstrePattern(tour int, joueur *Personnage, monstre *Monstre) {
	var damage int
	baseDamage := rand.Intn(monstre.PointsDAttaque)

	if tour%3 == 0 {
		damage = baseDamage * 2
		fmt.Println("Coup critique du monstre! Dégats doublés!")
	} else {
		damage = baseDamage
	}

	joueur.PvActuels -= damage
	fmt.Printf("%s attaque et inflige %d de dégâts ! %s Pv Restant %d/%d\n", monstre.Nom, damage, joueur.Nom, joueur.PvActuels, joueur.PvMax)
}
