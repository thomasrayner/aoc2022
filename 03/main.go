package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
    p1Score := 0
    p2Score := 0
    content, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    splitContent := strings.Split(string(content), "\n")

    Outter1:
    for _, line := range splitContent {
        first := line[:(len(line)/2)]
        second := line[(len(line)/2):]
        comparison := map[rune]bool{}

        for _, c := range first {
            comparison[c] = true
        }
        for _, c := range second {
            if comparison[c] {
                if unicode.IsUpper(c) {
                    p1Score += int(c) - 38
                } else {
                    p1Score += int(c) - 96
                }
                continue Outter1
            }
        }
    }

    Outter2:
    for i := 0; i < len(splitContent) - 2; i += 3 {
        for _, c := range splitContent[i] {
            if strings.ContainsRune(splitContent[i + 1], c) && strings.ContainsRune(splitContent[i + 2], c) {
                fmt.Printf("%d - %c - \n%s\n%s\n%s\n\n", i, c, splitContent[i], splitContent[i+1], splitContent[i+2])
                if unicode.IsUpper(c) {
                    p2Score += int(c) - 38
                } else {
                    p2Score += int(c) - 96
                }
                continue Outter2
            }
        }
    }

    fmt.Printf("P1: %d\n", p1Score)
    fmt.Printf("P2: %d\n", p2Score)
}