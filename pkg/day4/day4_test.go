package day4

import (
	"testing"

	"github.com/jannes-io/aoc2025/pkg/testlib"
)

const input = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func Test(t *testing.T) {
	testlib.Assert(t, Part1(input), "13")
	testlib.Assert(t, Part2(input), "43")
}
