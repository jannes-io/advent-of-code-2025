package day5

import (
	"strconv"
	"strings"
)

type ingredientRange struct {
	min uint64
	max uint64
}

func Part1(input string) string {
	ranges, ingredients := parseInput(input)
	ranges = reduceRanges(ranges)

	fresh := 0
	for _, i := range ingredients {
		for _, r := range ranges {
			if inRange(r, i) {
				fresh++
				break
			}
		}
	}

	return strconv.Itoa(fresh)
}

func Part2(input string) string {
	ranges, _ := parseInput(input)
	return strconv.FormatUint(addRanges(reduceRanges(ranges)), 10)
}

func addRanges(ranges []*ingredientRange) uint64 {
	var c uint64 = 0
	for _, r := range ranges {
		c += r.max - r.min
	}
	return c + uint64(len(ranges))
}

func reduceRanges(ranges []*ingredientRange) []*ingredientRange {
	resized := 1
	for resized > 0 {
		ranges, resized = reduceRange(ranges)
	}
	return ranges
}

func reduceRange(ranges []*ingredientRange) ([]*ingredientRange, int) {
	resized := 0
	newRanges := make([]*ingredientRange, 0)
	for _, r := range ranges {
		found := false
		for _, nr := range newRanges {
			if overlap(r, nr) {
				if r.min < nr.min {
					nr.min = r.min
				}
				if r.max > nr.max {
					nr.max = r.max
				}
				resized++
				found = true
				break
			}
		}
		if !found {
			newRanges = append(newRanges, r)
		}
	}

	return newRanges, resized
}

func overlap(r1 *ingredientRange, r2 *ingredientRange) bool {
	return inRange(r1, r2.min) || inRange(r1, r2.max) || inRange(r2, r1.min) || inRange(r2, r1.max)
}

func inRange(r *ingredientRange, v uint64) bool {
	return v >= r.min && v <= r.max
}

func parseInput(input string) ([]*ingredientRange, []uint64) {
	rangeStr, ingredientStr, _ := strings.Cut(input, "\n\n")

	r1 := strings.Split(rangeStr, "\n")
	ranges := make([]*ingredientRange, len(r1))
	for i, r := range r1 {
		minStr, maxStr, _ := strings.Cut(r, "-")
		min, _ := strconv.ParseUint(minStr, 10, 64)
		max, _ := strconv.ParseUint(maxStr, 10, 64)

		ranges[i] = &ingredientRange{min, max}
	}

	i1 := strings.Split(ingredientStr, "\n")
	ingredients := make([]uint64, len(i1))
	for i, in := range i1 {
		ingredients[i], _ = strconv.ParseUint(in, 10, 64)
	}

	return ranges, ingredients
}
