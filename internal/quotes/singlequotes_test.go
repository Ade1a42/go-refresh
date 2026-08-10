package quotes

import (
	"testing"
	"reflect"
)

type testCaseArrayFunc2 struct {
	CaseName string
	InputArray []string
	InputIndex int 
	Expected []string
}

func TestFixQuotes(t *testing.T){
	tests := []testCaseArrayFunc2{
		{		
			CaseName: "ex1",
			InputArray: []string{"I", "am", "exactly", "how", "they", "describe", "me:", "'", "awesome", "'"},
			InputIndex: 7, 
			Expected: []string{"I", "am", "exactly", "how", "they", "describe", "me:", "'awesome'"},
		},
		{		
			CaseName: "ex2",
			InputArray: []string{"As", "Elton", "John", "said:", "'", "I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world", "'"},
			InputIndex: 4, 
			Expected: []string{"As", "Elton", "John", "said:", "'I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world'"},
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := FixQuotes(ts.InputArray, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}




