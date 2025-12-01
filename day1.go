package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	day1InputFile     = "data/day1.txt"
	day1InputFileTest = "data/day1-test.txt"
)

func day01a() {
	inputBytes, err := os.ReadFile(day1InputFile)
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)

	zeroCount := 0
	pos := 50
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		rotation := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if value > 100 {
			value %= 100
		}
		switch rotation {
		case 'L':
			pos -= value
			if pos < 0 {
				pos += 100
			}
		case 'R':
			pos += value
			if pos > 99 {
				pos -= 100
			}
		}

		if pos == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}

func day01b() {
	inputBytes, err := os.ReadFile(day1InputFile)
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)

	zeroCount := 0
	pos := 50
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		rotation := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if value > 100 {
			zeroCount += value / 100
			value %= 100
		}
		// track prev position, so we don't count moving away from 0
		prevPos := pos
		switch rotation {
		case 'L':
			pos -= value
			if pos < 0 {
				pos += 100
				// don't double count if we start on zero or end on zero
				if prevPos != 0 && pos != 0 {
					zeroCount++
				}
			}
		case 'R':
			pos += value
			if pos > 99 {
				pos -= 100
				if prevPos != 0 && pos != 0 {
					zeroCount++
				}
			}
		}

		if pos == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
