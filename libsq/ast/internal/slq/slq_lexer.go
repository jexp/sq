// Code generated from SLQ.g4 by ANTLR 4.13.0. DO NOT EDIT.

package slq

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

type SLQLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SLQLexerLexerStaticData struct {
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

func slqlexerLexerInit() {
	staticData := &SLQLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "'*'", "'sum'", "'avg'", "'max'", "'min'", "'schema'", "'catalog'",
		"'rownum'", "'unique'", "'uniq'", "'count'", "'+'", "'-'", "'.['", "'||'",
		"'/'", "'%'", "'<<'", "'>>'", "'&'", "'&&'", "'~'", "'!'", "", "", "",
		"'group_by'", "'having'", "", "", "", "'null'", "", "", "'('", "')'",
		"'['", "']'", "','", "'|'", "':'", "", "", "'<='", "'<'", "'>='", "'>'",
		"'!='", "'=='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "PROPRIETARY_FUNC_NAME", "JOIN_TYPE",
		"WHERE", "GROUP_BY", "HAVING", "ORDER_BY", "ALIAS_RESERVED", "ARG",
		"NULL", "ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA", "COMMA", "PIPE",
		"COLON", "NN", "NUMBER", "LT_EQ", "LT", "GT_EQ", "GT", "NEQ", "EQ",
		"NAME", "HANDLE", "STRING", "LINECOMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "PROPRIETARY_FUNC_NAME",
		"JOIN_TYPE", "WHERE", "GROUP_BY", "HAVING", "ORDER_BY", "ALIAS_RESERVED",
		"ARG", "NULL", "ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA", "COMMA",
		"PIPE", "COLON", "NN", "NUMBER", "INTF", "EXP", "LT_EQ", "LT", "GT_EQ",
		"GT", "NEQ", "EQ", "NAME", "HANDLE", "STRING", "ESC", "UNICODE", "HEX",
		"DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"LINECOMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 54, 692, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 3, 25, 388, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 401, 8, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 434,
		8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 492, 8, 30, 1, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 504, 8, 33, 10,
		33, 12, 33, 507, 9, 33, 1, 34, 4, 34, 510, 8, 34, 11, 34, 12, 34, 511,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 3, 43,
		534, 8, 43, 1, 43, 1, 43, 1, 43, 4, 43, 539, 8, 43, 11, 43, 12, 43, 540,
		1, 43, 3, 43, 544, 8, 43, 1, 43, 3, 43, 547, 8, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 3, 43, 553, 8, 43, 1, 43, 3, 43, 556, 8, 43, 1, 44, 1, 44, 1, 44,
		5, 44, 561, 8, 44, 10, 44, 12, 44, 564, 9, 44, 3, 44, 566, 8, 44, 1, 45,
		1, 45, 3, 45, 570, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51,
		1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 594, 8, 52, 1, 53, 1, 53, 1,
		53, 1, 53, 5, 53, 600, 8, 53, 10, 53, 12, 53, 603, 9, 53, 1, 54, 1, 54,
		1, 54, 5, 54, 608, 8, 54, 10, 54, 12, 54, 611, 9, 54, 1, 54, 1, 54, 1,
		55, 1, 55, 1, 55, 3, 55, 618, 8, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1,
		61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66,
		1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 1, 70, 1, 71, 1, 71, 1,
		72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77,
		1, 77, 1, 78, 1, 78, 1, 79, 1, 79, 1, 80, 1, 80, 1, 81, 1, 81, 1, 82, 1,
		82, 1, 83, 1, 83, 1, 84, 1, 84, 1, 85, 1, 85, 5, 85, 684, 8, 85, 10, 85,
		12, 85, 687, 9, 85, 1, 85, 1, 85, 1, 85, 1, 85, 1, 685, 0, 86, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39,
		79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 0, 91, 0, 93, 45, 95, 46, 97,
		47, 99, 48, 101, 49, 103, 50, 105, 51, 107, 52, 109, 53, 111, 0, 113, 0,
		115, 0, 117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131, 0,
		133, 0, 135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149, 0,
		151, 0, 153, 0, 155, 0, 157, 0, 159, 0, 161, 0, 163, 0, 165, 0, 167, 0,
		169, 0, 171, 54, 1, 0, 35, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 1,
		0, 49, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 34, 34,
		92, 92, 8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114,
		114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 65, 65, 97, 97, 2,
		0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100, 2, 0,
		70, 70, 102, 102, 2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0,
		73, 73, 105, 105, 2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0,
		76, 76, 108, 108, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0,
		79, 79, 111, 111, 2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0,
		82, 82, 114, 114, 2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0,
		85, 85, 117, 117, 2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0,
		88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 701,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
		1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0,
		97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0,
		0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 171,
		1, 0, 0, 0, 1, 173, 1, 0, 0, 0, 3, 175, 1, 0, 0, 0, 5, 177, 1, 0, 0, 0,
		7, 181, 1, 0, 0, 0, 9, 185, 1, 0, 0, 0, 11, 189, 1, 0, 0, 0, 13, 193, 1,
		0, 0, 0, 15, 200, 1, 0, 0, 0, 17, 208, 1, 0, 0, 0, 19, 215, 1, 0, 0, 0,
		21, 222, 1, 0, 0, 0, 23, 227, 1, 0, 0, 0, 25, 233, 1, 0, 0, 0, 27, 235,
		1, 0, 0, 0, 29, 237, 1, 0, 0, 0, 31, 240, 1, 0, 0, 0, 33, 243, 1, 0, 0,
		0, 35, 245, 1, 0, 0, 0, 37, 247, 1, 0, 0, 0, 39, 250, 1, 0, 0, 0, 41, 253,
		1, 0, 0, 0, 43, 255, 1, 0, 0, 0, 45, 258, 1, 0, 0, 0, 47, 260, 1, 0, 0,
		0, 49, 262, 1, 0, 0, 0, 51, 387, 1, 0, 0, 0, 53, 400, 1, 0, 0, 0, 55, 402,
		1, 0, 0, 0, 57, 411, 1, 0, 0, 0, 59, 433, 1, 0, 0, 0, 61, 491, 1, 0, 0,
		0, 63, 493, 1, 0, 0, 0, 65, 496, 1, 0, 0, 0, 67, 501, 1, 0, 0, 0, 69, 509,
		1, 0, 0, 0, 71, 515, 1, 0, 0, 0, 73, 517, 1, 0, 0, 0, 75, 519, 1, 0, 0,
		0, 77, 521, 1, 0, 0, 0, 79, 523, 1, 0, 0, 0, 81, 525, 1, 0, 0, 0, 83, 527,
		1, 0, 0, 0, 85, 529, 1, 0, 0, 0, 87, 555, 1, 0, 0, 0, 89, 565, 1, 0, 0,
		0, 91, 567, 1, 0, 0, 0, 93, 573, 1, 0, 0, 0, 95, 576, 1, 0, 0, 0, 97, 578,
		1, 0, 0, 0, 99, 581, 1, 0, 0, 0, 101, 583, 1, 0, 0, 0, 103, 586, 1, 0,
		0, 0, 105, 589, 1, 0, 0, 0, 107, 595, 1, 0, 0, 0, 109, 604, 1, 0, 0, 0,
		111, 614, 1, 0, 0, 0, 113, 619, 1, 0, 0, 0, 115, 625, 1, 0, 0, 0, 117,
		627, 1, 0, 0, 0, 119, 629, 1, 0, 0, 0, 121, 631, 1, 0, 0, 0, 123, 633,
		1, 0, 0, 0, 125, 635, 1, 0, 0, 0, 127, 637, 1, 0, 0, 0, 129, 639, 1, 0,
		0, 0, 131, 641, 1, 0, 0, 0, 133, 643, 1, 0, 0, 0, 135, 645, 1, 0, 0, 0,
		137, 647, 1, 0, 0, 0, 139, 649, 1, 0, 0, 0, 141, 651, 1, 0, 0, 0, 143,
		653, 1, 0, 0, 0, 145, 655, 1, 0, 0, 0, 147, 657, 1, 0, 0, 0, 149, 659,
		1, 0, 0, 0, 151, 661, 1, 0, 0, 0, 153, 663, 1, 0, 0, 0, 155, 665, 1, 0,
		0, 0, 157, 667, 1, 0, 0, 0, 159, 669, 1, 0, 0, 0, 161, 671, 1, 0, 0, 0,
		163, 673, 1, 0, 0, 0, 165, 675, 1, 0, 0, 0, 167, 677, 1, 0, 0, 0, 169,
		679, 1, 0, 0, 0, 171, 681, 1, 0, 0, 0, 173, 174, 5, 59, 0, 0, 174, 2, 1,
		0, 0, 0, 175, 176, 5, 42, 0, 0, 176, 4, 1, 0, 0, 0, 177, 178, 5, 115, 0,
		0, 178, 179, 5, 117, 0, 0, 179, 180, 5, 109, 0, 0, 180, 6, 1, 0, 0, 0,
		181, 182, 5, 97, 0, 0, 182, 183, 5, 118, 0, 0, 183, 184, 5, 103, 0, 0,
		184, 8, 1, 0, 0, 0, 185, 186, 5, 109, 0, 0, 186, 187, 5, 97, 0, 0, 187,
		188, 5, 120, 0, 0, 188, 10, 1, 0, 0, 0, 189, 190, 5, 109, 0, 0, 190, 191,
		5, 105, 0, 0, 191, 192, 5, 110, 0, 0, 192, 12, 1, 0, 0, 0, 193, 194, 5,
		115, 0, 0, 194, 195, 5, 99, 0, 0, 195, 196, 5, 104, 0, 0, 196, 197, 5,
		101, 0, 0, 197, 198, 5, 109, 0, 0, 198, 199, 5, 97, 0, 0, 199, 14, 1, 0,
		0, 0, 200, 201, 5, 99, 0, 0, 201, 202, 5, 97, 0, 0, 202, 203, 5, 116, 0,
		0, 203, 204, 5, 97, 0, 0, 204, 205, 5, 108, 0, 0, 205, 206, 5, 111, 0,
		0, 206, 207, 5, 103, 0, 0, 207, 16, 1, 0, 0, 0, 208, 209, 5, 114, 0, 0,
		209, 210, 5, 111, 0, 0, 210, 211, 5, 119, 0, 0, 211, 212, 5, 110, 0, 0,
		212, 213, 5, 117, 0, 0, 213, 214, 5, 109, 0, 0, 214, 18, 1, 0, 0, 0, 215,
		216, 5, 117, 0, 0, 216, 217, 5, 110, 0, 0, 217, 218, 5, 105, 0, 0, 218,
		219, 5, 113, 0, 0, 219, 220, 5, 117, 0, 0, 220, 221, 5, 101, 0, 0, 221,
		20, 1, 0, 0, 0, 222, 223, 5, 117, 0, 0, 223, 224, 5, 110, 0, 0, 224, 225,
		5, 105, 0, 0, 225, 226, 5, 113, 0, 0, 226, 22, 1, 0, 0, 0, 227, 228, 5,
		99, 0, 0, 228, 229, 5, 111, 0, 0, 229, 230, 5, 117, 0, 0, 230, 231, 5,
		110, 0, 0, 231, 232, 5, 116, 0, 0, 232, 24, 1, 0, 0, 0, 233, 234, 5, 43,
		0, 0, 234, 26, 1, 0, 0, 0, 235, 236, 5, 45, 0, 0, 236, 28, 1, 0, 0, 0,
		237, 238, 5, 46, 0, 0, 238, 239, 5, 91, 0, 0, 239, 30, 1, 0, 0, 0, 240,
		241, 5, 124, 0, 0, 241, 242, 5, 124, 0, 0, 242, 32, 1, 0, 0, 0, 243, 244,
		5, 47, 0, 0, 244, 34, 1, 0, 0, 0, 245, 246, 5, 37, 0, 0, 246, 36, 1, 0,
		0, 0, 247, 248, 5, 60, 0, 0, 248, 249, 5, 60, 0, 0, 249, 38, 1, 0, 0, 0,
		250, 251, 5, 62, 0, 0, 251, 252, 5, 62, 0, 0, 252, 40, 1, 0, 0, 0, 253,
		254, 5, 38, 0, 0, 254, 42, 1, 0, 0, 0, 255, 256, 5, 38, 0, 0, 256, 257,
		5, 38, 0, 0, 257, 44, 1, 0, 0, 0, 258, 259, 5, 126, 0, 0, 259, 46, 1, 0,
		0, 0, 260, 261, 5, 33, 0, 0, 261, 48, 1, 0, 0, 0, 262, 263, 5, 95, 0, 0,
		263, 264, 3, 67, 33, 0, 264, 50, 1, 0, 0, 0, 265, 266, 5, 106, 0, 0, 266,
		267, 5, 111, 0, 0, 267, 268, 5, 105, 0, 0, 268, 388, 5, 110, 0, 0, 269,
		270, 5, 105, 0, 0, 270, 271, 5, 110, 0, 0, 271, 272, 5, 110, 0, 0, 272,
		273, 5, 101, 0, 0, 273, 274, 5, 114, 0, 0, 274, 275, 5, 95, 0, 0, 275,
		276, 5, 106, 0, 0, 276, 277, 5, 111, 0, 0, 277, 278, 5, 105, 0, 0, 278,
		388, 5, 110, 0, 0, 279, 280, 5, 108, 0, 0, 280, 281, 5, 101, 0, 0, 281,
		282, 5, 102, 0, 0, 282, 283, 5, 116, 0, 0, 283, 284, 5, 95, 0, 0, 284,
		285, 5, 106, 0, 0, 285, 286, 5, 111, 0, 0, 286, 287, 5, 105, 0, 0, 287,
		388, 5, 110, 0, 0, 288, 289, 5, 108, 0, 0, 289, 290, 5, 106, 0, 0, 290,
		291, 5, 111, 0, 0, 291, 292, 5, 105, 0, 0, 292, 388, 5, 110, 0, 0, 293,
		294, 5, 108, 0, 0, 294, 295, 5, 101, 0, 0, 295, 296, 5, 102, 0, 0, 296,
		297, 5, 116, 0, 0, 297, 298, 5, 95, 0, 0, 298, 299, 5, 111, 0, 0, 299,
		300, 5, 117, 0, 0, 300, 301, 5, 116, 0, 0, 301, 302, 5, 101, 0, 0, 302,
		303, 5, 114, 0, 0, 303, 304, 5, 95, 0, 0, 304, 305, 5, 106, 0, 0, 305,
		306, 5, 111, 0, 0, 306, 307, 5, 105, 0, 0, 307, 388, 5, 110, 0, 0, 308,
		309, 5, 108, 0, 0, 309, 310, 5, 111, 0, 0, 310, 311, 5, 106, 0, 0, 311,
		312, 5, 111, 0, 0, 312, 313, 5, 105, 0, 0, 313, 388, 5, 110, 0, 0, 314,
		315, 5, 114, 0, 0, 315, 316, 5, 105, 0, 0, 316, 317, 5, 103, 0, 0, 317,
		318, 5, 104, 0, 0, 318, 319, 5, 116, 0, 0, 319, 320, 5, 95, 0, 0, 320,
		321, 5, 106, 0, 0, 321, 322, 5, 111, 0, 0, 322, 323, 5, 105, 0, 0, 323,
		388, 5, 110, 0, 0, 324, 325, 5, 114, 0, 0, 325, 326, 5, 106, 0, 0, 326,
		327, 5, 111, 0, 0, 327, 328, 5, 105, 0, 0, 328, 388, 5, 110, 0, 0, 329,
		330, 5, 114, 0, 0, 330, 331, 5, 105, 0, 0, 331, 332, 5, 103, 0, 0, 332,
		333, 5, 104, 0, 0, 333, 334, 5, 116, 0, 0, 334, 335, 5, 95, 0, 0, 335,
		336, 5, 111, 0, 0, 336, 337, 5, 117, 0, 0, 337, 338, 5, 116, 0, 0, 338,
		339, 5, 101, 0, 0, 339, 340, 5, 114, 0, 0, 340, 341, 5, 95, 0, 0, 341,
		342, 5, 106, 0, 0, 342, 343, 5, 111, 0, 0, 343, 344, 5, 105, 0, 0, 344,
		388, 5, 110, 0, 0, 345, 346, 5, 114, 0, 0, 346, 347, 5, 111, 0, 0, 347,
		348, 5, 106, 0, 0, 348, 349, 5, 111, 0, 0, 349, 350, 5, 105, 0, 0, 350,
		388, 5, 110, 0, 0, 351, 352, 5, 102, 0, 0, 352, 353, 5, 117, 0, 0, 353,
		354, 5, 108, 0, 0, 354, 355, 5, 108, 0, 0, 355, 356, 5, 95, 0, 0, 356,
		357, 5, 111, 0, 0, 357, 358, 5, 117, 0, 0, 358, 359, 5, 116, 0, 0, 359,
		360, 5, 101, 0, 0, 360, 361, 5, 114, 0, 0, 361, 362, 5, 95, 0, 0, 362,
		363, 5, 106, 0, 0, 363, 364, 5, 111, 0, 0, 364, 365, 5, 105, 0, 0, 365,
		388, 5, 110, 0, 0, 366, 367, 5, 102, 0, 0, 367, 368, 5, 111, 0, 0, 368,
		369, 5, 106, 0, 0, 369, 370, 5, 111, 0, 0, 370, 371, 5, 105, 0, 0, 371,
		388, 5, 110, 0, 0, 372, 373, 5, 99, 0, 0, 373, 374, 5, 114, 0, 0, 374,
		375, 5, 111, 0, 0, 375, 376, 5, 115, 0, 0, 376, 377, 5, 115, 0, 0, 377,
		378, 5, 95, 0, 0, 378, 379, 5, 106, 0, 0, 379, 380, 5, 111, 0, 0, 380,
		381, 5, 105, 0, 0, 381, 388, 5, 110, 0, 0, 382, 383, 5, 120, 0, 0, 383,
		384, 5, 106, 0, 0, 384, 385, 5, 111, 0, 0, 385, 386, 5, 105, 0, 0, 386,
		388, 5, 110, 0, 0, 387, 265, 1, 0, 0, 0, 387, 269, 1, 0, 0, 0, 387, 279,
		1, 0, 0, 0, 387, 288, 1, 0, 0, 0, 387, 293, 1, 0, 0, 0, 387, 308, 1, 0,
		0, 0, 387, 314, 1, 0, 0, 0, 387, 324, 1, 0, 0, 0, 387, 329, 1, 0, 0, 0,
		387, 345, 1, 0, 0, 0, 387, 351, 1, 0, 0, 0, 387, 366, 1, 0, 0, 0, 387,
		372, 1, 0, 0, 0, 387, 382, 1, 0, 0, 0, 388, 52, 1, 0, 0, 0, 389, 390, 5,
		119, 0, 0, 390, 391, 5, 104, 0, 0, 391, 392, 5, 101, 0, 0, 392, 393, 5,
		114, 0, 0, 393, 401, 5, 101, 0, 0, 394, 395, 5, 115, 0, 0, 395, 396, 5,
		101, 0, 0, 396, 397, 5, 108, 0, 0, 397, 398, 5, 101, 0, 0, 398, 399, 5,
		99, 0, 0, 399, 401, 5, 116, 0, 0, 400, 389, 1, 0, 0, 0, 400, 394, 1, 0,
		0, 0, 401, 54, 1, 0, 0, 0, 402, 403, 5, 103, 0, 0, 403, 404, 5, 114, 0,
		0, 404, 405, 5, 111, 0, 0, 405, 406, 5, 117, 0, 0, 406, 407, 5, 112, 0,
		0, 407, 408, 5, 95, 0, 0, 408, 409, 5, 98, 0, 0, 409, 410, 5, 121, 0, 0,
		410, 56, 1, 0, 0, 0, 411, 412, 5, 104, 0, 0, 412, 413, 5, 97, 0, 0, 413,
		414, 5, 118, 0, 0, 414, 415, 5, 105, 0, 0, 415, 416, 5, 110, 0, 0, 416,
		417, 5, 103, 0, 0, 417, 58, 1, 0, 0, 0, 418, 419, 5, 111, 0, 0, 419, 420,
		5, 114, 0, 0, 420, 421, 5, 100, 0, 0, 421, 422, 5, 101, 0, 0, 422, 423,
		5, 114, 0, 0, 423, 424, 5, 95, 0, 0, 424, 425, 5, 98, 0, 0, 425, 434, 5,
		121, 0, 0, 426, 427, 5, 115, 0, 0, 427, 428, 5, 111, 0, 0, 428, 429, 5,
		114, 0, 0, 429, 430, 5, 116, 0, 0, 430, 431, 5, 95, 0, 0, 431, 432, 5,
		98, 0, 0, 432, 434, 5, 121, 0, 0, 433, 418, 1, 0, 0, 0, 433, 426, 1, 0,
		0, 0, 434, 60, 1, 0, 0, 0, 435, 436, 5, 58, 0, 0, 436, 437, 5, 99, 0, 0,
		437, 438, 5, 111, 0, 0, 438, 439, 5, 117, 0, 0, 439, 440, 5, 110, 0, 0,
		440, 492, 5, 116, 0, 0, 441, 442, 5, 58, 0, 0, 442, 443, 5, 99, 0, 0, 443,
		444, 5, 111, 0, 0, 444, 445, 5, 117, 0, 0, 445, 446, 5, 110, 0, 0, 446,
		447, 5, 116, 0, 0, 447, 448, 5, 95, 0, 0, 448, 449, 5, 117, 0, 0, 449,
		450, 5, 110, 0, 0, 450, 451, 5, 105, 0, 0, 451, 452, 5, 113, 0, 0, 452,
		453, 5, 117, 0, 0, 453, 492, 5, 101, 0, 0, 454, 455, 5, 58, 0, 0, 455,
		456, 5, 97, 0, 0, 456, 457, 5, 118, 0, 0, 457, 492, 5, 103, 0, 0, 458,
		459, 5, 58, 0, 0, 459, 460, 5, 103, 0, 0, 460, 461, 5, 114, 0, 0, 461,
		462, 5, 111, 0, 0, 462, 463, 5, 117, 0, 0, 463, 464, 5, 112, 0, 0, 464,
		465, 5, 95, 0, 0, 465, 466, 5, 98, 0, 0, 466, 492, 5, 121, 0, 0, 467, 468,
		5, 58, 0, 0, 468, 469, 5, 109, 0, 0, 469, 470, 5, 97, 0, 0, 470, 492, 5,
		120, 0, 0, 471, 472, 5, 58, 0, 0, 472, 473, 5, 109, 0, 0, 473, 474, 5,
		105, 0, 0, 474, 492, 5, 110, 0, 0, 475, 476, 5, 58, 0, 0, 476, 477, 5,
		111, 0, 0, 477, 478, 5, 114, 0, 0, 478, 479, 5, 100, 0, 0, 479, 480, 5,
		101, 0, 0, 480, 481, 5, 114, 0, 0, 481, 482, 5, 95, 0, 0, 482, 483, 5,
		98, 0, 0, 483, 492, 5, 121, 0, 0, 484, 485, 5, 58, 0, 0, 485, 486, 5, 117,
		0, 0, 486, 487, 5, 110, 0, 0, 487, 488, 5, 105, 0, 0, 488, 489, 5, 113,
		0, 0, 489, 490, 5, 117, 0, 0, 490, 492, 5, 101, 0, 0, 491, 435, 1, 0, 0,
		0, 491, 441, 1, 0, 0, 0, 491, 454, 1, 0, 0, 0, 491, 458, 1, 0, 0, 0, 491,
		467, 1, 0, 0, 0, 491, 471, 1, 0, 0, 0, 491, 475, 1, 0, 0, 0, 491, 484,
		1, 0, 0, 0, 492, 62, 1, 0, 0, 0, 493, 494, 5, 36, 0, 0, 494, 495, 3, 67,
		33, 0, 495, 64, 1, 0, 0, 0, 496, 497, 5, 110, 0, 0, 497, 498, 5, 117, 0,
		0, 498, 499, 5, 108, 0, 0, 499, 500, 5, 108, 0, 0, 500, 66, 1, 0, 0, 0,
		501, 505, 7, 0, 0, 0, 502, 504, 7, 1, 0, 0, 503, 502, 1, 0, 0, 0, 504,
		507, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0, 506, 68, 1,
		0, 0, 0, 507, 505, 1, 0, 0, 0, 508, 510, 7, 2, 0, 0, 509, 508, 1, 0, 0,
		0, 510, 511, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512,
		513, 1, 0, 0, 0, 513, 514, 6, 34, 0, 0, 514, 70, 1, 0, 0, 0, 515, 516,
		5, 40, 0, 0, 516, 72, 1, 0, 0, 0, 517, 518, 5, 41, 0, 0, 518, 74, 1, 0,
		0, 0, 519, 520, 5, 91, 0, 0, 520, 76, 1, 0, 0, 0, 521, 522, 5, 93, 0, 0,
		522, 78, 1, 0, 0, 0, 523, 524, 5, 44, 0, 0, 524, 80, 1, 0, 0, 0, 525, 526,
		5, 124, 0, 0, 526, 82, 1, 0, 0, 0, 527, 528, 5, 58, 0, 0, 528, 84, 1, 0,
		0, 0, 529, 530, 3, 89, 44, 0, 530, 86, 1, 0, 0, 0, 531, 556, 3, 85, 42,
		0, 532, 534, 5, 45, 0, 0, 533, 532, 1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534,
		535, 1, 0, 0, 0, 535, 536, 3, 89, 44, 0, 536, 538, 5, 46, 0, 0, 537, 539,
		7, 3, 0, 0, 538, 537, 1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540, 538, 1, 0,
		0, 0, 540, 541, 1, 0, 0, 0, 541, 543, 1, 0, 0, 0, 542, 544, 3, 91, 45,
		0, 543, 542, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 556, 1, 0, 0, 0, 545,
		547, 5, 45, 0, 0, 546, 545, 1, 0, 0, 0, 546, 547, 1, 0, 0, 0, 547, 548,
		1, 0, 0, 0, 548, 549, 3, 89, 44, 0, 549, 550, 3, 91, 45, 0, 550, 556, 1,
		0, 0, 0, 551, 553, 5, 45, 0, 0, 552, 551, 1, 0, 0, 0, 552, 553, 1, 0, 0,
		0, 553, 554, 1, 0, 0, 0, 554, 556, 3, 89, 44, 0, 555, 531, 1, 0, 0, 0,
		555, 533, 1, 0, 0, 0, 555, 546, 1, 0, 0, 0, 555, 552, 1, 0, 0, 0, 556,
		88, 1, 0, 0, 0, 557, 566, 5, 48, 0, 0, 558, 562, 7, 4, 0, 0, 559, 561,
		7, 3, 0, 0, 560, 559, 1, 0, 0, 0, 561, 564, 1, 0, 0, 0, 562, 560, 1, 0,
		0, 0, 562, 563, 1, 0, 0, 0, 563, 566, 1, 0, 0, 0, 564, 562, 1, 0, 0, 0,
		565, 557, 1, 0, 0, 0, 565, 558, 1, 0, 0, 0, 566, 90, 1, 0, 0, 0, 567, 569,
		7, 5, 0, 0, 568, 570, 7, 6, 0, 0, 569, 568, 1, 0, 0, 0, 569, 570, 1, 0,
		0, 0, 570, 571, 1, 0, 0, 0, 571, 572, 3, 89, 44, 0, 572, 92, 1, 0, 0, 0,
		573, 574, 5, 60, 0, 0, 574, 575, 5, 61, 0, 0, 575, 94, 1, 0, 0, 0, 576,
		577, 5, 60, 0, 0, 577, 96, 1, 0, 0, 0, 578, 579, 5, 62, 0, 0, 579, 580,
		5, 61, 0, 0, 580, 98, 1, 0, 0, 0, 581, 582, 5, 62, 0, 0, 582, 100, 1, 0,
		0, 0, 583, 584, 5, 33, 0, 0, 584, 585, 5, 61, 0, 0, 585, 102, 1, 0, 0,
		0, 586, 587, 5, 61, 0, 0, 587, 588, 5, 61, 0, 0, 588, 104, 1, 0, 0, 0,
		589, 593, 5, 46, 0, 0, 590, 594, 3, 63, 31, 0, 591, 594, 3, 67, 33, 0,
		592, 594, 3, 109, 54, 0, 593, 590, 1, 0, 0, 0, 593, 591, 1, 0, 0, 0, 593,
		592, 1, 0, 0, 0, 594, 106, 1, 0, 0, 0, 595, 596, 5, 64, 0, 0, 596, 601,
		3, 67, 33, 0, 597, 598, 5, 47, 0, 0, 598, 600, 3, 67, 33, 0, 599, 597,
		1, 0, 0, 0, 600, 603, 1, 0, 0, 0, 601, 599, 1, 0, 0, 0, 601, 602, 1, 0,
		0, 0, 602, 108, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 604, 609, 5, 34, 0, 0,
		605, 608, 3, 111, 55, 0, 606, 608, 8, 7, 0, 0, 607, 605, 1, 0, 0, 0, 607,
		606, 1, 0, 0, 0, 608, 611, 1, 0, 0, 0, 609, 607, 1, 0, 0, 0, 609, 610,
		1, 0, 0, 0, 610, 612, 1, 0, 0, 0, 611, 609, 1, 0, 0, 0, 612, 613, 5, 34,
		0, 0, 613, 110, 1, 0, 0, 0, 614, 617, 5, 92, 0, 0, 615, 618, 7, 8, 0, 0,
		616, 618, 3, 113, 56, 0, 617, 615, 1, 0, 0, 0, 617, 616, 1, 0, 0, 0, 618,
		112, 1, 0, 0, 0, 619, 620, 5, 117, 0, 0, 620, 621, 3, 115, 57, 0, 621,
		622, 3, 115, 57, 0, 622, 623, 3, 115, 57, 0, 623, 624, 3, 115, 57, 0, 624,
		114, 1, 0, 0, 0, 625, 626, 7, 9, 0, 0, 626, 116, 1, 0, 0, 0, 627, 628,
		7, 3, 0, 0, 628, 118, 1, 0, 0, 0, 629, 630, 7, 10, 0, 0, 630, 120, 1, 0,
		0, 0, 631, 632, 7, 11, 0, 0, 632, 122, 1, 0, 0, 0, 633, 634, 7, 12, 0,
		0, 634, 124, 1, 0, 0, 0, 635, 636, 7, 13, 0, 0, 636, 126, 1, 0, 0, 0, 637,
		638, 7, 5, 0, 0, 638, 128, 1, 0, 0, 0, 639, 640, 7, 14, 0, 0, 640, 130,
		1, 0, 0, 0, 641, 642, 7, 15, 0, 0, 642, 132, 1, 0, 0, 0, 643, 644, 7, 16,
		0, 0, 644, 134, 1, 0, 0, 0, 645, 646, 7, 17, 0, 0, 646, 136, 1, 0, 0, 0,
		647, 648, 7, 18, 0, 0, 648, 138, 1, 0, 0, 0, 649, 650, 7, 19, 0, 0, 650,
		140, 1, 0, 0, 0, 651, 652, 7, 20, 0, 0, 652, 142, 1, 0, 0, 0, 653, 654,
		7, 21, 0, 0, 654, 144, 1, 0, 0, 0, 655, 656, 7, 22, 0, 0, 656, 146, 1,
		0, 0, 0, 657, 658, 7, 23, 0, 0, 658, 148, 1, 0, 0, 0, 659, 660, 7, 24,
		0, 0, 660, 150, 1, 0, 0, 0, 661, 662, 7, 25, 0, 0, 662, 152, 1, 0, 0, 0,
		663, 664, 7, 26, 0, 0, 664, 154, 1, 0, 0, 0, 665, 666, 7, 27, 0, 0, 666,
		156, 1, 0, 0, 0, 667, 668, 7, 28, 0, 0, 668, 158, 1, 0, 0, 0, 669, 670,
		7, 29, 0, 0, 670, 160, 1, 0, 0, 0, 671, 672, 7, 30, 0, 0, 672, 162, 1,
		0, 0, 0, 673, 674, 7, 31, 0, 0, 674, 164, 1, 0, 0, 0, 675, 676, 7, 32,
		0, 0, 676, 166, 1, 0, 0, 0, 677, 678, 7, 33, 0, 0, 678, 168, 1, 0, 0, 0,
		679, 680, 7, 34, 0, 0, 680, 170, 1, 0, 0, 0, 681, 685, 5, 35, 0, 0, 682,
		684, 9, 0, 0, 0, 683, 682, 1, 0, 0, 0, 684, 687, 1, 0, 0, 0, 685, 686,
		1, 0, 0, 0, 685, 683, 1, 0, 0, 0, 686, 688, 1, 0, 0, 0, 687, 685, 1, 0,
		0, 0, 688, 689, 5, 10, 0, 0, 689, 690, 1, 0, 0, 0, 690, 691, 6, 85, 0,
		0, 691, 172, 1, 0, 0, 0, 22, 0, 387, 400, 433, 491, 505, 511, 533, 540,
		543, 546, 552, 555, 562, 565, 569, 593, 601, 607, 609, 617, 685, 1, 6,
		0, 0,
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

// SLQLexerInit initializes any static state used to implement SLQLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSLQLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SLQLexerInit() {
	staticData := &SLQLexerLexerStaticData
	staticData.once.Do(slqlexerLexerInit)
}

// NewSLQLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSLQLexer(input antlr.CharStream) *SLQLexer {
	SLQLexerInit()
	l := new(SLQLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SLQLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SLQ.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SLQLexer tokens.
const (
	SLQLexerT__0                  = 1
	SLQLexerT__1                  = 2
	SLQLexerT__2                  = 3
	SLQLexerT__3                  = 4
	SLQLexerT__4                  = 5
	SLQLexerT__5                  = 6
	SLQLexerT__6                  = 7
	SLQLexerT__7                  = 8
	SLQLexerT__8                  = 9
	SLQLexerT__9                  = 10
	SLQLexerT__10                 = 11
	SLQLexerT__11                 = 12
	SLQLexerT__12                 = 13
	SLQLexerT__13                 = 14
	SLQLexerT__14                 = 15
	SLQLexerT__15                 = 16
	SLQLexerT__16                 = 17
	SLQLexerT__17                 = 18
	SLQLexerT__18                 = 19
	SLQLexerT__19                 = 20
	SLQLexerT__20                 = 21
	SLQLexerT__21                 = 22
	SLQLexerT__22                 = 23
	SLQLexerT__23                 = 24
	SLQLexerPROPRIETARY_FUNC_NAME = 25
	SLQLexerJOIN_TYPE             = 26
	SLQLexerWHERE                 = 27
	SLQLexerGROUP_BY              = 28
	SLQLexerHAVING                = 29
	SLQLexerORDER_BY              = 30
	SLQLexerALIAS_RESERVED        = 31
	SLQLexerARG                   = 32
	SLQLexerNULL                  = 33
	SLQLexerID                    = 34
	SLQLexerWS                    = 35
	SLQLexerLPAR                  = 36
	SLQLexerRPAR                  = 37
	SLQLexerLBRA                  = 38
	SLQLexerRBRA                  = 39
	SLQLexerCOMMA                 = 40
	SLQLexerPIPE                  = 41
	SLQLexerCOLON                 = 42
	SLQLexerNN                    = 43
	SLQLexerNUMBER                = 44
	SLQLexerLT_EQ                 = 45
	SLQLexerLT                    = 46
	SLQLexerGT_EQ                 = 47
	SLQLexerGT                    = 48
	SLQLexerNEQ                   = 49
	SLQLexerEQ                    = 50
	SLQLexerNAME                  = 51
	SLQLexerHANDLE                = 52
	SLQLexerSTRING                = 53
	SLQLexerLINECOMMENT           = 54
)
