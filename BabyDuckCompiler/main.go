package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"

	"BabyDuckCompiler/parser"
)


type syntaxErrorListener struct {
	*antlr.DefaultErrorListener
}

func (l *syntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	fmt.Printf("Syntax error at line %d:%d - %s\n", line, column, msg)
	os.Exit(1) 
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err)
		return
	}

	lexer := parser.NewBabyDuckLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewBabyDuckParser(stream)

	
	p.RemoveErrorListeners()
	p.AddErrorListener(&syntaxErrorListener{})

	tree := p.Program() // la regla inicial en tu gramática

	fmt.Println("Parsed successfully!")
	fmt.Println(tree.ToStringTree(nil, p))
}
