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

func NewPersonnage() *Personnage {
	return &Personnage{}
}
func (joueur *Personnage) displayInfo() {
	clearScreen()
	fmt.Println("=======================================")
	fmt.Println("Informations du personnage:")
	fmt.Println("=======================================")
	fmt.Println("Nom:", joueur.Nom)
	fmt.Println("Classe:", joueur.Classe)
	fmt.Println("Niveau:", joueur.Niveau)
	fmt.Println("Points de vie maximum:", joueur.PvMax)
	fmt.Println("Points de vie actuels:", joueur.PvActuels)
	fmt.Println("Inventaire:", joueur.Inventaire)
	fmt.Println("Skill:", joueur.Sort)
	fmt.Println("Argent:", joueur.Argent)
	fmt.Println("=======================================")
}
