package pipeline

import (
	"testing"
	"reflect"
)

/* LOGIC OF TESTING TOKENIZER
Creating multiple testing cases on array of struct.
Struct (
	name of test case. 
	// eg: in tokenizer we are testing of that's how it will tokenize sentence containing super-puper
	input for processing
	expected output
)

Then through loop  that array and get input field and then process it then check wether it is as we expected. give error if not 
*/

type TestCaseTokenizer struct {
	CaseName string
	Input string
	Expected []string
}

func TestTokenizer(t *testing.T){
	// first I'm making template to fill it
	tests := []TestCaseTokenizer {
		{
			CaseName: "Ex1",
			Input: "1E (hex) files were added",
			Expected: []string{"1E", "(hex)", "files", "were", "added"},
		},
		{
			CaseName: "Ex2",
			Input: "It has been 10 (bin) years",
			Expected: []string{"It", "has", "been", "10", "(bin)", "years"},
		},
		{
			CaseName: "Ex3",
			Input: "Ready, set, go (up) !",
			Expected: []string{"Ready", ",", "set", ",", "go", "(up)", "!"},
		},
		{
			CaseName: "Ex4",
			Input: "I should stop SHOUTING (low)",
			Expected: []string{"I", "should", "stop", "SHOUTING", "(low)"},
		},
		{
			CaseName: "Ex5",
			Input: "Welcome to the Brooklyn bridge (cap)",
			Expected: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"},
		},
		{
			CaseName: "Ex6",
			Input: "This is so exciting (up, 2)",
			Expected: []string{"This", "is", "so", "exciting", "(up, 2)"},
		},
		{
			CaseName: "Ex7",
			Input: "I was sitting over there ,and then BAMM !!",
			Expected: []string{"I", "was", "sitting", "over", "there", ",", "and", "then", "BAMM", "!!"},
		},
		{
			CaseName: "Ex8",
			Input: "I was thinking ... You were right",
			Expected: []string{"I", "was", "thinking", "...", "You", "were", "right"},
		},
		{
			CaseName: "Ex9",
			Input: "I am exactly how they describe me: ' awesome '",
			Expected: []string{"I", "am", "exactly", "how", "they", "describe", "me", ":", "'", "awesome", "'"},
		},
		{
			CaseName: "Ex10",
			Input: "As Elton John said: ' I am the most well-known homosexual in the world '",
			Expected: []string{"As", "Elton", "John", "said", ":", "'", "I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world", "'"},
		},
		{
			CaseName: "Ex11",
			Input: "There it was. A amazing rock!",
			Expected: []string{"There", "it", "was", ".", "A", "amazing", "rock", "!"},
		},
	}
	

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := Tokenizer(ts.Input)
			if !(reflect.DeepEqual(ts.Expected, actual)){
				t.Errorf("Error on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}

type TestCaseBoolFunc struct {
	CaseName string
	Input string 
	Expected bool
}

func TestIsTag(t *testing.T){
	tests := []TestCaseBoolFunc{
		{
			CaseName: "hex1",
			Input: "(hex)",
			Expected: true,
		},
		{
			CaseName: "hex2",
			Input: "(hex, 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)",
			Expected: false,
		},
		{
			CaseName: "hex3",
			Input: "(hex, -789)",
			Expected: false,
		},
		{
			CaseName: "hex4",
			Input: "(hex, 789)",
			Expected: false,
		},	
		{
			CaseName: "hex5",
			Input: "(hex, @#$%)",
			Expected: false,
		},
		{
			CaseName: "hex6",
			Input: "(hex, @#$%)",
			Expected: false,
		},
		{
			CaseName: "hex7",
			Input: "( hex , 5 )",
			Expected: false,
		},
		{
			CaseName: "hex8",
			Input: "( hex )",
			Expected: true,
		},		
		{
			CaseName: "hex9",
			Input: "( hex , 0 )",
			Expected: false,
		},	
		{
			CaseName: "bin1",
			Input: "(bin)",
			Expected: true,
		},
		{
			CaseName: "bin2",
			Input: "(bin, 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)",
			Expected: false,
		},
		{
			CaseName: "bin3",
			Input: "(bin, -789)",
			Expected: false,
		},
		{
			CaseName: "bin4",
			Input: "(bin, 789)",
			Expected: false,
		},	
		{
			CaseName: "bin5",
			Input: "(bin, @#$%)",
			Expected: false,
		},
		{
			CaseName: "bin6",
			Input: "(bin, @#$%)",
			Expected: false,
		},
		{
			CaseName: "bin7",
			Input: "( bin , 5 )",
			Expected: false,
		},
		{
			CaseName: "bin8",
			Input: "( bin )",
			Expected: true,
		},		
		{
			CaseName: "bin9",
			Input: "( bin , 0 )",
			Expected: false,
		},	
		{
			CaseName: "up1",
			Input: "(up)",
			Expected: true,
		},
		{
			CaseName: "up2",
			Input: "(up, 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)",
			Expected: false,
		},
		{
			CaseName: "up3",
			Input: "(up, -789)",
			Expected: false,
		},
		{
			CaseName: "up4",
			Input: "(up, 789)",
			Expected: true,
		},	
		{
			CaseName: "up5",
			Input: "(up, @#$%)",
			Expected: false,
		},
		{
			CaseName: "up6",
			Input: "(up, @#$%)",
			Expected: false,
		},
		{
			CaseName: "up7",
			Input: "( up , 5 )",
			Expected: true,
		},
		{
			CaseName: "up8",
			Input: "( up )",
			Expected: true,
		},
		{
			CaseName: "up9",
			Input: "( up , 0 )",
			Expected: false,
		},
		{
			CaseName: "low1",
			Input: "(low)",
			Expected: true,
		},
		{
			CaseName: "low2",
			Input: "(low, 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)",
			Expected: false,
		},
		{
			CaseName: "low3",
			Input: "(low, -789)",
			Expected: false,
		},
		{
			CaseName: "low4",
			Input: "(low, 789)",
			Expected: true,
		},	
		{
			CaseName: "low5",
			Input: "(low, @#$%)",
			Expected: false,
		},
		{
			CaseName: "low6",
			Input: "(low, @#$%)",
			Expected: false,
		},
		{
			CaseName: "low7",
			Input: "( low , 5 )",
			Expected: true,
		},
		{
			CaseName: "low8",
			Input: "( low )",
			Expected: true,
		},
		{
			CaseName: "low9",
			Input: "( low , 0 )",
			Expected: false,
		},
		{
			CaseName: "cap1",
			Input: "(cap)",
			Expected: true,
		},
		{
			CaseName: "cap2",
			Input: "(cap, 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)",
			Expected: false,
		},
		{
			CaseName: "cap3",
			Input: "(cap, -789)",
			Expected: false,
		},
		{
			CaseName: "cap4",
			Input: "(cap, 789)",
			Expected: true,
		},	
		{
			CaseName: "cap5",
			Input: "(cap, @#$%)",
			Expected: false,
		},
		{
			CaseName: "cap6",
			Input: "(cap, @#$%)",
			Expected: false,
		},
		{
			CaseName: "cap7",
			Input: "( cap , 5 )",
			Expected: true,
		},
		{
			CaseName: "cap8",
			Input: "( cap )",
			Expected: true,
		},
		{
			CaseName: "cap9",
			Input: "( cap , 0 )",
			Expected: false,
		},
		{
			CaseName: "ex1",
			Input: "Ahex,5D",
			Expected: false,
		},
		{
			CaseName: "ex2",
			Input: "A low,5 A",
			Expected: false,
		},
		{
			CaseName: "ex3",
			Input: "( cap , 5.0 )",
			Expected: false,
		},
	}

	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := IsTag(ts.Input)
			if ts.Expected != actual {
				t.Errorf("Error on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	}
}

type TestCaseProcess struct {
	CaseName string
	InputArray []string
	Expected []string
}

func TestProcess(t *testing.T){
	tests := []TestCaseProcess{
		{
			CaseName: "ex1",
			InputArray: []string{"Ready", ",", "set", ",", "go", "(up)", "!"},
			Expected: []string{"Ready,", "set,", "GO!"},
		},
		{
			CaseName: "ex2",
			InputArray: []string{"Ready", ",", "set", ",", "gO", "(up, 2)", "!"},
			Expected: []string{"Ready,", "SET,", "GO!"},
		},
		{
			CaseName: "ex3",
			InputArray: []string{"I", "should", "stop", "SHOUTING", "(low)"},
			Expected: []string{"I", "should", "stop", "shouting"},
		},
		{
			CaseName: "ex4",
			InputArray: []string{"I", "SHOULD", "stop", "SHOUTING", "(low, 3)"},
			Expected: []string{"I", "should", "stop", "shouting"},
		},
		{
			CaseName: "ex5",
			InputArray: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"},
			Expected: []string{"Welcome", "to", "the", "Brooklyn", "Bridge"},
		},
		{
			CaseName: "ex6",
			InputArray: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap, 3)"},
			Expected: []string{"Welcome", "to", "The", "Brooklyn", "Bridge"},
		},
		{
			CaseName: "ex7",
			InputArray: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap, 0)"},
			Expected: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap, 0)"},
		},
		{
			CaseName: "ex8",
			InputArray: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap, -9)"},
			Expected: []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap, -9)"},
		},
		{
			CaseName: "ex9",
			InputArray: []string{"Simply", "add", "42", "(hex)", "and", "10", "(bin)", "and", "you", "will", "see", "the", "result", "is", "68", "."},
			Expected: []string{"Simply", "add", "66", "and", "2", "and", "you", "will", "see", "the", "result", "is", "68."},
		},
		{
			CaseName: "ex10",
			InputArray: []string{"There", "it", "was", ".", "A", "amazing", "rock", "!", "May", "be", "I", "forgot", "to", "add", "another", "a", "apple", ",", "an", "orange", "and", "others", "."},
			Expected: []string{"There", "it", "was.", "An", "amazing", "rock!", "May", "be", "I", "forgot", "to", "add", "another", "an", "apple,", "an", "orange", "and", "others."},
		},
		{
			CaseName: "ex11",
			InputArray: []string{"I", "was", "sitting", "over", "there", ",", "and", "then", "BAMM", "!!"},
			Expected: []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"},
		},
		{
			CaseName: "ex12",
			InputArray: []string{"I", "was", "thinking", "...", "You", "were", "right"},
			Expected: []string{"I", "was", "thinking...", "You", "were", "right"},
		},
		{
			CaseName: "ex13",
			InputArray: []string{"I", "am", "exactly", "how", "they", "describe", "me", ":", "'", "awesome", "'"},
			Expected: []string{"I", "am", "exactly", "how", "they", "describe", "me:", "'awesome'"},
		},
		{
			CaseName: "ex14",
			InputArray: []string{"As", "Elton", "John", "said", ":", "'", "I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world", "'"},
			Expected: []string{"As", "Elton", "John", "said:", "'I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world'"},
		},
		// {
		// 	CaseName: "ex",
		// 	InputArray: []string{},
		// 	Expected: []string{},
		// },
	}


	for _, ts := range tests {
		t.Run(ts.CaseName, func(t *testing.T){
			actual := Process(ts.InputArray)
			if !(reflect.DeepEqual(ts.Expected, actual)){
				t.Errorf("\nError on testing %v:\nExpected: %v;\nGot: %v\n", ts.CaseName, ts.Expected, actual)
			}
		})
	} 
}
