package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Rotation struct {
	Forward bool
	Steps   int
}

func part_1(rotations []Rotation) int {
	var current_position = 50
	var zero_count = 0
	for _, rotation := range rotations {
		if rotation.Forward {
			current_position = (current_position + rotation.Steps) % 100
		} else {
			current_position = (current_position - rotation.Steps + 100) % 100
		}

		if current_position == 0 {
			zero_count += 1
		}
	}

	return zero_count
}

// Why do maths when you can brute force it
func part_2(rotations []Rotation) int {
	var current_position = 50
	var zero_count = 0
	for _, rotation := range rotations {
		for step := 0; step < rotation.Steps; step++ {
			if rotation.Forward {
				current_position = (current_position + 1) % 100
			} else {
				current_position = (current_position - 1 + 100) % 100
			}
			if current_position == 0 {
				zero_count += 1
			}
		}
	}

	return zero_count
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []Rotation
	for scanner.Scan() {
		line := scanner.Text()
		is_forward := line[0] == 'R'
		steps, _ := strconv.Atoi(line[1:])
		lines = append(lines, Rotation{
			Forward: is_forward,
			Steps:   steps,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	print("Part 1:", part_1(lines))
	print("Part 2:", part_2(lines))
}
