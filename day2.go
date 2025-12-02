package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2a() {
	input := mustReadFile("day2.txt")
	count := 0
	for _, rs := range strings.Split(input, ",") {
		// find range, first as string
		rawStart, rawEnd := strings.Split(rs, "-")[0], strings.Split(rs, "-")[1]
		// then convert to ints (for iteration)
		start, err := strconv.Atoi(rawStart)
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(rawEnd)
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			// convert to string for manipulation
			strI := strconv.Itoa(i)
			// only numbers which are even in length, 12, 1234, 123456, etc.
			if len(strI)%2 != 0 {
				continue
			}
			// compare left and right parts for equality
			left := strI[:len(strI)/2]
			right := strI[len(strI)/2:]
			if left == right {
				count += i
			}
		}

	}

	fmt.Println(count)
}

func day2b() {
	input := mustReadFile("day2.txt")
	count := 0
	for _, rs := range strings.Split(input, ",") {
		// find range, first as string
		rawStart, rawEnd := strings.Split(rs, "-")[0], strings.Split(rs, "-")[1]
		// then convert to ints (for iteration)
		start, err := strconv.Atoi(rawStart)
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(rawEnd)
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			// dynamic sliding window?, eh, brute force

			// convert to string for manipulation
			strI := strconv.Itoa(i)
			// attempt to split string into parts equal upto it's length
			for j := 2; j <= len(strI); j++ {
				if invalid(strI, j) {
					count += i
					break
				}
			}
		}
	}

	fmt.Println(count)
}

func invalid(s string, numParts int) bool {
	// if s can't be split into equal parts, it must be valid
	if len(s) >= numParts && len(s)%numParts == 0 {
		size := len(s) / numParts
		var parts []string
		for i := 0; i < len(s); i += size {
			parts = append(parts, s[i:i+size])
		}
		for i := 1; i < len(parts); i++ {
			// if any part is not equal to the first part, this is valid, so return false
			if parts[0] != parts[i] {
				return false
			}
		}
		// if we made it here, it means each part was each to each other part, so invalid id (s)
		return true
	}

	return false
}
