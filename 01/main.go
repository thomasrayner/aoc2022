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
			overallMax = append(overallMax, runningMax)
			runningMax = 0
			continue
		}

		cur, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		runningMax += cur
	}

	sort.Slice(overallMax, func(i, j int) bool {
		return overallMax[i] > overallMax[j]
	})

	fmt.Printf("P1: %d \n", overallMax[0])
	fmt.Printf("P2: %d \n", overallMax[0]+overallMax[1]+overallMax[2])
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
