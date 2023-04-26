package pl

import "log"

type Conditional struct {
	antecedent Proposition
	consequent Proposition
}

func If(antecedent, consequent Proposition) *Conditional {
	return &Conditional{
		antecedent: antecedent,
		consequent: consequent,
	}
}

func (c Conditional) String() string {
	//TODO implement me
	panic("implement me")
}

func (c Conditional) evaluate() {
	log.Println("evaluating")
}
