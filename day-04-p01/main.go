package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string;

func (p passport) isValid() bool {
	for k, v := range p {
		switch k {
			case "byr":
				if !isNumberInRange(v, 1920, 2002) {
					return false
				}
			case "iyr":
				if !isNumberInRange(v, 2010, 2020) {
					return false
				}
			case "eyr":
				if !isNumberInRange(v, 2020, 2030) {
					return false
				}
			case "hgt":
				r, _ := regexp.Compile(`^\d\d\d?(cm|in)$`)
				if r.MatchString(v) {
					if strings.Contains(v, "cm") {
						if n, err := strconv.Atoi(v[:3]); err == nil {
							if n < 150 || n > 193 {
								return false
							}
						} else {
							return false
						}

					} else if strings.Contains(v, "in") {
						if n, err := strconv.Atoi(v[:2]); err == nil {
							if n < 59 || n > 76 {
								return false
							}
						} else {
							return false
						}

					} else {
						log.Panic("invalid value for hgt:", v)
					}
				} else {
					return false
				}
			case "hcl":
				// regex to match strings like #123abf
				if !isRegexValid(v, "^#[0-9a-f]{6}$") {
					return false
				}
			case "ecl":
				if !isRegexValid(v, "^(amb|blu|brn|gry|grn|hzl|oth)$") {
					return false
				}
			case "pid": 
				if !isRegexValid(v, `^\d{9}$`) {
					return false
				}
		}
	}
	return true
}

func isNumberInRange(value string, min, max int) bool {
	if n, err := strconv.Atoi(value); err == nil {
		if min <= n && n <= max {
			return true
		} else {
			return false
		}
	} else {
		log.Println(err.Error())
		return false
	}
}

func isRegexValid(value, regex string) bool {
	r, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}
	return r.MatchString(value)
}

func main() {

	passports := readPassportData("input.txt")
	valid := 0
	for _, p := range passports {
		if len(p) == 8 {
			if p.isValid() {
				valid++
				fmt.Printf("%#v [%d]\n", p, valid)
			}
		
		}
		if _, found := p["cid"]; !found && len(p) == 7 {
			if p.isValid() {
				valid++	
				fmt.Printf("%#v [%d]\n", p, valid)	
			}
			
		}
	}

	fmt.Println("result:", valid)
}

func readPassportData(fileName string) []passport {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)

	var passports = make([]passport, 0)
	var line string
	for s.Scan() {
		if s.Text() == "" {
			// we have a complete passport line from the input
			passports = append(passports, parsePassport(line))
			// reset the line so we can parse another passport
			line = ""
		} else {
			if line != "" {
				line += " "
			}
			line += s.Text()
		}
	}
	// check to see if we have another line after scanner reaches EOF
	if line != "" {
		passports = append(passports, parsePassport(line))
	}
	return passports
}

func parsePassport(line string) map[string]string {
	p := make(map[string]string)
	tokens := strings.Split(line, " ")
	//println(line, len(tokens))
	for _, t := range tokens {
		vals := strings.Split(t, ":")
		if len(vals) != 2 {
			log.Fatalf("unknown format: %s", t)
		}
		key, val := vals[0], vals[1]
		p[key] = val
	}
	return p
}
