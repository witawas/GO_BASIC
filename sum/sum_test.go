package sum

import (
	"testing"

	"github.com/fatih/color"
)

func TestSumEven(t *testing.T) {
	s := Sum(1, 3)
	//fmt.Println("Test sum even", s)

	if s != 4 {
		t.Error("Should return 4")
	} else {
		color.New(color.FgGreen).Println("PASS Test sum even", s)
	}
}

func TestSunOdd(t *testing.T) {
	s := Sum(3, 4)

	if s != 7 {
		t.Error("Should return 7")
	} else {
		color.New(color.FgGreen).Println("Test sum Odd", s)
	}
}

func TestSumNegative(t *testing.T) {
	s := Sum(-1, 0)
	//fmt.Println("Test sum negative", s)

	if s != -1 {
		t.Error("Should return 7")
	} else {
		color.New(color.FgGreen).Println("Test sum negative", s)
	}
}
