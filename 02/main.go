package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    p1Score := 0
    p1Poss := map[string]int{}
    p1Poss["A X"] = 4
    p1Poss["A Y"] = 8
    p1Poss["A Z"] = 3
    p1Poss["B X"] = 1
    p1Poss["B Y"] = 5
    p1Poss["B Z"] = 9
    p1Poss["C X"] = 7
    p1Poss["C Y"] = 2
    p1Poss["C Z"] = 6
    
    p2Score := 0
    p2Poss := map[string]int{}
    p2Poss["A X"] = 3
    p2Poss["A Y"] = 4
    p2Poss["A Z"] = 8
    p2Poss["B X"] = 1
    p2Poss["B Y"] = 5
    p2Poss["B Z"] = 9
    p2Poss["C X"] = 2
    p2Poss["C Y"] = 6
    p2Poss["C Z"] = 7

    content, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    for _, line := range strings.Split(string(content), "\n") {
        p1Score += p1Poss[line]
        p2Score += p2Poss[line]
    }

    fmt.Printf("P1: %d\n", p1Score)
    fmt.Printf("P2: %d\n", p2Score)
}
