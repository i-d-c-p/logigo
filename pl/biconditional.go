package pl

import "log"

type BiConditional struct {
	antecedent Proposition
	consequent Proposition
}

func Iff(antecedent, consequent Proposition) *BiConditional {
	return &BiConditional{
		antecedent: antecedent,
		consequent: consequent,
	}
}

func (c BiConditional) String() string {
	//TODO implement me
	panic("implement me")
}

func (c BiConditional) evaluate() {
	log.Println("evaluating")
}
