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
func (joueur *Personnage) AddToInventory(item Item) {
	joueur.Inventaire = append(joueur.Inventaire, item)
}

func (joueur *Personnage) SpellBook() {
	for _, v := range joueur.Sort {
		if v == "Boule de feu" {
			return
		}
	}
	joueur.Sort = append(joueur.Sort, "Boule de feu")
}

func (joueur *Personnage) dead() {
	if joueur.PvActuels <= 0 {
		clearScreen()
		fmt.Println("=======================================")
		fmt.Print("\033[31m") // Définir la couleur du texte en rouge
		fmt.Println("           VOUS ÊTES MORT !           ")
		fmt.Print("\033[0m") // Réinitialiser la couleur du texte
		fmt.Println("=======================================")
		joueur.PvActuels = joueur.PvMax / 2
		fmt.Printf("vous avez été réanimé avec %d point de vie grâce à une RedBull.\n", joueur.PvActuels)
		fmt.Println("Prenez garde et continuez votre aventure!")
	}
}
