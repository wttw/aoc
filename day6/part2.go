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

	groups := []map[rune]int{}
	answers := map[rune]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && len(answers) != 0 {
			groups = append(groups, answers)
			answers = map[rune]int{}
			continue
		}
		for _, r := range line {
			answers[r]++
		}
		answers[' ']++ // count group size as an always answered question ' '
	}
	if len(answers) != 0 {
		groups = append(groups, answers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, group := range groups {
		groupSize := group[' ']
		for k, v := range group {
			if k != ' ' && v == groupSize {
				sum++
			}
		}
	}
	log.Print("sum=", sum)
}
