package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.Open("./input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)
	// regex := regexp.MustCompile(`\d`)
	total := 0
	for scanner.Scan() {
		var lineDigit []string
		line := scanner.Text()
		// ln := 0
		for _, char := range line {
			if char >= '0' && char <= '9' {
				lineDigit = append(lineDigit, string(char))
			}

		}
		digit, _ := strconv.Atoi(strings.Join(lineDigit, ""))
		total += digit
		println(total)

	}
}
