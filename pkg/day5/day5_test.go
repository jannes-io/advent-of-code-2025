package day5

import (
	"testing"

	"github.com/jannes-io/aoc2025/pkg/testlib"
)

const input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Test(t *testing.T) {
	testlib.Assert(t, Part1(input), "3")
	testlib.Assert(t, Part2(input), "14")
}

// Our sample input works, but real data fails and is too high.
// So let's test some edge cases..

func TestEdgeOverlap(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 10, max: 20},
		{min: 8, max: 12},
		{min: 14, max: 22},
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 1)
	testlib.Assert(t, ranges[0].min, 8)
	testlib.Assert(t, ranges[0].max, 22)
	testlib.Assert(t, addRanges(ranges), 15)
}

func TestEdgeContinousMax(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 10, max: 20},
		{min: 20, max: 30},
		{min: 30, max: 40},
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 1)
	testlib.Assert(t, ranges[0].min, 10)
	testlib.Assert(t, ranges[0].max, 40)
	testlib.Assert(t, addRanges(ranges), 31)
}

func TestEdgeContinousMin(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 30, max: 40},
		{min: 20, max: 30},
		{min: 10, max: 20},
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 1)
	testlib.Assert(t, ranges[0].min, 10)
	testlib.Assert(t, ranges[0].max, 40)
	testlib.Assert(t, addRanges(ranges), 31)
}

func TestEdgeEngulfed(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 10, max: 20},
		{min: 10, max: 12}, // min edge match
		{min: 14, max: 20}, // max edge match
		{min: 14, max: 16}, // complete overlap
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 1)
	testlib.Assert(t, ranges[0].min, 10)
	testlib.Assert(t, ranges[0].max, 20)
	testlib.Assert(t, addRanges(ranges), 11)
}

func TestEdgeEngulfedLarger(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 12, max: 18},
		{min: 10, max: 20}, // the second one overlaps
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 1)
	testlib.Assert(t, ranges[0].min, 10)
	testlib.Assert(t, ranges[0].max, 20)
	testlib.Assert(t, addRanges(ranges), 11)
}

func TestEdgeOffByOne(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 10, max: 20},
		{min: 5, max: 9},
		{min: 21, max: 25},
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 3)
	testlib.Assert(t, addRanges(ranges), 21)
}

func TestZeroRange(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 1, max: 1},
		{min: 3, max: 3},
		{min: 5, max: 5},
		{min: 5, max: 6},
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 3)
	testlib.Assert(t, addRanges(ranges), 4)
}

func TestIntSize(t *testing.T) {
	ranges := []*ingredientRange{
		{min: 500789379744360, max: 501460137450447},
		{min: 493906974314704, max: 494688955465215},
	}

	ranges = reduceRanges(ranges)
	testlib.Assert(t, len(ranges), 2)
	testlib.Assert(t, addRanges(ranges), 1452738856600)
}

func TestParseInput(t *testing.T) {
	ranges, ingredients := parseInput(input)

	testlib.Assert(t, len(ranges), 4)
	testlib.Assert(t, len(ingredients), 6)
}
