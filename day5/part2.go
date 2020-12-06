package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seats := []int{}
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
		seats = append(seats, value)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(seats)
	for i, v := range seats {
		if seats[i+1] == v+2 {
			log.Print("seat:", v+1)
			break
		}
	}
}
