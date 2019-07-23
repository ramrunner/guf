// GUF provides a golang union find implementation that can track the connectivity
// of different elements in sets. It provides a way to the user to associate data
// with the set identifiers that GUF is using internally.
package guf

// SetElem represents a node in the forest of sets
// that can also hold a reference to user provided data.
type SetElem struct {
	parent   *SetElem
	id       int
	height   int
	size     int
	children map[int]*SetElem
	data     interface{}
}

// SetData is a convenience function that associates
// a reference to your data and returns back an integer,
// that is the identifier of that set.
func (t *SetElem) SetData(d interface{}) int {
	t.data = d
	return t.id
}

// ID returns the id of a register set.
func (t *SetElem) ID() int {
	return t.id
}

// height calculates the height of the node by issuing a call to all
// it's children and selecting the maximum.
func (t *SetElem) calcHeight() int {
	cheights := []int{}
	// if we have no children nodes we at the bottom of the tree.
	// in that case cheights will be empty and maxIntSlice will return
	// -1 which is the appropriate height for that case. Adding one
	// to that, would bring our height to 0.
	for _, v := range t.children {
		cheights = append(cheights, v.calcHeight())
	}
	cheight := maxIntSlice(cheights)
	return 1 + cheight
}

func (t *SetElem) calcSize() int {
	acc := 0
	for _, v := range t.children {
		acc = acc + v.calcSize()
	}
	return 1 + acc
}

func (t *SetElem) setParent(p *SetElem) {
	// we're already registered as a child of another parent.
	// go there and clear our reference from the children map.
	if t.parent != nil {
		t.parent.removeChild(t)
	}
	t.parent = p
	p.setChild(t)
}

func (t *SetElem) setChild(c *SetElem) {
	// update our size
	t.size += c.size
	// update our parents
	p := t.parent
	for p != nil {
		p.size += c.size
		p = p.parent
	}

	// update our height if it now smaller
	// than 1+child height
	if t.height-1 < c.height {
		t.height = 1 + c.height
		// propagate that up to our parents
		updateHeightUp(t)
	}

	t.children[c.id] = c
}

func updateHeightUp(t *SetElem) {
	p := t.parent
	curr := t
	for p != nil {
		if p.height-1 < curr.height {
			p.height = curr.height + 1
		}
		curr = p
		p = p.parent
	}
}

func (t *SetElem) removeChild(c *SetElem) {
	// update our size
	t.size -= c.size
	// if that child can be the one responsible for our height
	if c.height == t.height-1 {
		// see if we still have another child on the same height,
		// if that's the case we don't need to update anything.
		maxchildheight := 0
		for _, v := range t.children {
			if v.height > maxchildheight { // keep note of the tallest child
				maxchildheight = v.height
			}
			if v.ID() != c.ID() && v.height == t.height-1 {
				// we found another child with that height.
				goto justDelete
			}
		}
		// we now need to update our height and possibly our parents.
		t.height = 1 + maxchildheight
		updateHeightUp(t)
	}
justDelete:
	delete(t.children, c.id)
}

// Guf an instance of a Go Union Find.
// If memory to user data is held by it,
// to successfully collect back the memory, the caller
// should clear the variable by running a NewGuf() on it,
// or set it to nil in order to allow garbage collection.
type Guf struct {
	setForest []*SetElem
	maxElem   int
	unionfunc func(a, b, pa, pb *SetElem)
}

// NewGuf creates a new Guf object that holds a still empty forrest
// of sets.
func NewGuf() *Guf {
	g := &Guf{
		setForest: make([]*SetElem, 0),
		maxElem:   0,
	}
	g.SetUnionByHeight()
	return g
}

// RegisterNew returns a newly allocated Set Element in the set forest
// that has a unique ID in that Guf.
func (g *Guf) RegisterNew() *SetElem {
	t := &SetElem{
		parent:   nil,
		id:       g.maxElem,
		height:   0,
		size:     0,
		children: make(map[int]*SetElem),
	}
	g.maxElem++
	return t
}

// Union unifies two set elements.
func (g *Guf) Union(a, b *SetElem) {
	pa, pb := g.Find(a), g.Find(b)
	if pa.ID() == pb.ID() { // they belong in the same set, no need to union.
		return
	}
	g.unionfunc(a, b, pa, pb)
}

func (g *Guf) unionByHeight(a, b, pa, pb *SetElem) {
	ha, hb := a.height, b.height
	if ha <= hb {
		pa.setParent(pb)
	} else {
		pb.setParent(pa)
	}
}

func (g *Guf) unionBySize(a, b, pa, pb *SetElem) {
	sa, sb := a.size, b.size
	if sa <= sb {
		pa.setParent(pb)
	} else {
		pb.setParent(pa)
	}
}

// Find returns the parent SetElement of the whole set that it's
// argument belongs in.
func (g *Guf) Find(a *SetElem) *SetElem {
	p := a.parent
	curr := a
	for p != nil {
		curr = p
		p = p.parent
	}
	return curr
}

// SetUnionByHeight instructs this Guf to perform unions based on the
// the height of each subtree. Shorter subtrees will become children of
// taller ones, therefore minimizing updates.
func (g *Guf) SetUnionByHeight() {
	g.unionfunc = g.unionByHeight
}

// SetUnionBySize instructs this Guf to perform unions based on the number
// of elements that each subtree holds. Smaller subtrees will become the children
// of the larger ones.
func (g *Guf) SetUnionBySize() {
	g.unionfunc = g.unionBySize
}

// maxIntSlice returns the maximum number in that slice, or
// -1 if the slice is empty.
func maxIntSlice(a []int) int {
	if len(a) == 0 {
		return -1
	}
	max := a[0]
	for i := range a[1:] {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
