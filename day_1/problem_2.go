package day_1

import (
	"bufio"
	"os"
	"strings"
)

func Problem_2() {
	data, _ := os.Open("./input.txt")
	defer data.Close()

	scanner := bufio.NewScanner(data)

	nums := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	lastNumberIndex := -1
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range nums {
			// encontra o índicice do primeiro caractere do número por exetenso
			if y := strings.Index(line, v); y < lastNumberIndex {
				lastNumberIndex = y
			}
		}
	}
}
