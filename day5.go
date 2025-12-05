package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day5a() {
	input := mustReadFile("day5.txt")
	lines := strings.Split(input, "\n")
	fresh := true
	type limit struct {
		lower, upper int
	}
	var ingredients []limit
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			fresh = false
			continue
		}
		if fresh {
			limits := strings.Split(line, "-")
			start, _ := strconv.Atoi(limits[0])
			end, _ := strconv.Atoi(limits[1])
			ingredients = append(ingredients, limit{lower: start, upper: end})
		} else {
			ingredient, _ := strconv.Atoi(line)
			for _, ingredientLimit := range ingredients {
				if ingredient >= ingredientLimit.lower && ingredient <= ingredientLimit.upper {
					total++
					break
				}
			}
		}
	}

	fmt.Println(total)
}

func day5b() {
	input := mustReadFile("day5.txt")
	lines := strings.Split(input, "\n")
	type event struct {
		id    int
		delta int // +1 in range, -1 out of range
	}

	unique := map[int]int{}
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		limits := strings.Split(line, "-")
		start, _ := strconv.Atoi(limits[0])
		end, _ := strconv.Atoi(limits[1])
		end++ // end is inclusive
		unique[start]++
		unique[end]--

	}
	var events []event
	for id, count := range unique {
		if count == 0 {
			continue
		}

		events = append(events, event{id: id, delta: count})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].id < events[j].id
	})

	total := 0
	count := 0
	rangeStart := 0
	for _, e := range events {
		wasInRange := count > 0
		count += e.delta
		isInRange := count > 0

		// start new range
		if !wasInRange && isInRange {
			rangeStart = e.id
		} else if wasInRange && !isInRange {
			// hit end of range, count the difference
			total += e.id - rangeStart
		}
	}

	fmt.Println(total)
}
