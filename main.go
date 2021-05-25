// +build amd64
package main

import "fmt"
//import "github.com/vladbalmos/gosnake/states/initial"
//import "github.com/rthornton128/goncurses"

type P struct {
    x uint
}

type S struct {
    P
}

type S1 struct {
    P
}

func main() {
    a := &S{P{1}}
    b := &S1{P{2}}

    fmt.Println(a.P == b.P)
}
