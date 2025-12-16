package main

type State struct {
	name string
	dac, fft bool
}

func partTwo(devices map[string][]string) int {
	savedPaths := make(map[State]int)
	for name := range devices {
		savedPaths[State{name, false, false}] = -1
		savedPaths[State{name, true, false}] = -1
		savedPaths[State{name, false, true}] = -1
		savedPaths[State{name, true, true}] = -1
	}
	res := solverPartTwo("svr", devices, savedPaths, false, false)
	return res
}

func solverPartTwo(name string, devices map[string][]string, savedPaths map[State]int, hasVisitedDac, hasVisitedFFT bool) (res int) {
	for _, output := range devices[name] {
		if output == "out" {
			if hasVisitedDac && hasVisitedFFT {
				res++
			}
		} else {
			currentState := State{output, hasVisitedDac, hasVisitedFFT}
			switch output {
			case "dac":
				currentState.dac = true
			case "fft":
				currentState.fft = true
			}
			if savedPaths[currentState] != -1 {
				res += savedPaths[currentState]
			} else {
				res += solverPartTwo(output, devices, savedPaths, currentState.dac, currentState.fft)
			}
		}
	}
	savedPaths[State{name, hasVisitedDac, hasVisitedFFT}] = res
	return
}