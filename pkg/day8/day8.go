package day8

import (
	"slices"
	"strconv"
	"strings"
)

var P1Limit int = 1000

type junction struct {
	x int
	y int
	z int
}

type edge struct {
	a    *junction
	b    *junction
	dist uint
}

func Part1(input string) string {
	circuits, _ := buildCircuits(input, &P1Limit)

	acc := 1
	for _, circuit := range circuits[:3] {
		acc *= len(circuit)
	}

	return strconv.Itoa(acc)
}

func Part2(input string) string {
	_, edge := buildCircuits(input, nil)

	return strconv.Itoa(edge.a.x * edge.b.x)
}

func parseInput(input string) []*junction {
	rows := strings.Split(input, "\n")
	points := make([]*junction, len(rows))
	for i := range rows {
		xyz := strings.Split(rows[i], ",")
		p := &junction{x: 0, y: 0, z: 0}
		p.x, _ = strconv.Atoi(xyz[0])
		p.y, _ = strconv.Atoi(xyz[1])
		p.z, _ = strconv.Atoi(xyz[2])
		points[i] = p
	}

	return points
}

// Since we don't really care about the actual distance, just the "weight",
// we skip doing an expensive sqrt and we can use simpler integer math.
func distance(p1, p2 *junction) uint {
	dx := uint(p2.x - p1.x)
	dy := uint(p2.y - p1.y)
	dz := uint(p2.z - p1.z)
	return dx*dx + dy*dy + dz*dz
}

func buildEdges(junctions []*junction) []*edge {
	n := len(junctions)
	edges := make([]*edge, 0, n*(n-1)/2)
	for i := range junctions {
		for j := i + 1; j < len(junctions); j++ {
			dist := distance(junctions[i], junctions[j])
			edges = append(edges, &edge{
				a:    junctions[i],
				b:    junctions[j],
				dist: dist,
			})
		}
	}

	slices.SortFunc(edges, func(a, b *edge) int {
		return int(a.dist - b.dist)
	})

	return edges
}

func buildCircuits(input string, limit *int) ([]map[*junction]bool, *edge) {
	junctions := parseInput(input)
	edges := buildEdges(junctions)

	circuits := make([]map[*junction]bool, len(junctions))
	for i, j := range junctions {
		circuits[i] = make(map[*junction]bool)
		circuits[i][j] = true
	}

	i := 0
	var edge *edge
	for len(circuits) > 1 {
		aIdx := -1
		bIdx := -1
		edge = edges[i]
		for i, circuit := range circuits {
			_, ok := circuit[edge.a]
			if ok {
				aIdx = i
			}
			_, ok = circuit[edge.b]
			if ok {
				bIdx = i
			}
		}

		if aIdx > -1 && bIdx > -1 && aIdx != bIdx {
			for b := range circuits[bIdx] {
				circuits[aIdx][b] = true
			}
			circuits[bIdx] = circuits[len(circuits)-1]
			circuits = circuits[:len(circuits)-1]
		} else if aIdx > -1 {
			circuits[aIdx][edge.b] = true
		} else if bIdx > -1 {
			circuits[bIdx][edge.a] = true
		} else {
			newCirc := make(map[*junction]bool)
			newCirc[edge.a] = true
			newCirc[edge.b] = true
			circuits = append(circuits, newCirc)
		}

		i++
		if limit != nil && i == *limit {
			break
		}
	}

	slices.SortFunc(circuits, func(a, b map[*junction]bool) int {
		return len(b) - len(a)
	})

	return circuits, edge
}
