package main

import (
	"fmt"
	"time"
	"math/rand"
)

func combatEntrainement(joueur *Personnage) {

	tour := 1

	monstre := InitGoblin()
		fmt.Printf("Points de vie du personnage avant le combat: %d\n", joueur.PvActuels)
		time.Sleep(2 * time.Second)
		fmt.Println("============================================================")
		fmt.Printf("Points de vie du monstre avant le combat: %d\n", monstre.PvActuels)
		time.Sleep(2 * time.Second)
		fmt.Println("============================================================")

		rand.Seed(time.Now().UnixNano())

	for {
		damage := rand.Intn(joueur.PointsDAttaque)
		fmt.Printf("===== Tour %d =====\n", tour)
		monstre.PvActuels -= damage
		fmt.Printf("%s attaque et inflige %d de dégâts !\n", joueur.Nom, damage)
		time.Sleep(2 * time.Second)
		fmt.Println("============================================================")

		if monstre.PvActuels <= 0 {
			fmt.Println("Le monstre est vaincu!")
			time.Sleep(2 * time.Second)
			fmt.Println("============================================================")
			break
		}
		damage = rand.Intn(monstre.PointsDAttaque)
		joueur.PvActuels -= damage
		fmt.Printf("%s attaque et inflige %d de dégâts !\n", monstre.Nom, damage)
		time.Sleep(2 * time.Second)
		fmt.Println("============================================================")

		if joueur.PvActuels <= 0 {
			fmt.Println("Vous avez été vaincu!")
			time.Sleep(2 * time.Second)
			joueur.dead()
			break
		}
		tour++
	}
}
