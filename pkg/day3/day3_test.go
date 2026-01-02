package day3

import (
	"testing"

	"github.com/jannes-io/aoc2025/pkg/testlib"
)

const input = `987654321111111
811111111111119
234234234234278
818181911112111`

func Test(t *testing.T) {
	testlib.Assert(t, Part1(input), "357")
	testlib.Assert(t, Part2(input), "3121910778619")
}
