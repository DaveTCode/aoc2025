package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type JunctionBox struct {
	X int
	Y int
	Z int
}

type Distance struct {
	box_1    JunctionBox
	box_2    JunctionBox
	distance float64
}

func calculateDistance(box1 JunctionBox, box2 JunctionBox) float64 {
	return math.Sqrt(
		math.Pow(float64(box1.X-box2.X), 2) +
			math.Pow(float64(box1.Y-box2.Y), 2) +
			math.Pow(float64(box1.Z-box2.Z), 2))
}

func calculateDistances(boxes []JunctionBox) []Distance {
	distances := []Distance{}

	for i, box := range boxes {
		for _, other_box := range boxes[i:] {
			if box.X != other_box.X || box.Y != other_box.Y || box.Z != other_box.Z {
				distances = append(distances, Distance{
					box, other_box, calculateDistance(box, other_box),
				})
			}
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	return distances
}

func part_1(boxes []JunctionBox, count int) int {
	distances := calculateDistances(boxes)
	sets := map[int][]JunctionBox{}
	circuits := map[JunctionBox]int{}
	next_circuit := 0

	connections := 1
	for i := range count {
		next_shortest := distances[i]
		_, lhs_in := circuits[next_shortest.box_1]
		_, rhs_in := circuits[next_shortest.box_2]

		//fmt.Println("=====")
		fmt.Println(next_shortest)
		//fmt.Println(distances[x : x+4])
		fmt.Println(sets)
		//fmt.Println(circuits)
		fmt.Println("=====")

		if !lhs_in && !rhs_in {
			sets[next_circuit] = []JunctionBox{next_shortest.box_1, next_shortest.box_2}
			circuits[next_shortest.box_1] = next_circuit
			circuits[next_shortest.box_2] = next_circuit
			next_circuit++
			connections++
		} else if lhs_in && !rhs_in {
			circuit, _ := circuits[next_shortest.box_1]
			circuits[next_shortest.box_2] = circuit
			sets[circuit] = append(sets[circuit], next_shortest.box_2)
			connections++
		} else if !lhs_in && rhs_in {
			circuit, _ := circuits[next_shortest.box_2]
			circuits[next_shortest.box_1] = circuit
			sets[circuit] = append(sets[circuit], next_shortest.box_1)
			connections++
		} else {
			box_1_circuit, _ := circuits[next_shortest.box_1]
			box_2_circuit, _ := circuits[next_shortest.box_2]

			if box_1_circuit != box_2_circuit {
				for _, box := range sets[box_2_circuit] {
					circuits[box] = box_1_circuit
					sets[box_1_circuit] = append(sets[box_1_circuit], box)
				}

				delete(sets, box_2_circuit)
				connections++
			}
		}
	}

	sizes := make([]int, len(sets))
	for _, v := range sets {
		sizes = append(sizes, len(v))
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	fmt.Println(sizes)

	return sizes[0] * sizes[1] * sizes[2]
}

func part_2(boxes []JunctionBox) int {
	distances := calculateDistances(boxes)
	sets := map[int][]JunctionBox{}
	circuits := map[JunctionBox]int{}
	next_circuit := 0

	for _, next_shortest := range distances {
		_, lhs_in := circuits[next_shortest.box_1]
		_, rhs_in := circuits[next_shortest.box_2]

		//fmt.Println("=====")
		fmt.Println(next_shortest)
		//fmt.Println(distances[x : x+4])
		fmt.Println(sets)
		//fmt.Println(circuits)
		fmt.Println("=====")

		if !lhs_in && !rhs_in {
			sets[next_circuit] = []JunctionBox{next_shortest.box_1, next_shortest.box_2}
			circuits[next_shortest.box_1] = next_circuit
			circuits[next_shortest.box_2] = next_circuit
			next_circuit++
		} else if lhs_in && !rhs_in {
			circuit, _ := circuits[next_shortest.box_1]
			circuits[next_shortest.box_2] = circuit
			sets[circuit] = append(sets[circuit], next_shortest.box_2)
		} else if !lhs_in && rhs_in {
			circuit, _ := circuits[next_shortest.box_2]
			circuits[next_shortest.box_1] = circuit
			sets[circuit] = append(sets[circuit], next_shortest.box_1)
		} else {
			box_1_circuit, _ := circuits[next_shortest.box_1]
			box_2_circuit, _ := circuits[next_shortest.box_2]

			if box_1_circuit != box_2_circuit {
				for _, box := range sets[box_2_circuit] {
					circuits[box] = box_1_circuit
					sets[box_1_circuit] = append(sets[box_1_circuit], box)
				}

				delete(sets, box_2_circuit)
			}
		}

		if len(sets) == 1 && len(sets[circuits[next_shortest.box_1]]) == len(boxes) {
			// Got back to 1 set! Well done team
			return next_shortest.box_1.X * next_shortest.box_2.X
		}
	}

	fmt.Println(sets, len(sets), len(sets[5]), len(boxes))

	// Bad
	return -1
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	boxes := []JunctionBox{}
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		X, _ := strconv.Atoi(parts[0])
		Y, _ := strconv.Atoi(parts[1])
		Z, _ := strconv.Atoi(parts[2])
		boxes = append(boxes, JunctionBox{X, Y, Z})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := 10
	if os.Args[1] == "input" {
		count = 1000
	}

	println("Part 1:", part_1(boxes, count))
	println("Part 2:", part_2(boxes))
}
