package main

import (
	"math/rand"
    "fmt"
)

func menuMain() {

    fmt.Println("SELECT MODE : \n1 - Random graph\n2 - Read file")

	//declare un reader pour lire un input dans la console
	var input int
	var selectionErrorWarning = false
	//var inputErrorWarning = false

	//définit ce qui se passe en fonction de l'input
	for {
		//lit un input de l'utilisateur
		fmt.Scan(&input)
	    if (input == 1){
	    	menuRandomGraph()
	    	break
	    } else if (input == 2){
	    	fmt.Println("Menu 2 debug")
	    	break
	    } else {
	    	if (selectionErrorWarning != true){
	    		fmt.Println("ERROR. PLEASE SELECT ONE OF THE OPTIONS ABOVE.")
	    		selectionErrorWarning = false
	    	}
	    }
	}
}

func menuRandomGraph() {

	var nbSommets int

	fmt.Println("SELECT NODE QUANTITY :")
    fmt.Scan(&nbSommets)
    //fmt.Println(nbSommets)
    var grapheInit [][]int
    grapheInit = creerGrapheAleatoire(nbSommets)
    for i := 0; i<len(grapheInit); i++{
		for j := 0; j<len(grapheInit[0]); j++{
			fmt.Print(grapheInit[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}



}

func creerGrapheAleatoire(nbSommets int) [][]int{

	var graphe = make([][]int, nbSommets)
	for i := 0; i<nbSommets; i++{
		graphe[i] = make([]int, nbSommets)
	}

	//Generateur de matrice d'adjacence
	for i := 0; i<nbSommets; i++{
		for j := 0; j<nbSommets; j++{
			if j != i{
				var a = rand.Intn(5)
				if j >= i {
					graphe[i][j] = a
				} else {
					graphe[i][j] = graphe[j][i]
				}
			}else{
				graphe[i][j] = 0
			} 
			//fmt.Print(graphe[i][j])
			//fmt.Print(" ")
		}
		//fmt.Println("")
	}

	return graphe
}

func main() {

	menuMain()

}