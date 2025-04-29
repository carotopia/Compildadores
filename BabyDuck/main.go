package main

import (
	"fmt"
	"os"

	"babyduck/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Custom error listener para capturar errores de parsing
type syntaxErrorListener struct {
	*antlr.DefaultErrorListener
}

func (l *syntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	fmt.Printf("Syntax error at line %d:%d - %s\n", line, column, msg)
	os.Exit(1) // termina el programa si hay un error de sintaxis
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

	// Reemplazar el listener de errores por uno personalizado
	p.RemoveErrorListeners()
	p.AddErrorListener(&syntaxErrorListener{})

	tree := p.Program() // la regla inicial en tu gramática

	fmt.Println("Parsed successfully!")
	fmt.Println(tree.ToStringTree(nil, p))
}
