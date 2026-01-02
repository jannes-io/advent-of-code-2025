package day6

import (
	"regexp"
	"strconv"
	"strings"
)

type problem struct {
	operation string
	inputs    []uint64
}

func Part1(input string) string {
	total := solve(parseInput1(input))
	return strconv.FormatUint(total, 10)
}

func Part2(input string) string {
	total := solve(parseInput2(input))
	return strconv.FormatUint(total, 10)
}

func solve(problems []*problem) uint64 {
	var total uint64 = 0
	for _, p := range problems {
		if len(p.inputs) == 0 {
			continue
		}

		acc := p.inputs[0]
		for i := 1; i < len(p.inputs); i++ {
			if p.operation == "*" {
				acc *= p.inputs[i]
			} else {
				acc += p.inputs[i]
			}
		}
		total += acc
	}
	return total
}

func parseInput1(input string) []*problem {
	rows := strings.SplitSeq(input, "\n")

	inputRegex := regexp.MustCompile(`\s+`)

	var width int = 0
	cells := make([]string, 0)
	for row := range rows {
		inputs := strings.Split(inputRegex.ReplaceAllString(strings.Trim(row, " "), " "), " ")
		for _, i := range inputs {
			cells = append(cells, i)
		}

		if width == 0 {
			width = len(inputs)
		}
	}

	operations := cells[len(cells)-width:]
	values := cells[:len(cells)-width]
	height := len(values) / width

	problems := make([]*problem, width)
	for i, op := range operations {
		problems[i] = &problem{operation: op, inputs: make([]uint64, height)}
	}

	for i := 0; i < len(values); i += width {
		for j, p := range problems {
			p.inputs[i/width], _ = strconv.ParseUint(values[i+j], 10, 64)
		}
	}

	return problems
}

func parseInput2(input string) []*problem {
	rows := strings.Split(input, "\n")

	opRow := rows[len(rows)-1]

	opsRegex := regexp.MustCompile(`\s+`)
	ops := strings.Split(opsRegex.ReplaceAllString(strings.Trim(opRow, " "), " "), " ")
	problems := make([]*problem, len(ops))
	for i, op := range ops {
		problems[i] = &problem{operation: op, inputs: make([]uint64, 0)}
	}

	inputGrid := make([][]string, len(rows)-1)
	for i, row := range rows[:len(rows)-1] {
		inputGrid[i] = make([]string, len(row))
		for j, c := range row {
			inputGrid[i][j] = string(c)
		}
	}

	h := len(inputGrid)
	idx := len(problems) - 1
	for i := len(opRow) - 1; i >= 0; i-- {
		number := ""
		for j := range h {
			number += inputGrid[j][i]
		}

		number = strings.Trim(number, " ")
		if len(number) == 0 {
			idx--
			continue
		}

		asUint, _ := strconv.ParseUint(number, 10, 64)
		problems[idx].inputs = append(problems[idx].inputs, asUint)
	}

	return problems
}
