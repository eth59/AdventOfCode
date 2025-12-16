package main

func partOne(devices map[string][]string) int {
	savedPaths := make(map[string]int)
	for name := range devices {
		savedPaths[name] = -1
	}
	res := solverPartOne("you", devices, savedPaths)
	return res
}

func solverPartOne(name string, devices map[string][]string, savedPaths map[string]int) (res int) {
	for _, output := range devices[name] {
		if output == "out" {
			res++
		} else {
			if savedPaths[output] != -1 {
				res += savedPaths[output]
			} else {
				res += solverPartOne(output, devices, savedPaths)
			}
		}
	}
	savedPaths[name] = res
	return res
}