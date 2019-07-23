[![Go Report Card](https://goreportcard.com/badge/github.com/ramrunner/guf)](https://goreportcard.com/report/github.com/ramrunner/guf)
[![GoDoc](https://godoc.org/github.com/ramrunner/guf?status.svg)](https://godoc.org/github.com/ramrunner/guf)

## Golang Union Find (GUF)

guf is a library that provides connected component functionality for GO projects.
It maintains a forest of set elements and connects them throught Union() operations.
Find() operations return the parent element of the whole set.
Before any Union() operations each set element belongs to its own set.

## Usage

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
        //register the strings and populate the map
        registerMap(g, strElems)
        // union "each" with "string"
        g.Union(elemFromString("each"), elemFromString("string"))
        // union "starts" with "as"
        g.Union(elemFromString("starts"), elemFromString("as"))
        // union "as" with "its"
        g.Union(elemFromString("as"), elemFromString("its"))
        // find parent of "string"
        stringp := g.Find(elemFromString("string"))
        // find  parent of "its"
        itsp := g.Find(elemFromString("its"))
        fmt.Printf("parent of \"string\": %+v\n", stringp)
        fmt.Printf("parent of \"its\": %+v\n", itsp)
        // union "its" with "string"
        g.Union(elemFromString("its"), elemFromString("string"))
        itsp = g.Find(elemFromString("its"))
        fmt.Printf("parent of \"its\": %+v\n", itsp)
}
```
