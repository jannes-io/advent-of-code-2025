package day1

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	instructions := strings.SplitSeq(input, "\n")

	ptr := 50
	zeros := 0

	for i := range instructions {
		dir := string(i[0])
		amount, _ := strconv.Atoi(i[1:])

		if dir == "L" {
			ptr -= amount
		} else {
			ptr += amount
		}
		ptr %= 100

		if ptr < 0 {
			ptr += 100
		}
		if ptr == 0 {
			zeros++
		}
	}

	return strconv.Itoa(zeros)
}

func Part2(input string) string {
	instructions := strings.SplitSeq(input, "\n")

	ptr := 50
	zeros := 0

	for i := range instructions {
		dir := string(i[0])
		amount, _ := strconv.Atoi(i[1:])

		for range amount {
			if dir == "L" {
				ptr--
			} else {
				ptr++
			}
			ptr %= 100

			if ptr < 0 {
				ptr += 100
			}
			if ptr == 0 {
				zeros++
			}
		}
	}

	return strconv.Itoa(zeros)
}
