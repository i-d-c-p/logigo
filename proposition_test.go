package logigo

import (
	"github.com/i-d-c-p/logigo/pl"
	"testing"
)

func TestBasicApi(t *testing.T) {
	visit := pl.Atom{Sym: "visit", Desc: "I visit my mom"}
	tuesday := pl.Atom{Sym: "tuesday", Desc: "It is tuesday"}
	monday := pl.Atom{Sym: "monday", Desc: "It is monday"}
	rain := pl.Atom{Sym: "rain", Desc: "It rains"}

	proposition := pl.If(
		pl.And(
			pl.Or(tuesday, monday),
			pl.Not(rain),
		),
		visit,
	)

	if proposition == nil {
		t.Error("proposition is null")
	}
}

func TestPropositionString(t *testing.T) {
	visit := pl.Atom{Sym: "visit", Desc: "I visit my mom"}
	tuesday := pl.Atom{Sym: "tuesday", Desc: "It is tuesday"}
	monday := pl.Atom{Sym: "monday", Desc: "It is monday"}
	rain := pl.Atom{Sym: "rain", Desc: "It rains"}

	proposition := pl.If(
		pl.And(
			pl.Or(tuesday, monday),
			pl.Not(rain),
		),
		visit,
	)

	expected := "[(tuesday ∨ monday) ∧ ~rain] => visit"
	res := proposition.String()
	if res != expected {
		t.Errorf("expected: \"%s\", actual: \"%s\"", expected, res)
	}
}
