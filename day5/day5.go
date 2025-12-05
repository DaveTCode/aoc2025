package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	min int64
	max int64
}

type Cafe struct {
	ranges      []Range
	ingredients []int64
}

func part_1(cafe Cafe) int {
	total := 0
	for _, ingredient := range cafe.ingredients {
		for _, r := range cafe.ranges {
			if ingredient >= r.min && ingredient <= r.max {
				total++
				break
			}
		}
	}

	return total
}

func part_2(cafe Cafe) int64 {
	total := int64(0)

	sort.Slice(cafe.ranges, func(i, j int) bool {
		return cafe.ranges[i].min < cafe.ranges[j].min
	})

	current := int64(0)
	for _, r := range cafe.ranges {
		// Range doesn't overlap at all, so add the whole thing and move our pointer to the new max
		if r.min > current {
			current = r.max
			total += r.max - r.min + 1
		} else {
			if r.max <= current {
				// Skip range as totally superceded
				continue
			}
			min := current + 1
			current = r.max
			total += r.max - min + 1
		}
	}

	return total
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cafe := Cafe{
		[]Range{}, []int64{},
	}

	parsing_ranges := true
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			parsing_ranges = false
			continue
		}

		if parsing_ranges {
			parts := strings.Split(line, "-")
			min, _ := strconv.ParseInt(parts[0], 10, 64)
			max, _ := strconv.ParseInt(parts[1], 10, 64)
			cafe.ranges = append(cafe.ranges, Range{min, max})
		} else {
			ingredient, _ := strconv.ParseInt(line, 10, 64)
			cafe.ingredients = append(cafe.ingredients, ingredient)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	println("Part 1:", part_1(cafe))
	println("Part 2:", part_2(cafe))
}
