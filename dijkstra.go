package main

import (
	"fmt"
	"math/rand"
	"math"
)


//déclaration de la structure Sommet
type Sommet struct {
	poids[] float64			//Distance entre deux sommets
	dist float64 			//Distance du départ au sommet
	id int					//Identifiant du sommet
	listeVoisins []int		//Listes des sommets voisins
	pred int				//Identifiant du sommet précédent dans le déroulé de l'algorithme 

}

//fonction qui permet de retirer un élément d'une liste
func remove(slice []Sommet, s int) []Sommet {
    return append(slice[:s], slice[(s+1):]...)
}

// un noeud est représenté par une matrice d'adjacence (id : numéro colonne/ligne)
func dijstra(graph [][]int, depart int, arrivee int) {

//Initialisation

	//Liste des sommets
	var listeSommets = make([]Sommet, len(graph))

	//initialisation des attributs de chaque sommet
	for i := 0; i<len(graph); i++{
		listeSommets[i].dist = math.Inf(1)					//assigne une distance infinie à chaque sommet
		listeSommets[i].id = i
		listeSommets[i].poids = make([]float64, len(graph))	//il est nécessaire de créer toutes les cases du tableau de poids avant, car les indices seront importants
		listeSommets[i].pred = -1
	}
	for i := 0; i<len(graph); i++{
		for j := 0; j<len(graph); j++{
			if graph[i][j] != 0{
				listeSommets[i].listeVoisins = append(listeSommets[i].listeVoisins, j)	//création d'un tableau de voisins
				listeSommets[i].poids[j] = float64(graph[i][j])							//référencement du poids de chaque sommet avec ses voisins
			}
		}
		fmt.Printf("Liste des voisins de %d", i)
		fmt.Printf(" : %v\n", listeSommets[i].listeVoisins)
		fmt.Printf("Liste des poids de %d", i)
		fmt.Printf(" : %v\n", listeSommets[i].poids)
	}
	listeSommets[depart].dist = 0			//on assigne une distance nulle au sommet de départ 

	//Liste des sommets à traiter
	var listeSommetsATraiter = listeSommets

	//Liste des sommets déjà traités
	//var listeSommetsDejaTraites = make([]Sommet, 0)

	//ID sommet traité actuellement
	//var idX	int	//id point x
	//var idY	int	//id point y


//Traitement

	//Tant que la liste des sommets à traiter n'est pas vide, faire : 	
	for{
		fmt.Printf("Liste des sommets à traiter : %v\n", listeSommetsATraiter)

		//initialision des variables utiles 
		min := math.Inf(1)
		s1 := -1
		indexListeATraiter := -1

		//recherche du sommet qui possède la distance minimale, parmis les sommets encore à traiter
		for i := 0; i<len(listeSommetsATraiter); i++{
			if listeSommetsATraiter[i].dist < min{
				min = listeSommetsATraiter[i].dist
				s1 = listeSommetsATraiter[i].id
				indexListeATraiter = i					//important de conserver l'index du sommet dans la liste listeSommetsATraiter, pour y accéder plus tard
			}
		}
		fmt.Printf("min : %d\n", min)
		fmt.Printf("s1 : %d\n", s1)

		//Mise à jour des distances 
		fmt.Println("Avant la boucle")
		for i := 0; i<len(listeSommets[s1].listeVoisins); i++{
			s2 := listeSommets[s1].listeVoisins[i]
			if listeSommets[s2].dist > listeSommets[s1].dist + listeSommets[s1].poids[s2]{ 		//si la distance de début à s2 est plus grande que celle de début à s1 + celle de s1 à s2
				listeSommets[s2].dist = listeSommets[s1].dist + listeSommets[s1].poids[s2]		//alors on prend ce nouveau chemin qui est plus court
				listeSommets[s2].pred = s1														//et on note par où on passe
				listeSommetsATraiter[indexListeATraiter] = listeSommets[s2]						//on actualise aussi notre liste de sommets à traiter
				fmt.Printf("pred listeSommets : %d\n", listeSommets[s2].pred)
				fmt.Printf("pred listeSommetsATraiter : %d\n", listeSommetsATraiter[indexListeATraiter].pred)
			}
		}
		fmt.Println("Après la boucle")
			
		listeSommetsATraiter = remove(listeSommetsATraiter, indexListeATraiter)					//enfin, on peut supprimer le sommet sur lequel nous nous trouvons de la liste des sommets encore à traiter

		if len(listeSommetsATraiter) == 0 {				//permet de sortir de la boucle lorsque la liste des sommets à traiter est vide
			break
		}
	}


//Traçage du chemin

	//meilleur chemin
	var bestWay = make([]int, len(listeSommets))
	
	//déclaration des variables utiles
	s := arrivee
	fin := 100000

	//en partant de l'arrivée, nous remontons le chemin jusqu'au départ par les prédécesseurs
	for i := 0; i<fin; i++{
		bestWay[i] = s

		if listeSommets[s].pred != -1{ 		//si un sommet n'a pas de prédécesseur, nous retournons une erreur
			s = listeSommets[s].pred
		}else{
			fmt.Println("Erreur, il n'y a pas de prédécesseur disponible...")
			break
		}

		fmt.Printf("Chemin le plus court : %v\n", bestWay)

		//Lorsqu'on arrive au départ, on sort de la boucle
		if s == depart{
			fin = -1
		}
	}
	fmt.Printf("Finalement, le chemin le plus court est : %v\n", bestWay)

	/*for {
		fmt.Println(listeSommetsATraiter)
		for  i := 0; i<len(listeSommetsATraiter); i++ {															//on parcours chaque sommet a traiter
			idX = listeSommetsATraiter[i].id
			for j := 0; j<len(listeSommetsATraiter); j++{															
				if j != i && graph[i][j] != 0 && float64(graph[i][j]) <= listeSommetsATraiter[i].poids{
					listeSommets[listeSommetsATraiter[i].id].poids = listeSommetsATraiter[i].poids
					//listeSommetsATraiter[listeSommetsATraiter[i].id].poids = listeSommetsATraiter[i].poids
					//listeSommetsDejaTraites = append(listeSommetsDejaTraites, listeSommetsATraiter[i])
					listeSommetsATraiter = remove(listeSommetsATraiter, i)
				}

				for k := 0; k<len(listeSommets[idX].listeVoisins); k++{ 							//
					for l := 0; l<len(listeSommetsATraiter); l++{
						if listeSommets[idX].listeVoisins[k] == listeSommetsATraiter[l].id {
							idY = listeSommetsATraiter[l].id
							if listeSommets[idY].dist > listeSommets[idX].dist + float64(graph[idX][idY]) {
								listeSommets[idY].dist = listeSommets[idX].dist + float64(graph[idX][idY])
								listeSommets[idY].pred = idX
							}
							
						}
					}
				}

			}
		}
		if len(listeSommetsATraiter) == 0 {				//permet de sortir de la boucle lorsque la liste des sommets à traiter est vide
			break
		}
	}*/
}

func main() {
	const nbSommets int = 5

	var grapheNoeudsOriginal = make([][]int, nbSommets)
	for i := 0; i<nbSommets; i++{
		grapheNoeudsOriginal[i] = make([]int, nbSommets)
	}

	//Generateur de matrice d'adjacence
	for i := 0; i<nbSommets; i++{
		for j := 0; j<nbSommets; j++{
			if j != i{
				var a = rand.Intn(5)
				if j >= i {
					grapheNoeudsOriginal[i][j] = a
				} else {
					grapheNoeudsOriginal[i][j] = grapheNoeudsOriginal[j][i]
				}
			}else{
				grapheNoeudsOriginal[i][j] = 0
			} 
			fmt.Print(grapheNoeudsOriginal[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	dijstra(grapheNoeudsOriginal, 1, 4)
}