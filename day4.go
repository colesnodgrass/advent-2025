package main

import (
	"fmt"
	"strings"
)

func day4a() {
	input := mustReadFile("day4.txt")

	var matrix [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := make([]rune, 0, len(line))
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	hits := 0
	for r, row := range matrix {
		for c := range row {
			if matrix[r][c] == '.' {
				continue
			}
			count := 0
			for x := -1; x <= 1; x++ {
				if r+x < 0 || r+x >= len(row) {
					continue
				}

				for y := -1; y <= 1; y++ {
					if c+y < 0 || c+y >= len(matrix) {
						continue
					}
					if x == 0 && y == 0 {
						continue
					}
					if matrix[r+x][c+y] == '@' {
						count++
					}
				}
			}
			if count < 4 {
				hits++
			}
		}
	}

	fmt.Println(hits)
}

func day4b() {
	input := mustReadFile("day4.txt")

	var matrix [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := make([]rune, 0, len(line))
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	total := 0
	subTotal := -1
	for subTotal != 0 {
		type point struct{ x, y int }
		var remove []point
		subTotal = 0

		for r, row := range matrix {
			for c := range row {
				if matrix[r][c] == '.' || matrix[r][c] == 'x' {
					continue
				}
				count := 0
				for x := -1; x <= 1; x++ {
					if r+x < 0 || r+x >= len(row) {
						continue
					}

					for y := -1; y <= 1; y++ {
						if c+y < 0 || c+y >= len(matrix) {
							continue
						}
						if x == 0 && y == 0 {
							continue
						}
						if matrix[r+x][c+y] == '@' {
							count++
						}
					}
				}
				if count < 4 {
					remove = append(remove, point{r, c})
					subTotal++
					total++
				}
			}
		}

		for _, p := range remove {
			matrix[p.x][p.y] = 'x'
		}
	}

	fmt.Println(total)
}
