package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func routes(start string, end string, bad map[string]bool, connections map[string][]string, cache map[string]int) int {
	if start == end {
		return 1
	}
	if _, ok := bad[start]; ok {
		return 0
	}

	total := 0
	for _, connection := range connections[start] {
		if count, ok := cache[connection]; ok {
			total += count
		} else {
			result := routes(connection, end, bad, connections, cache)
			cache[connection] = result
			total += result
		}
	}
	return total
}

func part_1(connections map[string][]string) int {
	return routes("you", "out", map[string]bool{}, connections, map[string]int{})
}

func part_2(connections map[string][]string) int {
	svr_to_fft := routes("svr", "fft", map[string]bool{"dac": false, "out": false}, connections, map[string]int{})
	svr_to_dac := routes("svr", "dac", map[string]bool{"fft": false, "out": false}, connections, map[string]int{})
	fft_to_dac := routes("fft", "dac", map[string]bool{"out": false}, connections, map[string]int{})
	dac_to_fft := routes("dac", "fft", map[string]bool{"out": false}, connections, map[string]int{})
	fft_to_out := routes("fft", "out", map[string]bool{"dac": false}, connections, map[string]int{})
	dac_to_out := routes("dac", "out", map[string]bool{"fft": false}, connections, map[string]int{})
	return (svr_to_fft * fft_to_dac * dac_to_out) + (svr_to_dac * dac_to_fft * fft_to_out)
}

func read_file(file_name string) map[string][]string {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return map[string][]string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return map[string][]string{}
	}

	connections := map[string][]string{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		connections[parts[0]] = strings.Split(parts[1], " ")
	}

	return connections
}

func main() {
	test := read_file("test")
	test2 := read_file("test2")
	input := read_file("input")
	println("Part 1 (test):", part_1(test))
	println("Part 1 (real):", part_1(input))
	println("Part 2 (test):", part_2(test2))
	println("Part 2 (real):", part_2(input))
}
