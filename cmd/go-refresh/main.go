package main

import ( 
	"fmt"
	"os"
	"strings"
	"log"
	"goRefresh/internal/pipeline"
)

func main(){
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("==== To process text choose syntax based on your goal :) ====")
		fmt.Println("Note: You can rename your txt file as you want :)")
		fmt.Println("Text processing: go run . input.txt output.txt")
		fmt.Println("Identify language: go run . input.txt output.txt --lang")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]
	
	// only accepting txt files
	if !(strings.HasSuffix(input, ".txt")) || !(strings.HasSuffix(output, ".txt")){
		fmt.Println("Error: only .txt files accepted")
		os.Exit(1)
	}


	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	if len(content) == 0 {
		fmt.Println("Error: input file is empty")
		os.Exit(1)
	}

	tokens := pipeline.Tokenizer(string(content))
	tokens = pipeline.Process(tokens)

	fmt.Println(tokens)

	// text processing
		// tag + remove it
		// all punctuation
		// article 


	// check weather have language keyword 

	// output overwrite if need
}
