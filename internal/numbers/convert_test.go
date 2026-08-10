package numbers

import (
	"testing"
	"reflect"
)

type testCase struct {
	CaseName string
	InputArray []string
	InputIndex int 
	Expected []string
}

func TestHex(t *testing.T) {
	tests := []testCase{
		{
            CaseName: "ex1",
            InputArray: []string{"Convert", "1e", "(hex)"},
            InputIndex: 2,
            Expected: []string{"Convert", "30", "(hex)"},
        },
        {
            CaseName: "ex2",
            InputArray: []string{"Value", "FF", "(hex)"},
            InputIndex: 2,
            Expected: []string{"Value", "255", "(hex)"},
        },
        {
            CaseName: "ex3",
            InputArray: []string{"Result", "100", "(hex)"},
            InputIndex: 2,
            Expected: []string{"Result", "256", "(hex)"},
        },
        {
            CaseName: "ex4",
            InputArray: []string{"Error", "1G", "(hex)"},
            InputIndex: 2,
            Expected: []string{"Error", "1G", "(hex)"}, 
        },
		{
            CaseName: "ex5",
            InputArray: []string{"Error", "10", "-", "-", "(hex)"},
            InputIndex: 2,
            Expected: []string{"Error", "16", "-", "-", "(hex)"}, 
        },
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := Hex(ts.InputArray, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}


func TestBin(t *testing.T) {
	tests := []testCase{
		{
            CaseName: "ex1",
            InputArray: []string{"and", "10", "(bin)", "and", "you", "will", "see", "the", "result"},
            InputIndex: 2,
            Expected: []string{"and", "2", "(bin)", "and", "you", "will", "see", "the", "result"},
        },
		{
            CaseName: "ex2",
            InputArray: []string{"and", "10000", ".", ".", "(bin)", "and", "you", "will", "see", "the", "result"},
            InputIndex: 4,
            Expected: []string{"and", "16", ".", ".", "(bin)", "and", "you", "will", "see", "the", "result"},
        },
		{
            CaseName: "ex3",
            InputArray: []string{"and", "102", ".", ".", "(bin)", "and", "you", "will", "see", "the", "result"},
            InputIndex: 4,
            Expected: []string{"and", "102", ".", ".", "(bin)", "and", "you", "will", "see", "the", "result"},
        },
		// {
        //     CaseName: "",
        //     InputArray: []string{},
        //     InputIndex: 0,
        //     Expected: []string{},
        // },
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := Bin(ts.InputArray, ts.InputIndex)

			if !(reflect.DeepEqual(actual, ts.Expected)) {
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}
