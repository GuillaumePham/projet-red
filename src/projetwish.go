package main

import (
	"fmt"
	"time"
)

type monstre struct {
	nom             string
	pvmax           int
	pvmonstre       int
	pvmonstreactuel int
	attaque         int
}

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
	monstre    monstre
}
type equipement struct {
	price      int
	niveau_min int
	tête       []string
	torse      []string
	pied       []string
}

func (p *personnage) init(nom string, classe string, pvmax int, pvactuel int, niveau int, inventaire []string, skill string, money int, price int, niveau_min int, tête []string, torse []string, pied []string, pvmonstre int) { //initialise des personnages
	p.nom = nom
	p.classe = classe
	p.pvmax = pvmax
	p.niveau = niveau
	p.pvactuel = pvactuel
	p.inventaire = inventaire
	p.skill = []string{"coup de point"}
	p.money = money
	p.equipement.price = price
	p.equipement.niveau_min = niveau_min
	p.equipement.tête = tête
	p.equipement.torse = torse
	p.equipement.pied = pied
	p.monstre.pvmonstre = pvmonstre
}
func (m *monstre) initmonstre(nom string, pvmax int, attaque int, pvmonstre int) {
	m.nom = nom
	m.attaque = attaque
	m.pvmonstre = pvmonstre
	m.pvmax = pvmax
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
						p.pvactuel = p.pvactuel + 5
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
	fmt.Println()
	fmt.Println("☺----- Bienvenue dans Cristian's dungeon -----☺")
	fmt.Println()
	fmt.Println("↓----------------------------------------------↓")
	fmt.Println("Tapez inventaire si vous voulez voir votre inventaire")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez Information pour voir les specs(vie,talent...) de votre personnage")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez Marchand pour exercer votre Pouvoirs d'achat ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez Attaque Pour entrez dans l'arène")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez exit pour rage quit")
	fmt.Println()
	fmt.Print("→")
	fmt.Scan(&commande)
	switch commande {
	case "Information":
		p.displayInfo()
		time.Sleep(3 * time.Second)
		p.menu()
	case "information":
		p.displayInfo()
		time.Sleep(3 * time.Second)
		p.menu()
	case "Inventaire":
		p.accessInventory()
		time.Sleep(3 * time.Second)
		p.menu()
	case "inventaire":
		p.accessInventory()
		time.Sleep(3 * time.Second)
		p.menu()
	case "Marchand":
		var marchand int
		fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
		fmt.Println("Produit en Vente :")
		fmt.Println("→ Potion de Soins: 3$ (Tapez 0)")
		fmt.Println("→ Potion de Poison 6$(Tapez 1)")
		fmt.Println("→ Potion de Livre de Sort Boule de feu  !25 $(Tapez 2)")
		fmt.Println("→ Fourrure de loup  4$ (Tapez 3")
		fmt.Println("→ Peau de Troll  7$ (Tapez 4")
		fmt.Println("→ Peau de Sanglier 3$ (Tapez 5)")
		fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
		fmt.Println("→")
		fmt.Scan(&marchand)
		p.pnj((marchand))
		p.menu()
	case "marchand":
		var marchand int
		fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
		fmt.Println("Produit en Vente :")
		fmt.Println("→ Potion de Soins: 3$ (Tapez 0)")
		fmt.Println("→ Potion de Poison 6$(Tapez 1)")
		fmt.Println("→ Potion de Livre de Sort Boule de feu  !25 $(Tapez 2)")
		fmt.Println("→ Fourrure de loup  4$ (Tapez 3")
		fmt.Println("→ Peau de Troll  7$ (Tapez 4")
		fmt.Println("→ Peau de Sanglier 3$ (Tapez 5)")
		fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
		fmt.Println("→")
		fmt.Scan(&marchand)
		p.pnj((marchand))
		p.menu()
	case "Attaque":
		p.menucombat()
	case "attaque":
		p.menucombat()
	}

}
func (p *personnage) pnj(i int) { // pnj vendeurs qui vend pas
	if i == 0 && p.money >= 3 {
		if p.addInventory("popovie") == true { //ajoute l'objet si l'inventaire n'est pas plein
			p.money = p.money - 3
			p.displayInfo()
		} else {
			fmt.Println("Plus de Place ☺")
		}
	}
	if i == 1 && p.money >= 6 {
		if p.addInventory("poison") == true {
			p.money = p.money - 6
		} else {
			fmt.Println("Plus de Place ☺")
		}

	}
	if i == 2 && p.money >= 25 {
		if p.addInventory("Livre de Sort: Boule de feu") == true {
			p.money = p.money - 25
			p.removeInventory("Livre de Sort: Boule de feu")
			fmt.Println("Merci de votre Achat")
			p.spellBook("Boule de feu")
		} else {
			fmt.Println("Plus de Place ☺")
		}
	}
	if i == 3 && p.money >= 4 {
		if p.addInventory("Fourrure de loup") == true {
			fmt.Println("Merci de votre Achat")
			p.money = p.money - 4
		} else {
			fmt.Println("Plus de Place ☺")

		}
	}
	if i == 4 && p.money >= 7 {
		if p.addInventory("Peau de Troll") == true {
			p.money = p.money - 7
		} else {
			fmt.Println("Plus de Place ☺")

		}
	}
	if i == 5 && p.money >= 3 {
		if p.addInventory("Peau de Troll") == true {
			p.money = p.money - 7
		} else {
			fmt.Println("Plus de Place ☺")
		}
	}
	if i == 6 && p.money >= 1 {
		if p.addInventory("Plume de Corbeau") == true {
			p.money = p.money - 1
		} else {
			fmt.Println("Plus de Place ☺")
		}
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
						fmt.Print("Il te reste ") // donne le nombre d'or restant du joueur
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
func (p *personnage) menucombat() {
	var combat string
	fmt.Println("")
	fmt.Println("Bienvenue Dans L'arène")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez: attaque pour Attaquer (inflige des dégat aléatoire")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez: Soin pour vous soigner (cela fera passer votre tour)")
	fmt.Println("---------------------------------------------------")
	fmt.Print("→")
	fmt.Scan(&combat)

	switch combat {
	case "Attaque":
		p.attack()
		time.Sleep(3 * time.Second)
		p.menu()
	case "attaque":
		p.attack()
		time.Sleep(3 * time.Second)
		p.menu()
	case "Soin":
		p.popovie()
		time.Sleep(5 * time.Second)
		p.menu()
	case "soin":
		p.popovie()
		time.Sleep(5 * time.Second)
		p.menu()

	}
}
func (p *personnage) charTurn(n int) {
	fmt.Println("Attaquer :", " Pour Attaquer Tapez 0")
	fmt.Println("Inventaire :", " Pour ouvire l'Iventaire Tapez 1")
	for i := 0; i < len(p.skill); i++ {
		if p.skill[i] == "coup de point" {
			if n == 0 {
				p.goblinPattern()
			} else {
				for w := 0; w < len(p.inventaire); w++ {
					if n == 1 {
						p.accessInventory()
						// faire en sorte que le joueur utilise un objet et passe sont tours
						fmt.Println("Pour utiliser une popvie Tapez 2")
						fmt.Println("Pour utiliser un poisson Tapez 3")
						if n == 2 && p.inventaire[w] == "popvie" {
							p.popovie()
							time.Sleep(2)
							fmt.Print(p.monstre.nom, " attack ", p.nom, "et lui reste ")
							p.monstre.attaque = 5
							p.pvactuel = p.pvactuel - p.monstre.attaque
							fmt.Print(p.pvactuel)
							fmt.Print("/")
							fmt.Print(p.pvmax)
							fmt.Println(" PDV")
							fmt.Print("Fin du tour")
							break
						}
						if n == 3 && p.inventaire[w] == "poison" {
							p.poison()
							time.Sleep(2)
							fmt.Print(p.monstre.nom, " attack ", p.nom, "et lui reste ")
							p.monstre.attaque = 5
							p.pvactuel = p.pvactuel - p.monstre.attaque
							fmt.Print(p.pvactuel)
							fmt.Print("/")
							fmt.Print(p.pvmax)
							fmt.Println(" PDV")
							fmt.Print("Fin du tour")
							break
						}
					}
				}
			}
		}
	}
}
func (p *personnage) attack() {
	p.monstre.pvmonstre = p.monstre.pvmonstre - 10
	fmt.Println(p.monstre.pvmonstre)
	time.Sleep(2 * time.Second)
	fmt.Println(" Cristrian subit une attaque de : ", p.nom)
}
func (p *personnage) goblinPattern() {
	var tours = []int{}                      // ont re crée les tours comme avant cela reste les meme
	p.monstre.nom = "Gobelin d'entrainement" // je nome mon monstre
	p.monstre.pvmax = 40                     // je lui donne comme pvmax 40pdv
	p.monstre.attaque = 5
	p.monstre.pvmonstreactuel = p.monstre.pvmax // et je dit que par default le monstre comme nce avec c'est pv max
	for i := 0; i < len(tours); i++ {           // comme tout a l'heure pour i partant de zéro est inférieur a len de tours qui est  infinie, j'avance
		for n := 2; i == n; n += 2 { // a chaque fois que i atteint la valeur de n ont ajoute 2
			if i != tours[2] || i != tours[n] { // si i est différent de tours[2], soit le troixième tours étant donner que tours[0] == 1, la ont part de zéro et les tours dans un jeu eux commence a 1, soit le premier tour pas de tour zéro
				fmt.Println(tours[i+1])                                      // ont donne le tours
				fmt.Print(p.nom, " attack ", p.monstre.nom, "et lui reste ") // le nom du joueurs qui attack le nom du monstre
				p.attack()                                                   // attack du joueur fait sur le monstre
				fmt.Print(p.monstre.pvmonstreactuel)                         // et ont print les pv du monstre actuel
				fmt.Print("/")
				fmt.Print(p.pvmax)
				fmt.Println(" PDV")
				fmt.Print(p.monstre.nom, " attack ", p.nom, "et lui reste ")
				p.monstre.attaque = 5 // cette foit ci c'est le montre qui attcak
				p.pvactuel = p.pvactuel - p.monstre.attaque
				fmt.Print(p.pvactuel)
				fmt.Print("/")
				fmt.Print(p.monstre.pvmax)
				fmt.Println(" PDV")
				if i == 0 {
					fmt.Print("Fin du tour ")
					fmt.Println(tours[i+1]) // i +1 car tours part de zéro, pour dire fin du tour 1
				} else if i >= 1 {
					fmt.Print("Fin du tour ")
					fmt.Println(tours[i+1])
					break
				}
			} else {
				if i == tours[2] || i == tours[n] {
					fmt.Println(tours[i+1])
					fmt.Print(p.nom, " attack ", p.monstre.nom, "et lui reste ")
					p.attack()
					fmt.Print(p.monstre.pvmonstreactuel)
					fmt.Print("/")
					fmt.Print(p.pvmax)
					fmt.Println(" PDV")
					fmt.Print(p.monstre.nom, " attack ", p.nom, "et lui reste ")
					p.monstre.attaque *= 2
					p.pvactuel = p.pvactuel - p.monstre.attaque
					fmt.Print(p.pvactuel)
					fmt.Print("/")
					fmt.Print(p.monstre.pvmax)
					fmt.Println(" PDV")
					if i == 0 {
						fmt.Print("Fin du tour ")
						fmt.Println(tours[i+1])
					} else {
						fmt.Print("Fin du tours ")
						fmt.Println(tours[i+1])
						break
					}
				}
			}
		}
	}
}

func (p *personnage) upgradeInventorySlot(r int) {
	var upgrade = [10]string{}
	var z int
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == "Augmentation d'inventaire" {
			if len(p.inventaire) < 40 { // si l'inventaire n'est pas égale a 40
				if r == 0 { // et qu'il renvoie 0
					p.inventaire = append(p.inventaire, upgrade[z]) // et on ajoute a l'inventaire une upgrade
					p.removeInventory("Augmentation d'inventaire")
				}
			}
		} else if len(p.inventaire) >= 40 {
			fmt.Println("Vous ne pouvez plus agrandire votre inventaire")
		}
	}
	if len(p.inventaire) > 10 && len(p.inventaire) <= 20 {
		fmt.Println("Tu peut encore utilisé 2 Augmentation d'inventaire")
	}
	if len(p.inventaire) > 20 && len(p.inventaire) <= 30 {
		fmt.Println("Tu peut encore utilisé 1 Augmentation d'inventaire")
	}
	if len(p.inventaire) > 30 && len(p.inventaire) <= 40 {
		fmt.Println("Tu ne peut plus utiliser d'Augmentation d'inventaire")
	}
}
func main() {
	var p1 personnage
	p1.init("jackouille", "fripouille", 150, 10, 1, []string{"popovie", "poison", "popovie", "popovie", "popovie", "popovie", "popovie", "popovie"}, "coup de point", 100, 100, 10, []string{"vide"}, []string{"vide"}, []string{"vide"}, 100)
	var m1 monstre
	m1.initmonstre("Cristian", 100, 150, 5)
	//var p2 personnage
	//p2.init("Cristian ", "Cristian", 150, 1000, 1, []string{"poison", "poison", "poison", "poison", "poison", "poison", "poison", "popovie"}, "coup de point", 100)
	//p2.displayInfo()
	//p2.poison()
	fmt.Println()

	p1.menu()

}
