package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Split fields on \n\n
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		idx := bytes.Index(data, []byte("\n\n"))
		if idx != -1 {
			return idx + 2, data[:idx], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	valid := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		seen := map[string]string{}
		for _, field := range fields {
			kv := strings.SplitN(field, ":", 2)
			if len(kv) != 2 {
				log.Fatal("dodgy field:", field)
			}

			if kv[0] == "cid" {
				continue
			}

			seen[kv[0]] = kv[1]
		}
		if len(seen) == 7 {
			valid++
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print("valid=", valid)
}
