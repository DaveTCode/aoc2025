package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sum struct {
	numbers   []int
	operation string
}

func sum_of_sums(sums []Sum) int {
	total := 0
	for _, sum := range sums {
		fmt.Println(sum)

		sub_total := 0
		if sum.operation == "*" {
			sub_total = 1
		}
		for _, n := range sum.numbers {
			if sum.operation == "+" {
				sub_total += n
			} else {
				sub_total *= n
			}
		}
		total += sub_total
	}
	return total
}

func part_1(lines []string) int {
	start := 0
	sums := []Sum{}
	for ix := range len(lines[0]) {
		found_end := true
		for _, line := range lines {
			if line[ix] != ' ' {
				found_end = false
			}
		}

		if found_end || (ix == len(lines[0])-1) {
			nums := []int{}
			operation := ""
			for i, line := range lines {
				if i == len(lines)-1 {
					operation = strings.TrimSpace(line[start : ix+1])
				} else {
					num, _ := strconv.Atoi(strings.TrimSpace(line[start : ix+1]))
					nums = append(nums, num)
				}
			}
			sums = append(sums, Sum{
				nums,
				operation,
			})
			start = ix
		}
	}

	return sum_of_sums(sums)
}

func part_2(lines []string) int {
	sums := []Sum{}
	current_sum := Sum{}
	for ix := range len(lines[0]) {
		reverse_index := len(lines[0]) - ix - 1
		current_number := []string{}
		for line_ix, line := range lines {
			if line_ix == len(lines)-1 {
				if len(current_number) != 0 {
					number, _ := strconv.Atoi(strings.Join(current_number, ""))
					current_sum.numbers = append(current_sum.numbers, number)
				}

				if line[reverse_index] == '+' || line[reverse_index] == '*' {
					current_sum.operation = string(line[reverse_index])
					sums = append(sums, current_sum)
					current_sum = Sum{}
				}
			} else {
				if line[reverse_index] != ' ' {
					current_number = append(current_number, string(line[reverse_index]))
				}
			}
		}
	}
	sums = append(sums, current_sum)

	return sum_of_sums(sums)
}

func main() {
	file, err := os.Open(os.Args[1])
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

	println("Part 1:", part_1(lines))
	println("Part 2:", part_2(lines))
}
