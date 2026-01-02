package f_test

import (
	"testing"

	"github.com/jannes-io/aoc2025/pkg/testlib"
	"github.com/jannes-io/aoc2025/pkg/f"
)

func TestSplitEveryN(t *testing.T) {
	split := f.SplitEveryN("AABBC", 2)
	testlib.Assert(t, "AA", split[0])
	testlib.Assert(t, "BB", split[1])
	testlib.Assert(t, "C", split[2])
}

func TestAllEq(t *testing.T) {
	values := []string{}
	testlib.Assert(t, true, f.AllEq(values))

	values = append(values, "AA")
	testlib.Assert(t, true, f.AllEq(values))


	values = append(values, "AA")
	testlib.Assert(t, true, f.AllEq(values))


	values = append(values, "BB")
	testlib.Assert(t, false, f.AllEq(values))
}
