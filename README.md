# golang-project

## Préliminaires
Avant de démarrer la simulation, il est **important** de remplacer l'adresse IP mentionnée dans le fichier *client.go* à la **ligne 85** par la vôtre. Pour la trouver, il suffit de taper *ifconfig* dans votre terminal. 

## Connection
Pour élaborer la connection, lancer d'abord le serveur avec la commande *go run serveur.go num_port*.  
Puis, lancer le client avec la commande *go run client.go num_port*.  
Si l'adresse IP est bien modifiée, la connection devrait se faire. 

## Démarrage
Pour démarrer la simulation, il y a plusieurs possibilités (à indiquer dans le client):
* Simuler à partir d'un fichier *graphe.txt*. Pour le générer, il est tout à fait possible de remplir un tableau excel, puis de l'enregistrer au format *.txt*. 
* Simuler à partir d'un graphe aléatoire. Dans ce cas, il vous sera demandé le nombre de sommets à créer aléatoirement. Il est à noter que les distances entre chaque sommet sont choisies aléatoirement dans des valeurs allant de 0 à 20. 

## Quitter le programme
Pour quitter, il suffit de taper *ctrl+c* dans le serveur

## Points faibles
Nous ne sommes malheureusement pas parvenus à renvoyer au client les valeurs des chemins les plus courts, celui-ci attendant de quitter la connection pour envoyer ses propres valeurs. À défaut de cette fonctionnalité, le résultat est envoyé dans un fichier *resultat.txt*. 