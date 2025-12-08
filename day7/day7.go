package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part_1(lines []string) int {
	splits := 0

	start_ix := strings.Index(lines[0], "S")
	tachyon_ixs := map[int]bool{}
	tachyon_ixs[start_ix] = true
	new_tachyon_ixs := map[int]bool{}
	for _, line := range lines[1:] {
		fmt.Println(tachyon_ixs)
		for ix := range tachyon_ixs {
			if line[ix] == '^' {
				splits++
				new_tachyon_ixs[ix-1] = true
				new_tachyon_ixs[ix+1] = true
			} else {
				new_tachyon_ixs[ix] = true
			}
		}

		for k := range tachyon_ixs {
			delete(tachyon_ixs, k)
		}

		for k := range new_tachyon_ixs {
			tachyon_ixs[k] = true
			delete(new_tachyon_ixs, k)
		}
	}

	return splits
}

type CacheKey struct {
	line_num int
	line_ix  int
}

func part_2(line_num int, line_ix int, lines []string, cache map[CacheKey]int) int {
	cached_value, has_key := cache[CacheKey{line_num: line_num, line_ix: line_ix}]
	if has_key {
		return cached_value
	}

	result := 0
	if len(lines) == 0 {
		result = 1
	} else if line_ix < 0 || line_ix > len(lines[0]) {
		// Went off the line
		result = 0
	} else {
		if lines[0][line_ix] == '^' {
			result = part_2(line_num+1, line_ix-1, lines[1:], cache) + part_2(line_num+1, line_ix+1, lines[1:], cache)
		} else {
			result = part_2(line_num+1, line_ix, lines[1:], cache)
		}
	}

	cache[CacheKey{line_num, line_ix}] = result
	return result
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
	println("Part 2:", part_2(0, strings.Index(lines[0], "S"), lines, make(map[CacheKey]int)))
}
