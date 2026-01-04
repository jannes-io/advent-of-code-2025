package day7

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	splits, _ := run(input)
	return strconv.Itoa(splits)
}

func Part2(input string) string {
	_, timelines := run(input)
	return strconv.Itoa(timelines)
}

func run(input string) (int, int) {
	rows := strings.SplitSeq(input, "\n")

	splits := 0
	var beams []int
	for row := range rows {
		if beams == nil {
			beams = initBeam(row)
			continue
		}

		for i, c := range row {
			if c != '^' || beams[i] == 0 {
				continue
			}
			splits++
			beams[i-1] += beams[i]
			beams[i+1] += beams[i]
			beams[i] = 0
		}
	}

	timelines := 0
	for _, b := range beams {
		timelines += b
	}

	return splits, timelines
}

func initBeam(row string) []int {
	beams := make([]int, len(row))
	for i, c := range row {
		if c == 'S' {
			beams[i] = 1
		} else {
			beams[i] = 0
		}
	}
	return beams
}
