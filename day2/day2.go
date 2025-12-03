package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProductRange struct {
	min string
	max string
}

func part_1(ranges []ProductRange) int {
	sum := 0
	for _, r := range ranges {
		min, _ := strconv.Atoi(r.min)
		max, _ := strconv.Atoi(r.max)

		for id := min; id <= max; id++ {
			strid := strconv.FormatInt(int64(id), 10)
			if len(strid)%2 != 0 {
				continue
			}

			if strid[:len(strid)/2] == strid[len(strid)/2:] {
				sum += id
			}
		}
	}

	return sum
}

func part_2(ranges []ProductRange) int {
	sum := 0
	for _, r := range ranges {
		min, _ := strconv.Atoi(r.min)
		max, _ := strconv.Atoi(r.max)

		for id := min; id <= max; id++ {
			strid := strconv.FormatInt(int64(id), 10)

			for num_chunks := 2; num_chunks <= len(strid); num_chunks++ {
				if len(strid)%num_chunks != 0 {
					continue
				}
				valid := true

				chunk_size := len(strid) / num_chunks
				for chunk := 0; chunk <= num_chunks-1; chunk++ {
					if strid[chunk*chunk_size:(chunk+1)*chunk_size] != strid[0:chunk_size] {
						valid = false
						break
					}
				}

				if valid {
					println(id)
					sum += id
					break
				}
			}
		}
	}

	return sum
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(lines) > 1 {
		fmt.Println("Expected only one line in the input")
		return
	}

	ranges := strings.Split(lines[0], ",")
	product_ranges := []ProductRange{}
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		product_ranges = append(product_ranges, ProductRange{
			min: parts[0],
			max: parts[1],
		})
	}

	print("Part 1:", part_1(product_ranges))
	print("Part 2:", part_2(product_ranges))
}
