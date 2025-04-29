# Baby Duck
# BabyDuck Language

## Project Structure

- **`parser/`**: contains the `BabyDuck.g4` grammar file written in ANTLR, defining the lexical and syntactic rules of the language.
- **`ytest.bd`**: a test file with a sample BabyDuck program to validate the parser.
- **`main.go`**: the main program that:
  - Reads a `.bd` file
  - Uses the generated lexer and parser
  - Validates the syntax of the input

## Usage

1. Generate the parser with ANTLR:
   ```bash
   antlr -Dlanguage=Go -o parser BabyDuck.g4
## Program Output

When running `main.go` with `ytest.bd`:

- If the file is **syntactically correct**, the program completes successfully without printing errors. (No output by default unless you add a success message in `main.go`.)
- If there are **syntax errors** in the file, ANTLR automatically prints error messages indicating the line and type of error.


