package pl

import "log"

type Negation struct {
	operand Proposition
}

func Not(c Proposition) *Negation {
	return &Negation{
		operand: c,
	}
}

func (n Negation) String() string {
	//TODO implement me
	panic("implement me")
}

func (n Negation) evaluate() {
	log.Println("evaluating")
}
