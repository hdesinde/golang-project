package main

import {
	


}

func getArgs() {
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

func main() {
	port := getArgs()   
	portString :=fmt.Sprintf(":%s", strconv.Itos(port)) //Creation du port string
	ln, err := net.Listen("tcp", portString)
	if err != nil {  // gestion des erreurs géneré par Listen
		panic(errconn) // 
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

func handleConccetion(connection  next.Conn, connum int) {
	// partie à remplacer par l'application de l'algo de Dijksta

}
