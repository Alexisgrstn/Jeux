package main

import "fmt"

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	PvMax      int
	PvActuels  int
	Inventaire []Item
	Sort       []string
	Argent     int
}

