package pl

import (
	"fmt"
	"log"
)

type Proposition interface {
	fmt.Stringer

	evaluate()
}

type Atom struct {
	Sym  string
	Desc string
}

func (a Atom) String() string {
	return a.Sym
}

func (a Atom) evaluate() {
	log.Println("evaluating")
}
