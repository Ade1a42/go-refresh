# go-refresh

A Go-based text processing and auto-correction tool that applies transformation rules to text files. This project demonstrates string manipulation, text formatting, and simple rule-based language detection.

## Overview

`go-refresh` reads an input text file, applies a series of transformations (capitalization, punctuation fixing, number conversion, quote formatting, and article correction), and writes the result to an output file. It includes optional language detection and keyword extraction features using simple heuristics.

## Features

- **Number Conversion**: Convert hexadecimal `(hex)` and binary `(bin)` numbers to decimal
- **Case Transformation**: Apply `(up)`, `(low)`, and `(cap)` tags to modify word case
- **Multi-word Transformation**: Support numeric modifiers like `(up, 3)` to affect multiple previous words
- **Punctuation Fixing**: Correct spacing around punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`)
- **Quote Handling**: Properly format single quotes around text without extra spaces
- **Article Correction**: Replace `a` with `an` before vowels and the letter `h`
- **Language Detection** (`--lang` flag): Detect whether text is primarily English or French
- **Keyword Extraction** (`--keywords` flag): Extract top 5 keywords from processed text

## Architecture
```
go-refresh/
├─ cmd/
│  └─ go-refresh/
│     └─ main.go
├─ internal/
│  ├─ numbers/
│  │  ├─ convert.go
│  │  └─ convert_test.go
│  ├─ cases/
│  │  ├─ transform.go
│  │  └─ transform_test.go
│  ├─ punct/
│  │  ├─ spacing.go
│  │  └─ spacing_test.go
│  ├─ quotes/
│  │  ├─ singlequotes.go
│  │  └─ singlequotes_test.go
│  ├─ grammar/
│  │  ├─ articles.go
│  │  └─ articles_test.go
│  ├─ ai/
│  │  ├─ langdetect.go
│  │  └─ langdetect_test.go
│  └─ pipeline/
│     ├─ pipeline.go
│     └─ pipeline_test.go
├─ samples/
│  ├─ sample1.in.txt
│  ├─ sample1.out.txt
│  ├─ sample2.in.txt
│  ├─ sample2.out.txt
│  ├─ sample3.in.txt
│  └─ sample3.out.txt
├─ go.mod
├─ README.md
└─ Makefile
```

# How to use

## Installation

Clone this repository:
```
git clone https://01.tomorrow-school.ai/git/anurzhankyz/go-refresh.git
cd go-refresh
```

Ensure you have Go 1.26.4 or later installed.

## Basic Usage

From the project root directory:

```bash
go run ./cmd/go-refresh input.txt output.txt
```

### Examples

**Text processing only:**
```bash
go run ./cmd/go-refresh samples/sample1.in.txt samples/sample1.out.txt
```

**Identify language:**
```bash
go run ./cmd/go-refresh samples/sample1.in.txt samples/sample1.out.txt --lang
```

**Extract keywords:**
```bash
go run ./cmd/go-refresh samples/sample1.in.txt samples/sample1.out.txt --keywords
```

## Transformation Rules (in order)

1. **Number Tags**: `(hex)` and `(bin)` convert the previous word to decimal
2. **Case Tags**: `(up)`, `(low)`, `(cap)` with optional counts modify word case
3. **Quote Formatting**: Single quotes attach directly to first and last enclosed words
4. **Punctuation Spacing**: Attach punctuation to previous word, space before next
5. **Article Fixing**: Replace `a` with `an` before vowels and `h`

## Language Detection (`--lang`)

Counts signal words from English and French lists, plus accented characters.

**English signals**: the, and, of, to, is, in, you, it, that
**French signals**: le, la, les, de, et, un, une, est, dans, vous
**Accents counted as French**: é, è, à, ç, ô, î, û

Output format:
```
Language: English (67%)
Language: French (75%)
Language: Unknown
```

## Keyword Extraction (`--keywords`)

Extracts top 5 most frequent words after removing stopwords and sorting alphabetically on ties.

**Stopwords**: the, and, of, to, is, in, it, that, a, an, for, on, with, as, by, at, from, this, these, those, de, la, le, les, et, un, une, des

Output format:
```
Keywords: word1, word2, word3, word4, word5
Keywords: (none)
```

# How to test

Run all tests from the project root:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./internal/pipeline
go test ./internal/grammar
go test ./internal/numbers
go test ./internal/cases
go test ./internal/punct
go test ./internal/quotes
go test ./internal/ai
```

Run tests with coverage:
```bash
go test -cover ./...
```

## Test Files

Each module includes corresponding `_test.go` files:
- `internal/pipeline/pipeline_test.go` - Tokenizer, tag processing, full pipeline tests
- `internal/grammar/articles_test.go` - Article correction tests
- `internal/numbers/convert_test.go` - Hexadecimal and binary conversion tests
- `internal/cases/transform_test.go` - Case transformation tests
- `internal/punct/spacing_test.go` - Punctuation spacing and detection tests
- `internal/quotes/singlequotes_test.go` - Quote handling tests
- `internal/ai/langdetect_test.go` - Language detection tests

## Code Quality

Ensure consistent formatting:
```bash
go fmt ./...
```

Check for errors:
```bash
go build ./cmd/go-refresh
```

## Dependencies

This project uses only Go standard library packages:
- `fmt`, `os`, `strings` - file I/O and string manipulation
- `regexp` - tokenization
- `strconv` - number conversion
- `math` - rounding for percentages
- `sort` - keyword frequency sorting
- `unicode` - character classification

No external dependencies required.

## Project Structure

- `cmd/go-refresh/` - Main entry point
- `internal/numbers/` - Hexadecimal and binary conversion logic
- `internal/cases/` - Uppercase, lowercase, capitalize transformations
- `internal/punct/` - Punctuation spacing corrections
- `internal/quotes/` - Single quote handling
- `internal/grammar/` - Article correction (a/an)
- `internal/ai/` - Language detection and keyword extraction
- `internal/pipeline/` - Text tokenization and processing orchestration
- `samples/` - Example input and output files

## License

Educational project for learning Go text processing and testing.

