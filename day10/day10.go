package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Indicators          string
	WiringSchematics    [][]int
	JoltageRequirements string
}

type State struct {
	state string
	depth int
}

func bfs_machine(machine Machine) int {
	queue := []State{{strings.ReplaceAll(machine.Indicators, "#", "."), 0}}
	visited := map[string]bool{}

	for len(queue) > 0 {
		next_state := queue[0]
		next := next_state.state
		depth := next_state.depth
		queue = queue[1:]

		if next == machine.Indicators {
			return depth
		}

		for _, button := range machine.WiringSchematics {
			new_state := []byte(next)
			for _, wire := range button {
				chr := new_state[wire]
				new_chr := byte('#')
				if chr == '#' {
					new_chr = '.'
				}
				new_state[wire] = new_chr
			}
			if _, ok := visited[string(new_state)]; !ok {
				queue = append(queue, State{string(new_state), depth + 1})
			}
		}

		visited[next] = true
	}

	return 0
}

func part_1(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		queue := []State{{strings.ReplaceAll(machine.Indicators, "#", "."), 0}}
		visited := map[string]bool{}

		for len(queue) > 0 {
			next_state := queue[0]
			next := next_state.state
			depth := next_state.depth
			queue = queue[1:]

			if next == machine.Indicators {
				total += depth
				break
			}

			for _, button := range machine.WiringSchematics {
				new_state := []byte(next)
				for _, wire := range button {
					chr := new_state[wire]
					new_chr := byte('#')
					if chr == '#' {
						new_chr = '.'
					}
					new_state[wire] = new_chr
				}
				if _, ok := visited[string(new_state)]; !ok {
					queue = append(queue, State{string(new_state), depth + 1})
				}
			}

			visited[next] = true
		}
	}

	return total
}

func recurse_part_2(variables []int, mins []int, maxs []int, joltages []int) {
	new_variables := make([]int, variables[])
	for _, variable := range variables {
		if variable == -1 {
			variables[]
		}
	}
}

func part_2(machines []Machine) int {
	total := 0

	for _, machine := range machines {
		joltage_requirements := strings.Split(machine.JoltageRequirements, ",")
		joltages := make([]int, len(joltage_requirements))
		for i, x := range joltage_requirements {
			joltages[i], _ = strconv.Atoi(x)
		}

		mins := make([]int, len(machine.WiringSchematics)) // Maximum number of times a wire can be used
		maxs := make([]int, len(machine.WiringSchematics)) // Minimum number of times a wire can be used
		for i, wire := range machine.WiringSchematics {
			for _, j := range wire {
				if joltages[j] > maxs[i] {
					maxs[i] = joltages[j]
				}
			}
		}

		for i, wire := range machine.WiringSchematics {
			for x := mins[i]; x <= maxs[i]; x++ {
				
			}
		}

		fmt.Println(machine.WiringSchematics)
		fmt.Println(joltages)
		fmt.Println(mins)
		fmt.Println(maxs)
	}

	return total
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	machines := []Machine{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		machine := Machine{}
		for ix, part := range parts {
			if ix == 0 {
				machine.Indicators = strings.Trim(part, "[]")
			} else if ix == len(parts)-1 {
				machine.JoltageRequirements = strings.Trim(part, "{}")
			} else {
				wiringSchematics := []int{}
				for x := range strings.SplitSeq(strings.Trim(part, "()"), ",") {
					num, _ := strconv.Atoi(x)
					wiringSchematics = append(wiringSchematics, num)
				}
				machine.WiringSchematics = append(machine.WiringSchematics, wiringSchematics)
			}
		}
		machines = append(machines, machine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(machines)

	println("Part 1:", part_1(machines))
	println("Part 2:", part_2(machines))
}
