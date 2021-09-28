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
	monstre    monstre
}
type equipement struct {
	price      int
	niveau_min int
	tête       string
	torse      string
	pied       string
}
type monstre struct {
	nom             string
	pvmax           int
	attaque         int
	pvmonstreactuel int
}

func (p *personnage) init(nom string, classe string, pvmax int, pvactuel int, niveau int, inventaire []string, skill string) { //initialise des personnages
	p.nom = nom
	p.classe = classe
	p.pvmax = pvmax
	p.niveau = niveau
	p.pvactuel = pvactuel
	p.inventaire = inventaire
	p.skill = []string{"coup de point"}
}
func (b *equipement) init(price int, niveau_min int, tête string, torse string, pied string) {
	b.price = price
	b.niveau_min = niveau_min
	b.tête = tête
	b.torse = torse
	b.pied = pied
}
func (m *monstre) initmonstre(nom string, pvmax int, pvmonstreactuel int, attaque int) {
	m.nom = nom
	m.attaque = attaque
	m.pvmonstreactuel = pvmonstreactuel
	m.pvmax = pvmax
}
func (p *personnage) equiper(d int) {
	if d == 0 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Chapeau de l'aventurier" && p.equipement.tête == "" && p.niveau >= 9 { // si le personnage a dans l'inventaire Chapau de l'aventurier et a le niveau 9 minimum
				p.equipement.tête = "Chapeau de l'aventurier" // ont ajoute chapeau de l'aventurier a tete
				p.removeInventory("Chapeau de l'aventurier")  // et ont remove de l'inventaire chapeau de l'aventurier
				time.Sleep(2)
				p.pvmax = p.pvmax + 10
				fmt.Println("Chapeau de l'aventurier équiper")
				break
			} else {
				if !(p.equipement.tête == "") { // si tete n'est pas vide
					p.addInventory(p.equipement.tête) // ont ajoute l'item de tete dans inventaire
					p.equipement.tête = ""            // retire l'équipement de tete
				}
			}
		}
	}
	if d == 1 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Tunique de l'aventurier" && p.equipement.torse == "" && p.niveau >= 11 { // si le personnage a dans l'inventaire Chapau de l'aventurier et a le niveau 9 minimum
				p.equipement.torse = "Tunique de l'aventurier" // ont ajoute chapeau de l'aventurier a tete
				p.removeInventory("Tunique de l'aventurier")   // et ont remove de l'inventaire chapeau de l'aventurier
				time.Sleep(2)
				p.pvmax = p.pvmax + 10
				fmt.Println("Tunique de l'aventurier équiper")
				break
			} else {
				if !(p.equipement.torse == "") { // si tete n'est pas vide
					p.addInventory(p.equipement.torse) // ont ajoute l'item de tete dans inventaire
					p.equipement.torse = ""            // retire l'équipement de tete
				}
			}
		}
	}
	if d == 2 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Bottes de l'aventurier" && p.equipement.pied == "" && p.niveau >= 10 { // si le personnage a dans l'inventaire Chapau de l'aventurier et a le niveau 9 minimum
				p.equipement.pied = "Bottes de l'aventurier" // ont ajoute chapeau de l'aventurier a tete
				p.removeInventory("Bottes de l'aventurier")  // et ont remove de l'inventaire chapeau de l'aventurier
				time.Sleep(2)
				p.pvmax = p.pvmax + 10
				fmt.Println("Bottes de l'aventurier équiper")
				break
			} else {
				if !(p.equipement.pied == "") { // si tete n'est pas vide
					p.addInventory(p.equipement.pied) // ont ajoute l'item de tete dans inventaire
					p.equipement.pied = ""            // retire l'équipement de tete
				}
			}
		}
	}
}

/*func (p *personnage) ItemEffect() {
	for i := 0; i <= len(p.equipement.tête); i++ {
		if p.equipement.tête[i] == "Chapeau de l'aventurier" { // s'il y'a chapeau de l'aventurier dans tete ajoute 10pv max au personnage
			p.pvmax = p.pvmax + 10

		}
	}
	for i := 0; i <= len(p.equipement.torse); i++ {
		if p.equipement.torse[i] == "Tunique de l'aventurier" { // s'il y'a chapeau de l'aventurier dans tete ajoute 25pv max au personnage
			p.pvmax = p.pvmax + 25

		}
	}
	for i := 0; i <= len(p.equipement.pied); i++ {
		if p.equipement.pied[i] == "Bottes de l'aventurier" { // s'il y'a chapeau de l'aventurier dans tete ajoute 15pv max au personnage
			p.pvmax = p.pvmax + 15
		}
	}
}*/
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

/*func Color(){ // la couleur apparait que dans le terminale site de référence : https://newbedev.com/go-golang-color-print-code-example
	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	fmt.Println(string(colorRed), "test")
	fmt.Println(string(colorGreen), "test")
	fmt.Println(string(colorYellow), "test")
	fmt.Println(string(colorBlue), "test")
	fmt.Println(string(colorPurple), "test")
	fmt.Println(string(colorWhite), "test")
	fmt.Println(string(colorCyan), "test", string(colorReset))
	fmt.Println("next")
}*/
func (p *personnage) displayInfo() { // affiche les attribut des personnages
	fmt.Println("nom:", p.nom)
	fmt.Println("classe:", p.classe)
	fmt.Println("viemaximun:", p.pvmax)
	fmt.Println("PV:", p.pvactuel)
	fmt.Println("INVENTAIRE:", p.inventaire)
	fmt.Println("skill:", p.skill)
	fmt.Println("Argent :", p.money)
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
			fmt.Println("--------------------------------", "\nInventaire : \n")
			fmt.Println(p.inventaire)
		} else if len(p.inventaire) == 0 {
			fmt.Println("--------------------------------", "\nInventaire : \n")
			fmt.Println("L'inventaire est vide ")
		}
	}
}

func (p *personnage) removeInventory(itemremove string) { //supprime un objet de l'inventaire d'un personnage
	for i := 0; i <= len(p.inventaire); i++ {
		if itemremove == p.inventaire[i] {
			p.inventaire[i] = ""
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

/*func (p *personnage) addInventory(itemadd string) {
	var max_inventaire = [10]string{}             // ont fait un string qui comporte maximum 10 valeur a l'intérieur
	if len(p.inventaire) == len(max_inventaire) { // si le nombre d'ojet dans p.inventaire vaut plus que 10 ont renvoie l'inventaire est plein
		fmt.Println("inventaire plein")
	} else {
		if len(p.inventaire) <= len(max_inventaire) { // sinon si p.inventaire est inférieur ou égale a 10
			p.inventaire = append(p.inventaire, itemadd) // ont ajoute dans l'inventaire du personnage un nouvelle item pour l'istant inconnue
			fmt.Println(p.inventaire)                    // ont print l'inventaire
		}
	}
} */
func (p *personnage) Compteuritems(item_into_inventory string) { //juste du bonus  pour dire combien d'item x tu as dans ton inventaire si demmander
	item_into_inventory = ""
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == item_into_inventory {
			fmt.Print("tu as")
			fmt.Print(i)
			fmt.Println(item_into_inventory)
		}
	}
}
func (p *personnage) attack() {
	p.monstre.pvmonstreactuel = p.monstre.pvmonstreactuel - 10
	fmt.Println(p.monstre.pvmonstreactuel)
	time.Sleep(2 * time.Second)
	fmt.Println(" Cristrian subit une attaque de : ", p.nom)
}
func (p *personnage) menucombat() {
	var combat string
	fmt.Println("")
	fmt.Println("Bienvenue Dans L'arène")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez: attaque pour Attaquer ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez: Soin pour vous soigner (cela fera passer votre tour)")
	fmt.Println("---------------------------------------------------")
	fmt.Print("→")
	fmt.Scan(&combat)

	switch combat {
	case "Attaque":
		p.attack()
		time.Sleep(3 * time.Second)
		p.goblinPattern()
		p.menu()

	case "attaque":
		p.attack()
		time.Sleep(3 * time.Second)
		p.goblinPattern()
		p.menu()
	case "Soin":
		p.popovie()
		p.menu()
	case "soin":
		p.popovie()
		p.menu()

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
	fmt.Println("Tapez Marchand ou Forgeron pour exercer votre Pouvoirs d'achat ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez Attaque pour rentrer dans l'aréne")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Tapez exit  puis 0 pour rage quit")
	fmt.Println()
	fmt.Print("→")
	fmt.Scan(&commande)
	switch commande {
	case "Information":
		p.displayInfo()
		p.menu()
	case "information":
		p.displayInfo()
		p.menu()
	case "Inventaire":
		p.accessInventory()
		p.menu()
	case "inventaire":
		p.accessInventory()
		p.menu()
	case "équiper":
		var equiper int
		fmt.Println("→ Bienvene dans l'interface d'équippement")
		fmt.Println("→ Si vous souhaitez mettre un Chapeau (Tapez 0)")
		fmt.Println("→ Si vous souhaitez mettre un Plastron (Tapez 1)")
		fmt.Println("→ Si vous souhaitez mettre des Chaussures (Tapez 2)")
		fmt.Scan(&equiper)
		p.equiper(equiper)
		p.menu()
	case "UpgradeInventorySlot":
		var upgrade int
		fmt.Scan(&upgrade)
		p.upgradeInventorySlot((upgrade))
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
		fmt.Println("→ Peau de Sanglier 3$ (Tapez 5")
		fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
		fmt.Println("→ Augmentation d'inventaire 30 $ (Tapez 7) ")
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
		fmt.Println("→ Peau de Sanglier 3$ (Tapez 5")
		fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
		fmt.Println("→ Augmentation d'inventaire 30 $ (Tapez 7) ")
		fmt.Scan(&marchand)
		p.pnj((marchand))
		p.menu()
	case "Attaque":
		p.menucombat()
	case "attaque":
		p.menucombat()
	case "forgeron":
		var forgeron int
		fmt.Println("Pour craft le Chapeau de l'aventurier" + " écriver 0")
		fmt.Println("Pour craft la Tunique de l'aventurier" + " écriver 1")
		fmt.Println("Pour craft les Bottes de l'aventurier" + " écriver 2")
		fmt.Scan(&forgeron)
		p.forgeron((forgeron))
		p.menu()
	case "Forgeron":
		var forgeron int
		fmt.Println("Pour craft le Chapeau de l'aventurier" + " écriver 0")
		fmt.Println("Pour craft la Tunique de l'aventurier" + " écriver 1")
		fmt.Println("Pour craft les Bottes de l'aventurier" + " écriver 2")
		fmt.Scan(&forgeron)
		p.forgeron((forgeron))
		p.menu()
	case "Exit":
		var exit string
		fmt.Scan(&exit)
	}
	p.menu()
}
func (p *personnage) pnj(i int) { // pnj vendeurs qui vend pas
	if i == 0 && p.money >= 3 {
		if p.addInventory("popovie") == true {
			p.money = p.money - 3
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
	if i == 7 && p.money >= 30 {
		if p.addInventory("Augmentation d'inventaire") == true {
			p.money = p.money - 30
		} else {
			fmt.Println("Plus de Place ☺")
		}
	}
}
func (p *personnage) dead() { //verifie si le perso est mort
	if p.pvactuel < 0 {
		fmt.Println(p.nom, ": a succombé(e)")
		p.pvactuel = p.pvmax * 50 / 100
	}
}
func (p *personnage) charCreation(test string, s rune) {
	var result = []byte{}
	for i := 0; i >= 'A' && i <= 'Z'; i++ { // i commence a zéro et fera le tours de A a Z seulement ne dépassera pas
		if !(p.nom[0] == test[i]) { // si la premier caractère de p.nom n'appartient pas a test[i]qui comprend toutes les majuscules
			fmt.Println("il faut une majuscule en première lettre") // revoyez un message de prévention
		} else {
			result = append(result, p.nom[0]) // sinon on ajoute la première valeur de p.nom qui est une majuscule dans result qui était vide jusqu'a maintenant
		}
	}
	for compteur := 0; compteur < len(p.nom); compteur++ { // je crée un compteur qui part de zéro et ajoute plus 1 temp que compteur n'est pas plus grand que p.nom
		for b := 0; b >= 'a' && b <= 'z'; b++ { // b commence a zéro et fera le tours de a a z seulement ne dépassera pas
			if !(p.nom[1+compteur] == test[b]) { // si lettre n'est pas comprit entre A et Z renvoie "il faut une majuscule en première lettre"
				fmt.Println("seulement la première lettre doit etre en majuscule, il faut une minuscule")
			} else {
				result = append(result, p.nom[1+compteur])
			}
		}
	}
	var Classe2 = [4]string{"Humain", "Elfe", "Nain", "Cristian"}
	for i := 0; i < len(Classe2); i++ {
	}
}
func (p *personnage) spellBook(talentcaché string) { //attribue des compétemces en fonctions des livres achetés
	for i := 0; i < len(p.skill); i++ {
		if p.skill[i] == talentcaché {
			fmt.Println("Tu possédes déjà ce talent")
			break
		} else {
			p.skill = append(p.skill, talentcaché)
			break
		}
	}
}
func (p *personnage) forgeron(b int) {
	var Chapeau string = "Chapeau de l'aventurier" // crée chapeau qui contient Chapeau de l'aventurier
	var Tunique string = "Tunique de l'aventurier" // crée tunique qui contient tunique de l'aventurier
	var Bottes string = "Bottes de l'aventurier"   // crée bottes qui contient bottes de l'aventurier
	fmt.Println("Pour craft le Chapeau de l'aventurier" + " écriver 0")
	fmt.Println("Pour craft la Tunique de l'aventurier" + " écriver 1")
	fmt.Println("Pour craft les Bottes de l'aventurier" + " écriver 2")
	fmt.Print("→")
	if p.money == 0 {
		fmt.Println()
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
				break
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
				break
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
				break
			}
		}
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
func (p *personnage) goblinPattern() {
	fmt.Println("Rien")
}

/* var m1 monstre
	for i := 0; i < len(tours); i++ { // comme tout a l'heure pour i partant de zéro est inférieur a len de tours qui est  infinie, j'avance
		for n := 2; i == n; n += 2 { // a chaque fois que i atteint la valeur de n ont ajoute 2
			if i != tours[2] || i != tours[n] { // si i est différent de tours[2], soit le troixième tours étant donner que tours[0] == 1, la ont part de zéro et les tours dans un jeu eux commence a 1, soit le premier tour pas de tour zéro
				fmt.Println(tours[i+1])
				p.menucombat()
				fmt.Print(m1.pvmonstreactuel)
				fmt.Print("/")
				fmt.Print(p.pvmax)
				fmt.Println(" PDV")
				time.Sleep(2)
				fmt.Print(m1.nom, " attaque ", p.nom, "et lui reste ")
				m1.attaque = 5 // cette fois ci c'est le monstre qui attaque
				p.pvactuel = p.pvactuel - m1.attaque
				fmt.Print(p.pvactuel)
				fmt.Print("/")
				fmt.Print(m1.pvmonstreactuel)
				fmt.Println(" PDV")
				if i == 0 {
					fmt.Print("Fin du tour ")
					fmt.Println(tours[i+1]) // i +1 car tour part de zéro, pour dire fin du tour 1
				} else if i >= 1 {
					fmt.Print("Fin du tour ")
					fmt.Println(tours[i+1])
					break
				}
			} else {
				if i == tours[2] || i == tours[n] {
					fmt.Println(tours[i+1])
					p.menucombat()
					fmt.Print(m1.pvmonstreactuel)
					fmt.Print("/")
					fmt.Print(p.pvmax)
					fmt.Println(" PDV")
					time.Sleep(2)
					fmt.Print(m1.nom, " attaque ", p.nom, "et lui reste ")
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
} */

/* func (p *personnage) charTurn(n int) {
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
						fmt.Println("Pour utiliser un poison Tapez 3")
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
		if p.pvactuel <= 0 {
			p.dead()
		} else {
			if p.monstre.pvmonstreactuel == 0 {
				fmt.Println(p.monstre.nom, " a perdue contre ", p.nom)
				fmt.Println("Fin du combat !")
				break
			}
		}
	}
	p.menu()
} */

func main() {
	var p1 personnage
	p1.init("jackouille", "fripouille", 150, 10, 1, []string{"popovie", "poison", "popovie", "popovie", "popovie", "popovie", "popovie", "popovie"}, "coup de point")
	p1.menu()
	var m1 monstre
	m1.initmonstre("Cristian", 40, 40, 5)

}
