// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package ed_test

import (
	"fmt"

	"github.com/soniakeys/ed"
)

func ExampleAdjacencyList_Cyclic() {
	g := ed.AdjacencyList{
		0: {1, 2},
		1: {2},
		2: {3},
		3: {},
	}
	fmt.Println(g.Cyclic())
	g[3] = []int{2}
	fmt.Println(g.Cyclic())
	// Output:
	// false
	// true
}

func ExampleAdjacencyList_Topological() {
	g := ed.AdjacencyList{
		1: {2},
		3: {1, 2},
		4: {3, 2},
	}
	fmt.Println(g.Topological())
	g[2] = []int{3}
	fmt.Println(g.Topological())
	// Output:
	// [4 3 1 2 0]
	// []
}

func ExampleAdjacencyList_Tarjan() {
	g := ed.AdjacencyList{
		0: {1},
		1: {4, 2, 5},
		2: {3, 6},
		3: {2, 7},
		4: {5, 0},
		5: {6},
		6: {5},
		7: {3, 6},
	}
	scc := g.Tarjan()
	for _, c := range scc {
		fmt.Println(c)
	}
	fmt.Println(len(scc))
	// Output:
	// [6 5]
	// [7 3 2]
	// [4 1 0]
	// 3
}

/*
func ExampleAdjacencyList_Ok() {
	var g ed.AdjacencyList
	fmt.Println(g.Ok()) // zero value adjacency list is valid
	g = ed.AdjacencyList{
		0: {1},
	}
	fmt.Println(g.Ok()) // arc 0 to 1 invalid with only one node
	// Output:
	// true
	// false
}

func ExampleAdjacencyList_Undirected() {
	g := ed.AdjacencyList{
		0: {1, 2},
		2: {0},
	}
	fmt.Println(g.Undirected())
	g[1] = append(g[1], 0)
	fmt.Println(g.Undirected())
	// Output:
	// false 0 1
	// true -1 -1
}
func ExampleAdjacencyList_Simple() {
	g := ed.AdjacencyList{
		2: {0, 1},
	}
	fmt.Println(g.Simple())
	g[1] = []int{1} // loop
	fmt.Println(g.Simple())
	g[1] = nil
	g[2] = append(g[2], 0) // parallel arc
	fmt.Println(g.Simple())
	// Output:
	// true -1
	// false 1
	// false 2
}

func ExampleAdjacencyList_Bipartite() {
	g := ed.AdjacencyList{
		0: {3},
		1: {3},
		2: {3, 4},
		3: {0, 1, 2},
		4: {2},
	}
	b, c1, c2, oc := g.Bipartite(0)
	if b {
		fmt.Println(
			strconv.FormatInt(c1.Int64(), 2),
			strconv.FormatInt(c2.Int64(), 2))
	}
	g[3] = append(g[3], 4)
	g[4] = append(g[4], 3)
	b, c1, c2, oc = g.Bipartite(0)
	if !b {
		fmt.Println(oc)
	}
	// Output:
	// 111 11000
	// [3 4 2]
}

func ExampleAdjacencyList_BiconnectedComponents() {
	g := ed.AdjacencyList{
		0:  {1, 7},
		1:  {2, 4, 0},
		2:  {3, 1},
		3:  {2, 4},
		4:  {3, 1},
		5:  {6, 12},
		6:  {5, 12, 8},
		7:  {8, 0},
		8:  {6, 7, 9, 10},
		9:  {8},
		10: {8, 13, 11},
		11: {10},
		12: {5, 6, 13},
		13: {12, 10},
	}
	b := ed.NewBiconnectedComponents(g)
	b.Find(0)
	fmt.Println("n: cut from")
	for n, f := range b.From {
		fmt.Printf("%d: %d %d\n",
			n, b.Cuts.Bit(n), f)
	}
	fmt.Println("Leaves:", b.Leaves)
	// wip
	// Output:
}
*/