package day6

import (
	"testing"

	"github.com/jannes-io/aoc2025/pkg/testlib"
)

const input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func Test(t *testing.T) {
	testlib.Assert(t, Part1(input), "4277556")
	testlib.Assert(t, Part2(input), "3263827")
}
