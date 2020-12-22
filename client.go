package main

import {
	


}

func getArgs() int {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run client.go <portnumber>\n") // message d'erreur rappelant les attributs attendus
		os.Exit(1)
	} else {

		fmt>.Printf("DEBUG ARGS Port Number : %s\n", os.Args[1])
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

func main() {

	port := getArgs()
	
	portString := fmt.Sprintf("127.0.0.1:%s", )
}