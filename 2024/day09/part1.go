package main

func partOne(input string) (res int) {
	// déclaration des variables
	var (
		i int // parcours par la gauche
		j int = len(input)-3 // parcours par la droite
		iId int // id des fichiers à i
		jId int = len(input)/2 // id des fichiers à j
		index int // index du résultat pour calculer le checksum à renvoyer
		currJ int = int(input[j+2] - '0') // garde le nombre de fichiers avec comme id jId à mettre dans les trous
	)

	for {
		// on lit le caractère à i
		for k := 0; k < int(input[i] - '0'); k++ {
			res += iId * index
			index++
		}
		i++
		iId++

		// condition d'arrêt
		if iId == jId {
			for k := 0; k < currJ; k++ {
				res += jId * index
				index++
			}
			return
		}

		// on remplit le trou suivant par ce qui est à j
		for k := 0; k < int(input[i] - '0'); k++ {
			if currJ == 0 { // dans ce cas là il faut lire les fichiers à la fin de l'input
				currJ = int(input[j] - '0')
				j -= 2 // on décrémente de 2 car on saute l'espace vide
				jId--
			}

			// condition d'arrêt
			if jId < iId {
				return
			}

			res += jId * index
			index++
			currJ--
		}
		i++
	}
}