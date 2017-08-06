// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package graph

// adj_RO.go is code generated from adj_cg.go by directives in graph.go.
// Editing adj_cg.go is okay.
// DO NOT EDIT adj_RO.go.  The RO is for Read Only.

import (
	"math/rand"

	"github.com/soniakeys/bits"
)

// ArcDensity returns density for an simple directed graph.
//
// See also ArcDensity function.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) ArcDensity() float64 {
	return ArcDensity(len(g), g.ArcSize())
}

// ArcSize returns the number of arcs in g.
//
// Note that for an undirected graph without loops, the number of undirected
// edges -- the traditional meaning of graph size -- will be ArcSize()/2.
// On the other hand, if g is an undirected graph that has or may have loops,
// g.ArcSize()/2 is not a meaningful quantity.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) ArcSize() int {
	m := 0
	for _, to := range g {
		m += len(to)
	}
	return m
}

// BoundsOk validates that all arcs in g stay within the slice bounds of g.
//
// BoundsOk returns true when no arcs point outside the bounds of g.
// Otherwise it returns false and an example arc that points outside of g.
//
// Most methods of this package assume the BoundsOk condition and may
// panic when they encounter an arc pointing outside of the graph.  This
// function can be used to validate a graph when the BoundsOk condition
// is unknown.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) BoundsOk() (ok bool, fr NI, to NI) {
	for fr, to := range g {
		for _, to := range to {
			if to < 0 || to >= NI(len(g)) {
				return false, NI(fr), to
			}
		}
	}
	return true, -1, to
}

// BreadthFirstTraverse traverses a directed or undirected graph in breadth
// first order.
//
// There are equivalent labeled and unlabeled versions of this method.
//
// See also alt.BreadthFirst, a variant with more options, and
// alt.BreadthFirst2, a direction optimizing variant.
func (g AdjacencyList) BreadthFirstTraverse(start NI, visit func(NI)) {
	v := bits.New(len(g))
	v.SetBit(int(start), 1)
	visit(start)
	var next []NI
	for frontier := []NI{start}; len(frontier) > 0; {
		for _, n := range frontier {
			for _, nb := range g[n] {
				if v.Bit(int(nb)) == 0 {
					v.SetBit(int(nb), 1)
					visit(nb)
					next = append(next, nb)
				}
			}
		}
		frontier, next = next, frontier[:0]
	}
}

// Copy makes a deep copy of g.
// Copy also computes the arc size ma, the number of arcs.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) Copy() (c AdjacencyList, ma int) {
	c = make(AdjacencyList, len(g))
	for n, to := range g {
		c[n] = append([]NI{}, to...)
		ma += len(to)
	}
	return
}

// HasArc returns true if g has any arc from node `fr` to node `to`.
//
// Also returned is the index within the slice of arcs from node `fr`.
// If no arc from `fr` to `to` is present, HasArc returns false, -1.
//
// There are equivalent labeled and unlabeled versions of this method.
//
// See also the method ParallelArcs, which finds all parallel arcs from
// `fr` to `to`.
func (g AdjacencyList) HasArc(fr, to NI) (bool, int) {
	for x, h := range g[fr] {
		if h == to {
			return true, x
		}
	}
	return false, -1
}

// AnyLoop identifies if a graph contains a loop, an arc that leads from a
// a node back to the same node.
//
// If g contains a loop, the method returns true and an example of a node
// with a loop.  If there are no loops in g, the method returns false, -1.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) AnyLoop() (bool, NI) {
	for fr, to := range g {
		for _, to := range to {
			if NI(fr) == to {
				return true, to
			}
		}
	}
	return false, -1
}

// AnyParallelMap identifies if a graph contains parallel arcs, multiple arcs
// that lead from a node to the same node.
//
// If the graph has parallel arcs, the method returns true and
// results fr and to represent an example where there are parallel arcs
// from node `fr` to node `to`.
//
// If there are no parallel arcs, the method returns false, -1 -1.
//
// Multiple loops on a node count as parallel arcs.
//
// "Map" in the method name indicates that a Go map is used to detect parallel
// arcs.  Compared to method AnyParallelSort, this gives better asymtotic
// performance for large dense graphs but may have increased overhead for
// small or sparse graphs.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) AnyParallelMap() (has bool, fr, to NI) {
	for n, to := range g {
		if len(to) == 0 {
			continue
		}
		m := map[NI]struct{}{}
		for _, to := range to {
			if _, ok := m[to]; ok {
				return true, NI(n), to
			}
			m[to] = struct{}{}
		}
	}
	return false, -1, -1
}

// IsSimple checks for loops and parallel arcs.
//
// A graph is "simple" if it has no loops or parallel arcs.
//
// IsSimple returns true, -1 for simple graphs.  If a loop or parallel arc is
// found, simple returns false and a node that represents a counterexample
// to the graph being simple.
//
// See also separate methods AnyLoop and AnyParallel.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) IsSimple() (ok bool, n NI) {
	if lp, n := g.AnyLoop(); lp {
		return false, n
	}
	if pa, n, _ := g.AnyParallelSort(); pa {
		return false, n
	}
	return true, -1
}

// IsolatedNodes returns a bitmap of isolated nodes in receiver graph g.
//
// An isolated node is one with no arcs going to or from it.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) IsolatedNodes() (i bits.Bits) {
	i = bits.New(len(g))
	i.SetAll()
	for fr, to := range g {
		if len(to) > 0 {
			i.SetBit(fr, 0)
			for _, to := range to {
				i.SetBit(int(to), 0)
			}
		}
	}
	return
}

// Order is the number of nodes in receiver g.
//
// It is simply a wrapper method for the Go builtin len().
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) Order() int {
	// Why a wrapper for len()?  Mostly for Directed and Undirected.
	// u.Order() is a little nicer than len(u.LabeledAdjacencyList).
	return len(g)
}

// ParallelArcs identifies all arcs from node `fr` to node `to`.
//
// The returned slice contains an element for each arc from node `fr` to node `to`.
// The element value is the index within the slice of arcs from node `fr`.
//
// There are equivalent labeled and unlabeled versions of this method.
//
// See also the method HasArc, which stops after finding a single arc.
func (g AdjacencyList) ParallelArcs(fr, to NI) (p []int) {
	for x, h := range g[fr] {
		if h == to {
			p = append(p, x)
		}
	}
	return
}

// Permute permutes the node labeling of receiver g.
//
// Argument p must be a permutation of the node numbers of the graph,
// 0 through len(g)-1.  A permutation returned by rand.Perm(len(g)) for
// example is acceptable.
//
// The graph is permuted in place.  The graph keeps the same underlying
// memory but values of the graph representation are permuted to produce
// an isomorphic graph.  The node previously labeled 0 becomes p[0] and so on.
// See example (or the code) for clarification.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) Permute(p []int) {
	old := append(AdjacencyList{}, g...) // shallow copy
	for fr, arcs := range old {
		for i, to := range arcs {
			arcs[i] = NI(p[to])
		}
		g[p[fr]] = arcs
	}
}

// ShuffleArcLists shuffles the arc lists of each node of receiver g.
//
// For example a node with arcs leading to nodes 3 and 7 might have an
// arc list of either [3 7] or [7 3] after calling this method.  The
// connectivity of the graph is not changed.  The resulting graph stays
// equivalent but a traversal will encounter arcs in a different
// order.
//
// If Rand r is nil, the rand package default shared source is used.
//
// There are equivalent labeled and unlabeled versions of this method.
func (g AdjacencyList) ShuffleArcLists(r *rand.Rand) {
	ri := rand.Intn
	if r != nil {
		ri = r.Intn
	}
	// Knuth-Fisher-Yates
	for _, to := range g {
		for i := len(to); i > 1; {
			j := ri(i)
			i--
			to[i], to[j] = to[j], to[i]
		}
	}
}
