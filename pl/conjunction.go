package pl

import "log"

type Conjunction struct {
	conjunctions []Proposition
}

func And(c1 Proposition, c2 Proposition, cn ...Proposition) *Conjunction {
	return &Conjunction{
		conjunctions: append(cn, c1, c2),
	}
}

func (c Conjunction) String() string {
	//TODO implement me
	panic("implement me")
}

func (c Conjunction) evaluate() {
	log.Println("evaluating")
}
