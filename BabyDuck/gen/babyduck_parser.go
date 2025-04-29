// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // BabyDuck

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BabyDuckParser struct {
	*antlr.BaseParser
}

var BabyDuckParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func babyduckParserInit() {
	staticData := &BabyDuckParserStaticData
	staticData.LiteralNames = []string{
		"", "'program'", "';'", "'main'", "'end'", "'var'", "':'", "','", "'int'",
		"'float'", "'{'", "'}'", "'='", "'while'", "'('", "')'", "'do'", "'if'",
		"'else'", "'print'", "'>'", "'<'", "'!='", "'+'", "'-'", "'*'", "'/'",
		"'void'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "ID", "INT", "FLOAT", "STRING",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "vars", "var_list", "var_decl", "id_list", "type", "body",
		"statement", "assign", "cycle", "condition", "else_part", "print_stmt",
		"printexpr", "constant", "expression", "relop", "exp", "addop", "term",
		"mulop", "factor", "factorsign", "funcs", "func", "functype", "param_list",
		"param", "funcbody", "f_call", "arg_list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 269, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 77, 8, 1, 1, 2, 1, 2, 5, 2, 81, 8, 2, 10, 2, 12, 2, 84, 9, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 93, 8, 4, 10, 4, 12, 4,
		96, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 102, 8, 6, 10, 6, 12, 6, 105, 9,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 114, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 137, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 144, 8, 12, 10, 12, 12, 12, 147,
		9, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 154, 8, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 162, 8, 15, 1, 16, 3, 16, 165, 8,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 173, 8, 17, 10, 17,
		12, 17, 176, 9, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 184,
		8, 19, 10, 19, 12, 19, 187, 9, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 202, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 209, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 215, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 221, 8,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 228, 8, 25, 1, 26, 1, 26,
		1, 26, 5, 26, 233, 8, 26, 10, 26, 12, 26, 236, 9, 26, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 3, 28, 243, 8, 28, 1, 28, 5, 28, 246, 8, 28, 10, 28,
		12, 28, 249, 9, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29, 256, 8, 29,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 5, 30, 264, 8, 30, 10, 30, 12,
		30, 267, 9, 30, 1, 30, 0, 0, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 0, 5, 1, 0, 8, 9, 1, 0, 29, 30, 1, 0, 20, 21, 1, 0, 23, 24, 1,
		0, 25, 26, 265, 0, 62, 1, 0, 0, 0, 2, 76, 1, 0, 0, 0, 4, 78, 1, 0, 0, 0,
		6, 85, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 97, 1, 0, 0, 0, 12, 99, 1, 0,
		0, 0, 14, 113, 1, 0, 0, 0, 16, 115, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20,
		127, 1, 0, 0, 0, 22, 136, 1, 0, 0, 0, 24, 138, 1, 0, 0, 0, 26, 153, 1,
		0, 0, 0, 28, 155, 1, 0, 0, 0, 30, 157, 1, 0, 0, 0, 32, 164, 1, 0, 0, 0,
		34, 168, 1, 0, 0, 0, 36, 177, 1, 0, 0, 0, 38, 179, 1, 0, 0, 0, 40, 188,
		1, 0, 0, 0, 42, 201, 1, 0, 0, 0, 44, 208, 1, 0, 0, 0, 46, 214, 1, 0, 0,
		0, 48, 216, 1, 0, 0, 0, 50, 227, 1, 0, 0, 0, 52, 229, 1, 0, 0, 0, 54, 237,
		1, 0, 0, 0, 56, 240, 1, 0, 0, 0, 58, 252, 1, 0, 0, 0, 60, 260, 1, 0, 0,
		0, 62, 63, 5, 1, 0, 0, 63, 64, 5, 28, 0, 0, 64, 65, 5, 2, 0, 0, 65, 66,
		3, 2, 1, 0, 66, 67, 3, 46, 23, 0, 67, 68, 5, 3, 0, 0, 68, 69, 3, 12, 6,
		0, 69, 70, 5, 4, 0, 0, 70, 71, 5, 2, 0, 0, 71, 1, 1, 0, 0, 0, 72, 73, 5,
		5, 0, 0, 73, 74, 3, 4, 2, 0, 74, 75, 5, 2, 0, 0, 75, 77, 1, 0, 0, 0, 76,
		72, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 3, 1, 0, 0, 0, 78, 82, 3, 6, 3,
		0, 79, 81, 3, 6, 3, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 5, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0,
		85, 86, 3, 8, 4, 0, 86, 87, 5, 6, 0, 0, 87, 88, 3, 10, 5, 0, 88, 7, 1,
		0, 0, 0, 89, 94, 5, 28, 0, 0, 90, 91, 5, 7, 0, 0, 91, 93, 5, 28, 0, 0,
		92, 90, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1,
		0, 0, 0, 95, 9, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 7, 0, 0, 0, 98,
		11, 1, 0, 0, 0, 99, 103, 5, 10, 0, 0, 100, 102, 3, 14, 7, 0, 101, 100,
		1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 11, 0, 0,
		107, 13, 1, 0, 0, 0, 108, 114, 3, 16, 8, 0, 109, 114, 3, 18, 9, 0, 110,
		114, 3, 58, 29, 0, 111, 114, 3, 24, 12, 0, 112, 114, 3, 20, 10, 0, 113,
		108, 1, 0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 15, 1, 0, 0, 0, 115, 116, 5, 28,
		0, 0, 116, 117, 5, 12, 0, 0, 117, 118, 3, 34, 17, 0, 118, 119, 5, 2, 0,
		0, 119, 17, 1, 0, 0, 0, 120, 121, 5, 13, 0, 0, 121, 122, 5, 14, 0, 0, 122,
		123, 3, 34, 17, 0, 123, 124, 5, 15, 0, 0, 124, 125, 5, 16, 0, 0, 125, 126,
		3, 12, 6, 0, 126, 19, 1, 0, 0, 0, 127, 128, 5, 17, 0, 0, 128, 129, 5, 14,
		0, 0, 129, 130, 3, 30, 15, 0, 130, 131, 5, 15, 0, 0, 131, 132, 3, 12, 6,
		0, 132, 133, 3, 22, 11, 0, 133, 21, 1, 0, 0, 0, 134, 135, 5, 18, 0, 0,
		135, 137, 3, 12, 6, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		23, 1, 0, 0, 0, 138, 139, 5, 19, 0, 0, 139, 140, 5, 14, 0, 0, 140, 145,
		3, 26, 13, 0, 141, 142, 5, 7, 0, 0, 142, 144, 3, 26, 13, 0, 143, 141, 1,
		0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0,
		0, 146, 148, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 15, 0, 0, 149,
		150, 5, 2, 0, 0, 150, 25, 1, 0, 0, 0, 151, 154, 3, 34, 17, 0, 152, 154,
		5, 31, 0, 0, 153, 151, 1, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 27, 1, 0,
		0, 0, 155, 156, 7, 1, 0, 0, 156, 29, 1, 0, 0, 0, 157, 161, 3, 34, 17, 0,
		158, 159, 3, 32, 16, 0, 159, 160, 3, 34, 17, 0, 160, 162, 1, 0, 0, 0, 161,
		158, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 31, 1, 0, 0, 0, 163, 165, 7,
		2, 0, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0,
		0, 166, 167, 5, 22, 0, 0, 167, 33, 1, 0, 0, 0, 168, 174, 3, 38, 19, 0,
		169, 170, 3, 36, 18, 0, 170, 171, 3, 38, 19, 0, 171, 173, 1, 0, 0, 0, 172,
		169, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 35, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 178, 7, 3,
		0, 0, 178, 37, 1, 0, 0, 0, 179, 185, 3, 42, 21, 0, 180, 181, 3, 40, 20,
		0, 181, 182, 3, 42, 21, 0, 182, 184, 1, 0, 0, 0, 183, 180, 1, 0, 0, 0,
		184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		39, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 7, 4, 0, 0, 189, 41, 1,
		0, 0, 0, 190, 191, 5, 14, 0, 0, 191, 192, 3, 34, 17, 0, 192, 193, 5, 15,
		0, 0, 193, 194, 3, 44, 22, 0, 194, 202, 1, 0, 0, 0, 195, 196, 5, 28, 0,
		0, 196, 202, 3, 44, 22, 0, 197, 198, 5, 29, 0, 0, 198, 202, 3, 44, 22,
		0, 199, 200, 5, 30, 0, 0, 200, 202, 3, 44, 22, 0, 201, 190, 1, 0, 0, 0,
		201, 195, 1, 0, 0, 0, 201, 197, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202,
		43, 1, 0, 0, 0, 203, 204, 5, 25, 0, 0, 204, 209, 3, 42, 21, 0, 205, 206,
		5, 26, 0, 0, 206, 209, 3, 42, 21, 0, 207, 209, 1, 0, 0, 0, 208, 203, 1,
		0, 0, 0, 208, 205, 1, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 45, 1, 0, 0,
		0, 210, 211, 3, 48, 24, 0, 211, 212, 3, 46, 23, 0, 212, 215, 1, 0, 0, 0,
		213, 215, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215,
		47, 1, 0, 0, 0, 216, 217, 3, 50, 25, 0, 217, 218, 5, 28, 0, 0, 218, 220,
		5, 14, 0, 0, 219, 221, 3, 52, 26, 0, 220, 219, 1, 0, 0, 0, 220, 221, 1,
		0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 15, 0, 0, 223, 224, 3, 56,
		28, 0, 224, 49, 1, 0, 0, 0, 225, 228, 5, 27, 0, 0, 226, 228, 3, 10, 5,
		0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 51, 1, 0, 0, 0, 229,
		234, 3, 54, 27, 0, 230, 231, 5, 7, 0, 0, 231, 233, 3, 54, 27, 0, 232, 230,
		1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0,
		0, 0, 235, 53, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 238, 3, 10, 5, 0,
		238, 239, 5, 28, 0, 0, 239, 55, 1, 0, 0, 0, 240, 242, 5, 10, 0, 0, 241,
		243, 3, 2, 1, 0, 242, 241, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 247,
		1, 0, 0, 0, 244, 246, 3, 14, 7, 0, 245, 244, 1, 0, 0, 0, 246, 249, 1, 0,
		0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0,
		249, 247, 1, 0, 0, 0, 250, 251, 5, 11, 0, 0, 251, 57, 1, 0, 0, 0, 252,
		253, 5, 28, 0, 0, 253, 255, 5, 14, 0, 0, 254, 256, 3, 60, 30, 0, 255, 254,
		1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 5, 15,
		0, 0, 258, 259, 5, 2, 0, 0, 259, 59, 1, 0, 0, 0, 260, 265, 3, 34, 17, 0,
		261, 262, 5, 7, 0, 0, 262, 264, 3, 34, 17, 0, 263, 261, 1, 0, 0, 0, 264,
		267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 61, 1,
		0, 0, 0, 267, 265, 1, 0, 0, 0, 22, 76, 82, 94, 103, 113, 136, 145, 153,
		161, 164, 174, 185, 201, 208, 214, 220, 227, 234, 242, 247, 255, 265,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BabyDuckParserInit initializes any static state used to implement BabyDuckParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBabyDuckParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BabyDuckParserInit() {
	staticData := &BabyDuckParserStaticData
	staticData.once.Do(babyduckParserInit)
}

// NewBabyDuckParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBabyDuckParser(input antlr.TokenStream) *BabyDuckParser {
	BabyDuckParserInit()
	this := new(BabyDuckParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BabyDuckParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BabyDuck.g4"

	return this
}

// BabyDuckParser tokens.
const (
	BabyDuckParserEOF    = antlr.TokenEOF
	BabyDuckParserT__0   = 1
	BabyDuckParserT__1   = 2
	BabyDuckParserT__2   = 3
	BabyDuckParserT__3   = 4
	BabyDuckParserT__4   = 5
	BabyDuckParserT__5   = 6
	BabyDuckParserT__6   = 7
	BabyDuckParserT__7   = 8
	BabyDuckParserT__8   = 9
	BabyDuckParserT__9   = 10
	BabyDuckParserT__10  = 11
	BabyDuckParserT__11  = 12
	BabyDuckParserT__12  = 13
	BabyDuckParserT__13  = 14
	BabyDuckParserT__14  = 15
	BabyDuckParserT__15  = 16
	BabyDuckParserT__16  = 17
	BabyDuckParserT__17  = 18
	BabyDuckParserT__18  = 19
	BabyDuckParserT__19  = 20
	BabyDuckParserT__20  = 21
	BabyDuckParserT__21  = 22
	BabyDuckParserT__22  = 23
	BabyDuckParserT__23  = 24
	BabyDuckParserT__24  = 25
	BabyDuckParserT__25  = 26
	BabyDuckParserT__26  = 27
	BabyDuckParserID     = 28
	BabyDuckParserINT    = 29
	BabyDuckParserFLOAT  = 30
	BabyDuckParserSTRING = 31
	BabyDuckParserWS     = 32
)

// BabyDuckParser rules.
const (
	BabyDuckParserRULE_program    = 0
	BabyDuckParserRULE_vars       = 1
	BabyDuckParserRULE_var_list   = 2
	BabyDuckParserRULE_var_decl   = 3
	BabyDuckParserRULE_id_list    = 4
	BabyDuckParserRULE_type       = 5
	BabyDuckParserRULE_body       = 6
	BabyDuckParserRULE_statement  = 7
	BabyDuckParserRULE_assign     = 8
	BabyDuckParserRULE_cycle      = 9
	BabyDuckParserRULE_condition  = 10
	BabyDuckParserRULE_else_part  = 11
	BabyDuckParserRULE_print_stmt = 12
	BabyDuckParserRULE_printexpr  = 13
	BabyDuckParserRULE_constant   = 14
	BabyDuckParserRULE_expression = 15
	BabyDuckParserRULE_relop      = 16
	BabyDuckParserRULE_exp        = 17
	BabyDuckParserRULE_addop      = 18
	BabyDuckParserRULE_term       = 19
	BabyDuckParserRULE_mulop      = 20
	BabyDuckParserRULE_factor     = 21
	BabyDuckParserRULE_factorsign = 22
	BabyDuckParserRULE_funcs      = 23
	BabyDuckParserRULE_func       = 24
	BabyDuckParserRULE_functype   = 25
	BabyDuckParserRULE_param_list = 26
	BabyDuckParserRULE_param      = 27
	BabyDuckParserRULE_funcbody   = 28
	BabyDuckParserRULE_f_call     = 29
	BabyDuckParserRULE_arg_list   = 30
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Vars() IVarsContext
	Funcs() IFuncsContext
	Body() IBodyContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *ProgramContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *ProgramContext) Funcs() IFuncsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsContext)
}

func (s *ProgramContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *BabyDuckParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BabyDuckParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(BabyDuckParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(BabyDuckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Vars()
	}
	{
		p.SetState(66)
		p.Funcs()
	}
	{
		p.SetState(67)
		p.Match(BabyDuckParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Body()
	}
	{
		p.SetState(69)
		p.Match(BabyDuckParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(BabyDuckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_list() IVar_listContext

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_vars
	return p
}

func InitEmptyVarsContext(p *VarsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_vars
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) Var_list() IVar_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_listContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *BabyDuckParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BabyDuckParserRULE_vars)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__4 {
		{
			p.SetState(72)
			p.Match(BabyDuckParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Var_list()
		}
		{
			p.SetState(74)
			p.Match(BabyDuckParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_listContext is an interface to support dynamic dispatch.
type IVar_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext

	// IsVar_listContext differentiates from other interfaces.
	IsVar_listContext()
}

type Var_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_listContext() *Var_listContext {
	var p = new(Var_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_var_list
	return p
}

func InitEmptyVar_listContext(p *Var_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_var_list
}

func (*Var_listContext) IsVar_listContext() {}

func NewVar_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_listContext {
	var p = new(Var_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_var_list

	return p
}

func (s *Var_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_listContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *Var_listContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Var_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVar_list(s)
	}
}

func (s *Var_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVar_list(s)
	}
}

func (p *BabyDuckParser) Var_list() (localctx IVar_listContext) {
	localctx = NewVar_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BabyDuckParserRULE_var_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Var_decl()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserID {
		{
			p.SetState(79)
			p.Var_decl()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_list() IId_listContext
	Type_() ITypeContext

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_var_decl
	return p
}

func InitEmptyVar_declContext(p *Var_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_var_decl
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) Id_list() IId_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_listContext)
}

func (s *Var_declContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVar_decl(s)
	}
}

func (s *Var_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVar_decl(s)
	}
}

func (p *BabyDuckParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BabyDuckParserRULE_var_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Id_list()
	}
	{
		p.SetState(86)
		p.Match(BabyDuckParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IId_listContext is an interface to support dynamic dispatch.
type IId_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsId_listContext differentiates from other interfaces.
	IsId_listContext()
}

type Id_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_listContext() *Id_listContext {
	var p = new(Id_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_id_list
	return p
}

func InitEmptyId_listContext(p *Id_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_id_list
}

func (*Id_listContext) IsId_listContext() {}

func NewId_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_listContext {
	var p = new(Id_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_id_list

	return p
}

func (s *Id_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_listContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserID)
}

func (s *Id_listContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, i)
}

func (s *Id_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterId_list(s)
	}
}

func (s *Id_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitId_list(s)
	}
}

func (p *BabyDuckParser) Id_list() (localctx IId_listContext) {
	localctx = NewId_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BabyDuckParserRULE_id_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__6 {
		{
			p.SetState(90)
			p.Match(BabyDuckParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(BabyDuckParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *BabyDuckParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BabyDuckParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserT__7 || _la == BabyDuckParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitBody(s)
	}
}

func (p *BabyDuckParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BabyDuckParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(BabyDuckParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269099008) != 0 {
		{
			p.SetState(100)
			p.Statement()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
		p.Match(BabyDuckParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() IAssignContext
	Cycle() ICycleContext
	F_call() IF_callContext
	Print_stmt() IPrint_stmtContext
	Condition() IConditionContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatementContext) Cycle() ICycleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICycleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICycleContext)
}

func (s *StatementContext) F_call() IF_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IF_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IF_callContext)
}

func (s *StatementContext) Print_stmt() IPrint_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrint_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrint_stmtContext)
}

func (s *StatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *BabyDuckParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BabyDuckParserRULE_statement)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Cycle()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.F_call()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Print_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Condition()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Exp() IExpContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *AssignContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *BabyDuckParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BabyDuckParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(BabyDuckParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Exp()
	}
	{
		p.SetState(118)
		p.Match(BabyDuckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICycleContext is an interface to support dynamic dispatch.
type ICycleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	Body() IBodyContext

	// IsCycleContext differentiates from other interfaces.
	IsCycleContext()
}

type CycleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCycleContext() *CycleContext {
	var p = new(CycleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_cycle
	return p
}

func InitEmptyCycleContext(p *CycleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_cycle
}

func (*CycleContext) IsCycleContext() {}

func NewCycleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CycleContext {
	var p = new(CycleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_cycle

	return p
}

func (s *CycleContext) GetParser() antlr.Parser { return s.parser }

func (s *CycleContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *CycleContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *CycleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CycleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CycleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterCycle(s)
	}
}

func (s *CycleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitCycle(s)
	}
}

func (p *BabyDuckParser) Cycle() (localctx ICycleContext) {
	localctx = NewCycleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BabyDuckParserRULE_cycle)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(BabyDuckParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(BabyDuckParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Exp()
	}
	{
		p.SetState(123)
		p.Match(BabyDuckParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(BabyDuckParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Body()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Body() IBodyContext
	Else_part() IElse_partContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ConditionContext) Else_part() IElse_partContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_partContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_partContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *BabyDuckParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BabyDuckParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(BabyDuckParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(BabyDuckParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Expression()
	}
	{
		p.SetState(130)
		p.Match(BabyDuckParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Body()
	}
	{
		p.SetState(132)
		p.Else_part()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_partContext is an interface to support dynamic dispatch.
type IElse_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Body() IBodyContext

	// IsElse_partContext differentiates from other interfaces.
	IsElse_partContext()
}

type Else_partContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_partContext() *Else_partContext {
	var p = new(Else_partContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_else_part
	return p
}

func InitEmptyElse_partContext(p *Else_partContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_else_part
}

func (*Else_partContext) IsElse_partContext() {}

func NewElse_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_partContext {
	var p = new(Else_partContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_else_part

	return p
}

func (s *Else_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_partContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *Else_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterElse_part(s)
	}
}

func (s *Else_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitElse_part(s)
	}
}

func (p *BabyDuckParser) Else_part() (localctx IElse_partContext) {
	localctx = NewElse_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BabyDuckParserRULE_else_part)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__17 {
		{
			p.SetState(134)
			p.Match(BabyDuckParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Body()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrint_stmtContext is an interface to support dynamic dispatch.
type IPrint_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrintexpr() []IPrintexprContext
	Printexpr(i int) IPrintexprContext

	// IsPrint_stmtContext differentiates from other interfaces.
	IsPrint_stmtContext()
}

type Print_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_stmtContext() *Print_stmtContext {
	var p = new(Print_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_print_stmt
	return p
}

func InitEmptyPrint_stmtContext(p *Print_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_print_stmt
}

func (*Print_stmtContext) IsPrint_stmtContext() {}

func NewPrint_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_stmtContext {
	var p = new(Print_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_print_stmt

	return p
}

func (s *Print_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_stmtContext) AllPrintexpr() []IPrintexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintexprContext); ok {
			len++
		}
	}

	tst := make([]IPrintexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintexprContext); ok {
			tst[i] = t.(IPrintexprContext)
			i++
		}
	}

	return tst
}

func (s *Print_stmtContext) Printexpr(i int) IPrintexprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintexprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintexprContext)
}

func (s *Print_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterPrint_stmt(s)
	}
}

func (s *Print_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitPrint_stmt(s)
	}
}

func (p *BabyDuckParser) Print_stmt() (localctx IPrint_stmtContext) {
	localctx = NewPrint_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BabyDuckParserRULE_print_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(BabyDuckParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(BabyDuckParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Printexpr()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__6 {
		{
			p.SetState(141)
			p.Match(BabyDuckParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Printexpr()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
		p.Match(BabyDuckParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(BabyDuckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintexprContext is an interface to support dynamic dispatch.
type IPrintexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	STRING() antlr.TerminalNode

	// IsPrintexprContext differentiates from other interfaces.
	IsPrintexprContext()
}

type PrintexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintexprContext() *PrintexprContext {
	var p = new(PrintexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printexpr
	return p
}

func InitEmptyPrintexprContext(p *PrintexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printexpr
}

func (*PrintexprContext) IsPrintexprContext() {}

func NewPrintexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintexprContext {
	var p = new(PrintexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_printexpr

	return p
}

func (s *PrintexprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintexprContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *PrintexprContext) STRING() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSTRING, 0)
}

func (s *PrintexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterPrintexpr(s)
	}
}

func (s *PrintexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitPrintexpr(s)
	}
}

func (p *BabyDuckParser) Printexpr() (localctx IPrintexprContext) {
	localctx = NewPrintexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BabyDuckParserRULE_printexpr)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserT__13, BabyDuckParserID, BabyDuckParserINT, BabyDuckParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Exp()
		}

	case BabyDuckParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(BabyDuckParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserINT, 0)
}

func (s *ConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserFLOAT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *BabyDuckParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BabyDuckParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserINT || _la == BabyDuckParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext
	Relop() IRelopContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpressionContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BabyDuckParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BabyDuckParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Exp()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0 {
		{
			p.SetState(158)
			p.Relop()
		}
		{
			p.SetState(159)
			p.Exp()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relop
	return p
}

func InitEmptyRelopContext(p *RelopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relop
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }
func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *BabyDuckParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BabyDuckParserRULE_relop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__19 || _la == BabyDuckParserT__20 {
		{
			p.SetState(163)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BabyDuckParserT__19 || _la == BabyDuckParserT__20) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(166)
		p.Match(BabyDuckParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllAddop() []IAddopContext
	Addop(i int) IAddopContext

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpContext) AllAddop() []IAddopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddopContext); ok {
			len++
		}
	}

	tst := make([]IAddopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddopContext); ok {
			tst[i] = t.(IAddopContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Addop(i int) IAddopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddopContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddopContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *BabyDuckParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BabyDuckParserRULE_exp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Term()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__22 || _la == BabyDuckParserT__23 {
		{
			p.SetState(169)
			p.Addop()
		}
		{
			p.SetState(170)
			p.Term()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddopContext is an interface to support dynamic dispatch.
type IAddopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAddopContext differentiates from other interfaces.
	IsAddopContext()
}

type AddopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddopContext() *AddopContext {
	var p = new(AddopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_addop
	return p
}

func InitEmptyAddopContext(p *AddopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_addop
}

func (*AddopContext) IsAddopContext() {}

func NewAddopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddopContext {
	var p = new(AddopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_addop

	return p
}

func (s *AddopContext) GetParser() antlr.Parser { return s.parser }
func (s *AddopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterAddop(s)
	}
}

func (s *AddopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitAddop(s)
	}
}

func (p *BabyDuckParser) Addop() (localctx IAddopContext) {
	localctx = NewAddopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BabyDuckParserRULE_addop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserT__22 || _la == BabyDuckParserT__23) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllMulop() []IMulopContext
	Mulop(i int) IMulopContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMulop() []IMulopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulopContext); ok {
			len++
		}
	}

	tst := make([]IMulopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulopContext); ok {
			tst[i] = t.(IMulopContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Mulop(i int) IMulopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulopContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulopContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BabyDuckParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BabyDuckParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Factor()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__24 || _la == BabyDuckParserT__25 {
		{
			p.SetState(180)
			p.Mulop()
		}
		{
			p.SetState(181)
			p.Factor()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMulopContext is an interface to support dynamic dispatch.
type IMulopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMulopContext differentiates from other interfaces.
	IsMulopContext()
}

type MulopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulopContext() *MulopContext {
	var p = new(MulopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_mulop
	return p
}

func InitEmptyMulopContext(p *MulopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_mulop
}

func (*MulopContext) IsMulopContext() {}

func NewMulopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulopContext {
	var p = new(MulopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_mulop

	return p
}

func (s *MulopContext) GetParser() antlr.Parser { return s.parser }
func (s *MulopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterMulop(s)
	}
}

func (s *MulopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitMulop(s)
	}
}

func (p *BabyDuckParser) Mulop() (localctx IMulopContext) {
	localctx = NewMulopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BabyDuckParserRULE_mulop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserT__24 || _la == BabyDuckParserT__25) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	Factorsign() IFactorsignContext
	ID() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FactorContext) Factorsign() IFactorsignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorsignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorsignContext)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *FactorContext) INT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserINT, 0)
}

func (s *FactorContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserFLOAT, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *BabyDuckParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BabyDuckParserRULE_factor)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(BabyDuckParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Exp()
		}
		{
			p.SetState(192)
			p.Match(BabyDuckParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Factorsign()
		}

	case BabyDuckParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(BabyDuckParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Factorsign()
		}

	case BabyDuckParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Match(BabyDuckParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Factorsign()
		}

	case BabyDuckParserFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)
			p.Match(BabyDuckParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Factorsign()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorsignContext is an interface to support dynamic dispatch.
type IFactorsignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext

	// IsFactorsignContext differentiates from other interfaces.
	IsFactorsignContext()
}

type FactorsignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorsignContext() *FactorsignContext {
	var p = new(FactorsignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factorsign
	return p
}

func InitEmptyFactorsignContext(p *FactorsignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factorsign
}

func (*FactorsignContext) IsFactorsignContext() {}

func NewFactorsignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorsignContext {
	var p = new(FactorsignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_factorsign

	return p
}

func (s *FactorsignContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorsignContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorsignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorsignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorsignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFactorsign(s)
	}
}

func (s *FactorsignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFactorsign(s)
	}
}

func (p *BabyDuckParser) Factorsign() (localctx IFactorsignContext) {
	localctx = NewFactorsignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BabyDuckParserRULE_factorsign)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(BabyDuckParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Factor()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Match(BabyDuckParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Factor()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncsContext is an interface to support dynamic dispatch.
type IFuncsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Func_() IFuncContext
	Funcs() IFuncsContext

	// IsFuncsContext differentiates from other interfaces.
	IsFuncsContext()
}

type FuncsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsContext() *FuncsContext {
	var p = new(FuncsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcs
	return p
}

func InitEmptyFuncsContext(p *FuncsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcs
}

func (*FuncsContext) IsFuncsContext() {}

func NewFuncsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsContext {
	var p = new(FuncsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_funcs

	return p
}

func (s *FuncsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FuncsContext) Funcs() IFuncsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsContext)
}

func (s *FuncsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFuncs(s)
	}
}

func (s *FuncsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFuncs(s)
	}
}

func (p *BabyDuckParser) Funcs() (localctx IFuncsContext) {
	localctx = NewFuncsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BabyDuckParserRULE_funcs)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserT__7, BabyDuckParserT__8, BabyDuckParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Func_()
		}
		{
			p.SetState(211)
			p.Funcs()
		}

	case BabyDuckParserT__2:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Functype() IFunctypeContext
	ID() antlr.TerminalNode
	Funcbody() IFuncbodyContext
	Param_list() IParam_listContext

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) Functype() IFunctypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctypeContext)
}

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *FuncContext) Funcbody() IFuncbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FuncContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (p *BabyDuckParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BabyDuckParserRULE_func)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Functype()
	}
	{
		p.SetState(217)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Match(BabyDuckParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__7 || _la == BabyDuckParserT__8 {
		{
			p.SetState(219)
			p.Param_list()
		}

	}
	{
		p.SetState(222)
		p.Match(BabyDuckParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Funcbody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctypeContext is an interface to support dynamic dispatch.
type IFunctypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsFunctypeContext differentiates from other interfaces.
	IsFunctypeContext()
}

type FunctypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctypeContext() *FunctypeContext {
	var p = new(FunctypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functype
	return p
}

func InitEmptyFunctypeContext(p *FunctypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functype
}

func (*FunctypeContext) IsFunctypeContext() {}

func NewFunctypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctypeContext {
	var p = new(FunctypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_functype

	return p
}

func (s *FunctypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunctype(s)
	}
}

func (s *FunctypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunctype(s)
	}
}

func (p *BabyDuckParser) Functype() (localctx IFunctypeContext) {
	localctx = NewFunctypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BabyDuckParserRULE_functype)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.Match(BabyDuckParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BabyDuckParserT__7, BabyDuckParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Type_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param_list
	return p
}

func InitEmptyParam_listContext(p *Param_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param_list
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *Param_listContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParam_list(s)
	}
}

func (s *Param_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParam_list(s)
	}
}

func (p *BabyDuckParser) Param_list() (localctx IParam_listContext) {
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BabyDuckParserRULE_param_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Param()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__6 {
		{
			p.SetState(230)
			p.Match(BabyDuckParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Param()
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *BabyDuckParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BabyDuckParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Type_()
	}
	{
		p.SetState(238)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Vars() IVarsContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcbody
	return p
}

func InitEmptyFuncbodyContext(p *FuncbodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcbody
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *FuncbodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FuncbodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFuncbody(s)
	}
}

func (p *BabyDuckParser) Funcbody() (localctx IFuncbodyContext) {
	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BabyDuckParserRULE_funcbody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(BabyDuckParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(241)
			p.Vars()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269099008) != 0 {
		{
			p.SetState(244)
			p.Statement()
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(250)
		p.Match(BabyDuckParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IF_callContext is an interface to support dynamic dispatch.
type IF_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Arg_list() IArg_listContext

	// IsF_callContext differentiates from other interfaces.
	IsF_callContext()
}

type F_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_callContext() *F_callContext {
	var p = new(F_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_f_call
	return p
}

func InitEmptyF_callContext(p *F_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_f_call
}

func (*F_callContext) IsF_callContext() {}

func NewF_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_callContext {
	var p = new(F_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_f_call

	return p
}

func (s *F_callContext) GetParser() antlr.Parser { return s.parser }

func (s *F_callContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *F_callContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *F_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterF_call(s)
	}
}

func (s *F_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitF_call(s)
	}
}

func (p *BabyDuckParser) F_call() (localctx IF_callContext) {
	localctx = NewF_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BabyDuckParserRULE_f_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Match(BabyDuckParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1879064576) != 0 {
		{
			p.SetState(254)
			p.Arg_list()
		}

	}
	{
		p.SetState(257)
		p.Match(BabyDuckParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(BabyDuckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext

	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *Arg_listContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterArg_list(s)
	}
}

func (s *Arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitArg_list(s)
	}
}

func (p *BabyDuckParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BabyDuckParserRULE_arg_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Exp()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__6 {
		{
			p.SetState(261)
			p.Match(BabyDuckParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Exp()
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
