package pl

import "log"

type Disjunction struct {
	disjunctions []Proposition
}

func Or(c1 Proposition, c2 Proposition, cn ...Proposition) *Disjunction {
	return &Disjunction{
		disjunctions: append(cn, c1, c2),
	}
}

func (d Disjunction) String() string {
	//TODO implement me
	panic("implement me")
}

func (d Disjunction) evaluate() {
	log.Println("evaluating")
}
