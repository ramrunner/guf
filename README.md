[![Go Report Card](https://goreportcard.com/badge/github.com/ramrunner/guf)](https://goreportcard.com/report/github.com/ramrunner/guf)
[![GoDoc](https://godoc.org/github.com/ramrunner/guf?status.svg)](https://godoc.org/github.com/ramrunner/guf)

## Golang Union Find (GUF)

guf is a library that provides connected component functionality for GO projects.
It maintains a forest of set elements and connects them throught *Union()* operations.
*Find()* operations return the parent element of the whole set.

Before any *Union()* operations each set element belongs to its own set.

It also provides two functions: *Data()* and *SetData()* for the user to possibly
associate and retrieve data to/from set elements.

## Usage

Create a new guf, register a couple of set elements, connect them and see that they
have the same parent. A default guf unions elements by height.

```go
	g := guf.NewGuf()
	e1,e2 := g.RegisterNew(), g.RegisterNew()
	g.Union(e1,e2)
	setParent := g.Find(e2)
```

Or an full example that shows how to associate external data with the sets

```go
package main

import (
        "fmt"

        "github.com/ramrunner/guf"
)

var (
        // create a helper map to keep track of the set elements
        // associated with our local data (strings in this case).
        // warning. Only comparable types can be keys in a map.
        elemMap map[string]*guf.SetElem
)

// elemFromString is a helper function that doesn't deal with errors
// in case a string is not present in the map.
func elemFromString(a string) *guf.SetElem {
        if v, ok := elemMap[a]; ok {
                return v
        }
        return nil
}

// stringFromElem is a helper that expects the saved data in a SetElem to be
// of the string type.
func stringFromElem(a *guf.SetElem) string {
	return a.Data().(string)
}

// registerMap is a helper function will register as many new SetElements with the guf
// and keep track of the association with the user provided data in the
// map.
func registerMap(g *guf.Guf, strs []string) {
        elemMap = make(map[string]*guf.SetElem)
        for i := range strs {
                e := g.RegisterNew()
                e.SetData(strs[i])
                elemMap[strs[i]] = e
        }
}
func main() {
        strElems := []string{"each", "string", "starts", "as", "its", "own", "set"}
        // create a new guf instance
        g := guf.NewGuf()
        // instruct the guf to union by size
        g.SetUnionBySize()
        // register the strings and populate the map
        registerMap(g, strElems)
        fmt.Println("union 'each' with 'string'")
        g.Union(elemFromString("each"), elemFromString("string"))
        fmt.Println("union 'starts' with 'as'")
        g.Union(elemFromString("starts"), elemFromString("as"))
        fmt.Println("union 'as' with 'its'")
        g.Union(elemFromString("as"), elemFromString("its"))
        fmt.Println("finding parent of 'string'")
        stringp := g.Find(elemFromString("string"))
        fmt.Println("finding  parent of 'its'")
        itsp := g.Find(elemFromString("its"))
        fmt.Printf("parent of %s: %s\n", stringFromElem(stringp), stringp)
        fmt.Printf("parent of %s: %s\n", stringFromElem(itsp), itsp)
        fmt.Println("union 'its' with 'string'")
        g.Union(elemFromString("its"), elemFromString("string"))
        itsp = g.Find(elemFromString("its"))
        fmt.Printf("parent of %s: %s\n", stringFromElem(itsp), itsp)
}
```
