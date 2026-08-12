package punct

import (
	"reflect"
	"testing"
)

type testCaseArrayFunc2 struct {
	CaseName   string
	InputArray []string
	InputIndex int
	Expected   []string
}

func TestFixSpacing(t *testing.T) {
	tests := []testCaseArrayFunc2{
		{
			CaseName:   "ex1",
			InputArray: []string{"I", "was", "sitting", "over", "there", ",", "and", "then", "BAMM!!"},
			InputIndex: 5,
			Expected:   []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"},
		},
		{
			CaseName:   "ex2",
			InputArray: []string{"I", "was", "thinking", "...", "You", "were", "right"},
			InputIndex: 3,
			Expected:   []string{"I", "was", "thinking...", "You", "were", "right"},
		},
		// {
		// 	CaseName: "",
		// 	InputArray: []string{},
		// 	InputIndex: 0,
		// 	Expected: []string{},
		// },
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T) {
			actual := FixSpacing(ts.InputArray, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}

type TestCaseBoolFunc struct {
	CaseName string
	Input    string
	Expected bool
}

func TestIsPunctuations(t *testing.T) {
	tests := []TestCaseBoolFunc{
		{
			CaseName: "ex1",
			Input:    ".",
			Expected: true,
		},
		{
			CaseName: "ex2",
			Input:    "....-",
			Expected: false,
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T) {
			actual := IsPunctuations(ts.Input)
			if ts.Expected != actual {
				t.Errorf("Error on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}
