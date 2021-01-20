package main

import (
	"os"
	"fmt"
	"strconv"
	"net"
	"math"
	"io/ioutil"
	"strings"
	"bufio"
	"io"
	"sync"
)

func getArgs() int {
	if len(os.Args) != 2 {
	fmt.Printf("Usage: go run serveur.go <portnumber> \n") // message d'erreur rappelant les attributs attendus
	os.Exit(1)
	} else {
		portNum, err := strconv.Atoi(os.Args[1])
		if err != nil {  // gestion des erreurs géneré par strconv
			fmt.Printf("Usage: go run serveur.go <portnumber> \n") 
			os.Exit(1)
		} else {
			return portNum
		}
	}
	//ne devrait jamais être atteint:
	return -1
}

func handleConnection(connection  net.Conn, connum int) {

	defer connection.Close()
	
	connReader := bufio.NewReader(connection)
	inputLine, err := connReader.ReadString('x')

	if err != nil && err != io.EOF{
		fmt.Printf("Error", err.Error())
	}
	graphe := readGraphe(string(inputLine))
	
	//lancement des calculs grace à Dijkstra
	res := make([][]string, len(graphe))
	for i := 0; i<len(graphe); i++{
		res[i] = make([]string, len(graphe))
	}
	var wg sync.WaitGroup
	for i := 0; i < len(res); i++{
		for j := 0; j < len(res[i]); j++{
			if i != j{
				wg.Add(1)						// Increment the WaitGroup counter
				depart:=i
				arrivee:=j
				go func(){
					defer wg.Done()				// Decrement the counter when the goroutine completes
					res[depart][arrivee] = dijkstra(graphe,depart,arrivee)
				}()
			}
		}
	}
	wg.Wait()

	var result_str string
	for i:=0; i<len(res); i++{
		result_str += fmt.Sprint("\r")
		for j:=0; j<len(res[i]); j++{
			result_str += fmt.Sprint(res[i][j])
			result_str += fmt.Sprint("	")
		}
		result_str += fmt.Sprintln("")
	}
	result_str += fmt.Sprint("x")

	//Affichage du résultat dans un fichier resultat.txt
	resultat := []byte(result_str)
	err = ioutil.WriteFile("resultat.txt", resultat, 0644)
	if err != nil{
		panic(err)
	}

	 // sample process for string received  
     msg_return := strings.ToUpper(result_str) 
     // send new string back to client  
     connection.Write([]byte(msg_return)) 
}


func main() {
	port := getArgs()   
	portString :=fmt.Sprintf(":%s", strconv.Itoa(port)) //Creation du port string
	ln, err := net.Listen("tcp", portString)
	if err != nil {  // gestion des erreurs génerées par Listen
		panic(err) 
	}

	// si sortie du if précedent, le programme n'a pas panic et ln est valide.

	connum := 1

	for {
		fmt.Printf("#DEBUG MAIN Accepting next connection\n")
		conn, errcon := ln.Accept()

		if errcon != nil {
			fmt.Printf("#DEBUG MAIN Error when accepting next connection\n")
			panic(errcon)
		}

		//Si on arrive ici, la connexion con est valide et le programme n'a pas panic

		go handleConnection(conn, connum)
		connum += 1

	}
}

func readGraphe(graphe string) [][]int{

	ln := strings.Split(graphe, "\n")				//traitement ligne par ligne

	//graphe au format integer
	nbSommets := len(ln)-1
	grapheInt := make([][]int, nbSommets)

	//conversion du graphe
	for i := 0; i<nbSommets; i++{
		nb := strings.Split(ln[i], "	")				//traitement valeur par valeur
		grapheInt[i] = make([]int, nbSommets)
		for j := 0; j<nbSommets; j++{
			value_str := strings.Split(nb[j], "")		//récupération de la valeur seule, sans \r ou autres choses cachées 
			var value int
			if len(value_str) > 1{			//si nous sommes en lecture de fichier
				value, _ = strconv.Atoi(value_str[1])		//la valeur qui nous intéresse se trouve en deuxième position du tableau value_str
			}else if len(value_str) != 0{	//si nous sommes en graphe aléatoire
				value, _ = strconv.Atoi(value_str[0])
			}
			grapheInt[i][j] = value
		}
	}

	return grapheInt
}

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


func dijkstra(graph [][]int, depart int, arrivee int) string{
//Initialisation
	//Liste des sommets
	var listeSommets = make([]Sommet, len(graph))

	//Liste des sommets à traiter
	var listeSommetsATraiter = make([]Sommet, len(graph))

	//initialisation des attributs de chaque sommet
	for i := 0; i<len(graph); i++{
		listeSommets[i].dist = math.Inf(1)					//assigne une distance infinie à chaque sommet
		listeSommets[i].id = i
		listeSommets[i].poids = make([]float64, len(graph))	//il est nécessaire de créer toutes les cases du tableau de poids avant, car les indices seront importants
		listeSommets[i].pred = -1

		//on initialise de la même façon la liste de sommets à traiter
		listeSommetsATraiter[i].dist = math.Inf(1)
		listeSommetsATraiter[i].id = i
		listeSommetsATraiter[i].poids = make([]float64, len(graph))
		listeSommetsATraiter[i].pred = -1
	}
	for i := 0; i<len(graph); i++{
		for j := 0; j<len(graph); j++{
			if graph[i][j] != 0{
				listeSommets[i].listeVoisins = append(listeSommets[i].listeVoisins, j)	//création d'un tableau de voisins
				listeSommets[i].poids[j] = float64(graph[i][j])							//référencement du poids de chaque sommet avec ses voisins
				
				//de même avec la liste de sommets à traiter
				listeSommetsATraiter[i].listeVoisins = append(listeSommetsATraiter[i].listeVoisins, j)
				listeSommetsATraiter[i].poids[j] = float64(graph[i][j])
			}
		}
	}
	listeSommets[depart].dist = 0			//on assigne une distance nulle au sommet de départ 
	listeSommetsATraiter[depart].dist = 0


//Traitement

	//Tant que la liste des sommets à traiter n'est pas vide, faire : 	
	for{

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

		//Mise à jour des distances 
		for i := 0; i<len(listeSommets[s1].listeVoisins); i++{

			s2 := listeSommets[s1].listeVoisins[i]

			if listeSommets[s2].dist > listeSommets[s1].dist + listeSommets[s1].poids[s2]{ 		//si la distance de début à s2 est plus grande que celle de début à s1 + celle de s1 à s2
				listeSommets[s2].dist = listeSommets[s1].dist + listeSommets[s1].poids[s2]		//alors on prend ce nouveau chemin qui est plus court
				listeSommets[s2].pred = listeSommets[s1].id										//et on note par où on passe

				//on actualise aussi notre liste de sommets à traiter
				for j := 0; j < len(listeSommets); j++{
					if listeSommetsATraiter[j].id == s2{							//recherche de la position du sommet à modifier dans la liste de sommets à traiter
						listeSommetsATraiter[j].dist = listeSommets[s2].dist
						listeSommetsATraiter[j].pred = listeSommets[s2].pred
						break
					}
				}
			}
		}
		
		//enfin, on peut supprimer le sommet sur lequel nous nous trouvons de la liste des sommets encore à traiter
		listeSommetsATraiter = remove(listeSommetsATraiter, indexListeATraiter)

		if len(listeSommetsATraiter) == 0 {				//permet de sortir de la boucle lorsque la liste des sommets à traiter est vide
			break
		}
	}


//Traçage du chemin

	//meilleur chemin de l'arrivée au départ
	var bestWayUpsideDown []int
	
	//déclaration des variables utiles
	s := arrivee
	fin := 100000

	//en partant de l'arrivée, nous remontons le chemin jusqu'au départ par les prédécesseurs
	for i := 0; i<fin; i++{
		bestWayUpsideDown = append(bestWayUpsideDown, s)

		if listeSommets[s].pred != -1{ 		//si un sommet n'a pas de prédécesseur, nous retournons une erreur
			s = listeSommets[s].pred
		}else{
			fmt.Println("Erreur, il n'y a pas de prédécesseur disponible...")
			break
		}

		//Lorsqu'on arrive au départ, on sort de la boucle
		if s == depart || len(bestWayUpsideDown)>len(listeSommets){
			bestWayUpsideDown = append(bestWayUpsideDown, s) 		//on ajoute quand même le dernier sommet, i.e le sommet de départ
			fin = -1
		}
	}
	
	//finalement, on remet le chemin à l'endroit
	var bestWay = make([]int, len(bestWayUpsideDown))

	for i := 0; i<len(bestWay); i++{
		bestWay[i] = bestWayUpsideDown[len(bestWayUpsideDown)-1-i]
	}
	
	//et on renvoie le résultat dans un string
	res  := "[" + strconv.Itoa(bestWay[0])
	for i:=1; i<len(bestWay); i++{
		res += ";" + strconv.Itoa(bestWay[i])
	}
	res += "] d=" + strconv.FormatFloat(listeSommets[arrivee].dist, 'g', 1, 64)

	return res
}