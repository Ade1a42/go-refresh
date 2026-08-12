package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"goRefresh/internal/ai"
	"goRefresh/internal/pipeline"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("==== To process text choose syntax based on your goal :) ====")
		fmt.Println("Note: You can rename your txt file as you want :)")
		fmt.Println("Text processing: go run . ../../samples/sample1.in.txt ../../samples/sample1.out.txt")
		fmt.Println("Identify language: go run . ../../samples/sample1.in.txt ../../samples/sample1.out.txt --lang")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]

	if !(strings.HasSuffix(input, ".txt")) || !(strings.HasSuffix(output, ".txt")) {
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
	outputText := strings.Join(tokens, " ")

	if err := os.WriteFile(output, []byte(outputText), 0644); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 4 {
		if os.Args[3] != "--lang" {
			fmt.Printf("Error: unknown flag %s\n", os.Args[3])
			os.Exit(1)
		}

		fmt.Println(ai.DetectLanguage(outputText))
	}
}
