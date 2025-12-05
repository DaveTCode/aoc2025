package main

import (
	"bufio"
	"fmt"
	"os"
)

func part_1(grid [][]bool) int {
	adjacent_positions := [][]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	result := 0
	for row_ix, row := range grid {
		for col_ix, is_paper := range row {
			if is_paper {
				count_adjacent_paper := 0

				for _, adjacencies := range adjacent_positions {
					new_x := adjacencies[0] + col_ix
					new_y := adjacencies[1] + row_ix

					if new_x >= 0 && new_x < len(row) && new_y >= 0 && new_y < len(grid) {
						if grid[new_y][new_x] {
							count_adjacent_paper++
						}
					}
				}

				if count_adjacent_paper < 4 {
					println(row_ix, col_ix)
					result++
				}
			}
		}
	}
	return result
}

func part_2(grid [][]bool) int {
	adjacent_positions := [][]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	result := 0
	for {
		last_result := result
		for row_ix, row := range grid {
			for col_ix, is_paper := range row {
				if is_paper {
					count_adjacent_paper := 0

					for _, adjacencies := range adjacent_positions {
						new_x := adjacencies[0] + col_ix
						new_y := adjacencies[1] + row_ix

						if new_x >= 0 && new_x < len(row) && new_y >= 0 && new_y < len(grid) {
							if grid[new_y][new_x] {
								count_adjacent_paper++
							}
						}
					}

					if count_adjacent_paper < 4 {
						grid[row_ix][col_ix] = false
						result++
					}
				}
			}
		}

		if last_result == result {
			break
		}
	}

	return result
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		grid_line := []bool{}
		for _, chr := range line {
			grid_line = append(grid_line, chr == '@')
		}
		grid = append(grid, grid_line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	println("Part 1:", part_1(grid))
	println("Part 2:", part_2(grid))
}
