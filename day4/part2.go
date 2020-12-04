package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
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

	validator := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`^(19[2-9][0-9]|200[0-2])$`),
		"iyr": regexp.MustCompile(`^(201[0-9]|2020)$`),
		"eyr": regexp.MustCompile(`^(202[0-9]|2030)$`),
		"hgt": regexp.MustCompile(`^(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)$`),
		"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
		"ecl": regexp.MustCompile(`^(amb|blu|brn|gr[yn]|hzl|oth)$`),
		"pid": regexp.MustCompile(`^[0-9]{9}$`),
	}

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
			
			val, ok := validator[kv[0]]
			if ok && val.MatchString(kv[1]) {
				seen[kv[0]] = kv[1]
			}
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
