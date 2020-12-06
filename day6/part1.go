package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	groups := []map[rune]struct{}{}
	answers := map[rune]struct{}{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && len(answers) != 0 {
			groups = append(groups, answers)
			answers = map[rune]struct{}{}
		}
		for _, r := range line {
			answers[r] = struct{}{}
		}
	}
	if len(answers) != 0 {
		groups = append(groups, answers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, group := range groups {
		sum += len(group)
	}
	log.Print("sum=", sum)
}
