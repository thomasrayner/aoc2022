package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1Score := 0
	p2Score := 0
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	splitContent := strings.Split(string(content), "\n")

	for _, line := range splitContent {
		assignments := strings.Split(line, ",")
		elf1Start, _ := strconv.Atoi(string(strings.Split(assignments[0], "-")[0]))
		elf2Start, _ := strconv.Atoi(string(strings.Split(assignments[1], "-")[0]))
		elf1End, _ := strconv.Atoi(string(strings.Split(assignments[0], "-")[1]))
		elf2End, _ := strconv.Atoi(string(strings.Split(assignments[1], "-")[1]))

		if (elf1Start <= elf2Start && elf1End >= elf2End) || (elf2Start <= elf1Start && elf2End >= elf1End) {
			p1Score++
		}
		if (elf2Start >= elf1Start && elf2Start <= elf1End) || (elf1Start >= elf2Start && elf1Start <= elf2End) {
			p2Score++
		}
	}

	fmt.Printf("P1: %d\n", p1Score)
	fmt.Printf("P2: %d\n", p2Score)
}