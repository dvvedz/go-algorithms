package module01

import (
	"testing"
)

type numInListTest struct {
	argSlice []int
	argInt   int
	expected bool
}

func TestNumInList(t *testing.T) {
	var numInListTests = []numInListTest{
		{[]int{1, 3, 4, 5, 33, 3}, 3, true},
		{[]int{3, 3, 3, 3, 3, 3}, 3, true},
		{[]int{2, 1, 2, 3, 3, 3}, 9, false},
		{[]int{33, 33, 33, 33, 33, 33}, 3, false},
		{[]int{-3, 33, 33, 33, 33, 33}, -3, true},
		{[]int{-3, 33, 33, 33, 33, 33}, 3, false},
		{[]int{}, 1, false},
		{nil, 1, false},
	}
	for i, te := range numInListTests {
		if output := NumInList(te.argSlice, te.argInt); output != te.expected {
			t.Errorf("output %t not equal to expected %t, data %v on line %d in dataset", output, te.expected, te, i+1)
		}
	}
}

type sumListTest struct {
	sumSlice []int
	expected int
}

func TestSum(t *testing.T) {
	var sumListTests = []sumListTest{
		{[]int{1, 1, 5, 7}, 14},
		{[]int{5, 1, 5, 10}, 21},
		{[]int{1, 1, -5, 10}, 7},
		{[]int{-1, -1, -5, -1}, -8},
		{[]int{}, 0},
		{nil, 0},
	}

	for i, test := range sumListTests {
		if output := Sum(test.sumSlice); output != test.expected {
			t.Errorf("output %d is not equal to expected %d | test data: %v on line %d in dataset", output, test.expected, test, i+1)
		}
	}
}

type revStringTest struct {
	str      string
	expected string
}

func TestRevString(t *testing.T) {
	var revStringTests = []revStringTest{
		{"Test", "tseT"},
		{"String", "gnirtS"},
		{"Test String", "gnirtS tseT"},
	}

	for i, test := range revStringTests {
		if output := RevString(test.str); output != test.expected {
			t.Errorf("output %s is not equal to expected %s | test data %v on line %d in dataset", output, test.expected, test, i-1)
		}
	}
}

type fizzBuzzTest struct {
	number   int
	expected string
}

func TestFizzBuzz(t *testing.T) {
	var FizBuzzTests = []fizzBuzzTest{
		{3, "1 2 Fizz"},
		{5, "1 2 Fizz 4 Buzz"},
		{15, "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"},
	}

	for i, test := range FizBuzzTests {
		if output := FizzBuzz(test.number); output != test.expected {
			t.Errorf("output %s is not equal to expected %s | test data %v on line %d in dataset", output, test.expected, test, i-1)
		}
	}
}

type decToBaseTest struct {
	dec      int
	base     int
	expected string
}

func TestDecToBase(t *testing.T) {
	var DecToBaseTests = []decToBaseTest{
		{1, 2, "1"},
		{2, 2, "10"},
		{7, 3, "21"},
		{14, 2, "1110"},
		{14, 16, "E"},
		{17, 16, "11"},
		{3735928559, 2, "11011110101011011011111011101111"},
		{3735928559, 3, "100122100210211112102"},
		{3735928559, 4, "3132223123323233"},
		{3735928559, 5, "30122344203214"},
		{3735928559, 6, "1414413525315"},
		{3735928559, 7, "161402603666"},
		{3735928559, 8, "33653337357"},
		{3735928559, 9, "10570724472"},
		{3735928559, 10, "3735928559"},
		{3735928559, 11, "164791A470"},
		{3735928559, 12, "8831A383B"},
		{3735928559, 13, "476CC321C"},
		{3735928559, 14, "276253DDD"},
		{3735928559, 15, "16CEB1BDE"},
		{3735928559, 16, "DEADBEEF"},
	}
	for i, test := range DecToBaseTests {
		if output := DecToBase(test.dec, test.base); output != test.expected {
			t.Errorf("output %s is not equal to expected %s | test data %v on line %d in dataset", output, test.expected, test, i-1)
		}
	}
}
