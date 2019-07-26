package guf

import (
	"fmt"
	"testing"
)

func TestUnionByHeight(t *testing.T) {
	g := NewGuf()
	unionFind10(t, g)
}

func TestUnionBySize(t *testing.T) {
	g := NewGuf()
	g.SetUnionBySize()
	unionFind10(t, g)
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
	for i := 0; i < 10; i++ {
		root := g.Find(elems[i])
		fmt.Printf("root of elem %d : %+v\n", i, root)
	}
}
