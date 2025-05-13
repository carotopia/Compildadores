package semantic

import (
	"BabyDuckCompiler/parser"
	"BabyDuckCompiler/symbols"
	"fmt"

	_ "github.com/antlr4-go/antlr/v4"
)

// BabyDuckListener implementa la interfaz del listener generado por ANTLR
type BabyDuckListener struct {
	parser.BaseBabyDuckListener
	SymbolTable *symbols.SymbolTable
	Errors      []string
}

// NewBabyDuckListener crea una nueva instancia del listener
func NewBabyDuckListener() *BabyDuckListener {
	return &BabyDuckListener{
		SymbolTable: symbols.NewSymbolTable(),
		Errors:      make([]string, 0),
	}
}

// AddError agrega un error semántico a la lista de errores
func (l *BabyDuckListener) AddError(msg string) {
	l.Errors = append(l.Errors, msg)
}

// EnterProgram se llama cuando el parser entra en la regla 'program'
func (l *BabyDuckListener) EnterProgram(ctx *parser.ProgramContext) {
	// Inicializar el compilador
	fmt.Println("Compilando programa BabyDuck:", ctx.ID().GetText())
}

// ExitProgram se llama cuando el parser sale de la regla 'program'
func (l *BabyDuckListener) ExitProgram(ctx *parser.ProgramContext) {
	// Verificar si hubo errores
	if len(l.Errors) > 0 {
		fmt.Println("Compilación fallida. Se encontraron errores:")
		for _, err := range l.Errors {
			fmt.Println("- " + err)
		}
	} else {
		fmt.Println("Compilación exitosa!")
	}
}

// EnterVars se llama cuando el parser entra en la sección de variables
func (l *BabyDuckListener) EnterVars(ctx *parser.VarsContext) {
	// No necesitamos crear un nuevo scope aquí, ya que las variables globales
	// ya están en el scope global creado por NewSymbolTable()
}

// ExitVar_decl se llama cuando el parser sale de una declaración de variables
func (l *BabyDuckListener) ExitVar_decl(ctx *parser.Var_declContext) {
	varType := ctx.Type_().GetText()

	// Procesar cada ID en la declaración
	for _, idNode := range ctx.AllID() {
		varName := idNode.GetText()

		// Verificar que la variable no esté ya declarada en el scope actual
		err := CheckVariableNotRedeclared(varName, l.SymbolTable)
		if err != nil {
			l.AddError(err.Error())
			continue
		}

		// Agregar la variable a la tabla de símbolos
		variable := symbols.Variable{
			Name: varName,
			Type: varType,
		}
		err = l.SymbolTable.AddVariable(variable)
		if err != nil {
			l.AddError(err.Error())
		}
	}
}

// EnterFunc se llama cuando el parser entra en la definición de una función
func (l *BabyDuckListener) EnterFunc(ctx *parser.FuncContext) {
	funcName := ctx.ID(0).GetText()

	// Verificar que la función no esté ya declarada
	err := CheckFunctionNotRedeclared(funcName, l.SymbolTable)
	if err != nil {
		l.AddError(err.Error())
		return
	}

	// Preparar los parámetros
	parameters := make([]symbols.Variable, 0)

	// Si hay parámetros
	if ctx.Param_list() != nil {
		for i := 0; i < len(ctx.Param_list().AllParam()); i++ {
			paramCtx := ctx.Param_list().Param(i)
			paramName := paramCtx.ID().GetText()
			paramType := paramCtx.Type_().GetText()
			parameters = append(parameters, symbols.Variable{
				Name: paramName,
				Type: paramType,
			})
		}
	}

	// Registrar la función
	function := symbols.Function{
		Name:       funcName,
		Parameters: parameters,
	}
	err = l.SymbolTable.AddFunction(function)
	if err != nil {
		l.AddError(err.Error())
		return
	}

	// Entrar en un nuevo scope para los parámetros y variables locales
	l.SymbolTable.EnterScope()

	// Agregar los parámetros al scope actual
	for _, param := range parameters {
		err := l.SymbolTable.AddVariable(param)
		if err != nil {
			l.AddError(err.Error())
		}
	}
}

// ExitFunc se llama cuando el parser sale de la definición de una función
func (l *BabyDuckListener) ExitFunc(ctx *parser.FuncContext) {
	// Salir del scope de la función
	err := l.SymbolTable.ExitScope()
	if err != nil {
		l.AddError(err.Error())
	}
}

// EnterF_call se llama cuando el parser entra en una llamada a función
func (l *BabyDuckListener) EnterF_call(ctx *parser.F_callContext) {
	funcName := ctx.ID().GetText()

	// Contar el número de argumentos
	argCount := 0
	if ctx.Arg_list() != nil {
		argCount = len(ctx.Arg_list().AllExpression())
	}

	// Verificar que la función esté declarada y que tenga el número correcto de argumentos
	err := CheckFunctionCallArguments(funcName, argCount, l.SymbolTable)
	if err != nil {
		l.AddError(err.Error())
	}
}

// EnterAssign se llama cuando el parser entra en una asignación
func (l *BabyDuckListener) EnterAssign(ctx *parser.AssignContext) {
	varName := ctx.ID().GetText()

	// Verificar que la variable esté declarada
	err := CheckVariableDeclared(varName, l.SymbolTable)
	if err != nil {
		l.AddError(err.Error())
	}
}

// EnterPrint_stmt se llama cuando el parser entra en una sentencia print
func (l *BabyDuckListener) EnterPrint_stmt(ctx *parser.Print_stmtContext) {
	// Aquí podrías verificar cada expresión en el print
	// Pero esto depende de cómo hayas definido las reglas en tu gramática
}

// EnterFactor se llama cuando el parser entra en un factor
func (l *BabyDuckListener) EnterFactor(ctx *parser.FactorContext) {
	// Si el factor es un ID, verificar que esté declarado
	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		err := CheckVariableDeclared(varName, l.SymbolTable)
		if err != nil {
			l.AddError(err.Error())
		}
	}
}

// EnterBody se llama cuando el parser entra en el bloque body
func (l *BabyDuckListener) EnterBody(ctx *parser.BodyContext) {
	// Si este es el body principal (main), crear un nuevo scope
	// Esto depende de tu gramática, puede que necesites ajustarlo
	if ctx.GetParent().GetRuleIndex() == parser.BabyDuckParserRULE_program {
		l.SymbolTable.EnterScope()
	}
}

// ExitBody se llama cuando el parser sale del bloque body
func (l *BabyDuckListener) ExitBody(ctx *parser.BodyContext) {
	// Si este es el body principal (main), salir del scope
	// Esto depende de tu gramática, puede que necesites ajustarlo
	if ctx.GetParent().GetRuleIndex() == parser.BabyDuckParserRULE_program {
		err := l.SymbolTable.ExitScope()
		if err != nil {
			l.AddError(err.Error())
		}
	}
}
