package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jannes-io/aoc2025/pkg/day1"
	"github.com/jannes-io/aoc2025/pkg/day2"
	"github.com/jannes-io/aoc2025/pkg/day3"
	"github.com/jannes-io/aoc2025/pkg/day4"
	"github.com/jannes-io/aoc2025/pkg/day5"
	"github.com/jannes-io/aoc2025/pkg/day6"
	"github.com/jannes-io/aoc2025/pkg/day7"
	"github.com/jannes-io/aoc2025/pkg/day8"
	"github.com/jannes-io/aoc2025/pkg/f"
)

func main() {
	var day, part int
	flag.IntVar(&day, "d", 0, "Which day")
	flag.IntVar(&part, "p", 0, "Part 1 or 2")
	flag.Parse()

	fmt.Printf("Running day %d, part %d...\n", day, part)

	if day <= 0 || part <= 0 {
		log.Fatal("Invalid day/part provided.")
	}

	input, err := f.ReadInput(day, false)
	if err != nil {
		log.Fatal(err)
	}

	days := map[int]map[int]f.DayFn{
		1: {
			1: day1.Part1,
			2: day1.Part2,
		},
		2: {
			1: day2.Part1,
			2: day2.Part2,
		},
		3: {
			1: day3.Part1,
			2: day3.Part2,
		},
		4: {
			1: day4.Part1,
			2: day4.Part2,
		},
		5: {
			1: day5.Part1,
			2: day5.Part2,
		},
		6: {
			1: day6.Part1,
			2: day6.Part2,
		},
		7: {
			1: day7.Part1,
			2: day7.Part2,
		},
		8: {
			1: day8.Part1,
			2: day8.Part2,
		},
	}
	result := days[day][part](string(input))
	fmt.Println(result)
}
