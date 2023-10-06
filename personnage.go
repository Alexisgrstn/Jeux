package main

import (
	"fmt"
	"time"
)

type Personnage struct {
	Nom                    string
	Classe                 string
	Niveau                 int
	PvMax                  int
	PvActuels              int
	Inventaire             []Item
	Sort                   []string
	Argent                 int
	Equipement             map[string]int
	AmeliorationInventaire int
	MaxSlotsInventaire     int
	PointsDAttaque         int
	MaxExp                 int
	Level                  int
	Experience             int
	Mana                   int
	MaxMana                int
}

type Bouclier struct {
	Head  string
	Torso string
	Feet  string
}

type Item struct {
	Name  string
	Value int
}

type Monstre struct {
	Nom            string
	PvMax          int
	PvActuels      int
	PointsDAttaque int
}

func InitMonstre() Monstre {
	return Monstre{
		Nom:            "Gobelin d'entrainement",
		PvMax:          40,
		PvActuels:      40,
		PointsDAttaque: 20,
	}
}

func (p *Personnage) Init(nom string, classe string, niveau int, pvMax int, pvActuels int, sort string, Argent int, Equipement map[string]int, PointsDAttaque int, MaxExp int, Level int, Experience int, Mana int, MaxMana int) {
	p.Nom = nom
	p.Classe = classe
	p.Niveau = niveau
	p.PvMax = pvMax
	p.PvActuels = pvActuels
	p.Inventaire = []Item{}
	p.Sort = []string{sort}
	p.Argent = 100
	p.Equipement = make(map[string]int)
	p.AmeliorationInventaire = 0
	p.MaxSlotsInventaire = 10
	p.PointsDAttaque = 20
	p.MaxExp = 100
	p.Level = 1
	p.Experience = 0
	p.Mana = 80
	p.MaxMana = 100

}

func NewPersonnage() *Personnage {
	return &Personnage{}
}

func (joueur *Personnage) displayInfo() {
	fmt.Println("=======================================")
	fmt.Printf("Nom: %s\n", joueur.Nom)
	fmt.Printf("Classe: %s\n", joueur.Classe)
	fmt.Printf("Niveau: %d\n", joueur.Niveau)
	fmt.Printf("Points de vie maximum: %d\n", joueur.PvMax)
	fmt.Printf("Points de vie actuels: %d\n", joueur.PvActuels)
	fmt.Printf("Points d'attaque: %d\n", joueur.PointsDAttaque)
	fmt.Printf("Skill: %+v\n", joueur.Sort)
	fmt.Printf("Inventaire: %+v\n", joueur.Inventaire)
	fmt.Printf("Argent: %d\n", joueur.Argent)
	fmt.Printf("Equipement: %+v\n", joueur.Equipement)
	fmt.Printf("Amélioration de l'inventaire: %d\n", joueur.AmeliorationInventaire)
	fmt.Printf("Nombre de slots de l'inventaire: %d\n", joueur.MaxSlotsInventaire)
	fmt.Printf("Expérience: %d\n", joueur.Experience)
	fmt.Println("=======================================")
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
		time.Sleep(5 * time.Second)
		clearScreen()
	}
}
