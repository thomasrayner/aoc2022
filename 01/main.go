package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	splitContent := strings.Split(string(content), "\n")
	overallMax := []int{}
	runningMax := 0

	for _, s := range splitContent {
		if s == "" {
			// New line means new elf, evaluate the last elf
			// Find where the new elf fits in, shuffle the others around if needed
			if len(overallMax) > 2 {
				if runningMax > overallMax[0] {
					overallMax[2] = overallMax[1]
					overallMax[1] = overallMax[0]
					overallMax[0] = runningMax
				} else if runningMax > overallMax[1] {
					overallMax[2] = overallMax[1]
					overallMax[1] = runningMax
				} else if runningMax > overallMax[2] {
					overallMax[2] = runningMax
				}
			} else {
				// Not up to 3 elves yet, don't bother doing the comparison
				overallMax = append(overallMax, runningMax)
			}

			// Reset and move on to the next elf
			runningMax = 0
			sort.Slice(overallMax, func(i, j int) bool {
				return overallMax[i] > overallMax[j]
			})
			continue
		}

		// Figure out the number on the current line and append it to the total
		// carried by the current elf
		cur, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		runningMax += cur
	}

	fmt.Printf("P1: %d \n", overallMax[0])
	fmt.Printf("P2: %d \n", overallMax[0]+overallMax[1]+overallMax[2])
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
