package semantic

import (
	"BabyDuckCompiler/symbols"
	"fmt"
)

// Checks if a variable has been declared before use.
func CheckVariableDeclared(name string, st *symbols.SymbolTable) error {
	if _, found := st.GetVariable(name); !found {
		return fmt.Errorf("variable '%s' is not declared", name)
	}
	return nil
}

// Ensures a variable is not redeclared in the same scope.
func CheckVariableNotRedeclared(name string, st *symbols.SymbolTable) error {
	currentScope := st.Scopes[len(st.Scopes)-1]
	if _, exists := currentScope[name]; exists {
		return fmt.Errorf("variable '%s' is already declared in this scope", name)
	}
	return nil
}

// Ensures a function is not declared more than once.
func CheckFunctionNotRedeclared(name string, st *symbols.SymbolTable) error {
	if _, exists := st.GetFunction(name); exists {
		return fmt.Errorf("function '%s' is already declared", name)
	}
	return nil
}

// Validates that the number of arguments matches the number of parameters in a function call.
func CheckFunctionCallArguments(name string, argsCount int, st *symbols.SymbolTable) error {
	fn, found := st.GetFunction(name)
	if !found {
		return fmt.Errorf("function '%s' is not declared", name)
	}
	if len(fn.Parameters) != argsCount {
		return fmt.Errorf("function '%s' expects %d argument(s), but got %d", name, len(fn.Parameters), argsCount)
	}
	return nil
}
