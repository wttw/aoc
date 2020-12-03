package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vals := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		vals = append(vals, i);
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, i := range vals {
		for _, j := range vals {
			for _, k := range vals {
				if i+j+k == 2020 {
					fmt.Println(i, j, k, i*j*k)
				}
			}
		}
	}
}
