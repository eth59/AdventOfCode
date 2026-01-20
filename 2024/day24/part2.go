package main

import (
	"fmt"
	"slices"
	"strings"
)

func partTwo(input string) string {
	parts := strings.Split(input, "\n\n")
	nbBitsInput := len(strings.Split(parts[0], "\n")) / 2
	gateMap := mapGatesByOutput(parts[1])
	gatesToSwitch := make([]string, 0)

	// check z00
	if !checkGateTypeA(gateMap, "z00", 0) {
		gatesToSwitch = append(gatesToSwitch, "z00")
	}

	// check z01
	if nbBitsInput > 0 {
		gatesToSwitch = append(gatesToSwitch, checkZ01(gateMap)...)
	}

	// check zi
	for i := 2; i < nbBitsInput; i++ {
		gatesToSwitch = append(gatesToSwitch, checkZi(gateMap, i)...)
	}

	// check last carry
	zlast := fmt.Sprintf("z%02d", nbBitsInput)
	if !checkOpGate(gateMap, zlast, "OR") {
		gatesToSwitch = append(gatesToSwitch, zlast)
	}

	// remove x & y
	filteredGatesToSwitch := make([]string, 0, len(gatesToSwitch))
	for _, g := range gatesToSwitch {
		if !(strings.HasPrefix(g, "x") || strings.HasPrefix(g, "y")) {
			filteredGatesToSwitch = append(filteredGatesToSwitch, g)
		}
	}

	slices.Sort(filteredGatesToSwitch)

	return strings.Join(filteredGatesToSwitch, ",")
}

func mapGatesByOutput(input string) (gateMap map[string]Gate) {
	gateMap = make(map[string]Gate)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		opParts := strings.Fields(parts[0])
		gateMap[parts[1]] = Gate{opParts[0], opParts[1], opParts[2], parts[1]}
	}
	return
}

func checkZ01(gateMap map[string]Gate) (gatesToSwitch []string) {
	gatesToSwitch = make([]string, 0)
	entryA := gateMap["z01"].entryA
	entryB := gateMap["z01"].entryB
	// une entrée doit être type B à i=0 (Cin) 
	// l'autre c'est une type A
	// les deux sont Xorées
	AisA := checkGateTypeA(gateMap, entryA, 1)
	AisCin := checkGateTypeB(gateMap, entryA, 0)
	BisA := checkGateTypeA(gateMap, entryB, 1)
	BisCin := checkGateTypeB(gateMap, entryB, 0)

	if !(AisA || AisCin) {
		gatesToSwitch = append(gatesToSwitch, entryA)
	}
	if !(BisA || BisCin) {
		gatesToSwitch = append(gatesToSwitch, entryB)
	}
	if !checkOpGate(gateMap, "z01", "XOR") {
		gatesToSwitch = append(gatesToSwitch, "z01")
	}
	return
}

func checkZi(gateMap map[string]Gate, i int) (gatesToSwitch []string) {
    gatesToSwitch = make([]string, 0)
    zi := fmt.Sprintf("z%02d", i)

	if !checkOpGate(gateMap, zi, "XOR") {
        gatesToSwitch = append(gatesToSwitch, zi)
		return
    }

    entryA := gateMap[zi].entryA
    entryB := gateMap[zi].entryB

    AisA := checkGateTypeA(gateMap, entryA, i)
    AisOr := checkOpGate(gateMap, entryA, "OR")
	BisA := checkGateTypeA(gateMap, entryB, i)
    BisOr := checkOpGate(gateMap, entryB, "OR")

    if AisOr {
        gatesToSwitch = append(gatesToSwitch, checkCin(gateMap, entryA, i)...)
    }
    if BisOr {
        gatesToSwitch = append(gatesToSwitch, checkCin(gateMap, entryB, i)...)
    }

    if !(AisA || AisOr) {
        gatesToSwitch = append(gatesToSwitch, entryA)
    }
    if !(BisA || BisOr) {
        gatesToSwitch = append(gatesToSwitch, entryB)
    }

    return
}

func checkCin(gateMap map[string]Gate, output string, i int) (gatesToSwitch []string) {
	// B[i-1] OR (Cin[i-1] AND A[i-1])
	// on check pas les valeurs dans le AND
	gatesToSwitch = make([]string, 0)

	entryA := gateMap[output].entryA
	entryB := gateMap[output].entryB

	AisB := checkGateTypeB(gateMap, entryA, i-1)
	AisAnd := checkOpGate(gateMap, entryA, "AND")
	BisB := checkGateTypeB(gateMap, entryB, i-1)
	BisAnd := checkOpGate(gateMap, entryB, "AND")

	if !(AisB || AisAnd) {
		gatesToSwitch = append(gatesToSwitch, entryA)
	}
	if !(BisB || BisAnd) {
		gatesToSwitch = append(gatesToSwitch, entryB)
	}
	if !checkOpGate(gateMap, output, "OR") {
		gatesToSwitch = append(gatesToSwitch, output)
	}
	return
}

// x XOR y
func checkGateTypeA(gateMap map[string]Gate, output string, i int) bool {
	x := fmt.Sprintf("x%02d", i)
	y := fmt.Sprintf("y%02d", i)
	return checkEntriesGate(gateMap, output, x, y) && checkOpGate(gateMap, output, "XOR")
}

// x AND y
func checkGateTypeB(gateMap map[string]Gate, output string, i int) bool {
	x := fmt.Sprintf("x%02d", i)
	y := fmt.Sprintf("y%02d", i)
	return checkEntriesGate(gateMap, output, x, y) && checkOpGate(gateMap, output, "AND")
}

func checkEntriesGate(gateMap map[string]Gate, output, entryA, entryB string) bool {
	g := gateMap[output]
	return (g.entryA == entryA && g.entryB == entryB) || (g.entryA == entryB && g.entryB == entryA)
}

func checkOpGate(gateMap map[string]Gate, output, op string) bool {
	return gateMap[output].operation == op
}