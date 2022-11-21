package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Input() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println()
	value := input.Text()

	return value
}

func CToF(num string) {

	celsius, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	output := (celsius * (9 / 5)) + 32

	fmt.Printf("Celsius: %.2f  ->  Farenheight: %.2f \n", celsius, output)
}

func FToC(num string) {

	faren, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	output := (faren - 32) * (5.0 / 9.0)

	fmt.Printf("Farenheight: %.2f  ->  Celsius: %.2f \n", faren, output)
}
