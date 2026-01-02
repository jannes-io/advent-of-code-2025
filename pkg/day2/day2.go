package day2

import (
	"strconv"
	"strings"

	"github.com/jannes-io/aoc2025/pkg/f"
)

func Part1(input string) string {
	return run(input, checkNumber)
}

func Part2(input string) string {
	return run(input, checkNumber2)
}

type CheckFn func(int) bool

func run(input string, check CheckFn) string {
	ranges := strings.SplitSeq(input, ",")

	var i uint64 = 0

	for r := range ranges {
		rs := strings.Split(r, "-")
		start, _ := strconv.Atoi(rs[0])
		end, _ := strconv.Atoi(rs[1])

		for n := start; n <= end; n++ {
			if check(n) {
				i += uint64(n)
			}
		}
	}

	return strconv.FormatUint(i, 10)
}

func checkNumber(n int) bool {
	s := strconv.Itoa(n)
	sl := len(s)
	if sl%2 != 0 {
		return false
	}
	sl /= 2
	return s[sl:] == s[:sl]
}

func checkNumber2(n int) bool {
	s := strconv.Itoa(n)
	sl := len(s)

	for r := 1; r <= sl; r++ {
		if sl%r != 0 {
			continue
		}

		slices := f.SplitEveryN(s, r)
		if len(slices) < 2 {
			continue
		}

		if f.AllEq(slices) {
			return true
		}
	}
	return false
}
