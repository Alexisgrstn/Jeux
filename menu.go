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

