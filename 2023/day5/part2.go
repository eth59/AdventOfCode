package main

type SeedRange struct {
	start, end int
}

func partTwo(input string) int {
	seeds, maps := parse(input)
	nbSeeds := len(seeds)

	seedRanges := make([]SeedRange, 0, nbSeeds/2)
	for i := 0; i < nbSeeds; i += 2 {
		seedRanges = append(seedRanges, SeedRange{seeds[i], seeds[i]+seeds[i+1]})
	}

	for _, m := range maps {
		var nextRanges []SeedRange

		for len(seedRanges) > 0 {
			curr := seedRanges[0]
			seedRanges = seedRanges[1:]

			validInter := false

			for _, r := range m {
				interStart := max(curr.start, r.sourceStart)
				interEnd := min (curr.end, r.sourceStart + r.rangeLength)

				if interStart < interEnd {
					validInter = true

					// partie qui match
					offset := r.destStart - r.sourceStart
					nextRanges = append(nextRanges, SeedRange{interStart+offset, interEnd+offset})

					// partie avant l'inter
					if curr.start < interStart {
						seedRanges = append(seedRanges, SeedRange{curr.start, interStart})
					}

					// partie après l'inter
					if curr.end > interEnd {
						seedRanges = append(seedRanges, SeedRange{interEnd, curr.end})
					}

					break
				}
			}

			if !validInter {
				nextRanges = append(nextRanges, curr)
			}
		}
		seedRanges = nextRanges
	}

	minVal := seedRanges[0].start
	for _, r := range seedRanges {
		if r.start < minVal {
			minVal = r.start
		}
	}
	
	return minVal
}
