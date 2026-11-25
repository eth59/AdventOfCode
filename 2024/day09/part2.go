package main

import "slices"

// structure pour stocker une partition
// on change par rapport à la partie 1, on utilisera un []Partition pour stocker le problème
type Partition struct {
	Size, Content int
}

func partTwo(input string) (res int) {
	originalFilesystem := make([]Partition, 0, len(input)) // original, on y touche pas
	newFilesystem := make([]Partition, 0, len(input)) // le résultat

	// parsing
	id := 0
	for i, d := range input {
		if i % 2 == 0 {
			originalFilesystem = append(originalFilesystem, Partition{int(d-'0'), id})
			newFilesystem = append(newFilesystem, Partition{int(d-'0'), id})
			id++
		} else {
			originalFilesystem = append(originalFilesystem, Partition{int(d-'0'), -1})
			newFilesystem = append(newFilesystem, Partition{int(d-'0'), -1})
		}
	}

	// on parcourt à l'envers le filesystem pour essayer de caser les fichiers dans des trous
	for i := len(originalFilesystem)-1; i > 0; i-=2 {
		// on parcourt à l'endroit le nouveau filesystem pour trouver un trou assez grand
		hasMoved := false
		for j := 0; j < len(newFilesystem); j++ {
			if newFilesystem[j].Content != -1 || newFilesystem[j].Size < originalFilesystem[i].Size {
				continue // c pas un trou ou pas assez grand
			}

			if newFilesystem[j].Content == originalFilesystem[i].Content {
				// pas de trou assez grand dispo avant la place originale du fichier
				continue
			}

			newFilesystem[j].Size -= originalFilesystem[i].Size // on réduit la taille du trou
			newFilesystem = slices.Insert(newFilesystem, j, originalFilesystem[i]) // on insère le fichier
			hasMoved = true
			break
		}

		// si le fichier a été déplacé, on supprime l'ancien, c'est le premier en partant de la fin
		if hasMoved {
			for j := len(newFilesystem)-1 ; j >= 0 ; j-- {
				if newFilesystem[j].Content == originalFilesystem[i].Content {
					newFilesystem[j].Content = -1
					break
				}
			}
		}
	}

	// on calcule le résultat
	index := 0
	for _, p := range newFilesystem {
		if p.Content != -1 {
			for i := 0; i < p.Size; i++ {
				res += p.Content * index
				index++
			}
		} else {
			index += p.Size // on avance juste l'index quand c'est un trou
		}
	}

	return
}