package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	for true {
		fmt.Println("--------------- Temperature Converter ---------------")
		fmt.Println("1: C to F")
		fmt.Println("2: F to C")
		fmt.Print("3: Exit \n")

		option := Input()

		switch option {
		case "1":
			fmt.Println("--------------- Celsius To Farenheight Converter ---------------")
			fmt.Println()
			fmt.Print("Enter the temperature you want to convert: ")
			CToF(Input())
		case "2":
			fmt.Println("--------------- Farenheight To Celsius Converter ---------------")
			fmt.Println()
			fmt.Print("Enter the temperature you want to convert: ")
			FToC(Input())
		case "3":
			fmt.Println("--------------- Exit ---------------")
			os.Exit(1)
		}

		fmt.Println("Do you want to make another conversion? Y/N")
		input := Input()

		if strings.ToLower(input) != "y" {
			break
		}

	}

}
