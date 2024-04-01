package main

import (
	"bufio"
	"os"
)

func main() {
	data, _ := os.Open("./input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)
	// regex := regexp.MustCompile(`\d`)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		// println(line)
		for _, char := range line {
			if char >= '0' && char <= '9' {
				digit := int(char - '0')
				total += digit
				println(total)

			}
		}

	}
}
