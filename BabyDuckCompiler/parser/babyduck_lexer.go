// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BabyDuckLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BabyDuckLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func babyducklexerLexerInit() {
	staticData := &BabyDuckLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'program'", "'main'", "'end'", "'var'", "'{'", "'}'", "'='", "'while'",
		"'do'", "'if'", "'else'", "'print'", "'>'", "'<'", "'!='", "'+'", "'-'",
		"'*'", "'/'", "'void'", "'int'", "'float'", "'('", "')'", "'['", "']'",
		"':'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "VOID", "INTTYPE", "FLOATTYPE", "LPAREN", "RPAREN", "LBRACKET",
		"RBRACKET", "COLON", "COMMA", "SEMICOLON", "ID", "INT", "FLOAT", "STRING",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "VOID", "INTTYPE", "FLOATTYPE", "LPAREN", "RPAREN",
		"LBRACKET", "RBRACKET", "COLON", "COMMA", "SEMICOLON", "ID", "INT",
		"FLOAT", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 202, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 5, 29, 166, 8, 29, 10, 29, 12, 29, 169, 9, 29, 1, 30,
		4, 30, 172, 8, 30, 11, 30, 12, 30, 173, 1, 31, 4, 31, 177, 8, 31, 11, 31,
		12, 31, 178, 1, 31, 1, 31, 4, 31, 183, 8, 31, 11, 31, 12, 31, 184, 1, 32,
		1, 32, 5, 32, 189, 8, 32, 10, 32, 12, 32, 192, 9, 32, 1, 32, 1, 32, 1,
		33, 4, 33, 197, 8, 33, 11, 33, 12, 33, 198, 1, 33, 1, 33, 0, 0, 34, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 1, 0, 5, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 34,
		34, 3, 0, 9, 10, 13, 13, 32, 32, 207, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 82, 1,
		0, 0, 0, 7, 86, 1, 0, 0, 0, 9, 90, 1, 0, 0, 0, 11, 92, 1, 0, 0, 0, 13,
		94, 1, 0, 0, 0, 15, 96, 1, 0, 0, 0, 17, 102, 1, 0, 0, 0, 19, 105, 1, 0,
		0, 0, 21, 108, 1, 0, 0, 0, 23, 113, 1, 0, 0, 0, 25, 119, 1, 0, 0, 0, 27,
		121, 1, 0, 0, 0, 29, 123, 1, 0, 0, 0, 31, 126, 1, 0, 0, 0, 33, 128, 1,
		0, 0, 0, 35, 130, 1, 0, 0, 0, 37, 132, 1, 0, 0, 0, 39, 134, 1, 0, 0, 0,
		41, 139, 1, 0, 0, 0, 43, 143, 1, 0, 0, 0, 45, 149, 1, 0, 0, 0, 47, 151,
		1, 0, 0, 0, 49, 153, 1, 0, 0, 0, 51, 155, 1, 0, 0, 0, 53, 157, 1, 0, 0,
		0, 55, 159, 1, 0, 0, 0, 57, 161, 1, 0, 0, 0, 59, 163, 1, 0, 0, 0, 61, 171,
		1, 0, 0, 0, 63, 176, 1, 0, 0, 0, 65, 186, 1, 0, 0, 0, 67, 196, 1, 0, 0,
		0, 69, 70, 5, 112, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5, 111, 0, 0, 72,
		73, 5, 103, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75, 5, 97, 0, 0, 75, 76, 5,
		109, 0, 0, 76, 2, 1, 0, 0, 0, 77, 78, 5, 109, 0, 0, 78, 79, 5, 97, 0, 0,
		79, 80, 5, 105, 0, 0, 80, 81, 5, 110, 0, 0, 81, 4, 1, 0, 0, 0, 82, 83,
		5, 101, 0, 0, 83, 84, 5, 110, 0, 0, 84, 85, 5, 100, 0, 0, 85, 6, 1, 0,
		0, 0, 86, 87, 5, 118, 0, 0, 87, 88, 5, 97, 0, 0, 88, 89, 5, 114, 0, 0,
		89, 8, 1, 0, 0, 0, 90, 91, 5, 123, 0, 0, 91, 10, 1, 0, 0, 0, 92, 93, 5,
		125, 0, 0, 93, 12, 1, 0, 0, 0, 94, 95, 5, 61, 0, 0, 95, 14, 1, 0, 0, 0,
		96, 97, 5, 119, 0, 0, 97, 98, 5, 104, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100,
		5, 108, 0, 0, 100, 101, 5, 101, 0, 0, 101, 16, 1, 0, 0, 0, 102, 103, 5,
		100, 0, 0, 103, 104, 5, 111, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106, 5, 105,
		0, 0, 106, 107, 5, 102, 0, 0, 107, 20, 1, 0, 0, 0, 108, 109, 5, 101, 0,
		0, 109, 110, 5, 108, 0, 0, 110, 111, 5, 115, 0, 0, 111, 112, 5, 101, 0,
		0, 112, 22, 1, 0, 0, 0, 113, 114, 5, 112, 0, 0, 114, 115, 5, 114, 0, 0,
		115, 116, 5, 105, 0, 0, 116, 117, 5, 110, 0, 0, 117, 118, 5, 116, 0, 0,
		118, 24, 1, 0, 0, 0, 119, 120, 5, 62, 0, 0, 120, 26, 1, 0, 0, 0, 121, 122,
		5, 60, 0, 0, 122, 28, 1, 0, 0, 0, 123, 124, 5, 33, 0, 0, 124, 125, 5, 61,
		0, 0, 125, 30, 1, 0, 0, 0, 126, 127, 5, 43, 0, 0, 127, 32, 1, 0, 0, 0,
		128, 129, 5, 45, 0, 0, 129, 34, 1, 0, 0, 0, 130, 131, 5, 42, 0, 0, 131,
		36, 1, 0, 0, 0, 132, 133, 5, 47, 0, 0, 133, 38, 1, 0, 0, 0, 134, 135, 5,
		118, 0, 0, 135, 136, 5, 111, 0, 0, 136, 137, 5, 105, 0, 0, 137, 138, 5,
		100, 0, 0, 138, 40, 1, 0, 0, 0, 139, 140, 5, 105, 0, 0, 140, 141, 5, 110,
		0, 0, 141, 142, 5, 116, 0, 0, 142, 42, 1, 0, 0, 0, 143, 144, 5, 102, 0,
		0, 144, 145, 5, 108, 0, 0, 145, 146, 5, 111, 0, 0, 146, 147, 5, 97, 0,
		0, 147, 148, 5, 116, 0, 0, 148, 44, 1, 0, 0, 0, 149, 150, 5, 40, 0, 0,
		150, 46, 1, 0, 0, 0, 151, 152, 5, 41, 0, 0, 152, 48, 1, 0, 0, 0, 153, 154,
		5, 91, 0, 0, 154, 50, 1, 0, 0, 0, 155, 156, 5, 93, 0, 0, 156, 52, 1, 0,
		0, 0, 157, 158, 5, 58, 0, 0, 158, 54, 1, 0, 0, 0, 159, 160, 5, 44, 0, 0,
		160, 56, 1, 0, 0, 0, 161, 162, 5, 59, 0, 0, 162, 58, 1, 0, 0, 0, 163, 167,
		7, 0, 0, 0, 164, 166, 7, 1, 0, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1, 0,
		0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 60, 1, 0, 0, 0,
		169, 167, 1, 0, 0, 0, 170, 172, 7, 2, 0, 0, 171, 170, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 62, 1,
		0, 0, 0, 175, 177, 7, 2, 0, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0,
		0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		182, 5, 46, 0, 0, 181, 183, 7, 2, 0, 0, 182, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 64, 1, 0,
		0, 0, 186, 190, 5, 34, 0, 0, 187, 189, 8, 3, 0, 0, 188, 187, 1, 0, 0, 0,
		189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191,
		193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 5, 34, 0, 0, 194, 66,
		1, 0, 0, 0, 195, 197, 7, 4, 0, 0, 196, 195, 1, 0, 0, 0, 197, 198, 1, 0,
		0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 201, 6, 33, 0, 0, 201, 68, 1, 0, 0, 0, 7, 0, 167, 173, 178, 184, 190,
		198, 1, 6, 0, 0,
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

// BabyDuckLexerInit initializes any static state used to implement BabyDuckLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBabyDuckLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BabyDuckLexerInit() {
	staticData := &BabyDuckLexerLexerStaticData
	staticData.once.Do(babyducklexerLexerInit)
}

// NewBabyDuckLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBabyDuckLexer(input antlr.CharStream) *BabyDuckLexer {
	BabyDuckLexerInit()
	l := new(BabyDuckLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BabyDuckLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "BabyDuck.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BabyDuckLexer tokens.
const (
	BabyDuckLexerT__0      = 1
	BabyDuckLexerT__1      = 2
	BabyDuckLexerT__2      = 3
	BabyDuckLexerT__3      = 4
	BabyDuckLexerT__4      = 5
	BabyDuckLexerT__5      = 6
	BabyDuckLexerT__6      = 7
	BabyDuckLexerT__7      = 8
	BabyDuckLexerT__8      = 9
	BabyDuckLexerT__9      = 10
	BabyDuckLexerT__10     = 11
	BabyDuckLexerT__11     = 12
	BabyDuckLexerT__12     = 13
	BabyDuckLexerT__13     = 14
	BabyDuckLexerT__14     = 15
	BabyDuckLexerT__15     = 16
	BabyDuckLexerT__16     = 17
	BabyDuckLexerT__17     = 18
	BabyDuckLexerT__18     = 19
	BabyDuckLexerVOID      = 20
	BabyDuckLexerINTTYPE   = 21
	BabyDuckLexerFLOATTYPE = 22
	BabyDuckLexerLPAREN    = 23
	BabyDuckLexerRPAREN    = 24
	BabyDuckLexerLBRACKET  = 25
	BabyDuckLexerRBRACKET  = 26
	BabyDuckLexerCOLON     = 27
	BabyDuckLexerCOMMA     = 28
	BabyDuckLexerSEMICOLON = 29
	BabyDuckLexerID        = 30
	BabyDuckLexerINT       = 31
	BabyDuckLexerFLOAT     = 32
	BabyDuckLexerSTRING    = 33
	BabyDuckLexerWS        = 34
)
