package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape [3][3]int

type Region struct {
	Width    int
	Length   int
	Presents []int
}

type Puzzle struct {
	Shapes  []Shape
	Regions []Region
}

// This is lame
func part_1(puzzle Puzzle) int {
	valid_regions := 0
	for _, region := range puzzle.Regions {
		width_fit := region.Width / 3
		length_fit := region.Length / 3

		total := width_fit * length_fit

		shape_total := 0
		for _, region_shape := range region.Presents {
			shape_total += region_shape
		}

		if total >= shape_total {
			valid_regions++
		}

		fmt.Println(region, "= ", width_fit, length_fit, total, shape_total, " = ", total >= valid_regions)
	}

	return valid_regions
}

func read_file(file_name string) Puzzle {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Puzzle{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return Puzzle{}
	}

	puzzle := Puzzle{}
	line_num := 0
	parsing_shapes := true
	for line_num < len(lines) {
		if parsing_shapes {
			shape := Shape{}
			for shape_line_num := range 3 {
				line := lines[line_num+1+shape_line_num]
				shape_line := [3]int{0, 0, 0}

				for chr_num, chr := range line {
					switch chr {
					case '.':
						shape_line[chr_num] = 0
					case '#':
						shape_line[chr_num] = 1
					}
				}

				shape[shape_line_num] = shape_line
			}

			puzzle.Shapes = append(puzzle.Shapes, shape)
			line_num += 5

			if strings.Contains(lines[line_num], "x") {
				parsing_shapes = false
			}
		} else {
			parts := strings.Split(lines[line_num], ":")
			width_length := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(width_length[0])
			length, _ := strconv.Atoi(width_length[1])
			presents := []int{}
			for present_str := range strings.SplitSeq(strings.TrimLeft(parts[1], " "), " ") {
				present, _ := strconv.Atoi(present_str)
				presents = append(presents, present)
			}

			puzzle.Regions = append(puzzle.Regions, Region{
				Width:    width,
				Length:   length,
				Presents: presents,
			})
			line_num++

			if line_num >= len(lines) {
				break
			}
		}
	}

	return puzzle
}

func main() {
	test := read_file("test")
	input := read_file("input")
	println("Part 1 (test):", part_1(test))
	println("Part 1 (real):", part_1(input))
}
