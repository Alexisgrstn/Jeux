package main

import (
	"fmt"
	"time"
)

// Structure représentant un personnage
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

// Structure pour un bouclier
type Bouclier struct {
	Head  string
	Torso string
	Feet  string
}

// Structure pour un item
type Item struct {
	Name  string
	Value int
}

// Structure représentant un monstre
type Monstre struct {
	Nom            string
	PvMax          int
	PvActuels      int
	PointsDAttaque int
}

// Fonction pour initialiser un monstre
func InitMonstre() Monstre {
	return Monstre{
		Nom:            "Gobelin d'entrainement",
		PvMax:          40,
		PvActuels:      40,
		PointsDAttaque: 20,
	}
}

// Méthode pour initialiser un personnage
func (p *Personnage) Init(nom string, classe string, niveau int, pvMax int, pvActuels int, sort string, Argent int, Equipement map[string]int, PointsDAttaque int, MaxExp int, Level int, Experience int, Mana int, MaxMana int) {
	p.Nom = nom
	p.Classe = classe
	p.Niveau = niveau
	p.PvMax = pvMax
	p.PvActuels = pvActuels
	p.Inventaire = []Item{}
	p.Sort = []string{sort}
	p.Argent = Argent
	p.Equipement = make(map[string]int)
	p.AmeliorationInventaire = 0
	p.MaxSlotsInventaire = 10
	p.PointsDAttaque = PointsDAttaque
	p.MaxExp = MaxExp
	p.Level = Level
	p.Experience = Experience
	p.Mana = Mana
	p.MaxMana = MaxMana
}

// Fonction pour créer un nouveau personnage
func NewPersonnage() *Personnage {
	return &Personnage{}
}

// Méthode pour afficher les informations d'un personnage
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

// Méthode pour gérer le livre de sorts du personnage
func (joueur *Personnage) SpellBook() {
	// Parcourir tous les sorts du personnage
	for _, v := range joueur.Sort {
		// Si le sort "Boule de feu" est déjà dans le livre de sorts, sortir de la fonction
		if v == "Boule de feu" {
			return
		}
	}
	// Ajouter le sort "Boule de feu" dans le livre de sorts si ce n'est pas déjà le cas
	joueur.Sort = append(joueur.Sort, "Boule de feu")
}

// Méthode pour gérer la mort du personnage
func (joueur *Personnage) dead() {
	// Vérifier si le personnage est mort (points de vie actuels <= 0)
	if joueur.PvActuels <= 0 {
		clearScreen()
		fmt.Println("=======================================")
		fmt.Print("\033[31m") // Changer la couleur du texte en rouge
		fmt.Println("           VOUS ÊTES MORT !           ")
		fmt.Print("\033[0m") // Réinitialiser la couleur du texte
		fmt.Println("=======================================")
		fmt.Println("Da Vinci")

		// Ressusciter le personnage avec la moitié de ses points de vie max
		joueur.PvActuels = joueur.PvMax / 2
		fmt.Printf("vous avez été réanimé avec %d point de vie grâce à une RedBull.\n", joueur.PvActuels)

		// Message pour continuer le jeu
		fmt.Println("Prenez garde et continuez votre aventure!")

		// Pause pour laisser le joueur lire le message
		time.Sleep(5 * time.Second)

		// Effacer l'écran
		clearScreen()
	}
}
