package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day3a() {
	input := mustReadFile("day3.txt")
	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		firstNum := line[0]
		firstPos := 0
		for i := 1; i < len(line)-1; i++ {
			if line[i] > firstNum {
				firstNum = line[i]
				firstPos = i
			}
		}
		secondNum := line[firstPos+1]
		for i := firstPos + 2; i < len(line); i++ {
			if line[i] > secondNum {
				secondNum = line[i]
			}
		}
		tens, err := strconv.Atoi(string(firstNum))
		if err != nil {
			panic(err)
		}
		ones, err := strconv.Atoi(string(secondNum))
		if err != nil {
			panic(err)
		}
		result += (tens * 10) + ones
	}

	fmt.Println(result)
}

func day3b() {
	//input := mustReadFile("day3.txt")
	//lines := strings.Split(input, "\n")

	total := 0
	lines := []string{
		"1112121122222223222222222222112242222232323212322222213112322622222132213212123253423222223242122232",
	}
	for _, line := range lines {
		result := ""
		start := 0
		for i := 11; i >= 0; i-- {
			num := "0"
			subLine := line[start : len(line)-i]
			innerStart := 0
			for j, ch := range subLine {
				if string(ch) > num {
					num = string(ch)
					innerStart = j
				}
			}

			result += num
			start += innerStart + 1
		}

		if asInt, err := strconv.Atoi(result); err != nil {
			panic(err)
		} else {
			total += asInt
		}
	}

	fmt.Println(total)
}
