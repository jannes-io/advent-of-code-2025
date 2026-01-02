package day1

import (
	"testing"

	"github.com/jannes-io/aoc2025/pkg/testlib"
)

const input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func Test(t *testing.T) {
	testlib.Assert(t, Part1(input), "3")
	testlib.Assert(t, Part2(input), "6")
}
