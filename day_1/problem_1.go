package day_1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Problem_1() {
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
		s := append(lineDigit[0:1], lineDigit[len(lineDigit)-1])
		digit, _ := strconv.Atoi(strings.Join(s, ""))
		total += digit
		println(total)

	}
}
