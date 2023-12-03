package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var solution1Regexp, _ = regexp.Compile("\\d")

func solution1(txt string) (int) {
	match := solution1Regexp.FindAllString(txt,-1)
	
	result,_ :=strconv.Atoi(match[0]+match[len(match)-1])

	return result;
}

var solution2Regexp, _ = regexp.Compile("\\d|one|two|three|four|five|six|seven|eight|nine")

func transformToNumber(txt string) (string) {
	switch txt {
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		case "four":
			return "4"
		case "five":
			return "5"
		case "six":
			return "6"
		case "seven":
			return "7"
		case "eight":
			return "8"
		case "nine":
			return "9"
		default:
			return txt
	}
}

func solution2(txt string) (int) {
	var match []string;

	if(len(txt) <= 5) {
		match = solution2Regexp.FindAllString(txt,-1)
	} else {
		match = []string{solution2Regexp.FindString(txt)};

		for i := len(txt)-5; i>0; i-- {
			var last = solution2Regexp.FindAllString(txt[i:i+5], -1)
			if(len(last) > 0) {
				match = append(match, last[len(last)-1])
				break;
			}
		}
	}

	result,_ :=strconv.Atoi(transformToNumber(match[0])+transformToNumber(match[len(match)-1]))

	return result;
}

func main() {
	absfile, _ := filepath.Abs("./day1/input.txt");

	file, err := os.Open(absfile)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sol1 int = 0
	var sol2 int = 0

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
			var txt string = scanner.Text();

			sol1 += solution1(txt);
			sol2 += solution2(txt);
	}

	fmt.Println("solution1:", sol1, "solution2:", sol2)
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}