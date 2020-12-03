package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	valid := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineRe := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]*)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := lineRe.FindStringSubmatch(scanner.Text())
		if fields == nil {
			log.Fatal("failed to read line", scanner.Text())
		}
		pos1, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal("failed to parse min", scanner.Text())
		}
		pos2, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatal("failed to parse max", scanner.Text())
		}
		character := fields[3][0]
		password := fields[4]

		if (password[pos1-1] == character) != (password[pos2-1] == character) {
			valid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("valid", valid)
}