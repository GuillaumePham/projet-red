package main

import "fmt"

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

func main() {
	var p1 personnage
	p1.init("jackouille", "fripouille", 12, 53, []string{"kdjrb", "hjbdbhehu", "popovie"})
	fmt.Println(p1)
}

func (p *personnage) displayInfo() {
	fmt.Print("nom:", p.nom)
	fmt.Print("nom:", p.classe)
	fmt.Print("nom:", p.pvmax)
	fmt.Print("nom:", p.pvactuel)
	fmt.Print("nom:", p.inventaire)
}
func (p *personnage) popovie() {
	if p.pvmax == p.pvactuel {
		fmt.Print("tu a deja toutes tes vies")
	} else {
		for i := 1; i < len(p.inventaire); i++ {
			if len(p.inventaire) == 0 {
				fmt.Print("ton inventaire est vide")
			} else {
				if p.inventaire[i] == "popovie" {
					if p.pvactuel+2 > p.pvmax {
						p.pvactuel = p.pvactuel + 1
					} else {
						p.pvactuel = p.pvactuel + 2
					}
				}
			}
		}
	}
}
