package main

import (
	"fmt"
	"math/rand"
	"math"
)


//déclaration de la structure Sommet
type Sommet struct {
	poids float64			//Distance entre deux sommets
	dist float64 			//Distance du départ au sommet
	id int					//Identifiant du sommet
	listeVoisins []int		//Listes des sommets voisins
	pred int				//Identifiant du sommet précédent dans le déroulé de l'algorithme 

}

// un noeud est représenté par une matrice d'adjacence (id : numéro colonne/ligne)
func dijstra(graph [][]int, depart int, arrivee int) {

//Initialisation

	//Liste des sommets
	var listeSommets = make([]Sommet, len(graph))
	for i := 0; i<len(graph); i++{
		listeSommets[i].poids = math.Inf(1)					//assigne un poids infini à chaque sommet
		listeSommets[i].id = i
		/*if listeSommets[i].poids == float64(depart) {			//listeSommets[i].id plutôt que listeSommets[i].poids non ? Obligé de mettre float64 du coup ?
			listeSommets[i].poids = 0							//assigne un poids nul au sommet de départ
		}*/
		for j := 0; j<len(graph); j++{
			if graph[i][j] != 0{
				listeSommets[i].listeVoisins = append (listeSommets[i].listeVoisins, j)
			}
		}
		fmt.Println(listeSommets[i].listeVoisins)
	}
	listeSommets[depart].poids = 0

	//Liste des sommets à traiter
	var listeSommetsATraiter = listeSommets

	//Liste des sommets déjà traités
	//var listeSommetsDejaTraites = make([]Sommet, 0)

	//ID sommet traité actuellement
	var idX	int	//id point x
	var idY	int	//id point y


//Traitement

	for {
		fmt.Println(listeSommetsATraiter)
		for  i := 0; i<len(listeSommetsATraiter); i++ {															//on parcours chaque sommet a traiter
			idX = listeSommetsATraiter[i].id
			for j := 0; j<len(listeSommetsATraiter); j++{															
				if j != i && graph[i][j] != 0 && float64(graph[i][j]) <= listeSommetsATraiter[i].poids{
					listeSommets[listeSommetsATraiter[i].id].poids = listeSommetsATraiter[i].poids
					//listeSommetsATraiter[listeSommetsATraiter[i].id].poids = listeSommetsATraiter[i].poids
					//listeSommetsDejaTraites = append(listeSommetsDejaTraites, listeSommetsATraiter[i])
					remove(listeSommetsATraiter, i)
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
	}


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
	dijstra(grapheNoeudsOriginal, 1, 5)
}

func remove(slice []Sommet, s int) []Sommet {
    return append(slice[:s], slice[s+1:]...)
}