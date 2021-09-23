package main

import (
	"fmt"
)

type personnage struct {
	nom        string
	classe     string
	niveau     int
	pvmax      int
	pvactuel   int
	inventaire []string
}

func (p *personnage) init(nom string, classe string, pvmax int, pvactuel int, inventaire []string) {
	p.nom = nom
	p.classe = classe
	p.pvmax = pvmax
	p.pvactuel = pvactuel
	p.inventaire = inventaire
}
func (p *personnage) displayInfo() {
	fmt.Println("nom:", p.nom)
	fmt.Println("classe:", p.classe)
	fmt.Println("viemaximun:", p.pvmax)
	fmt.Println("PV:", p.pvactuel)
	fmt.Println("INVENTAIRE:", p.inventaire)
}
func (p *personnage) popovie() {
	if p.pvmax == p.pvactuel {
		fmt.Print("tu a deja toutes tes vies")
	} else {
		if len(p.inventaire) == 0 {
			fmt.Print("ton inventaire est vide")
		} else {
			for i := 0; i < len(p.inventaire); i++ {
				if p.inventaire[i] == "popovie" {
					if p.pvactuel+50 > p.pvmax {
						p.pvactuel = p.pvactuel + p.pvmax - p.pvactuel
						heal := 0
						for i := 0; i < len(p.inventaire); i++ {
							if p.inventaire[i] == "popovie" {
								heal++
							}
						}
						removeInventory()
						/* p.inventaire = []string{}
						for i := 0; len(p.inventaire) < heal-1; i++ {
							p.inventaire = append(p.inventaire, "popovie") */
						//}

					} else {
						p.pvactuel = p.pvactuel + 50
						heal := 0
						for i := 0; i < len(p.inventaire); i++ {
							if p.inventaire[i] == "popovie" {
								heal++
							}
						}
						/* p.inventaire = []string{}
						for i := 0; len(p.inventaire) < heal-1; i++ {
							p.inventaire = append(p.inventaire, "popovie") */
						//}
					}
					break
				} /* else {
					fmt.Println("Plus de popo")
				} */
			}
		}
	}
}
func (p *personnage) accessInventory() {
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

func (p *personnage) removeInventory(test string) {
	var itemremove string = ""
	for i := 0; i <= len(p.inventaire); i++ {
		if itemremove == p.inventaire[i] {
			itemremove = test[i-1:i] + test[i+1:i]
			p.inventaire = append(p.inventaire, itemremove)
			fmt.Print(p.inventaire)
		}
	}
}
func main() {
	var p1 personnage
	p1.init("jackouille", "fripouille", 150, 120, []string{"popovie", "popovie", "popovie"})
	p1.popovie()
	p1.displayInfo()
	fmt.Println()
}