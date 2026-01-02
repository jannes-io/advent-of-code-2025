package day3

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input string) string {
	return run(input, 2)
}

func Part2(input string) string {
	return run(input, 12)
}

func run(input string, bc int) string {
	var total uint64 = 0
	for bank := range strings.SplitSeq(input, "\n") {
		total += processBank(bank, bc)
	}

	return strconv.FormatUint(total, 10)
}

func processBank(bank string, bc int) uint64 {
	l := len(bank)
	batteries := make([]uint64, l)
	for i := range l {
		batteries[i] = uint64(bank[i] - 48)
	}

	var counter uint64 = 0
	for i := range bc {
		bci := bc-i
		pickable := batteries[:len(batteries)-bci+1]

		var largest uint64 = 0
		nextOffset := 0
		for pos, v := range pickable {
			if v > largest {
				largest = v
				nextOffset = pos + 1
			}
		}

		batteries = batteries[nextOffset:]
		counter += uint64(math.Pow10(bci-1)) * largest
	}

	return counter
}
