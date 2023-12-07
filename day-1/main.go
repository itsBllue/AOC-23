package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extractDigitsAndSpeltStrings(str string) []int {

	// this is the map that will be used to convert the spelled out numbers to digits
	spellingMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digits := []int{}
	for i, char := range str {
		// convert to unicode and try to convert to int ( this will  store the number into the array)
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		} else {
			// if the character is not a number, then try to match it to the spelling map
			for spelling, number := range spellingMap {
				if strings.HasPrefix(str[i:], spelling) {
					digits = append(digits, number)
					break
				}
			}
		}
	}

	return digits
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	total := 0
	i := 0
	for scanner.Scan() {
		i = i + 1
		// // fmt.Println(scanner.Text())
		var intArray = extractDigitsAndSpeltStrings(scanner.Text())
		combineFirstLast, _ := strconv.Atoi(strconv.Itoa(intArray[0]) + strconv.Itoa(intArray[len(intArray)-1]))
		total += combineFirstLast
	}
	// fmt.Println("total:", total)
	fmt.Println(total)

	log.SetFlags(0)
	log.SetPrefix("Hello!\n")

}
