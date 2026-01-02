package f

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadInput(day int, fromTestFile bool) (string, error) {
	path := "/home/developer/projects/aoc2025/inputs/day" + strconv.Itoa(day)
	if fromTestFile {
		path += "_test"
	}
	path += ".txt"

	input, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	normalized := strings.ReplaceAll(string(input), "\r", "")
	normalized, _ = strings.CutSuffix(normalized, "\n")
	return normalized, nil
}

type DayFn func(string) string

/*
 * Split every Nth character, for example AABBCC,2 becomes [AA, BB, CC]
 */
func SplitEveryN(s string, n int) []string {
	sl := math.Ceil(float64(len(s)) / float64(n))
	result := make([]string, int(sl))

	i := 0
	for _, c := range s {
		result[i/n] += string(c)
		i++
	}

	return result
}

func AllEq[T comparable](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	first := slice[0]
	for _, v := range slice {
		if v != first {
			return false
		}
	}

	return true
}
