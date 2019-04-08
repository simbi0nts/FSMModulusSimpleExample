package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// StrbinZero == str(bin(0)) == 48
	StrbinZero = 48
	// StrbinOne == str(bin(1)) == 49
	StrbinOne = 49
)

func main() {
	// Parse args from console
	if len(os.Args) < 3 {
		fmt.Printf("Not enough arguments\n")
		return
	}

	// First argument: numerator
	origNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("First argument is not a number\n")
		return
	}

	// Second argument: denominator
	divRoot, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Second argument is not a number\n")
		return
	}

	// Making FSM dict, like
	// fsmDict = {
	// 	0: {0: "0", 1: "1",}
	// 	1: {2: "0", 3: "1",}
	// 	2: {4: "0", 5: "1",}
	// 	...
	// 	n: {((n*2) mod n): "0", ((n*2 + 1) mod n): "1",}
	// }

	fsmDct := map[string]map[string]string{}
	nextIdx := 0

	for i := 0; i < divRoot; i++ {
		idx := strconv.Itoa(i)
		idx1 := strconv.Itoa(nextIdx)
		idx2 := strconv.Itoa(nextIdx + 1)

		fsmDct[idx] = map[string]string{
			idx1: "0",
			idx2: "1",
		}

		nextIdx += 2
		if nextIdx-divRoot == 0 {
			nextIdx = 0
		} else if nextIdx-divRoot == 1 {
			nextIdx = 1
		}
	}

	// int -> bin -> str
	origNumBinStr := fmt.Sprintf("%b", origNum)

	// Go through all of the bits
	curPos := fsmDct["0"]
	lastPosIdx := "0"
	rLit := "0"

	for _, lit := range origNumBinStr {
		// rune code check
		if lit == StrbinZero {
			rLit = "0"
		} else if lit == StrbinOne {
			rLit = "1"
		}

		for k, v := range curPos {
			if v == rLit {
				lastPosIdx = k
				curPos = fsmDct[k]
				break
			}
		}
	}

	// Print result
	s := fmt.Sprintf("%d mod %d = %s \n", origNum, divRoot, lastPosIdx)
	fmt.Printf(s)
}
