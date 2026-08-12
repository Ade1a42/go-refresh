package cases

import (
	"reflect"
	"testing"
)

type testCaseArrayFunc struct {
	CaseName   string
	InputArray []string
	InputNum   int
	InputIndex int
	Expected   []string
}

func TestUpN(t *testing.T) {
	tests := []testCaseArrayFunc{
		{
			CaseName:   "ex1",
			InputArray: []string{"This", "is", "so", "exciting", "(up, 2)", "to", "see", "it"},
			InputNum:   2,
			InputIndex: 4,
			Expected:   []string{"This", "is", "SO", "EXCITING", "(up, 2)", "to", "see", "it"},
		},
		{
			CaseName:   "ex2",
			InputArray: []string{"This", "is", "so", "exciting", "(up, 1)", "to", "see", "it"},
			InputNum:   1,
			InputIndex: 4,
			Expected:   []string{"This", "is", "so", "EXCITING", "(up, 1)", "to", "see", "it"},
		},
		{
			CaseName:   "ex3",
			InputArray: []string{"Ready", ",", "set", ",", "go", ",", "(up)", "!"},
			InputNum:   1,
			InputIndex: 6,
			Expected:   []string{"Ready", ",", "set", ",", "GO", ",", "(up)", "!"},
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T) {
			actual := UpN(ts.InputArray, ts.InputNum, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}

func TestLowN(t *testing.T) {
	tests := []testCaseArrayFunc{
		{
			CaseName:   "ex1",
			InputArray: []string{"I", "should", "stop", "SHOUTING", "(low)"},
			InputNum:   1,
			InputIndex: 4,
			Expected:   []string{"I", "should", "stop", "shouting", "(low)"},
		},
		{
			CaseName:   "ex2",
			InputArray: []string{"I", "SHOULD", "stop", "SHOUTING", "(low, 3)"},
			InputNum:   3,
			InputIndex: 4,
			Expected:   []string{"I", "should", "stop", "shouting", "(low, 3)"},
		},
		{
			CaseName:   "ex3",
			InputArray: []string{"I", "should", "stop", "SHOUTING", "(low , 1)"},
			InputNum:   1,
			InputIndex: 4,
			Expected:   []string{"I", "should", "stop", "shouting", "(low , 1)"},
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T) {
			actual := LowN(ts.InputArray, ts.InputNum, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}

func TestCapN(t *testing.T) {
	tests := []testCaseArrayFunc{
		{
			CaseName:   "ex1",
			InputArray: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"},
			InputNum:   1,
			InputIndex: 5,
			Expected:   []string{"Welcome", "to", "the", "Brooklyn", "Bridge", "(cap)"},
		},
		{
			CaseName:   "ex2",
			InputArray: []string{"Welcome", "TO", "the", "BroOklyn", "bridge", "(cap, 4)"},
			InputNum:   4,
			InputIndex: 5,
			Expected:   []string{"Welcome", "To", "The", "Brooklyn", "Bridge", "(cap, 4)"},
		},
		{
			CaseName:   "ex3",
			InputArray: []string{"Welcome", "to", "the", "our", "Brooklyn", "bridge", "(cap , 1)"},
			InputNum:   1,
			InputIndex: 6,
			Expected:   []string{"Welcome", "to", "the", "our", "Brooklyn", "Bridge", "(cap , 1)"},
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T) {
			actual := CapN(ts.InputArray, ts.InputNum, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}
