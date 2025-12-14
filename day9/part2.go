package main

import (
	"strconv"
	"strings"
)

func partTwo(input string) int {
	// parsing
	lines := strings.Split(input, "\n")
	var redTiles []coords
	for _, line := range lines {
		c := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(c[0], 64)
		y, _ := strconv.ParseFloat(c[1], 64)
		redTiles = append(redTiles, coords{x, y})
	}

	nbRedTiles := len(redTiles)
	var maxArea float64
	var minX, maxX, minY, maxY float64

	// on parcourt les couples de red tiles
	for i := 0; i < nbRedTiles; i++ {
		for j := i+1; j < nbRedTiles; j++ {
			tile1 := redTiles[i]
			tile2 := redTiles[j]

			// calcul des points du rectangle
			if tile1.x < tile2.x {
				minX, maxX = tile1.x, tile2.x
			} else {
				minX, maxX = tile2.x, tile1.x
			}

			if tile1.y < tile2.y {
				minY, maxY = tile1.y, tile2.y
			} else {
				minY, maxY = tile2.y, tile1.y
			}

			// calcul de l'aire
			area := (maxX - minX + 1)*(maxY - minY + 1)

			// vérification du rectangle
			if isValidRectangle(minX, maxX, minY, maxY, redTiles) {
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return int(maxArea)
}

func isValidRectangle(minX, maxX, minY, maxY float64, redTiles []coords) bool{
	// on vérifie que le centre soit dans le polygone
	midX := (maxX + minX) / 2
	midY := (maxY + minY) / 2
	if !isPointInPolygon(midX, midY, redTiles) {
		return false
	}

	// on teste tous les bords du polygone
	nbRedTiles := len(redTiles)
	for i := 0; i < nbRedTiles; i++ {
		point1 := redTiles[i]
		point2 := redTiles[(i+1)%nbRedTiles]

		if segmentIntersectsRectangleInterior(point1, point2, minX, maxX, minY, maxY) {
			return false
		}
	}

	return true
}

// ray casting
func isPointInPolygon(x, y float64, poly []coords) (inside bool) {
	n := len(poly)

	for i := 0; i < n; i++ {
		j := (i+1)%n
		xi, yi := poly[i].x, poly[i].y
		xj, yj := poly[j].x, poly[j].y

		intersect := ((yi > y) != (yj > y)) && // check mur vertical
					 (x < (xj-xi)*(y-yi)/(yj-yi)+xi) // check mur horizontal :
					 // ratio : (y-yi)/(yj-yi)
					 // par proportionnalité, on a le décalage en horizontal : (xj-xi)*ratio
					 // on y ajoute xi pour avoir la coordonnée absolue
		
		if intersect {
			inside = !inside
		}
	}
	
	return
}

func segmentIntersectsRectangleInterior(pA, pB coords, minX, maxX, minY, maxY float64) bool {
	// segment vertical
	if pA.x == pB.x {
		var segMinY, segMaxY float64
		segX := pA.x
		if segX > minX && segX < maxX {
			if pA.y < pB.y {
				segMinY, segMaxY = pA.y, pB.y
			} else {
				segMinY, segMaxY = pB.y, pA.y
			}
			
			overlapMin := max(segMinY, minY)
			overlapMax := min(segMaxY, maxY)
			
			if overlapMin < overlapMax {
				return true
			}
		}
	} else if pA.y == pB.y { // segment horizontal
		var segMinX, segMaxX float64
		if pA.y > minY && pA.y < maxY {
			if pA.x < pB.x {
				segMinX, segMaxX = pA.x, pB.x
			} else {
				segMinX, segMaxX = pB.x, pA.x
			}

			overlapMin := max(segMinX, minX)
			overlapMax := min(segMaxX, maxX)
			
			if overlapMin < overlapMax {
				return true
			}
		}
	}
	return false
}