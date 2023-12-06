package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getIntsFromString(s string) (int, error) {

	var numbers []int
	// println(s)

	for _, char := range s {
		// println(char)
		if unicode.IsDigit(char) {
			digit := int(char - '0')
			numbers = append(numbers, digit)
		}
	}

	if len(numbers) < 2 {
		numbers = append(numbers, numbers[0])
	}

	conc_strr := strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1])
	intToReturn, err := strconv.Atoi(conc_strr)
	if err != nil {
		return 0, err
	}
	return intToReturn, nil
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	// intsFromLine, err := getIntsFromString("eight8fivexsbkzcthree")

	// fmt.Println(intsFromLine)
	scanner := bufio.NewScanner(inputFile)
	total := 0
	i := 0
	for scanner.Scan() {
		i = i + 1
		fmt.Println(scanner.Text())

		intsFromLine, err := getIntsFromString(scanner.Text())
		total += intsFromLine
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(intsFromLine)

	}
	fmt.Println(total)
	log.SetFlags(0)
	log.SetPrefix("Hello!\n")
	// fmt.Println(string(content))

	// message, err := getHelloMessage("Gopher")
	// if err != nil {
	// log.Fatal(err)
	// }
	// fmt.Println(message)
}
