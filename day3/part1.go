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

	trees := 0
	x := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		terrain := row[x % len(row)]
		if terrain == '#' {
			trees++
		}
		x += 3
 	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("trees", trees)
}
