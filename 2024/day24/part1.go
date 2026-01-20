package main

import (
	"slices"
	"strconv"
	"strings"
)

type Gate struct {
	entryA, operation, entryB, output string
}

func partOne(input string) (res int64) {
	parts := strings.Split(input, "\n\n")
	wires := parseWires(parts[0])
	gates := parseGates(parts[1])
	wires = solve(wires, gates)
	zWires := make([]string, 0)
	for wire := range wires {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}
	slices.Sort(zWires)
	var resStr string
	for i := len(zWires)-1; i >= 0; i-- {
		if wires[zWires[i]] {
			resStr += "1"
		} else {
			resStr += "0"
		}
	}
	res, _ = strconv.ParseInt(resStr, 2, 64)
	return
}

func parseWires(input string) (wires map[string]bool) {
	wires = make(map[string]bool)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		if parts[1] == "0" {
			wires[parts[0]] = false
		} else {
			wires[parts[0]] = true
		}
	}
	return
}

func parseGates(input string) (gates []Gate) {
	lines := strings.Split(input, "\n")
	gates = make([]Gate, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		opParts := strings.Fields(parts[0])
		gates = append(gates, Gate{opParts[0], opParts[1], opParts[2], parts[1]})
	}
	return
}

func solve(wires map[string]bool, gates []Gate) map[string]bool {
	for len(gates) > 0 {
		var remainingGates []Gate
		progress := false
		for _, gate := range gates {
			valA, okA := wires[gate.entryA]
			valB, okB := wires[gate.entryB]
			if okA && okB {
				switch gate.operation {
				case "XOR":
					wires[gate.output] = (valA || valB) && !(valA && valB)
				case "OR":
					wires[gate.output] = valA || valB
				case "AND":
					wires[gate.output] = valA && valB
				}
				progress = true
			} else {
				remainingGates = append(remainingGates, gate)
			}
		}
		gates = remainingGates
		if !progress {
			break
		}
	}
	return wires
}