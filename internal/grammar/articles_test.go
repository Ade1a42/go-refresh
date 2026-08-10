package grammar 

import (
	"testing"
	"reflect"
)

type testCase struct {
	CaseName string;
	InputArray []string;
	InputIndex int;
	Expected []string;
}

func TestFixArticle(t *testing.T){
	tests := []testCase{
		{
			CaseName: "ex1",
			InputArray: []string{"There", "it", "was", ".", "A", "amazing", "rock", "!"},
			InputIndex: 4,
			Expected: []string{"There", "it", "was", ".", "An", "amazing", "rock", "!"},
		},
		{
			CaseName: "ex2",
			InputArray: []string{"a", "house"},
			InputIndex: 0,
			Expected: []string{"an", "house"},
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := FixArticle(ts.InputArray, ts.InputIndex)
			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("Error on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}