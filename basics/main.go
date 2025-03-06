package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	for {
		fmt.Println("Please think of a number between", low, "and", high)
		fmt.Println("Press ENTER when ready")
		scanner.Scan()

		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatalf("Tried to parse input value, err: %v", err)
		}

		numCast := int(num)

		if numCast < low || numCast > high {
			fmt.Println("Number out of range, try again.")
			continue
		}
		break
	}

	numGuesses := 1
	for {
		// Binary search log(n)
		guess := (low + high) / 2
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) just right?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
			numGuesses++
		} else if response == "b" {
			low = guess + 1
			numGuesses++
		} else if response == "c" {
			fmt.Printf("I won in %d turns!\n", numGuesses)
			break
		} else {
			fmt.Println("Invalid response, try again")
		}
	}
}
