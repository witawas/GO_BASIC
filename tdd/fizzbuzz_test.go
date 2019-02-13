package tdd

import "testing"

func TestSayInput1AP1(t *testing.T) {
	s := Say(1)

	if s != "1" {
		t.Error("TestSayInput1AP1 Should be 1 but got", s)
	}
}

func TestSayInput2(t *testing.T) {
	s := Say(2)
	if s != "2" {
		t.Error("TestSayInput2 Should be 2 but got", s)
	}
}
