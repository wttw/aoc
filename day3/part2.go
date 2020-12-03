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

	slopes := []*struct{x, y, trees int}{
		{1, 1, 0},
		{3, 1, 0},
		{5, 1, 0},
		{7, 1, 0},
		{1, 2, 0},
	}

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		for _, s := range slopes {
			if y % s.y == 0 {
				x := s.x * y / s.y
				if row[x % len(row)] == '#' {
					s.trees++
				}
			}
		}
		y++
 	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	product := 1
	for _, s := range slopes {
		product *= s.trees
	}
	log.Println("product", product)
}
