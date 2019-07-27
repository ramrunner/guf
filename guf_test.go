package guf

import (
	"fmt"
	"testing"
)

func TestUnionByHeight(t *testing.T) {
	t.Logf("Running with union by height")
	g := NewGuf()
	unionFind10(t, g)
}

func TestUnionBySize(t *testing.T) {
	t.Logf("Running with union by size")
	g := NewGuf()
	g.SetUnionBySize()
	unionFind10(t, g)
}

func printElemSlice(sname string, elems []*SetElem) {
	for i := range elems {
		fmt.Printf("%s: %s\n", sname, elems[i])
	}
}

func unionFind10(t *testing.T, g *Guf) {
	elems := []*SetElem{}
	for i := 0; i < 10; i++ {
		elems = append(elems, g.RegisterNew())
	}
	g.Union(elems[0], elems[1])
	g.Union(elems[2], elems[3])
	g.Union(elems[2], elems[4])
	g.Union(elems[2], elems[5])
	g.Union(elems[5], elems[6])
	// in the set of 2 now we should have: 5 elements
	// and in the set of 1: 2 elements
	for i := 0; i < 10; i++ {
		root := g.Find(elems[i])
		fmt.Printf("root of elem %d : %+v\n", i, root)
	}
	allIn2 := elems[2].AllInSet()
	allIn1 := elems[1].AllInSet()
	printElemSlice("In set of 2", allIn2)
	printElemSlice("In set of 1", allIn1)
	if len(allIn2) != 5 || len(allIn1) != 2 {
		t.Fatalf("Sets don't have the correct number of elements")
	}
}
