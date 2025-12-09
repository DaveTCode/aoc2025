package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	A Point
	B Point
}

type Rectangle struct {
	A    Point
	B    Point
	Size int
}

func line_in_rectangle(r Rectangle, l Line) bool {
	xmin := min(r.A.X, r.B.X)
	xmax := max(r.A.X, r.B.X)
	ymin := min(r.A.Y, r.B.Y)
	ymax := max(r.A.Y, r.B.Y)

	if l.A.X == l.B.X { // vertical line
		x := l.A.X
		if x > xmin && x < xmax { // inside rectangle horizontally
			lymin := min(l.A.Y, l.B.Y)
			lymax := max(l.A.Y, l.B.Y)
			if lymax > ymin && lymin < ymax { // intersects interior vertically
				return true
			}
		}
	} else { // horizontal line
		y := l.A.Y
		if y > ymin && y < ymax { // inside rectangle vertically
			lxmin := min(l.A.X, l.B.X)
			lxmax := max(l.A.X, l.B.X)
			if lxmax > xmin && lxmin < xmax { // intersects interior horizontally
				return true
			}
		}
	}

	return false
}

func create_sorted_rectangles(points []Point) []Rectangle {
	rectangles := []Rectangle{}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			size := (max(p1.X, p2.X) - min(p1.X, p2.X) + 1) * (max(p1.Y, p2.Y) - min(p1.Y, p2.Y) + 1)
			rectangles = append(rectangles, Rectangle{p1, p2, size})
		}
	}

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Size > rectangles[j].Size
	})

	return rectangles
}

func part_1(points []Point) int {
	return create_sorted_rectangles(points)[0].Size
}

func part_2(points []Point) int {
	rectangles := create_sorted_rectangles(points)
	lines := make([]Line, len(points))

	// Construct the list of lines
	for ix, point := range points {
		next_point := points[0]
		if ix != len(points)-1 {
			next_point = points[ix+1]
		}

		// Ensure the lines are ordered specifically to make it easier to calculate based on them
		if (point.X == next_point.X && point.Y < next_point.Y) || (point.Y == next_point.Y && point.X < next_point.X) {
			lines[ix] = Line{point, next_point}
		} else {
			lines[ix] = Line{next_point, point}
		}
	}

	for _, rectangle := range rectangles {
		valid := true
		// For each rectangle check if the 4 lines cross any other line
		for _, line := range lines {
			if line_in_rectangle(rectangle, line) {
				fmt.Println(line, " --cross-- ", rectangle)
				valid = false
				break
			}
		}

		if valid {
			fmt.Println(rectangle)
			return rectangle.Size
		}
	}

	return 0
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := []Point{}
	for scanner.Scan() {
		line_parts := strings.Split(scanner.Text(), ",")

		x, _ := strconv.Atoi(line_parts[0])
		y, _ := strconv.Atoi(line_parts[1])
		points = append(points, Point{x, y})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	println("Part 1:", part_1(points))
	println("Part 2:", part_2(points))
}
