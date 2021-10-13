package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type personnage struct { //creation de la classe personnnage
	nom        string
	classe     string
	niveau     int
	pvmax      int
	pvactuel   int
	inventaire []string
	skill      []string
	money      int
	initiative int
	manamax    int
	manaactuel int
	experience int
	equipement equipement
	monstre    monstre
	invocation invocation
}
type equipement struct { //classe equipement
	tête  string
	torse string
	pied  string
}
type monstre struct { //classe monstre
	nom             string
	pvmax           int
	attaque         int
	pvmonstreactuel int
	initiative2     int
	mana            int
}
type invocation struct {// classe invocation pour le perso nain
	attaque int
	nom     string
}

func (p *personnage) init(nom string, classe string, pvmax int, pvactuel int, niveau int, inventaire []string, skill string, money int, tête string, torse string, pied string, initiative int, manamax int, manaactuel int, experience int) { //initialise des personnages
	p.nom = nom
	p.classe = classe
	p.pvmax = pvmax
	p.niveau = niveau
	p.pvactuel = pvactuel
	p.inventaire = inventaire
	p.skill = []string{"coup de poing"}
	p.money = money
	p.equipement.tête = tête
	p.equipement.torse = torse
	p.equipement.pied = pied
	p.initiative = initiative
	p.manamax = manamax
	p.manaactuel = manaactuel
	p.experience = experience
}
func (b *equipement) init(tête string, torse string, pied string) { //init de l'équipement
	b.tête = tête
	b.torse = torse
	b.pied = pied
}
func (m *monstre) initmonstre(nom string, pvmax int, pvmonstreactuel int, attaque int, initiative2 int, mana int) {//init des monstres
	m.nom = nom
	m.attaque = attaque
	m.pvmonstreactuel = pvmonstreactuel
	m.pvmax = pvmax
	m.initiative2 = initiative2
	m.mana = mana
}

var m1 monstre
var tour int // compteur de tour
var comptetour int// compteur de tour spéciaux pour les attaques *2
func (p *personnage) TrainingFight(m *monstre) { // Combat d'entrainement contre Cristian un gobelin
	if p.pvactuel <= 0 || m.pvmonstreactuel <= 0 {
		tour++
	} else {
		if m.Mdead() == false {
			if comptetour%3 == 2 {//verifie si on est bien tous les deux tours
				m.Mdead()
				if m.Mdead() == false { // si monstre mort
					m.attaque = 5 * 2
					if m.pvmonstreactuel > 0 {
						p.pvactuel = p.pvactuel - m.attaque
						fmt.Println(p.nom, "subit une attaque de", m.nom)
						time.Sleep(3 * time.Second)
						fmt.Println(p.pvactuel, "/", p.pvmax)
						comptetour = 0
						m.Mdead()
					}
				}
			} else {
				m.Mdead()
				if m.Mdead() == false {// si il est pas mort le monstre riposte
					m.attaque = 5
					p.pvactuel = p.pvactuel - m.attaque
					fmt.Println(p.nom, "subit une attaque de", m.nom)
					time.Sleep(3 * time.Second)
					fmt.Println(p.pvactuel, "/", p.pvmax)
					fmt.Println()
					time.Sleep(3 * time.Second)
					comptetour++
					m.Mdead()
				}
			}
		} else {
			if m.Mdead() == false {
				p.menucombat()
			}
		}
	}
}
func (p *personnage) déséquiper(e int) { // permet de retirer son équipement
	for i := 0; i < len(p.equipement.tête); i++ {
		if !(p.equipement.tête == "") { // si tete a chapeau de l'avanturier
			if e == 0 { // si l'utilisateur envoie 0
				p.inventaire = append(p.inventaire, p.equipement.tête) // chapeau de l'aventurier ajouter a l'inventaire
				p.equipement.tête = ""                                 // est remove de chapeau de l'aventurier dans tete
			}
		}
	}
	for i := 0; i < len(p.equipement.torse); i++ {
		if !(p.equipement.torse == "") { // si torse a tunique de l'aventurier
			if e == 1 { // si l'utilisateur envoie 0
				p.inventaire = append(p.inventaire, p.equipement.torse) // tunique de l'aventurier ajouter a l'inventaire
				p.equipement.torse = ""                                 // est remove de tunique de l'aventurier dans tete
			}
		}
	}
	for i := 0; i < len(p.equipement.pied); i++ {
		if !(p.equipement.pied == "") { // si pied n'est pas vide
			if e == 2 { // et que l'utilisateur clique sur 2
				p.inventaire = append(p.inventaire, p.inventaire[i]) // ont ajoute a l'inventaire l'item de pied
				p.equipement.pied = ""                               // et l'ont le supprime de pied
			}
		}
	}
}
func (p *personnage) équiper(d int) { // permet d'équiper ce que l'on veut
	if d == 0 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Chapeau de l'aventurier" && p.equipement.tête == "" { // si le personnage a dans l'inventaire Chapau de l'aventurier et a le niveau 9 minimum
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
					fmt.Println("Tu n'a pas d'item a équiper ")
				}
			}
		}
	}
	if d == 1 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Tunique de l'aventurier" && p.equipement.torse == "" { // si le personnage a dans l'inventaire Chapau de l'aventurier et a le niveau 9 minimum
				p.equipement.torse = "Tunique de l'aventurier" // ont ajoute chapeau de l'aventurier a tete
				p.removeInventory("Tunique de l'aventurier")   // et ont remove de l'inventaire chapeau de l'aventurier
				time.Sleep(2)
				p.pvmax = p.pvmax + 25
				fmt.Println("Tunique de l'aventurier équiper")
				break
			} else {
				if !(p.equipement.torse == "") { // si tete n'est pas vide
					p.addInventory(p.equipement.torse) // ont ajoute l'item de tete dans inventaire
					p.equipement.torse = ""            // retire l'équipement de tete
					fmt.Println("Tu n'a pas d'item a équiper ")
				}
			}
		}
	}
	if d == 2 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Bottes de l'aventurier" && p.equipement.pied == "" { // si le personnage a dans l'inventaire Chapau de l'aventurier et a le niveau 9 minimum
				p.equipement.pied = "Bottes de l'aventurier" // ont ajoute chapeau de l'aventurier a tete
				p.removeInventory("Bottes de l'aventurier")  // et ont remove de l'inventaire chapeau de l'aventurier
				time.Sleep(2)
				p.pvmax = p.pvmax + 15
				fmt.Println("Bottes de l'aventurier équiper")
				break
			} else {
				if !(p.equipement.pied == "") { // si tete n'est pas vide
					p.addInventory(p.equipement.pied) // ont ajoute l'item de tete dans inventaire
					p.equipement.pied = ""            // retire l'équipement de tete
					fmt.Println("Tu n'a pas d'item a équiper ")
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
var try2 int = 0
var tours2 int = 0 //	sert pour le talent de la classe nain
func (p *personnage) chosespell(t int) {
	var count2 int
	var count3 int
	var count4 int
	var count5 int
	if !(p.pvactuel <= 0) || !(m1.pvmonstreactuel <= 0) {
		if t == 0 {
			p.attack(&m1)
			if tours2 == 1 {
				for tours2 < 2 {
					m1.pvmonstreactuel -= p.invocation.attaque
					fmt.Println(p.invocation.nom, "fait", p.invocation.attaque, "dégats a", m1.nom)
					fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
					m1.Mdead()
					tours2 += 1
					count5++
					break
				}
			} else {
				fmt.Print("")
			}
		}
		if p.classe != "Cristian" {//cristian est un cheat code
			for i := 0; i < len(p.skill); i++ {
				if t == 1 {// attaque au point selectionné par le joueur		
					if p.skill[i] == "coup de poing" {// verifie si le kill est possédé par le perso
						if p.manaactuel >= 5 { // verifie si le joueur posséde la mana requise
							if count3 < 1 { // attaque spécial tout les trois tour			
								fmt.Println(p.nom, " utilise coup de poing sur ", m1.nom)
								m1.pvmonstreactuel = m1.pvmonstreactuel - 8
								fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
								time.Sleep(2 * time.Second)
								fmt.Println(m1.nom, " subit une attaque de : ", p.nom)
								p.manaactuel -= 5
								fmt.Print("Il te reste ")
								fmt.Print(p.manaactuel)
								fmt.Println(" mana")
								m1.Mdead()
								count3++
								if tours2 == 1 { // invocation du joueur si il est humain
									for tours2 < 2 {
										m1.pvmonstreactuel -= p.invocation.attaque
										fmt.Println(p.invocation.nom, "fait", p.invocation.attaque, "dégats a", m1.nom)
										fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
										m1.Mdead()
										tours2 += 1
										count5++
										break
									}
								} else {
									fmt.Print("")
								}
							}
						} else {
							fmt.Println("Ta réserve de mana est trop faible pour lancer le sort coup de poing")
							p.attack(&m1)
							break
						}
					}
				} else {
					if t == 2 { //attaque boule de feu
						if expl == 1 {
							if p.manaactuel >= 20 {// verifie si la mana est suffisante
								if count2 < 1 {
									fmt.Println(p.nom, " utilise une boule de feu sur ", m1.nom)
									m1.pvmonstreactuel = m1.pvmonstreactuel - 18
									fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
									time.Sleep(2 * time.Second)
									fmt.Println(m1.nom, " subit une attaque de : ", p.nom)
									p.manaactuel -= 20
									fmt.Print("Il te reste ")
									fmt.Print(p.manaactuel)
									fmt.Println(" mana")
									m1.Mdead()
									count2 += 1
									if tours2 == 1 {
										for tours2 < 2 {
											m1.pvmonstreactuel -= p.invocation.attaque
											fmt.Println(p.invocation.nom, "fait", p.invocation.attaque, "dégats a", m1.nom)
											fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
											m1.Mdead()
											tours2 += 1
											count5++
											break
										}
									} else {
										fmt.Print("")
									}
									break
								}
							} else {
								fmt.Println("Ta réserve de mana est trop faible pour lancer le sort boule de feu")
								p.attack(&m1)
								break
							}
						} else {
							if expl == 0 {
								fmt.Println("tu ne possède pas la compétence Boule de feu")
								p.attack(&m1)
								break
							}
						}
					} else {
						if t == 3 {
							if expl2 == 1 {
								if p.classe == "Elfe" {
									if p.niveau >= 9 {
										if count4 < 1 {
											fmt.Println(p.nom, "Utilise Arbre de vie pour ce soigner")
											p.pvactuel += u
											p.manaactuel -= u
											fmt.Println("Tu a récuprer ", u, "vie")
											fmt.Print("Il te reste ")
											fmt.Print(p.manaactuel)
											fmt.Println(" mana")
											m1.Mdead()
											count4++
										}
									}
								}
							} else {
								if expl2 == 0 {
									fmt.Println("Tu ne possède pas le sort Arbre de vie")
								}
							}
						} else {
							if t == 4 {
								if expl3 == 1 {
									if p.classe == "Nain" && p.niveau >= 3 {
										if count5 < 1 {
											try = 1
											if p.manaactuel >= 20 {
												if try2 == 0 {
													p.manaactuel -= 20
													try2 += 1
													fmt.Println(p.nom, "Utilise Invocation de squelette")
													p.invocation.attaque += rand.Intn(10)
													if i <= 1 {
														p.invocation.nom = "Squeltte"
													} else {
														if i > 1 {
															p.invocation.nom = "Squelttes"
														}
													}
												} else {
													if try2 == 1 {
														fmt.Print("")
													}
												}
												if p.invocation.attaque > 1 {
													fmt.Println("Tu a invoqué une armée", p.invocation.attaque, "Squelettes")
													fmt.Print("Il te reste ")
													fmt.Print(p.manaactuel)
													fmt.Println(" mana")
												} else {
													if p.invocation.attaque == 1 {
														fmt.Println("Tu a invoqué", p.invocation.attaque, "Squelette")
														fmt.Print("Il te reste ")
														fmt.Print(p.manaactuel)
														fmt.Println(" mana")
													} else {
														for tours2 < 2 {
															if p.invocation.attaque <= 0 {
																fmt.Println("L'invocation a échouer")
																fmt.Print("Il te reste ")
																fmt.Print(p.manaactuel)
																fmt.Println(" mana")
															}
														}
													}
												}
												for tours2 < 2 {
													m1.pvmonstreactuel -= p.invocation.attaque
													fmt.Println(p.invocation.nom, "fait", p.invocation.attaque, "dégats a", m1.nom)
													fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
													m1.Mdead()
													tours2 += 1
													count5++
													break
												}
											}
										}
									} else {
										if expl3 == 0 {
											fmt.Println("Tu ne possède pas le sort Invocation de squelette")
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			if p.classe == "Cristian" {
				for i := 0; i < len(p.skill); i++ {
					if t == 1 {
						if p.skill[i] == "coup de poing" {
							if count3 < 1 {
								fmt.Println(p.nom, " utilise coup de poing sur ", m1.nom)
								m1.pvmonstreactuel = m1.pvmonstreactuel - 8
								fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
								time.Sleep(2 * time.Second)
								fmt.Println(m1.nom, " subit une attaque de : ", p.nom)
								fmt.Print("Il te reste ")
								fmt.Print(p.manaactuel)
								fmt.Println(" mana")
								m1.Mdead()
								count3++
							}
						}
					} else {
						if t == 2 {
							if expl == 1 {
								if count2 < 1 {
									fmt.Println(p.nom, " utilise une boule de feu sur ", m1.nom)
									m1.pvmonstreactuel = m1.pvmonstreactuel - 18
									fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
									time.Sleep(2 * time.Second)
									fmt.Println(m1.nom, " subit une attaque de : ", p.nom)
									fmt.Print("Il te reste ")
									fmt.Print(p.manaactuel)
									fmt.Println(" mana")
									m1.Mdead()
									count2++
								}
							} else {
								if expl == 0 {
									fmt.Println("tu ne possède pas la compétence Boule de feu")
									p.attack(&m1)
								}
							}
						} else {//cheat code de la classe dieu qui one shoot
							if t == 5 {
								fmt.Println(p.nom, " UN ONE SHOT SUR ", m1.nom)
								m1.pvmonstreactuel -= m1.pvmonstreactuel
								fmt.Println(m1.pvmonstreactuel, "/", m1.pvmax)
								time.Sleep(2 * time.Second)
								fmt.Println(m1.nom, " subit le jugement dernier de : ", p.nom)
								m1.Mdead()
							} else {
								if t == 6 {
									if expl4 == 1 {
										if newspelloui1 == oui {
											p.manaactuel += newspellint
											fmt.Println("Tu récupère", newspellint, "pv")
										} else {
											if newspelloui1 == non {
												fmt.Print("")
											}
										}
										if newspelloui2 == oui {
											p.pvactuel -= newspell2
											fmt.Println("Tu récupère", newspell2, "mana")
										} else {
											if newspelloui2 == non {
												fmt.Print("")
											}
										}
										if newspelloui3 == oui {
											m1.pvmonstreactuel -= newspell3
											fmt.Println(p.nom, "fait", newspell3, "dégâts")
											fmt.Println("Tu a utilisé le sort", newspellname)
										} else {
											if newspelloui3 == non {
												fmt.Print("")
											}
										}
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

var u int

func (p *personnage) popmana() { //  elles augmentent la mana du joueur (même principe que les popos de soin)
	if p.manamax == p.manaactuel {
		fmt.Println("tu a deja toutes ta mana")// verifie la mana max du joueur
	} else {
		if len(p.inventaire) == 0 {
			fmt.Print("ton inventaire est vide")// verifie si le joueur a un inventaire vide
		} else {
			for i := 0; i < len(p.inventaire); i++ { // cherche une popo de mana dans l'inventaire
				if p.inventaire[i] == "popmana" {
					if p.manaactuel+50 > p.manamax { // verifie si l'utilisation d'une popo ne fera pas dépasser la mana max 
						p.manaactuel = p.manaactuel + p.manamax - p.manaactuel
						p.removeInventory("popmana")
						fmt.Println(p.nom, ":", p.manaactuel, "/", p.manamax)
						break
					} else {
						p.manaactuel = p.manaactuel + 50 
						p.removeInventory("popmana")
						fmt.Println(p.nom, ":", p.manaactuel, "/", p.manamax)
						break
					}
				} else {
					fmt.Println()
				}
			}
		}
	}
}
func (p *personnage) upgradeInventorySlot(r int) { // permet d'augmenter la taille de l'inventaire
	count := 0
	if r == 0 {
		if count == 0 {
			for i := 0; i < len(p.inventaire); i++ {
				if p.inventaire[i] == "Augmentation d'inventaire" {
					s += 10
					p.removeInventory("Augmentation d'inventaire")
					count++
					fmt.Println("Vous avez augmenter votre inventaire de +10")
					break
				}
			}
		}
		if count == 1 {
			for i := 0; i < len(p.inventaire); i++ {
				if p.inventaire[i] == "Augmentation d'inventaire" {
					s += 10
					p.removeInventory("Augmentation d'inventaire")
					fmt.Println("Vous avez augmenter votre inventaire de +10")
					count++
					break
				}
			}
		}
		if count == 2 {
			for i := 0; i < len(p.inventaire); i++ {
				if p.inventaire[i] == "Augmentation d'inventaire" {
					s += 10
					p.removeInventory("Augmentation d'inventaire")
					fmt.Println("Vous avez augmenter votre inventaire de +10")
					count++
					break
				}
			}
		}
		if count >= 3 {
			fmt.Println("Vous avez atteint la limite possible d'augmentation d'inventaire")
			p.menu()
		}
	}
}

/*var upgrade string = ""
	if r == 0 {
		for i := 0; i < len(p.inventaire); i++ {
			if p.inventaire[i] == "Augmentation d'inventaire" {
				if len(p.inventaire) < 40 { // si l'inventaire n'est pas égale a 40
					p.inventaire = append(p.inventaire, upgrade)
					p.removeInventory("Augmentation d'inventaire")
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
}*/

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
	time.Sleep(2 * time.Second)
	fmt.Println("viemaximun:", p.pvmax)
	fmt.Println("PV:", p.pvactuel)
	time.Sleep(2 * time.Second)
	fmt.Println("INVENTAIRE:", p.inventaire)
	fmt.Println("skill:", p.skill)
	fmt.Println("Argent :", p.money)
	fmt.Println("Niveau :", p.niveau)
	time.Sleep(2 * time.Second)
}

func (p *personnage) popovie() { // soigne le perso
	if p.pvmax == p.pvactuel { //si le personnage a toutes ses vies il ne peut pas se soigner
		fmt.Println("tu a deja toutes tes vies")
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
					fmt.Println()
				}
			}
		}
	}
}

func (p *personnage) accessInventory(x int) { // permet d'affiché le contenu d'un inventaireet en supprimer des items
	fmt.Println("Bienvenue dans votre Inventaire")
	for i := len(p.inventaire); i <= len(p.inventaire); i++ {//affiche chaque element de l'inventaire
		if len(p.inventaire) != 0 {
			fmt.Println("--------------------------------", "\nInventaire : \n")
			fmt.Println(p.inventaire)
		} else if len(p.inventaire) == 0 {
			fmt.Println("-------------------------------", "\nInventaire : \n")
			fmt.Println("L'inventaire est vide ")
		}
	}
	var supprimer string
	fmt.Println()
	fmt.Println("Si vous voulez supprimer un élément de l'inventaire tapez Supprimer sinon tapez Fermer")
	fmt.Print("→")// supprime ou permer de revenir au menu de combat ou le menu normal en fonction si on est en combat
	fmt.Scan(&supprimer)
	switch supprimer {
	case "Supprimer":
		var result75 string //sert a  supprimer l'item ou les items demandé(s) c'est pour cela qu'il ya bcp de variable  d
		var result76 string
		var result77 string
		var result78 string
		var result79 string
		var result80 string
		var result81 string
		var result82 string
		var resultall string
		if len(p.inventaire) == 0 { // verifie si l'inventaire  est vide 	
			fmt.Println("Ton inventaire est vide")
		}
		if len(p.inventaire) == 1 {
			result75 += p.inventaire[0] //
			fmt.Println("tapez 0 pour supprimer ", result75)
		}
		if len(p.inventaire) == 2 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
		}
		if len(p.inventaire) == 3 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
		}
		if len(p.inventaire) == 4 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
		}
		if len(p.inventaire) == 5 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
		}
		if len(p.inventaire) == 6 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
		}
		if len(p.inventaire) == 7 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
			result81 += p.inventaire[6]
			fmt.Println("tapez 6 pour supprimer ", result81)
		}
		if len(p.inventaire) == 7 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
			result82 += p.inventaire[6]
			fmt.Println("tapez 6 pour supprimer ", result82)
		}
		if len(p.inventaire) > 7 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
			result82 += p.inventaire[6]
			fmt.Println("tapez 6 pour supprimer ", result82)
			for i := 7; i < len(p.inventaire); i++ {
				resultall = p.inventaire[i]
				if resultall != "" {
					fmt.Println("tapez", i, "pour supprimer", resultall)
				} else {
					if resultall == "" {
						fmt.Print("")
					}
				}
			}
		}
		fmt.Println("→ Tapez 999 pour je sais pas")
		fmt.Print("→")
		var suppr int
		fmt.Scan(&suppr)
		if suppr == 0 {
			fmt.Println("Tu a supprimer", p.inventaire[0], "x1")
			p.removeInventory(p.inventaire[0])
			fmt.Println(p.inventaire)
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 1 {
			fmt.Println("Tu a supprimer", p.inventaire[1], "x1")
			p.removeInventory(p.inventaire[1])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 2 {
			fmt.Println("Tu a supprimer", p.inventaire[2], "x1")
			p.removeInventory(p.inventaire[2])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 3 {
			fmt.Println("Tu a supprimer", p.inventaire[3], "x1")
			p.removeInventory(p.inventaire[3])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 4 {
			fmt.Println("Tu a supprimer", p.inventaire[4], "x1")
			p.removeInventory(p.inventaire[4])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 5 {
			fmt.Println("Tu a supprimer", p.inventaire[5], "x1")
			p.removeInventory(p.inventaire[5])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 6 {
			fmt.Println("Tu a supprimer", p.inventaire[6], "x1")
			p.removeInventory(p.inventaire[6])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 7 {
			fmt.Println("Tu a supprimer", p.inventaire[7], "x1")
			p.removeInventory(p.inventaire[7])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		for i := 8; i < len(p.inventaire); i++ {
			if suppr == i {
				fmt.Println("Tu a supprimer", p.inventaire[i], "x1")
				p.removeInventory(p.inventaire[i])
				p.accessInventory(x)
				time.Sleep(2 * time.Second)
				break
			}
		}
		if suppr == 999 {
			fmt.Println("→ Pourquoi vous l'avez fait?")
			p.accessInventory(x)
		}
	case "supprimer":
		var result75 string
		var result76 string
		var result77 string
		var result78 string
		var result79 string
		var result80 string
		var result81 string
		var result82 string
		var resultall string
		if len(p.inventaire) == 0 {
			fmt.Println("Ton inventaire est vide")
		}
		if len(p.inventaire) == 1 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
		}
		if len(p.inventaire) == 2 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
		}
		if len(p.inventaire) == 3 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
		}
		if len(p.inventaire) == 4 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
		}
		if len(p.inventaire) == 5 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
		}
		if len(p.inventaire) == 6 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
		}
		if len(p.inventaire) == 7 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
			result81 += p.inventaire[6]
			fmt.Println("tapez 6 pour supprimer ", result81)
		}
		if len(p.inventaire) == 7 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
			result82 += p.inventaire[6]
			fmt.Println("tapez 6 pour supprimer ", result82)
		}
		if len(p.inventaire) > 7 {
			result75 += p.inventaire[0]
			fmt.Println("tapez 0 pour supprimer ", result75)
			result76 += p.inventaire[1]
			fmt.Println("tapez 1 pour supprimer ", result76)
			result77 += p.inventaire[2]
			fmt.Println("tapez 2 pour supprimer ", result77)
			result78 += p.inventaire[3]
			fmt.Println("tapez 3 pour supprimer ", result78)
			result79 += p.inventaire[4]
			fmt.Println("tapez 4 pour supprimer ", result79)
			result80 += p.inventaire[5]
			fmt.Println("tapez 5 pour supprimer ", result80)
			result82 += p.inventaire[6]
			fmt.Println("tapez 6 pour supprimer ", result82)
			for i := 7; i < len(p.inventaire); i++ {
				resultall = p.inventaire[i]
				if resultall != "" {
					fmt.Println("tapez", i, "pour supprimer", resultall)
				} else {
					if resultall == "" {
						fmt.Print("")
					}
				}
			}
		}
		fmt.Println("→ Tapez 999 pour je sais pas")
		fmt.Print("→")
		var suppr int
		fmt.Scan(&suppr)
		if suppr == 0 {
			fmt.Println("Tu a supprimer", p.inventaire[0], "x1")
			p.removeInventory(p.inventaire[0])
			fmt.Println(p.inventaire)
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 1 {
			fmt.Println("Tu a supprimer", p.inventaire[1], "x1")
			p.removeInventory(p.inventaire[1])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 2 {
			fmt.Println("Tu a supprimer", p.inventaire[2], "x1")
			p.removeInventory(p.inventaire[2])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 3 {
			fmt.Println("Tu a supprimer", p.inventaire[3], "x1")
			p.removeInventory(p.inventaire[3])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 4 {
			fmt.Println("Tu a supprimer", p.inventaire[4], "x1")
			p.removeInventory(p.inventaire[4])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 5 {
			fmt.Println("Tu a supprimer", p.inventaire[5], "x1")
			p.removeInventory(p.inventaire[5])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 6 {
			fmt.Println("Tu a supprimer", p.inventaire[6], "x1")
			p.removeInventory(p.inventaire[6])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		if suppr == 7 {
			fmt.Println("Tu a supprimer", p.inventaire[7], "x1")
			p.removeInventory(p.inventaire[7])
			p.accessInventory(x)
			time.Sleep(2 * time.Second)
		}
		for i := 8; i < len(p.inventaire); i++ {
			if suppr == i {
				fmt.Println("Tu a supprimer", p.inventaire[i], "x1")
				p.removeInventory(p.inventaire[i])
				p.accessInventory(x)
				time.Sleep(2 * time.Second)
				break
			}
		}
		if suppr == 999 {
			fmt.Println("→ Pourquoi vous l'avez fait?")
			p.accessInventory(x)
		}
	case "fermer":
		if x == 0 {
			p.menu()
		} else {
			p.menucombat()
		}

	case "Fermer":
		if x == 0 {
			p.menu()
		} else {
			p.menucombat()
		}
	}
}

func (p *personnage) removeInventory(itemremove string) { //supprime un objet de l'inventaire d'un personnage
	for i := 0; i < len(p.inventaire); i++ {// cherche l'élément a supprimer
		if itemremove == p.inventaire[i] { // si il est trouvé il est remplacé par un vide a son indice du tableau
			p.inventaire[i] = ""
			break
		}
	}
}

var s int = 10

func (p *personnage) addInventory(itemadd string) bool { //rajoute un élément dans l'inventaire
	if len(p.inventaire) <= s {// s correspond à la limite de l'inventaire 
		p.inventaire = append(p.inventaire, itemadd) //rajoute l'objet
	} else {
		fmt.Println("inventaire complet")
		return false
	}
	return true
}
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

var try int

func (p *personnage) attack(m *monstre) { //attaque standard du poing au cas ou le joueur n'a pas de skill 
	m.pvmonstreactuel = m.pvmonstreactuel - 10
	fmt.Println(m.pvmonstreactuel, "/", m.pvmax)
	time.Sleep(2 * time.Second)
	fmt.Println(m.nom, " subit une attaque de : ", p.nom)
	if comptetour < 3 && try == 1 {
		p.chosespell(4) // invocation du trelnt de la classe
	}
	m.Mdead()
}
func (p *personnage) bouledefeu() { // verifie si le joueur possséde une boule de feu
	for i := 0; i < len(p.skill); i++ {
		if p.skill[i] == "Boule de feu" {
			expl += 1
		}
	}
	if expl == 1 { //modiife les choix d'attaque du joueur en fonction de ses talents 
		fmt.Println("→ Pour faire un coup classique (Tapez 0)")
		fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
		fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
	} else {
		if expl == 0 {
			fmt.Println("→ Pour faire un coup classique (Tapez 0)")
			fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
		}
	}
}

var expl int = 0
var expl2 int = 0
var expl3 int = 0

func (p *personnage) menucombat() { // menu mais en combat certaines fonctionnalités sont limité
	var combat string
	if !(p.pvactuel <= 0) || !(p.monstre.pvmonstreactuel <= 0) {
		fmt.Println("")
		fmt.Println("Bienvenue Dans L'arène")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Tapez: attaque pour Attaquer ")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Tapez: Soin pour vous soigner (cela fera passer votre tour)")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Tapez: Mana pour récupérer votre mana (cela fera passer votre tour)")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Tapez: Inventaire Pour accéder à ce dernier")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Tapez: exit pour fuir comme un cafard  !")
		fmt.Println()
		fmt.Print("→")
		fmt.Scan(&combat)
	} else {
		p.menu()
	}
	switch combat {
	case "Attaque": //attaque par defaut le gobelin
		var chosesort int
		for i := 0; i < len(p.skill); i++ {
			if p.skill[i] == "Boule de feu" {
				expl = 1
			}
			if p.skill[i] == "Arbre de vie" {
				expl2 = 1
			}
			if p.skill[i] == "Invocation de squelette" {
				expl3 = 1
			}
		}
		if expl4 == 1 && expl == 1 { 
			if p.classe == "Cristian" {
				fmt.Println("→ Pour faire un coup classique (Tapez 0)")
				fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
				fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
				fmt.Println("→ Utilisez votre sort personnalisé, du nom de :", newspellname, "(Tapez 6)")
			}
		}
		if expl == 1 && expl4 == 0 {
			if p.classe == "Cristian" {
				fmt.Println("→ Pour faire un coup classique (Tapez 0)")
				fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
				fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
			} else {
				if p.classe != "Cristian" && p.classe != "Elfe" && p.classe != "Nain" {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				} else {
					if p.classe == "Elfe" && p.niveau < 9 {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
					}
				}
				if expl2 == 1 {
					if p.classe == "Elfe" && p.niveau >= 9 {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
						fmt.Println("→ Pour lancer le sort Arbre de vie (Tapez 3)")
					}
				}
			}
			if expl3 == 0 {
				if p.classe == "Nain" && p.niveau < 3 {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				}
			}
			if expl3 == 1 {
				if p.classe == "Nain" && p.niveau >= 3 {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
					fmt.Println("→ Pour lancer le sort Invocation de Squelette (Tapez 4)")
				}
			}
		} else {
			if expl == 0 && expl4 == 0 {
				if p.classe == "Cristian" {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
				} else {
					if p.classe != "Cristian" && p.classe != "Elfe" && p.classe != "Nain" {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					} else {
						if p.classe == "Elfe" && p.niveau < 9 {
							fmt.Println("→ Pour faire un coup classique (Tapez 0)")
							fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						}
						if expl2 == 1 {
							if p.classe == "Elfe" && p.niveau >= 9 {
								fmt.Println("→ Pour faire un coup classique (Tapez 0)")
								fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
								fmt.Println("→ Pour lancer le sort Arbre de vie (Tapez 3)")
							}
						}
						if expl3 == 0 {
							if p.classe == "Nain" && p.niveau < 3 {
								fmt.Println("→ Pour faire un coup classique (Tapez 0)")
								fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
							}
						}
						if expl3 == 1 {
							if p.classe == "Nain" && p.niveau >= 3 {
								fmt.Println("→ Pour faire un coup classique (Tapez 0)")
								fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
								fmt.Println("→ Pour lancer le sort Invocation de Squelette (Tapez 4)")
							}
						}
					}
				}
			} else {
				if p.classe == "Cristian" {
					if expl4 == 1 && expl == 0 {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
						fmt.Println("→ Utilisez votre sort personnalisé, du nom de :", newspellname, "(Tapez 6)")
					}
				}
			}
		}
		if p.initiative >= m1.initiative2 { // l'initiative definbit qui attaque en premier
			if p.classe == "Elfe" && p.niveau >= 9 {
				fmt.Scan(&chosesort)
				if chosesort == 3 {
					fmt.Println("Tu as ", p.manaactuel, " de mana")
					fmt.Scan(&u)
					p.chosespell(chosesort)
					time.Sleep(1 * time.Second)
					p.TrainingFight(&m1)
					p.dead()
					m1.Mdead()
					p.menucombat()
				} else {
					p.chosespell(chosesort)
					time.Sleep(1 * time.Second)
					p.TrainingFight(&m1)
					m1.Mdead()
					p.dead()
					p.menucombat()
				}
			} else {
				fmt.Scan(&chosesort)
				p.chosespell(chosesort)
				time.Sleep(1 * time.Second)
				p.TrainingFight(&m1)
				m1.Mdead()
				p.dead()
				p.menucombat()
			}
		} else {
			if m1.initiative2 > p.initiative {
				if p.classe == "Elfe" && p.niveau >= 9 {
					fmt.Scan(&chosesort)
					if chosesort == 3 {
						fmt.Println("Tu as ", p.manaactuel, " de mana")
						fmt.Scan(&u)
						p.TrainingFight(&m1)
						m1.Mdead()
						time.Sleep(1 * time.Second)
						p.chosespell(chosesort)
						p.dead()
						p.menucombat()
					} else {
						p.TrainingFight(&m1)
						m1.Mdead()
						time.Sleep(1 * time.Second)
						p.chosespell(chosesort)
						p.dead()
						p.menucombat()
					}
				} else {
					fmt.Scan(&chosesort)
					p.TrainingFight(&m1)
					m1.Mdead()
					time.Sleep(1 * time.Second)
					p.chosespell(chosesort)
					p.dead()
					p.menucombat()
				}
			}
		}
	case "attaque":
		var chosesort int
		for i := 0; i < len(p.skill); i++ {
			if p.skill[i] == "Boule de feu" {
				expl = 1
			}
			if p.skill[i] == "Arbre de vie" {
				expl2 = 1
			}
			if p.skill[i] == "Invocation de squelette" {
				expl3 = 1
			}
		}
		if expl4 == 1 && expl == 1 {
			if p.classe == "Cristian" {
				fmt.Println("→ Pour faire un coup classique (Tapez 0)")
				fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
				fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
				fmt.Println("→ Utilisez votre sort personnalisé, du nom de :", newspellname, "(Tapez 6)")
			}
		}
		if expl == 1 && expl4 == 0 { // un niveau d'experience est requis pour utiliser certain talents
			if p.classe == "Cristian" {
				fmt.Println("→ Pour faire un coup classique (Tapez 0)")
				fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
				fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
			} else {
				if p.classe != "Cristian" && p.classe != "Elfe" && p.classe != "Nain" {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				} else {
					if p.classe == "Elfe" && p.niveau < 9 {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
					}
				}
				if expl2 == 1 {
					if p.classe == "Elfe" && p.niveau >= 9 {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
						fmt.Println("→ Pour lancer le sort Arbre de vie (Tapez 3)")
					}
				}
			}
			if expl3 == 0 {
				if p.classe == "Nain" && p.niveau < 3 {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
				}
			}
			if expl3 == 1 {
				if p.classe == "Nain" && p.niveau >= 3 {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ Pour lancer une Boule de feu (Tapez 2)")
					fmt.Println("→ Pour lancer le sort Invocation de Squelette (Tapez 4)")
				}
			}
		} else {
			if expl == 0 && expl4 == 0 {
				if p.classe == "Cristian" {
					fmt.Println("→ Pour faire un coup classique (Tapez 0)")
					fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
				} else {
					if p.classe != "Cristian" && p.classe != "Elfe" && p.classe != "Nain" {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
					} else {
						if p.classe == "Elfe" && p.niveau < 9 {
							fmt.Println("→ Pour faire un coup classique (Tapez 0)")
							fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						}
						if expl2 == 1 {
							if p.classe == "Elfe" && p.niveau >= 9 {
								fmt.Println("→ Pour faire un coup classique (Tapez 0)")
								fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
								fmt.Println("→ Pour lancer le sort Arbre de vie (Tapez 3)")
							}
						}
						if expl3 == 0 {
							if p.classe == "Nain" && p.niveau < 3 {
								fmt.Println("→ Pour faire un coup classique (Tapez 0)")
								fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
							}
						}
						if expl3 == 1 {
							if p.classe == "Nain" && p.niveau >= 3 {
								fmt.Println("→ Pour faire un coup classique (Tapez 0)")
								fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
								fmt.Println("→ Pour lancer le sort Invocation de Squelette (Tapez 4)")
							}
						}
					}
				}
			} else {
				if p.classe == "Cristian" {
					if expl4 == 1 && expl == 0 {
						fmt.Println("→ Pour faire un coup classique (Tapez 0)")
						fmt.Println("→ Pour faire un coup de poing (Tapez 1)")
						fmt.Println("→ POUR PULVERISER TON ENNEMI (Tapez 5)")
						fmt.Println("→ Utilisez votre sort personnalisé, du nom de :", newspellname, "(Tapez 6)")
					}
				}
			}
		}
		if p.initiative >= m1.initiative2 {
			if p.classe == "Elfe" && p.niveau >= 9 {
				fmt.Scan(&chosesort)
				if chosesort == 3 {
					fmt.Println("Tu as ", p.manaactuel, " de mana")
					fmt.Scan(&u)
					p.chosespell(chosesort)
					time.Sleep(1 * time.Second)
					p.TrainingFight(&m1)
					p.dead()
					m1.Mdead()
					p.menucombat()
				} else {
					p.chosespell(chosesort)
					time.Sleep(1 * time.Second)
					p.TrainingFight(&m1)
					m1.Mdead()
					p.dead()
					p.menucombat()
				}
			} else {
				fmt.Scan(&chosesort)
				p.chosespell(chosesort)
				time.Sleep(1 * time.Second)
				p.TrainingFight(&m1)
				m1.Mdead()
				p.dead()
				p.menucombat()
			}
		} else {
			if m1.initiative2 > p.initiative {
				if p.classe == "Elfe" && p.niveau >= 9 {
					fmt.Scan(&chosesort)
					if chosesort == 3 {
						fmt.Println("Tu as ", p.manaactuel, " de mana")
						fmt.Scan(&u)
						p.TrainingFight(&m1)
						m1.Mdead()
						time.Sleep(1 * time.Second)
						p.chosespell(chosesort)
						p.dead()
						p.menucombat()
					} else {
						p.TrainingFight(&m1)
						m1.Mdead()
						time.Sleep(1 * time.Second)
						p.chosespell(chosesort)
						p.dead()
						p.menucombat()
					}
				} else {
					fmt.Scan(&chosesort)
					p.TrainingFight(&m1)
					m1.Mdead()
					time.Sleep(1 * time.Second)
					p.chosespell(chosesort)
					p.dead()
					p.menucombat()
				}
			}
		}
	case "Soin":
		if p.initiative > m1.initiative2 {
			p.popovie()
			p.TrainingFight(&m1)
			p.menucombat()
		} else {
			if m1.initiative2 > p.initiative {
				p.TrainingFight(&m1)
				p.popovie()
				p.menucombat()
			}
		}
	case "soin":
		if p.initiative >= m1.initiative2 {
			p.popovie()
			p.TrainingFight(&m1)
			p.menucombat()
		} else {
			if m1.initiative2 > p.initiative {
				p.TrainingFight(&m1)
				p.popovie()
				p.menucombat()
			}
		}
	case "Mana":
		if p.initiative >= m1.initiative2 {
			p.popmana()
			p.TrainingFight(&m1)
			p.menucombat()
		} else {
			if m1.initiative2 > p.initiative {
				p.TrainingFight(&m1)
				p.popmana()
				p.menucombat()
			}
		}
	case "mana":
		if p.initiative >= m1.initiative2 {
			p.popmana()
			p.TrainingFight(&m1)
			p.menucombat()
		} else {
			if m1.initiative2 > p.initiative {
				p.TrainingFight(&m1)
				p.popmana()
				p.menucombat()
			}
		}
	case "Inventaire":
		p.accessInventory(1)
	case "inventaire":
		p.accessInventory(1)
	case "Exit":
		fmt.Println("Fin du combat !")
		time.Sleep(2 * time.Second)
		p.menu()
	case "exit":
		fmt.Println("Fin du combat !")
		time.Sleep(2 * time.Second)
		p.menu()
	}
}

var test27 int
var test28 int
var test29 int
var test30 int
var test31 int
var test32 int
var mdp string
var newspell string
var newspelloui1 string
var newspelloui2 string
var newspelloui3 string
var newspell2 int
var newspell3 int
var newspellint int
var newspellname string
var expl4 int
var oui string = "oui"
var non string = "non"

func (p *personnage) newspellcreate() { //permet de créer un sort quand on est en mod dieu
	if p.classe == "Cristian" {
		fmt.Println("Ici tu peut crée ton propre sort :")
		fmt.Println("Quelle sera le nom de ce nouveau sort, OH MON ROI!!")
		fmt.Scan(&newspellname)
		p.skill = append(p.skill, newspellname)
		fmt.Println("Tu a crée un sort du nom de", newspellname)
		fmt.Println("Veut tu crée un sort qui coute de la mana", oui, "ou", non)
		fmt.Scan(&newspelloui1)
		if newspelloui1 == non {
			fmt.Println("Ton sort ne coûte aucun point de mana")
			time.Sleep(1 * time.Second)
		} else {
			if newspelloui1 == oui {
				fmt.Println("Combien de mana lui faut-il ?")
				fmt.Scan(&newspellint)
				fmt.Println("Ton sort coûte", newspellint, "points de mana par utilisations")
				time.Sleep(2 * time.Second)
			}
		}
		fmt.Println("Es que ton sort peut soigné", oui, "ou", non)
		fmt.Scan(&newspelloui2)
		if newspelloui2 == non {
			fmt.Println("Ce sort n'est pas un sort de soin")
			time.Sleep(1 * time.Second)
		} else {
			if newspelloui2 == oui {
				fmt.Println("Combien de pv peut-il soigné ?")
				fmt.Scan(&newspell2)
				fmt.Println("Ton sort peut soigner", newspell2, "de pv")
				time.Sleep(2 * time.Second)
			}
		}
		fmt.Println("Es que tu peut attaquer", oui, "ou", non)
		fmt.Scan(&newspelloui3)
		if newspelloui3 == non {
			fmt.Println("Ce sort n'est pas fait pour tabasser sont adversaire")
			time.Sleep(1 * time.Second)
		} else {
			if newspelloui3 == oui {
				fmt.Println("Combien de dégats peut tu faire ?")
				fmt.Scan(&newspell3)
				fmt.Println("Tu fait", newspell3, "dégats par attaque")
				time.Sleep(2 * time.Second)
			}
		}
		fmt.Println("Bravo tu a crée un nouveau sort")
	}
	expl4 += 1
}

func (p *personnage) removestats(w int) { // permet en mode dieu de modifier les pv,l'attque de son perso genial pour tester les failles du jeu
	if p.classe == "Cristian" {
		if w == 0 {
			p.money = test27
			fmt.Print("tu a reçu  ")
			fmt.Print(test27)
			fmt.Println(" PO")
			time.Sleep(2 * time.Second)
		}
		if w == 1 {
			p.pvmax = test28
			fmt.Print("tu a maintenant  ")
			fmt.Print(test28)
			fmt.Println(" en pv max")
			time.Sleep(2 * time.Second)
		}
		if w == 2 {
			p.pvactuel = test29
			fmt.Print("tu a maintenant  ")
			fmt.Print(test29)
			fmt.Println(" en pv actuel")
			time.Sleep(2 * time.Second)
		}
		if w == 3 {
			p.manamax = test30
			fmt.Print("tu a maintenant  ")
			fmt.Print(test30)
			fmt.Println(" en mana max")
			time.Sleep(2 * time.Second)
		}
		if w == 4 {
			p.manaactuel = test31
			fmt.Print("tu a maintenant  ")
			fmt.Print(test31)
			fmt.Println(" en mana actuel")
			time.Sleep(2 * time.Second)

		}
		if w == 5 {
			p.niveau = test32
			fmt.Print("tu es maintenant niveau ")
			fmt.Println(test32)
			time.Sleep(2 * time.Second)
		}
		if w == 6 {
			p.newspellcreate()
		}
	} else {
		fmt.Println("Tu ne possède pas les droits nécessaires pour utiliser cet commande")
	}
}
func (p *personnage) menu() {
	var commande string
	if p.classe != "Cristian" {
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
		fmt.Println("Tapez Name pour rentrer Votre Nom d'utilisateur et choisir votre classe ")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Tapez exit pour RAGE QUIT !!!")
		fmt.Println("---------------------------------------------------")
		fmt.Println()
		fmt.Print("→")
		fmt.Scan(&commande)
	} else {
		if p.classe == "Cristian" { //un menu différent apparait en fonction du cheatcode
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
			fmt.Println("Tapez Name pour rentrer Votre Nom d'utilisateur et choisir votre classe ")
			fmt.Println("---------------------------------------------------")
			fmt.Println("MODIFIER VOS STATS TAPEZ removestats")
			fmt.Println("---------------------------------------------------")
			fmt.Println("Tapez exit pour RAGE QUIT !!!")
			fmt.Println("---------------------------------------------------")
			fmt.Println()
			fmt.Print("→")
			fmt.Scan(&commande)
		}
	}
	switch commande {
	case "Information":
		p.displayInfo()
		p.menu()
	case "information":
		p.displayInfo()
		p.menu()
	case "Removestats":
		var n int// modifie les stats du joueur 
		var test int
		var test2 int
		var test3 int
		var test4 int
		var test5 int
		var test6 int
		fmt.Println("→ Si vous souhaitez modifier votre argent (Tapez 0)")
		fmt.Println("→ Si vous souhaitez modifier vos pv max (Tapez 1)")
		fmt.Println("→ Si vous souhaitez modifier pv actuel (Tapez 2)")
		fmt.Println("→ Si vous souhaitez modifier votre mana max (Tapez 3)")
		fmt.Println("→ Si vous souhaitez modifier votre mana actuel (Tapez 4)")
		fmt.Println("→ Si vous souhaitez modifier votre niveau actuel (Tapez 5)")
		fmt.Println("→ Si vous souhaitez crée un nouveau sort (Tapez 6)")
		fmt.Scan(&n)
		if n == 0 {
			fmt.Scan(&test)
			p.removestats(test)
			test27 = test
		}
		if n == 1 {
			fmt.Scan(&test2)
			p.removestats(test2)
			test28 = test2
		}
		if n == 2 {
			fmt.Scan(&test3)
			p.removestats(test3)
			test29 = test3
		}
		if n == 3 {
			fmt.Scan(&test4)
			p.removestats(test4)
			test30 = test4
		}
		if n == 4 {
			fmt.Scan(&test5)
			p.removestats(test5)
			test31 = test5
		}
		if n == 5 {
			fmt.Scan(&test6)
			p.removestats(test6)
			test32 = test6
		}
		p.removestats(n)
		p.menu()
	case "removestats":
		var n int
		var test int
		var test2 int
		var test3 int
		var test4 int
		var test5 int
		var test6 int
		fmt.Println("→ Si vous souhaitez modifier votre argent (Tapez 0)")
		fmt.Println("→ Si vous souhaitez modifier vos pv max (Tapez 1)")
		fmt.Println("→ Si vous souhaitez modifier pv actuel (Tapez 2)")
		fmt.Println("→ Si vous souhaitez modifier votre mana max (Tapez 3)")
		fmt.Println("→ Si vous souhaitez modifier votre mana actuel (Tapez 4)")
		fmt.Println("→ Si vous souhaitez modifier votre niveau actuel (Tapez 5)")
		fmt.Println("→ Si vous souhaitez crée un nouveau sort (Tapez 6)")
		fmt.Scan(&n)
		if n == 0 {
			fmt.Scan(&test)
			p.removestats(test)
			test27 = test
		}
		if n == 1 {
			fmt.Scan(&test2)
			p.removestats(test2)
			test28 = test2
		}
		if n == 2 {
			fmt.Scan(&test3)
			p.removestats(test3)
			test29 = test3
		}
		if n == 3 {
			fmt.Scan(&test4)
			p.removestats(test4)
			test30 = test4
		}
		if n == 4 {
			fmt.Scan(&test5)
			p.removestats(test5)
			test31 = test5
		}
		if n == 5 {
			fmt.Scan(&test6)
			p.removestats(test6)
			test32 = test6
		}
		p.removestats(n)
		p.menu()
	case "Inventaire":
		p.accessInventory(0)
	case "inventaire":
		p.accessInventory(0)
	case "équiper":
		var equiper int
		fmt.Println("→ Bienvene dans l'interface d'équippement")
		fmt.Println("→ Si vous souhaitez mettre un Chapeau (Tapez 0)")
		fmt.Println("→ Si vous souhaitez mettre un Plastron (Tapez 1)")
		fmt.Println("→ Si vous souhaitez mettre des Chaussures (Tapez 2)")
		fmt.Print("→")
		fmt.Scan(&equiper)
		p.équiper(equiper)
		p.menu()
	case "UpgradeInventorySlot":
		var upgrade int
		fmt.Scan(&upgrade)
		p.upgradeInventorySlot((upgrade))
		p.menu()
	case "Marchand": // mArchand
		var marchand int
		fmt.Println("↓----------------------------------------------↓")
		fmt.Println("→ Tapez acheter pour accéder a l'interface d'achat")
		fmt.Println("→ Tapez vendre pour accéder a l'interface de vente")
		fmt.Println("---------------------------------------------------")
		fmt.Scan(&l)
		if l == "acheter" || l == "Acheter" || l == "ACHETER" {
			if p.classe != "Cristian" {
				fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
				fmt.Println("Produit en Vente :")
				fmt.Println("→ Potion de Soins: 3$ (Tapez 0)")
				fmt.Println("→ Potion de Poison 6$(Tapez 1)")
				fmt.Println("→ Potion de Livre de Sort Boule de feu  !25 $(Tapez 2)")
				fmt.Println("→ Fourrure de loup  4$ (Tapez 3)")
				fmt.Println("→ Peau de Troll  7$ (Tapez 4)")
				fmt.Println("→ Peau de Sanglier 3$ (Tapez 5)")
				fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
				fmt.Println("→ Augmentation d'inventaire 30 $ (Tapez 7) ")
				fmt.Println("→ Potion de Mana 11 $ (Tapez 8) ")
			} else {
				if p.classe == "Cristian" {
					fmt.Println("→ Bienvenue dans le mode marchand caché pour les modérateurs ← ")
					fmt.Println("Produit en Vente :")
					fmt.Println("→ Potion de Soins (Tapez 0)")
					fmt.Println("→ Potion de Poison (Tapez 1)")
					fmt.Println("→ Potion de Livre de Sort Boule de feu  !(Tapez 2)")
					fmt.Println("→ Fourrure de loup   (Tapez 3)")
					fmt.Println("→ Peau de Troll   (Tapez 4)")
					fmt.Println("→ Peau de Sanglier (Tapez 5)")
					fmt.Println("→ Plume de Corbeau  (Tapez 6) ")
					fmt.Println("→ Augmentation d'inventaire  (Tapez 7) ")
					fmt.Println("→ Potion de Mana  (Tapez 8) ")
					fmt.Println("→ Pour doubler votre argent  (Tapez 9)")
					fmt.Println("→ Pour Choisir votre argent (Tapez 10)")
				}
			}
			fmt.Scan(&marchand)
			p.pnj(marchand)
			p.menu()
		} else {
			if l == "vendre" || l == "Vendre" || l == "VENDRE" || l == "vente" {
				if p.classe != "Cristian" {
					fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
					fmt.Println("Produit a Vendre :")
					fmt.Println("→ Potion de Soins: +1$ (Tapez 0)")
					fmt.Println("→ Potion de Poison +4$(Tapez 1)")
					fmt.Println("→ Fourrure de loup  +2$ (Tapez 2)")
					fmt.Println("→ Peau de Troll  +3$ (Tapez 3)")
					fmt.Println("→ Peau de Sanglier +1$ (Tapez 4)")
					fmt.Println("→ Plume de Corbeau +1$ (Tapez 5) ")
					fmt.Println("→ Augmentation d'inventaire +12$ (Tapez 6) ")
					fmt.Println("→ Potion de Mana +7$ (Tapez 7) ")
				}
			}
			fmt.Scan(&marchand)
			p.pnj(marchand)
			p.menu()
		}
		p.menu()
	case "marchand":
		var marchand int
		fmt.Println("↓----------------------------------------------↓")
		fmt.Println("→ Tapez acheter pour accéder a l'interface d'achat")
		fmt.Println("→ Tapez vendre pour accéder a l'interface de vente")
		fmt.Println("---------------------------------------------------")
		fmt.Scan(&l)
		if l == "acheter" || l == "Acheter" || l == "ACHETER" {
			if p.classe != "Cristian" {
				fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
				fmt.Println("Produit en Vente :")
				fmt.Println("→ Potion de Soins: 3$ (Tapez 0)")
				fmt.Println("→ Potion de Poison 6$(Tapez 1)")
				fmt.Println("→ Potion de Livre de Sort Boule de feu  !25 $(Tapez 2)")
				fmt.Println("→ Fourrure de loup  4$ (Tapez 3)")
				fmt.Println("→ Peau de Troll  7$ (Tapez 4)")
				fmt.Println("→ Peau de Sanglier 3$ (Tapez 5)")
				fmt.Println("→ Plume de Corbeau 1 $ (Tapez 6) ")
				fmt.Println("→ Augmentation d'inventaire 30 $ (Tapez 7) ")
				fmt.Println("→ Potion de Mana 11 $ (Tapez 8) ")
			} else {
				if p.classe == "Cristian" {
					fmt.Println("→ Bienvenue dans le mode marchand caché pour les modérateurs ← ")
					fmt.Println("Produit en Vente :")
					fmt.Println("→ Potion de Soins (Tapez 0)")
					fmt.Println("→ Potion de Poison (Tapez 1)")
					fmt.Println("→ Potion de Livre de Sort Boule de feu  !(Tapez 2)")
					fmt.Println("→ Fourrure de loup   (Tapez 3)")
					fmt.Println("→ Peau de Troll   (Tapez 4)")
					fmt.Println("→ Peau de Sanglier (Tapez 5)")
					fmt.Println("→ Plume de Corbeau  (Tapez 6) ")
					fmt.Println("→ Augmentation d'inventaire  (Tapez 7) ")
					fmt.Println("→ Potion de Mana  (Tapez 8) ")
					fmt.Println("→ Pour doubler votre argent  (Tapez 9)")
					fmt.Println("→ Pour Choisir votre argent (Tapez 10)")
				}
			}
			fmt.Scan(&marchand)
			p.pnj(marchand)
			p.menu()
		} else {
			if l == "vendre" || l == "Vendre" || l == "VENDRE" || l == "vente" {
				if p.classe != "Cristian" {
					fmt.Println("→ Bienvenue chez Jacquie Farce & Attrape ← ")
					fmt.Println("Produit a Vendre :")
					fmt.Println("→ Potion de Soins: +1$ (Tapez 0)")
					fmt.Println("→ Potion de Poison +4$(Tapez 1)")
					fmt.Println("→ Fourrure de loup  +2$ (Tapez 2)")
					fmt.Println("→ Peau de Troll  +3$ (Tapez 3)")
					fmt.Println("→ Peau de Sanglier +1$ (Tapez 4)")
					fmt.Println("→ Plume de Corbeau +1$ (Tapez 5) ")
					fmt.Println("→ Augmentation d'inventaire +12$ (Tapez 6) ")
					fmt.Println("→ Potion de Mana +7$ (Tapez 7) ")
				}
			}
			fmt.Scan(&marchand)
			p.pnj(marchand)
			p.menu()
		}
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
		fmt.Println("Pour craft le Chapeau de l'aventurier" + " écrivez 0")
		fmt.Println("Pour craft la Tunique de l'aventurier" + " écrivez 1")
		fmt.Println("Pour craft les Bottes de l'aventurier" + " écrivez 2")
		fmt.Scan(&forgeron)
		p.forgeron((forgeron))
		p.menu()
	case "name":
		var choice string
		var clase int
		fmt.Println("Veuillez Ecrire le nom de votre Personnage")
		fmt.Println("---------------------------------------------------")
		fmt.Println()
		fmt.Print("→")
		fmt.Scan(&choice)
		fmt.Println("Puis Choisissez votre Classe :")
		fmt.Println("♪---------------------------------------------------")
		fmt.Println("-Tapez 0 pour devenir un Humain")
		fmt.Println("---------------------------------------------------")
		fmt.Println("-Tapez 1 pour devenir un Elfe")
		fmt.Println("---------------------------------------------------")
		fmt.Println("-Tapez 2 pour devenir un Nain")
		fmt.Println("---------------------------------------------------")
		fmt.Println()
		fmt.Print("→")
		fmt.Scan(&clase)
		p.charCreation(choice, clase)
		p.menu()
	case "Name":
		var choice string
		var clase int
		fmt.Println("Veuillez Ecrire le nom de votre Personnage")
		fmt.Println("---------------------------------------------------")
		fmt.Println()
		fmt.Print("→")
		fmt.Scan(&choice)
		fmt.Println("Puis Choisissez votre Classe :")
		fmt.Println("-Tapez 0 pour devenir un Humain")
		fmt.Println("---------------------------------------------------")
		fmt.Println("-Tapez 1 pour devenir un Elfe")
		fmt.Println("---------------------------------------------------")
		fmt.Println("-Tapez 2 pour devenir un Nain")
		fmt.Println("---------------------------------------------------")
		fmt.Println()
		fmt.Print("→")
		fmt.Scan(&clase)
		p.charCreation(choice, clase)
		p.menu()
	case "Exit":
		os.Exit(0)
	case "exit":
		os.Exit(0)
	}
	fmt.Println("La commande exécuter n'est pas enregistrer")
	time.Sleep(2 * time.Second)
	p.menu()
}

var t int
var l string

func (p *personnage) pnj(i int) { // pnj vendeurs qui vend des items en fonction des entrés utilisateurs et du portefeuille du joueur  
	if l == "acheter" || l == "Acheter" || l == "ACHETER" {
		if p.classe != "Cristian" {
			if i == 0 && p.money >= 3 {
				if p.addInventory("popovie") == true {
					p.money = p.money - 3
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")
				}
			}
			if i == 8 && p.money >= 11 {
				if p.addInventory("popmana") == true {
					p.money = p.money - 11
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")
				}
			}
			if i == 1 && p.money >= 6 {
				if p.addInventory("poison") == true {
					p.money = p.money - 6
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")
				}

			}
			if i == 2 && p.money >= 25 {
				if p.addInventory("Livre de Sort: Boule de feu") == true {
					if t == 0 {
						p.money = p.money - 25
						p.removeInventory("Livre de Sort: Boule de feu")
						fmt.Println("Merci de votre Achat")
						p.spellBook("Boule de feu", t)
						t += 1
					} else {
						fmt.Println("Tu possède déja ce sort")
					}
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
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")

				}
			}
			if i == 5 && p.money >= 3 {
				if p.addInventory("Peau de Sanglier") == true {
					p.money = p.money - 3
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")
				}
			}
			if i == 6 && p.money >= 1 {
				if p.addInventory("Plume de Corbeau") == true { //verifie la place dans l'inventaire 
					p.money = p.money - 1 // l'argent 
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")
				}
			}
			if i == 7 && p.money >= 30 {
				if p.addInventory("Augmentation d'inventaire") == true {
					p.money = p.money - 30
					fmt.Println("Merci de votre Achat")
				} else {
					fmt.Println("Plus de Place ☺")
				}
			}
		} else {
			if p.classe == "Cristian" {
				if i == 0 {
					if p.addInventory("popovie") == true {
					} else {
						fmt.Println("Plus de Place ☺")
					}
				}
				if i == 8 {
					if p.addInventory("popmana") == true {
					} else {
						fmt.Println("Plus de Place ☺")
					}
				}
				if i == 1 {
					if p.addInventory("poison") == true {
					} else {
						fmt.Println("Plus de Place ☺")
					}

				}
				if i == 2 {
					if p.addInventory("Livre de Sort: Boule de feu") == true {
						if t == 0 {
							p.removeInventory("Livre de Sort: Boule de feu")
							fmt.Println("Merci de votre Achat")
							p.spellBook("Boule de feu", t)
							t += 1
						} else {
							fmt.Println("Tu possède déja ce sort")
						}
					} else {
						fmt.Println("Plus de Place ☺")
					}
				}
				if i == 3 {
					if p.addInventory("Fourrure de loup") == true {
						fmt.Println("Merci de votre Achat")
					} else {
						fmt.Println("Plus de Place ☺")

					}
				}
				if i == 4 {
					if p.addInventory("Peau de Troll") == true {
					} else {
						fmt.Println("Plus de Place ☺")

					}
				}
				if i == 5 {
					if p.addInventory("Peau de Sanglier") == true {
					} else {
						fmt.Println("Plus de Place ☺")
					}
				}
				if i == 6 {
					if p.addInventory("Plume de Corbeau") == true {
					} else {
						fmt.Println("Plus de Place ☺")
					}
				}
				if i == 7 {
					if p.addInventory("Augmentation d'inventaire") == true {
					} else {
						fmt.Println("Plus de Place ☺")
					}
				}

			}
		}
	} else {
		var montant int
		if l == "vendre" || l == "Vendre" || l == "VENDRE" || l == "vente" {
			for m := 0; m < len(p.inventaire); m++ {
				if i == 0 {
					if montant < 1 {
						if "popvie" == p.inventaire[m] {
							montant++
							fmt.Println("Tu a vendue une popvie")
							p.removeInventory("popvie")
							p.money = p.money + 1
							fmt.Println("Tu as reçu 1 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 7 {
					if montant < 1 {
						if p.inventaire[m] == "popmana" {
							p.removeInventory("popmana")
							p.money = p.money + 7
							montant++
							fmt.Println("Tu a vendue une popmana")
							fmt.Println("Tu a reçu 7 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 1 {
					if montant < 1 {
						if p.inventaire[m] == "poison" {
							p.removeInventory("poison")
							p.money = p.money + 4
							montant++
							fmt.Println("Tu a vendue une potion de poison")
							fmt.Println("Tu a reçu 4 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 2 {
					if montant < 1 {
						if p.inventaire[m] == "Fourrure de loup" {
							p.removeInventory("Fourrure de loup")
							p.money = p.money + 2
							montant++
							fmt.Println("Tu a vendue une Fourrure de loup")
							fmt.Println("Tu a reçu 2 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 3 {
					if montant < 1 {
						if p.inventaire[m] == "Peau de Troll" {
							p.removeInventory("Peau de Troll")
							p.money = p.money + 3
							montant++
							fmt.Println("Tu a vendue une Peau de Troll")
							fmt.Println("Tu a reçu 3 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 4 {
					if montant < 1 {
						if p.inventaire[m] == "Peau de Sanglier" {
							p.removeInventory("Peau de Sanglier")
							p.money = p.money + 1
							montant++
							fmt.Println("Tu a vendue une Peau de Sanglier")
							fmt.Println("Tu a reçu 1 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 5 {
					if montant < 1 {
						if p.inventaire[m] == "Plume de Corbeau" {
							p.removeInventory("Plume de Corbeau")
							p.money = p.money + 1
							montant++
							fmt.Println("Tu a vendue une Plume de Corbeau")
							fmt.Println("Tu a reçu 1 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
				if i == 6 {
					if montant < 1 {
						if p.inventaire[m] == "Augmentation d'inventaire" {
							p.removeInventory("Augmentation d'inventaire")
							p.money = p.money + 12
							montant++
							fmt.Println("Tu a vendue une Augmentation d'inventaire")
							fmt.Println("Tu a reçu 12 PO")
						}
					} else {
						fmt.Println("Tu na pas l'item!!")
						break
					}
				}
			}
		}
	}
}

func (p *personnage) dead() bool { //verifie si le perso est mort
	if p.pvactuel <= 0 {
		fmt.Println(p.nom, ": a succombé(e)")
		fmt.Println("Défaite")
		p.pvactuel = p.pvmax * 50 / 100
		p.menu()
		return true
	} else {
		return false
	}

}
func (m *monstre) Mdead() bool { //verifie si le monstre est mort
	if m.pvmonstreactuel <= 0 {
		fmt.Println(m.nom, ": a succombé(e)")
		fmt.Println("C'est gagné")
		m.pvmonstreactuel = m.pvmax
		m1.niveaujoueur(&p1, 28000)
		m1.Deadv2(&p1)
		return true
	} else {
		return false
	}
}
func (m *monstre) Deadv2(p *personnage) {
	p.dropend()
	time.Sleep(2 * time.Second)
	p.menu()
}
func (p *personnage) dropend() {
	p.luckofdrop([]string{"Fourrure de loup", "popvie", "Cuir de Sanglier", "Peau de Gobelin", "Plume de Corbeau", "Peau de Troll"})
	p.lootmoneyrandom(7)
}
func (p *personnage) charCreation(s string, r int) {
	count := 1
	p.nom = ""
	result := ""
	for z := 0; z < len(s); z++ {
		if z <= 0 {
			if s[0] >= 'A' && s[0] <= 'Z' {
				result += string(s[0])
			} else if s[0] >= 'a' && s[0] <= 'z' {
				result += string(s[0] - 32)
			}
		} else {
			if z >= 1 {
				if s[count] >= 'A' && s[count] <= 'Z' {
					result += string(s[count] + 32)
					count++
				} else if s[count] >= 'a' && s[count] <= 'z' {
					result += string(s[count])
					count++
				}
			}
		}
	}
	p.nom += result
	time.Sleep(2 * time.Second)
	if r == 0 {
		p.classe = "Humain"
		p.pvmax = 100
		p.pvactuel = p.pvmax / 2
		p.niveau = 1
		p.inventaire = nil
		p.skill = []string{"coup de poing"}
		fmt.Print("Tu a crée un Humain du nom de ")
		fmt.Println(p.nom)
		p.displayInfo()
	}
	if r == 1 {
		p.classe = "Elfe"
		p.pvmax = 80
		p.pvactuel = p.pvmax / 2
		p.niveau = 1
		p.inventaire = nil
		p.skill = []string{"coup de poing"}
		fmt.Print("Tu a crée un Elfe du nom de ")
		fmt.Println(p.nom)
		p.displayInfo()
	}
	if r == 2 {
		p.classe = "Nain"
		p.pvmax = 120
		p.pvactuel = p.pvmax / 2
		p.niveau = 1
		p.inventaire = nil
		p.skill = []string{"coup de poing"}
		fmt.Print("Tu a crée un Nain du nom de ")
		fmt.Println(p.nom)
		p.displayInfo()
	}
	if r == 3 {
		fmt.Println("Mot de passe Requis :")
		fmt.Scan(&mdp)
		if mdp == "1234" {
			p.classe = "Cristian"
			p.pvmax = 99999
			p.pvactuel = p.pvmax / 2
			p.niveau = 999
			p.money = 10000
			p.skill = []string{"coup de poing", "one shot"}
			p.inventaire = []string{"GOD SWORD", "GOD PICKAXE", "GOD AXE", "GOD HELMET", "GOD LEGGINS"}
			fmt.Print("Tu a crée un DIEU du nom de ")
			fmt.Println(p.nom)
			p.displayInfo()
		} else {
			fmt.Println("Mauvais mot de passe")
		}
	}
}

/*for z := 0; z < len(s); z++ {
		for i := 0; i >= 'A' && i <= 'Z'; i++ {
			if !(p.nom[0] == s[i]) {
				result2 += string(s[z] - 32)
				p.nom += result2
			} else {
				p.nom += result2
			}
		}
		for compteur := 1; compteur < len(s); compteur++ {
			for b := 0; b >= 'a' && b <= 'z'; b++ {
				if !(p.nom[compteur] == s[b]) {
					result2 += string(s[z] + 32)
					p.nom += result2
				} else {
					p.nom += result2
				}
			}
		}
	}
}*/

/*var Classe2 = [4]string{"Humain", "Elfe", "Nain", "Cristian"}
	for i := 0; i < len(Classe2); i++ {
		fmt.Println(p.nom)
		fmt.Println("Choisit ta Classe :")
		fmt.Println("-Tapez 0 pour devenir un Humain")
		fmt.Println("-Tapez 1 pour devenir un Elfe")
		fmt.Println("-Tapez 2 pour devenir un Nain")
		if r == 0 {
			p.classe = "Humain"
			p.pvmax = 100
			p.pvactuel = p.pvmax / 2
			p.niveau = 1
			p.skill = append(p.skill, "coup de point")
			fmt.Print("Tu a crée un Humain du nom de ")
			fmt.Println(p.nom)
		}
		if r == 1 {
			p.classe = "Elfe"
			p.pvmax = 80
			p.pvactuel = p.pvmax / 2
			p.niveau = 1
			p.skill = append(p.skill, "coup de point")
			fmt.Print("Tu a crée un Elfe du nom de ")
			fmt.Println(p.nom)
		}
		if r == 2 {
			p.classe = "Nain"
			p.pvmax = 120
			p.pvactuel = p.pvmax / 2
			p.niveau = 1
			p.skill = append(p.skill, "coup de point")
			fmt.Print("Tu a crée un Nain du nom de ")
			fmt.Println(p.nom)
		}
		if r == 3 {
			p.classe = "Cristian"
			p.pvmax = 99999
			p.pvactuel = p.pvmax / 2
			p.niveau = 999
			p.skill = append(p.skill, "coup de point")
			fmt.Print("Tu a crée un DIEU du nom de ")
			fmt.Println(p.nom)
		}
	}
}*/

func (p *personnage) spellBook(talentcaché string, count int) { //attribue des compétemces en fonctions des livres achetés
	for i := 0; i < len(p.skill); i++ {
		if count == 0 {
			if !(p.skill[i] == talentcaché) {
				p.skill = append(p.skill, talentcaché)
			}
		}
		if count == 1 {
			if p.skill[i] == talentcaché {
				fmt.Println("Tu possédes déjà ce talent")
				break

			}
		}
	}
}
func (p *personnage) forgeron(b int) {
	if p.money == 0 {
		fmt.Println("vous n'avez plus d'argent sur vous!")
	}
	var count int = 0
	for i := 0; i < len(p.inventaire); i++ {
		for q := 0; q < len(p.inventaire); q++ {
			if b == 0 && p.money >= 5 {
				if p.inventaire[i] == "Plume de Corbeau" {
					if p.inventaire[q] == "Cuir de Sanglier" {
						if count < 1 {
							p.removeInventory("Plume de Corbeau")
							p.removeInventory("Cuir de Sanglier")
							count++
							p.money = p.money - 5
							p.addInventory("Chapeau de l'aventurier")
							fmt.Println("Tu as craft un Chapeau de l'aventurier")
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(p.inventaire); i++ {
		for q := 0; q < len(p.inventaire); q++ {
			if b == 1 && p.money >= 5 {
				if p.inventaire[i] == "Fourrure de loup" {
					if p.inventaire[q] == "Peau de Troll" {
						if count < 1 {
							p.removeInventory("Fourrure de loup")
							p.removeInventory("Peau de Troll")
							p.removeInventory("Fourrure de loup")
							count++
							p.money = p.money - 5
							p.addInventory("Tunique de l'aventurier")
							fmt.Println("Tu as craft une Tunique de l'aventurier")
							break
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(p.inventaire); i++ {
		for q := 0; q < len(p.inventaire); q++ {
			if b == 2 && p.money >= 5 {
				if p.inventaire[i] == "Fourrure de loup" {
					if p.inventaire[q] == "Cuir de Sanglier" {
						if count < 1 {
							p.removeInventory("Fourrure de loup")
							p.removeInventory("Cuir de Sanglier")
							count++
							p.money = p.money - 5
							p.addInventory("Bottes de l'aventurier")
							fmt.Println("Tu as craft les Bottes de l'aventurier")
							break
						}
					}
				}
			}
		}
	}
}

/*func (m *monstre) attackgoblinPattern() {
	var tours = []int{}               // je crée une variable int qui va prendre en compte plsusieurs chiffre pour compter les tours jusqu'a l'infinie
	m.attaque = 5                     // l'attaque du monstre est de basse vaut 5
	for i := 0; i < len(tours); i++ { // si i inférieur a len de tours on continue d'avancer, est tours étant infinie
		if tours[i] == tours[2] { // si tours[i] vaut tours[2], c'est a dire le troixième tours vue que i part de zéro, et que les tours eux partent de 1 (1er tours, 2ème, ect...)
			m.attaque *= 2 // l'attaque du monstre est multiplié par deux
			break          // ont casse la boucle
		} else {
			m.attaque = 5 // sinon il est normale
		}
	}
}

func (m *monstre) initmonstre(nom string, pvmax int, attaque int, pvmonstre int) {
    m.nom = nom
    m.attaque = attaque
    m.pvmonstre = pvmonstre
    m.pvmax = pvmax
}*/
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

/*func (p *personnage) goblinPattern() {
	var tours = []int{}                      // ont re crée les tours comme avant cela reste les meme
	p.monstre.nom = "Gobelin d'entrainement" // je nome mon monstre
	p.monstre.pvmax = 40                     // je lui donne comme pvmax 40pdv
	p.monstre.attaque = 5
	p.monstre.pvmonstreactuel = p.monstre.pvmax // et je dit que par default le monstre comme nce avec c'est pv max
	for i := 0; i < len(tours); i++ {           // comme tout a l'heure pour i partant de zéro est inférieur a len de tours qui est  infinie, j'avance
		for n := 2; i == n; n += 2 { // a chaque fois que i atteint la valeur de n ont ajoute 2
			if i != tours[2] || i != tours[n] { // si i est différent de tours[2], soit le troixième tours étant donner que tours[0] == 1, la ont part de zéro et les tours dans un jeu eux commence a 1, soit le premier tour pas de tour zéro
				fmt.Println(tours[i+1])                                      // ont donne le tours
				fmt.Print(p.nom, " attack ", p.monstre.nom, "et lui reste ") // lrye nom du joueurs qui attack le nom du monstre
				p.attack()                                                   // attack du joueur fait sur le monstre
				fmt.Print(p.monstre.pvmonstreactuel)                         // et ont print les pv du monstre actuel
				fmt.Print("/")
				fmt.Print(p.pvmax)
				fmt.Println(" PDV")
				time.Sleep(2)
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
					time.Sleep(2)
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
}*/
var nivsup int = 100

func (m *monstre) niveaujoueur(p *personnage, x int) { // experience & niveau du joueur (il gagne des récompemses )
	p.experience += x
	fmt.Println("Tu a gagné ", x, "expérience")
	for p.experience >= nivsup {
		p.experience = p.experience - nivsup // exp du joueur
		nivsup = nivsup * 2                  // exp nécessaire pour le prochain niveau
		p.niveau += 1                        // niveau supplémentaire
		fmt.Println(p.experience, "/", nivsup)
		time.Sleep(3 * time.Second)
		fmt.Println("vous avez atteint le niveau supérieur !!")
		if p.classe == "Elfe" {
			if p.niveau == 9 {
				p.skill = append(p.skill, "Arbre de vie")
				fmt.Println("Oh toi jeune âme, qui vient de naître dans ce monde, Dieu t'offre une compétence nommé : Arbre de vie")
				fmt.Println("Cette compétence te permet de convertire ta mana en pv")
				time.Sleep(5 * time.Second)
			}
		}
		if p.classe == "Nain" {
			if p.niveau == 3 {
				p.skill = append(p.skill, "Invocation de squelette")
				fmt.Println("Oh toi jeune âme, qui vient de naître dans ce monde, Dieu t'offre une compétence nommé : Invocation de Squelette")
				fmt.Println("Tu pourra invoquer une armer de squelette pour t'aider durant un combat, et restera t'aidez tours")
				time.Sleep(5 * time.Second)
			}
		}
		if p.niveau%2 == 0 {
			fmt.Println("Félicitation vous avez gagné un présent du Dieu Cristian •́‿•̀  ") //tout les deux niveaux une popo est donné au joueur
			if p.addInventory("popovie") == false {
				fmt.Println("Veuillez à ce que votre inventaire ne soit pas plein la prochaine fois ಠ‿ಠ ")
			}
		}
	} // affiche le nb d'exp manquant si il y'a pas de gain de niveau
	var manque = 0
	manque = nivsup - p.experience
	fmt.Println("Gain de ", p.experience, "exp, il ne manque plus que :", manque, "exp pour atteindre le prochain niveau")
}
func (p *personnage) luckofdrop(list []string) {
	var result string
	rand.Seed(time.Now().UnixNano())
	if len(p.inventaire) <= 10 {
		if len(list) == 1 {
			test := rand.Intn(1)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 2 {
			test := rand.Intn(2)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 3 {
			test := rand.Intn(3)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 4 {
			test := rand.Intn(4)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 5 {
			test := rand.Intn(5)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 6 {
			test := rand.Intn(6)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 7 {
			test := rand.Intn(7)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 8 {
			test := rand.Intn(8)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
		if len(list) == 9 {
			test := rand.Intn(9)
			result += list[test]
			time.Sleep(1 * time.Second)
			p.inventaire = append(p.inventaire, result)
			fmt.Println("Tu as drop", result, "x1")
		}
	}
}
func (p *personnage) lootmoneyrandom(g int) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(g)
	p.money += v
	fmt.Println("Tu a gagné", v, "PO")
}

var p1 personnage

func main() {
	p1.init("jackouille", "fripouille", 150, 150, 1, []string{"Fourrure de loup", "Cuir de Sanglier", "popvie"}, "coup de poing", 1000, "", "", "", 100, 100, 100, 0)
	//var p2 personnage
	//p2.init("Cristian ", "Cristian", 150, 1, 1, []string{"poison", "poison", "poison", "poison", "poison", "poison", "poison", "popovie"}, "coup de point", 100)
	//p1.displayInfo()
	//p2.displayInfo()
	//p2.poison()
	m1.initmonstre("Cristian", 40, 40, 5, 50, 100)
	p1.menu()
}
