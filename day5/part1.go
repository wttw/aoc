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
	max := 0
	for scanner.Scan() {
		value := 0
		for _, digit := range scanner.Text() {
			value <<= 1
			switch digit {
			case 'F', 'L':
			case 'B', 'R':
				value |= 1
			default:
				log.Fatal("unexpected digit", digit)
			}
		}
		if value > max {
			max = value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print("max=", max)
}
