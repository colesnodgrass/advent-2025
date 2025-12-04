package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("day01:")
	//day01a()
	//day01b()

	//fmt.Println("day02:")
	//day2a()
	//day2b()

	fmt.Println("day03:")
	day3a()
	day3b()
}

func mustReadFile(file string) string {
	inputBytes, err := os.ReadFile("data/" + file)
	if err != nil {
		panic(err)
	}
	return string(inputBytes)
}
