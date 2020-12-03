package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		min, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal("failed to parse min", scanner.Text())
		}
		max, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatal("failed to parse max", scanner.Text())
		}
		matches := strings.Count(fields[4], fields[3])
		if matches >= min && matches <= max {
			valid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("valid", valid)
}