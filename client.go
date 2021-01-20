package main

import (
	"os"
	"fmt"
	"strconv"
	"net"
	"io/ioutil"
	"io"
	"math/rand"
	"time"
	"bufio"
)

func getArgs() int {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run client.go <portnumber>\n") // message d'erreur rappelant les attributs attendus
		os.Exit(1)
	} else {

		fmt.Printf("DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNum, err := strconv.Atoi(os.Args[1])
		if err != nil {   // gestion des erreurs géneré par strconv
			fmt.Printf("Usage: go run serveur.go <portnumber> \n")
			os.Exit(1)
		} else {
			return portNum
		}
	}
	//ne devrait jamais être atteint:
	return -1
}

func handleConnection(connection  net.Conn) {
	
	defer connection.Close()
	fmt.Printf("#DEBUG MAIN Connected\n")
	fmt.Println("SELECT MODE : \n1 - Random graph\n2 - Read file")
		
	//declare un reader pour lire un input dans la console
	var input int
	var selectionErrorWarning = false

	//définit ce qui se passe en fonction de l'input
	for {
		//lit un input de l'utilisateur
		fmt.Scan(&input)
		if (input == 1){
			fmt.Println("Construction et envoi d'un graphe aléatoire")
			fmt.Println("SELECT NODE QUANTITY :")
			var nbSommets int
			fmt.Scan(&nbSommets)
			graphe := menuRandomGraph(nbSommets)
			fmt.Println(graphe)
			io.WriteString(connection, graphe)
			break
		} else if (input == 2){
			fmt.Println("Envoi du fichier graphe.txt")
			graphe := readFile( "graphe.txt")
			io.WriteString(connection, graphe)
			break
		} else {
			if (selectionErrorWarning != true){
				fmt.Println("ERROR. PLEASE SELECT ONE OF THE OPTIONS ABOVE.")
				selectionErrorWarning = false
			}
		}
	}

	time.Sleep(100)
	fmt.Println("Après le sleep")
	Reader := bufio.NewReader(connection)
	inputLine, err := Reader.ReadString('x')
	fmt.Println(inputLine)

	if err != nil && err != io.EOF{
		fmt.Printf("Error", err.Error())
	}

}


func main() {

	port := getArgs()
	
	portString := fmt.Sprintf("192.168.144.21:%s", strconv.Itoa(port))

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n")
		os.Exit(1)
	} else {
		// ici on rajoute ce qu'on veut envoyer 
		handleConnection(conn)
	}
}

func readFile(fn string ) string {
	//ouverture du fichier
	file, err := ioutil.ReadFile("graphe.txt")
	if err != nil{
		fmt.Println(err)
	}

	//récupération du graphe en format string
	graphestr := string(file)
	return graphestr
}

func menuRandomGraph(nbSommets int) string{

	var grapheInit [][]int
	var res string
    grapheInit = creerGrapheAleatoire(nbSommets)
    for i := 0; i<len(grapheInit); i++{
		res += fmt.Sprint("\r")
		for j := 0; j<len(grapheInit[i]); j++{
			res += fmt.Sprint(grapheInit[i][j])
			res += fmt.Sprint("	")
		}
		res += fmt.Sprintln("")
	}
	res += fmt.Sprint("x")
	return res
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
				var a = rand.Intn(20)
				if j >= i {
					graphe[i][j] = a
				} else {
					graphe[i][j] = graphe[j][i]
				}
			}else{
				graphe[i][j] = 0
			} 
		}
	}

	return graphe
}