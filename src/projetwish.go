package main

import (
	"fmt"
	"time"
)

type personnage struct { //creation d'une classe
	nom        string
	classe     string
	niveau     int
	pvmax      int
	pvactuel   int
	inventaire []string
	skill      []string
	money      int
	equipement equipement
}
type equipement struct {
	tête  string
	torse string
	pied  string
}

func (p *personnage) init(nom string, classe string, pvmax int, pvactuel int, niveau int, inventaire []string, skill string, money int) { //initialise des personnages
	p.nom = nom
	p.classe = classe
	p.pvmax = pvmax
	p.niveau = niveau
	p.pvactuel = pvactuel
	p.inventaire = inventaire
	p.skill = []string{"coup de point"}
	p.money = money
}
func (p *personnage) displayInfo() { // affiche les attribut des personnages
	fmt.Println("nom:", p.nom)
	fmt.Println("classe:", p.classe)
	fmt.Println("viemaximun:", p.pvmax)
	fmt.Println("PV:", p.pvactuel)
	fmt.Println("INVENTAIRE:", p.inventaire)
	fmt.Println("skill:", p.skill)
	fmt.Print("Or:", p.money)
}
func (p *personnage) popovie() { // soigne le perso
	if p.pvmax == p.pvactuel { //si le personnage a toutes ses vies il ne peut pas se soigner
		fmt.Print("tu a deja toutes tes vies")
	} else {
		if len(p.inventaire) == 0 { //si l'inventaire du personnage est vide il peut pas se soigner
			fmt.Print("ton inventaire est vide")
		} else {
			for i := 0; i < len(p.inventaire); i++ { // parcours de l'inventaire du personnage a la recherche de popo de soin
				if p.inventaire[i] == "popovie" {
					if p.pvactuel+50 > p.pvmax { // verifie si la santé du personnage ne sera pas superieur a celle maximun autorisé lors du heal
						p.pvactuel = p.pvactuel + p.pvmax - p.pvactuel
						p.removeInventory("popovie") //appelle de la fonction remove qui supprimme la popo de soin consommé
						fmt.Println(p.nom, ":", p.pvactuel, "/", p.pvmax)
						break
					} else {
						p.pvactuel = p.pvactuel + 50
						p.removeInventory("popovie")
						fmt.Println(p.nom, ":", p.pvactuel, "/", p.pvmax)
						break
					}
				} else {
					fmt.Println("Plus de popo")
				}
			}
		}
	}
}

func (p *personnage) accessInventory() { // permet d'affiche le contenu d'un inventaire
	for i := len(p.inventaire); i <= len(p.inventaire); i++ {
		if len(p.inventaire) != 0 {
			fmt.Println("Inventaire : \n", "--------------------------------")
			fmt.Println(p.inventaire)
		} else if len(p.inventaire) == 0 {
			fmt.Println("Inventaire : \n", "--------------------------------")
			fmt.Println("L'inventaire est vide ")
		}
	}
}

func (p *personnage) removeInventory(itemremove string) { //supprime un objet de l'inventaire d'un personnage
	for j := 0; j < len(p.inventaire); j++ {
		if itemremove == p.inventaire[j] {
			p.inventaire[j] = ""
			break
		}
	}
}
func (p *personnage) addInventory(itemadd string) bool {
	if len(p.inventaire) < 10 {
		p.inventaire = append(p.inventaire, itemadd) // ont ajoute dans l'inventaire du personnage un nouvelle item pour l'instant inconnue
	} else {
		fmt.Println("inventaire complet")
		return false
	}
	return true
}

/* func (p *personnage) pnj(i int) { // pnj vendeurs qui vend pas
	if i == 0 {
		p.addInventory("popovie")
	} else if i == 1 {
		p.addInventory("poison")
	} else {
		if p.addInventory("Livre de Sort: Boule de feu") == false {
			fmt.Println(" ! Plus de Place !")

		} else {
			p.removeInventory("Livre de Sort: Boule de feu")
			p.spellBook("Boule de feu")
		}
	}
} */

func (p *personnage) dead() { //verifie si le perso est mort
	if p.pvactuel < 0 {
		fmt.Println(p.nom, ": a succombé(e)")
		p.pvactuel = p.pvmax * 50 / 100
	}
}
func (p *personnage) poison() { // retire 30 hp aux personnages
	for i := 0; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(p.nom, ":", p.pvactuel)
		p.pvactuel = p.pvactuel - 10
		if p.pvactuel < 0 {
			time.Sleep(1 * time.Second)
			p.dead()
			break
		}

	}
}
func (p *personnage) spellBook(talentcaché string) { //attribue des compétemces en fonctions des livres achetés
	for i := 0; i < len(p.skill); i++ {
		if p.skill[i] == talentcaché {
			fmt.Println("_____Tu possédes déjà ce talent______")
			break
		} else {
			p.skill = append(p.skill, talentcaché)
			break
		}
	}
}
func (p *personnage) menu() {
	var commande string
	fmt.Println("☺----- Bienvenue dans Cristian's dungeon -----☺")
	fmt.Println()
	fmt.Println("↓----------------------------------------------↓")
	fmt.Println("Tapez inventaire si vous voulez voir votre inventaire")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez Information pour voir les specs(vie,talent...) de votre personnage")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez Marchand pour exercer votre Pouvoirs d'achat ")
	fmt.Println()
	fmt.Print("→")
	fmt.Scan(&commande)
	switch commande {
	case "Information":
		p.displayInfo()
	case "information":
		p.displayInfo()
	case "Inventaire":
		p.accessInventory()
	case "inventaire":
		p.accessInventory()
	case "Marchand":
		var marchand int
		fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
		fmt.Println("Produit en Vente :")
		fmt.Println("→ Potion de Soins: 3$ (Tapez 0)")
		fmt.Println("→ Potion de Poison 6$(Tapez 1)")
		fmt.Println("→ Potion de Livre de Sort Boule de feu  !25 $(Tapez 2)")
		fmt.Println("→ Fourrure de loup  4$ (Tapez 3")
		fmt.Println("→ Peau de Troll  7$ (Tapez 4")
		fmt.Println("→ Peau de Sanglier 3$ (Tapez 5")
		fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
		fmt.Scan(&marchand)
		p.pnj((marchand))
	case "marchand":
		var marchand int
		fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
		fmt.Println("Produit en Vente :")
		fmt.Println("→ Potion de Soins: 3$ (Tapez 0)")
		fmt.Println("→ Potion de Poison 6$(Tapez 1)")
		fmt.Println("→ Potion de Livre de Sort Boule de feu  !25 $(Tapez 2)")
		fmt.Println("→ Fourrure de loup  4$ (Tapez 3")
		fmt.Println("→ Peau de Troll  7$ (Tapez 4")
		fmt.Println("→ Peau de Sanglier 3$ (Tapez 5")
		fmt.Scan(&marchand)
		p.pnj((marchand))
	}
}
func (p *personnage) pnj(i int) { // pnj vendeurs qui vend pas
	if i == 0 && p.money >= 3 {
		p.addInventory("popovie")
		p.money = p.money - 3
	} else if i == 1 && p.money >= 6 {
		p.addInventory("poison")
		p.money = p.money - 6
	}
	if i == 2 && p.money >= 25 {
		if p.addInventory("Livre de Sort: Boule de feu") == true {
			p.money = p.money - 25
			p.removeInventory("Livre de Sort: Boule de feu")
			p.spellBook("Boule de feu")
		} else {
			fmt.Println("Plus de Place ☺")
		}
	}
	if i == 3 && p.money >= 4 {
		p.addInventory("Fourrure de loup")
		p.money = p.money - 4
	}
	if i == 4 && p.money >= 7 {
		p.addInventory("Peau de Troll")
		p.money = p.money - 7
	}
	if i == 5 && p.money >= 3 {
		p.addInventory("Cuir de Sanglier")
		p.money = p.money - 3
	}
	if i == 6 && p.money >= 1 {
		p.addInventory("Plume de Corbeau")
		p.money = p.money - 1
	}
}
func (p *personnage) forgeron(b int) {
	var Chapeau string = "Chapeau de l'aventurier" // crée chapeau qui contient Chapeau de l'aventurier
	var Tunique string = "Tunique de l'aventurier" // crée tunique qui contient tunique de l'aventurier
	var Bottes string = "Bottes de l'aventurier"   // crée bottes qui contient bottes de l'aventurier
	fmt.Println("Pour craft le Chapeau de l'aventurier" + " écriver Chapeau")
	fmt.Println("Pour craft la Tunique de l'aventurier" + " écriver Tunique")
	fmt.Println("Pour craft les Bottes de l'aventurier" + " écriver Bottes")
	var negative int = 0
	if p.money == negative {
		fmt.Println("vous n'avez plus d'argent sur vous!")
	}
	for i := 0; i <= len(p.inventaire); i++ {
		if b == 0 && p.money >= 5 {
			if p.inventaire[i] == "Plume de Corbeau" {
				if p.inventaire[i] == "Cuir de Sanglier" {
					p.removeInventory("Plume de Corbeau")
					p.removeInventory("Cuir de Sanglier")
					p.money = p.money - 5
					p.addInventory(Chapeau)
				}
			} else if p.money < 5 {
				fmt.Println("Vous n'avez plus assez d'agent!!")
				for c := 0; c < p.money; c++ {
					if p.money < 5 && c > 1 {
						fmt.Print("Il te reste ") // donne le nombre de po restant du joueur
						fmt.Print(c)
						fmt.Print(" pièces d'or")

					} else if p.money < 5 && c <= 1 {
						fmt.Print("Il te reste ")
						fmt.Print(c)
						fmt.Print(" pièce d'or")
					}
				}
			} else {
				fmt.Println("Il vous manque un item pour craft")
			}
		}

		//if Tunique && p.money >= 5 && p.inventaire[i+2] == "Fourrure de loup" && p.inventaire[i] == "Peau de Troll" {
		if b == 1 && p.money >= 5 {
			if p.inventaire[i] == "Fourrure de loup"+"Fourrure de loup" {
				if p.inventaire[i] == "Peau de Troll" {
					p.removeInventory("Fourrure de loup" + "Fourrure de loup")
					p.removeInventory("Peau de Troll")
					p.money = p.money - 5
					p.addInventory(Tunique)
				}
			} else if p.money < 5 {
				fmt.Println("Vous n'avez plus assez d'agent!!")
				for c := 0; c < p.money; c++ {
					if p.money < 5 && c > 1 {
						fmt.Print("Il te reste ") // donne le nombre de po restant du joueur
						fmt.Print(c)
						fmt.Print(" pièces d'or")

					} else if p.money < 5 && c <= 1 {
						fmt.Print("Il te reste ")
						fmt.Print(c)
						fmt.Print(" pièce d'or")
					}
				}
			} else {
				fmt.Println("Il vous manque un item pour craft")
			}
		}
		if b == 2 && p.money >= 5 {
			if p.inventaire[i] == "Fourrure de loup" {
				if p.inventaire[i] == "Peau de Troll" {
					p.removeInventory("Fourrure de loup")
					p.removeInventory("Cuir de Sanglier")
					p.money = p.money - 5
					p.addInventory(Bottes)
				}
			} else if p.money < 5 {
				fmt.Println("Vous n'avez plus assez d'agent!!")
				for c := 0; c < p.money; c++ {
					if p.money < 5 && c > 1 {
						fmt.Print("Il te reste ") // donne le nombre de po restant du joueur
						fmt.Print(c)
						fmt.Print(" pièces d'or")

					} else if p.money < 5 && c <= 1 {
						fmt.Print("Il te reste ")
						fmt.Print(c)
						fmt.Print(" pièce d'or")
					}
				}
			} else {
				fmt.Println("Il vous manque un item pour craft")
			}
		}
	}
}
func main() {
	var p1 personnage
	p1.init("jackouille", "fripouille", 150, 10, 1, []string{"popovie", "poison", "popovie", "popovie", "popovie", "popovie", "popovie", "popovie"}, "coup de point", 100)
	var p2 personnage
	p2.init("Cristian ", "Cristian", 150, 1, 1, []string{"poison", "poison", "poison", "poison", "poison", "poison", "poison", "popovie"}, "coup de point", 100)
	//p2.displayInfo()
	//p2.poison()
	fmt.Println()

	p1.menu()

}
