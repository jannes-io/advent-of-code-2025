package day4

import (
	"strconv"
	"strings"
)

func Part1(input string) string {
	grid := gridify(input)
	_, removed := removePaper(grid)

	return strconv.Itoa(removed)
}

func Part2(input string) string {
	grid := gridify(input)

	removed := 0
	removedStep := 1
	for removedStep > 0 {
		grid, removedStep = removePaper(grid)
		removed += removedStep
	}

	return strconv.Itoa(removed)
}

func removePaper(grid [][]string) ([][]string, int) {
	h, w := len(grid), len(grid[0])
	grid2 := make([][]string, h)

	removed := 0
	for y := range h {
		row := make([]string, w)
		copy(row, grid[y])
		grid2[y] = row

		for x := range w {
			if grid[y][x] != "@" {
				continue
			}

			c := 0
			for dy := -1; dy < 2; dy++ {
				yy := y + dy
				if yy < 0 || yy >= h {
					continue
				}

				for dx := -1; dx < 2; dx++ {
					if dy == 0 && dx == 0 {
						continue
					}

					xx := x + dx
					if xx < 0 || xx >= w {
						continue
					}

					if grid[yy][xx] == "@" {
						c++
					}
				}
			}

			if c < 4 {
				removed++
				grid2[y][x] = "X"
			}
		}
	}
	return grid2, removed
}

func gridify(input string) [][]string {
	rows := strings.SplitSeq(input, "\n")

	grid := make([][]string, 0)
	for row := range rows {
		gridRow := make([]string, len(row))
		for p, c := range row {
			gridRow[p] = string(c)
		}
		grid = append(grid, gridRow)
	}
	return grid
}
