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

type Item struct {
	Name  string
	Value int
}

func (p *Personnage) Init(nom string, classe string, niveau int, pvMax int, pvActuels int, sort string, Argent int) {
	p.Nom = nom
	p.Classe = classe
	p.Niveau = niveau
	p.PvMax = pvMax
	p.PvActuels = pvActuels
	p.Inventaire = []Item{}
	p.Sort = []string{sort}
	p.Argent = 100
}

