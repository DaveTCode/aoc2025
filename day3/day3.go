package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func part_1(banks [][]byte) int {
	result := 0

	for _, bank := range banks {
		firstDigit := 0
		secondDigit := 0
		for ix, chr := range bank {
			i := int(chr - 48) // Convert from ascii byte to int
			if i > firstDigit && len(bank)-1 != ix {
				firstDigit = i
				secondDigit = 0
			} else if i > secondDigit {
				secondDigit = i
			}
		}

		println(firstDigit, secondDigit)

		result += secondDigit + (firstDigit * 10)
	}

	return result
}

func part_2(banks [][]byte) int {
	result := 0

	for _, bank := range banks {
		total := 0
		offset := 0

		for ix := range 12 {
			best := 0
			new_offset := offset
			for bank_ix, chr := range bank[offset : len(bank)-(11-ix)] {
				voltage := int(chr - 48) // Convert ascii to int

				if voltage > best {
					best = voltage
					new_offset = offset + bank_ix + 1
				}
			}

			offset = new_offset
			total += best * int(math.Pow10(11-ix))
		}
		println(total)
		result += total
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

	banks := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	println("Part 1:", part_1(banks))
	println("Part 2:", part_2(banks))
}
