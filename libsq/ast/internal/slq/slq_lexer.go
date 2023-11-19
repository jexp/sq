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
		"'unique'", "'uniq'", "'count'", "'.['", "'||'", "'/'", "'%'", "'<<'",
		"'>>'", "'&'", "'&&'", "'~'", "'!'", "", "", "", "'group_by'", "'+'",
		"'-'", "", "", "", "'null'", "", "", "'('", "')'", "'['", "']'", "','",
		"'|'", "':'", "", "", "'<='", "'<'", "'>='", "'>'", "'!='", "'=='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "PROPRIETARY_FUNC_NAME", "JOIN_TYPE", "WHERE", "GROUP_BY",
		"ORDER_ASC", "ORDER_DESC", "ORDER_BY", "ALIAS_RESERVED", "ARG", "NULL",
		"ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA", "COMMA", "PIPE", "COLON",
		"NN", "NUMBER", "LT_EQ", "LT", "GT_EQ", "GT", "NEQ", "EQ", "NAME", "HANDLE",
		"STRING", "LINECOMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "PROPRIETARY_FUNC_NAME", "JOIN_TYPE",
		"WHERE", "GROUP_BY", "ORDER_ASC", "ORDER_DESC", "ORDER_BY", "ALIAS_RESERVED",
		"ARG", "NULL", "ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA", "COMMA",
		"PIPE", "COLON", "NN", "NUMBER", "INTF", "EXP", "LT_EQ", "LT", "GT_EQ",
		"GT", "NEQ", "EQ", "NAME", "HANDLE", "STRING", "ESC", "UNICODE", "HEX",
		"DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"LINECOMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 52, 674, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		83, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 373,
		8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 386, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 3, 27, 416, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 474, 8, 28,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5,
		31, 486, 8, 31, 10, 31, 12, 31, 489, 9, 31, 1, 32, 4, 32, 492, 8, 32, 11,
		32, 12, 32, 493, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		41, 1, 41, 3, 41, 516, 8, 41, 1, 41, 1, 41, 1, 41, 4, 41, 521, 8, 41, 11,
		41, 12, 41, 522, 1, 41, 3, 41, 526, 8, 41, 1, 41, 3, 41, 529, 8, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 3, 41, 535, 8, 41, 1, 41, 3, 41, 538, 8, 41, 1,
		42, 1, 42, 1, 42, 5, 42, 543, 8, 42, 10, 42, 12, 42, 546, 9, 42, 3, 42,
		548, 8, 42, 1, 43, 1, 43, 3, 43, 552, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 576, 8, 50,
		1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 582, 8, 51, 10, 51, 12, 51, 585, 9,
		51, 1, 52, 1, 52, 1, 52, 5, 52, 590, 8, 52, 10, 52, 12, 52, 593, 9, 52,
		1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 3, 53, 600, 8, 53, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58,
		1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1,
		63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68,
		1, 69, 1, 69, 1, 70, 1, 70, 1, 71, 1, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1,
		74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77, 1, 78, 1, 78, 1, 79,
		1, 79, 1, 80, 1, 80, 1, 81, 1, 81, 1, 82, 1, 82, 1, 83, 1, 83, 5, 83, 666,
		8, 83, 10, 83, 12, 83, 669, 9, 83, 1, 83, 1, 83, 1, 83, 1, 83, 1, 667,
		0, 84, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 0, 87, 0, 89, 43, 91, 44, 93,
		45, 95, 46, 97, 47, 99, 48, 101, 49, 103, 50, 105, 51, 107, 0, 109, 0,
		111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127, 0,
		129, 0, 131, 0, 133, 0, 135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0,
		147, 0, 149, 0, 151, 0, 153, 0, 155, 0, 157, 0, 159, 0, 161, 0, 163, 0,
		165, 0, 167, 52, 1, 0, 35, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
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
		88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 683,
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
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 89,
		1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0,
		97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0,
		0, 0, 105, 1, 0, 0, 0, 0, 167, 1, 0, 0, 0, 1, 169, 1, 0, 0, 0, 3, 171,
		1, 0, 0, 0, 5, 173, 1, 0, 0, 0, 7, 177, 1, 0, 0, 0, 9, 181, 1, 0, 0, 0,
		11, 185, 1, 0, 0, 0, 13, 189, 1, 0, 0, 0, 15, 196, 1, 0, 0, 0, 17, 204,
		1, 0, 0, 0, 19, 211, 1, 0, 0, 0, 21, 216, 1, 0, 0, 0, 23, 222, 1, 0, 0,
		0, 25, 225, 1, 0, 0, 0, 27, 228, 1, 0, 0, 0, 29, 230, 1, 0, 0, 0, 31, 232,
		1, 0, 0, 0, 33, 235, 1, 0, 0, 0, 35, 238, 1, 0, 0, 0, 37, 240, 1, 0, 0,
		0, 39, 243, 1, 0, 0, 0, 41, 245, 1, 0, 0, 0, 43, 247, 1, 0, 0, 0, 45, 372,
		1, 0, 0, 0, 47, 385, 1, 0, 0, 0, 49, 387, 1, 0, 0, 0, 51, 396, 1, 0, 0,
		0, 53, 398, 1, 0, 0, 0, 55, 415, 1, 0, 0, 0, 57, 473, 1, 0, 0, 0, 59, 475,
		1, 0, 0, 0, 61, 478, 1, 0, 0, 0, 63, 483, 1, 0, 0, 0, 65, 491, 1, 0, 0,
		0, 67, 497, 1, 0, 0, 0, 69, 499, 1, 0, 0, 0, 71, 501, 1, 0, 0, 0, 73, 503,
		1, 0, 0, 0, 75, 505, 1, 0, 0, 0, 77, 507, 1, 0, 0, 0, 79, 509, 1, 0, 0,
		0, 81, 511, 1, 0, 0, 0, 83, 537, 1, 0, 0, 0, 85, 547, 1, 0, 0, 0, 87, 549,
		1, 0, 0, 0, 89, 555, 1, 0, 0, 0, 91, 558, 1, 0, 0, 0, 93, 560, 1, 0, 0,
		0, 95, 563, 1, 0, 0, 0, 97, 565, 1, 0, 0, 0, 99, 568, 1, 0, 0, 0, 101,
		571, 1, 0, 0, 0, 103, 577, 1, 0, 0, 0, 105, 586, 1, 0, 0, 0, 107, 596,
		1, 0, 0, 0, 109, 601, 1, 0, 0, 0, 111, 607, 1, 0, 0, 0, 113, 609, 1, 0,
		0, 0, 115, 611, 1, 0, 0, 0, 117, 613, 1, 0, 0, 0, 119, 615, 1, 0, 0, 0,
		121, 617, 1, 0, 0, 0, 123, 619, 1, 0, 0, 0, 125, 621, 1, 0, 0, 0, 127,
		623, 1, 0, 0, 0, 129, 625, 1, 0, 0, 0, 131, 627, 1, 0, 0, 0, 133, 629,
		1, 0, 0, 0, 135, 631, 1, 0, 0, 0, 137, 633, 1, 0, 0, 0, 139, 635, 1, 0,
		0, 0, 141, 637, 1, 0, 0, 0, 143, 639, 1, 0, 0, 0, 145, 641, 1, 0, 0, 0,
		147, 643, 1, 0, 0, 0, 149, 645, 1, 0, 0, 0, 151, 647, 1, 0, 0, 0, 153,
		649, 1, 0, 0, 0, 155, 651, 1, 0, 0, 0, 157, 653, 1, 0, 0, 0, 159, 655,
		1, 0, 0, 0, 161, 657, 1, 0, 0, 0, 163, 659, 1, 0, 0, 0, 165, 661, 1, 0,
		0, 0, 167, 663, 1, 0, 0, 0, 169, 170, 5, 59, 0, 0, 170, 2, 1, 0, 0, 0,
		171, 172, 5, 42, 0, 0, 172, 4, 1, 0, 0, 0, 173, 174, 5, 115, 0, 0, 174,
		175, 5, 117, 0, 0, 175, 176, 5, 109, 0, 0, 176, 6, 1, 0, 0, 0, 177, 178,
		5, 97, 0, 0, 178, 179, 5, 118, 0, 0, 179, 180, 5, 103, 0, 0, 180, 8, 1,
		0, 0, 0, 181, 182, 5, 109, 0, 0, 182, 183, 5, 97, 0, 0, 183, 184, 5, 120,
		0, 0, 184, 10, 1, 0, 0, 0, 185, 186, 5, 109, 0, 0, 186, 187, 5, 105, 0,
		0, 187, 188, 5, 110, 0, 0, 188, 12, 1, 0, 0, 0, 189, 190, 5, 115, 0, 0,
		190, 191, 5, 99, 0, 0, 191, 192, 5, 104, 0, 0, 192, 193, 5, 101, 0, 0,
		193, 194, 5, 109, 0, 0, 194, 195, 5, 97, 0, 0, 195, 14, 1, 0, 0, 0, 196,
		197, 5, 99, 0, 0, 197, 198, 5, 97, 0, 0, 198, 199, 5, 116, 0, 0, 199, 200,
		5, 97, 0, 0, 200, 201, 5, 108, 0, 0, 201, 202, 5, 111, 0, 0, 202, 203,
		5, 103, 0, 0, 203, 16, 1, 0, 0, 0, 204, 205, 5, 117, 0, 0, 205, 206, 5,
		110, 0, 0, 206, 207, 5, 105, 0, 0, 207, 208, 5, 113, 0, 0, 208, 209, 5,
		117, 0, 0, 209, 210, 5, 101, 0, 0, 210, 18, 1, 0, 0, 0, 211, 212, 5, 117,
		0, 0, 212, 213, 5, 110, 0, 0, 213, 214, 5, 105, 0, 0, 214, 215, 5, 113,
		0, 0, 215, 20, 1, 0, 0, 0, 216, 217, 5, 99, 0, 0, 217, 218, 5, 111, 0,
		0, 218, 219, 5, 117, 0, 0, 219, 220, 5, 110, 0, 0, 220, 221, 5, 116, 0,
		0, 221, 22, 1, 0, 0, 0, 222, 223, 5, 46, 0, 0, 223, 224, 5, 91, 0, 0, 224,
		24, 1, 0, 0, 0, 225, 226, 5, 124, 0, 0, 226, 227, 5, 124, 0, 0, 227, 26,
		1, 0, 0, 0, 228, 229, 5, 47, 0, 0, 229, 28, 1, 0, 0, 0, 230, 231, 5, 37,
		0, 0, 231, 30, 1, 0, 0, 0, 232, 233, 5, 60, 0, 0, 233, 234, 5, 60, 0, 0,
		234, 32, 1, 0, 0, 0, 235, 236, 5, 62, 0, 0, 236, 237, 5, 62, 0, 0, 237,
		34, 1, 0, 0, 0, 238, 239, 5, 38, 0, 0, 239, 36, 1, 0, 0, 0, 240, 241, 5,
		38, 0, 0, 241, 242, 5, 38, 0, 0, 242, 38, 1, 0, 0, 0, 243, 244, 5, 126,
		0, 0, 244, 40, 1, 0, 0, 0, 245, 246, 5, 33, 0, 0, 246, 42, 1, 0, 0, 0,
		247, 248, 5, 95, 0, 0, 248, 249, 3, 63, 31, 0, 249, 44, 1, 0, 0, 0, 250,
		251, 5, 106, 0, 0, 251, 252, 5, 111, 0, 0, 252, 253, 5, 105, 0, 0, 253,
		373, 5, 110, 0, 0, 254, 255, 5, 105, 0, 0, 255, 256, 5, 110, 0, 0, 256,
		257, 5, 110, 0, 0, 257, 258, 5, 101, 0, 0, 258, 259, 5, 114, 0, 0, 259,
		260, 5, 95, 0, 0, 260, 261, 5, 106, 0, 0, 261, 262, 5, 111, 0, 0, 262,
		263, 5, 105, 0, 0, 263, 373, 5, 110, 0, 0, 264, 265, 5, 108, 0, 0, 265,
		266, 5, 101, 0, 0, 266, 267, 5, 102, 0, 0, 267, 268, 5, 116, 0, 0, 268,
		269, 5, 95, 0, 0, 269, 270, 5, 106, 0, 0, 270, 271, 5, 111, 0, 0, 271,
		272, 5, 105, 0, 0, 272, 373, 5, 110, 0, 0, 273, 274, 5, 108, 0, 0, 274,
		275, 5, 106, 0, 0, 275, 276, 5, 111, 0, 0, 276, 277, 5, 105, 0, 0, 277,
		373, 5, 110, 0, 0, 278, 279, 5, 108, 0, 0, 279, 280, 5, 101, 0, 0, 280,
		281, 5, 102, 0, 0, 281, 282, 5, 116, 0, 0, 282, 283, 5, 95, 0, 0, 283,
		284, 5, 111, 0, 0, 284, 285, 5, 117, 0, 0, 285, 286, 5, 116, 0, 0, 286,
		287, 5, 101, 0, 0, 287, 288, 5, 114, 0, 0, 288, 289, 5, 95, 0, 0, 289,
		290, 5, 106, 0, 0, 290, 291, 5, 111, 0, 0, 291, 292, 5, 105, 0, 0, 292,
		373, 5, 110, 0, 0, 293, 294, 5, 108, 0, 0, 294, 295, 5, 111, 0, 0, 295,
		296, 5, 106, 0, 0, 296, 297, 5, 111, 0, 0, 297, 298, 5, 105, 0, 0, 298,
		373, 5, 110, 0, 0, 299, 300, 5, 114, 0, 0, 300, 301, 5, 105, 0, 0, 301,
		302, 5, 103, 0, 0, 302, 303, 5, 104, 0, 0, 303, 304, 5, 116, 0, 0, 304,
		305, 5, 95, 0, 0, 305, 306, 5, 106, 0, 0, 306, 307, 5, 111, 0, 0, 307,
		308, 5, 105, 0, 0, 308, 373, 5, 110, 0, 0, 309, 310, 5, 114, 0, 0, 310,
		311, 5, 106, 0, 0, 311, 312, 5, 111, 0, 0, 312, 313, 5, 105, 0, 0, 313,
		373, 5, 110, 0, 0, 314, 315, 5, 114, 0, 0, 315, 316, 5, 105, 0, 0, 316,
		317, 5, 103, 0, 0, 317, 318, 5, 104, 0, 0, 318, 319, 5, 116, 0, 0, 319,
		320, 5, 95, 0, 0, 320, 321, 5, 111, 0, 0, 321, 322, 5, 117, 0, 0, 322,
		323, 5, 116, 0, 0, 323, 324, 5, 101, 0, 0, 324, 325, 5, 114, 0, 0, 325,
		326, 5, 95, 0, 0, 326, 327, 5, 106, 0, 0, 327, 328, 5, 111, 0, 0, 328,
		329, 5, 105, 0, 0, 329, 373, 5, 110, 0, 0, 330, 331, 5, 114, 0, 0, 331,
		332, 5, 111, 0, 0, 332, 333, 5, 106, 0, 0, 333, 334, 5, 111, 0, 0, 334,
		335, 5, 105, 0, 0, 335, 373, 5, 110, 0, 0, 336, 337, 5, 102, 0, 0, 337,
		338, 5, 117, 0, 0, 338, 339, 5, 108, 0, 0, 339, 340, 5, 108, 0, 0, 340,
		341, 5, 95, 0, 0, 341, 342, 5, 111, 0, 0, 342, 343, 5, 117, 0, 0, 343,
		344, 5, 116, 0, 0, 344, 345, 5, 101, 0, 0, 345, 346, 5, 114, 0, 0, 346,
		347, 5, 95, 0, 0, 347, 348, 5, 106, 0, 0, 348, 349, 5, 111, 0, 0, 349,
		350, 5, 105, 0, 0, 350, 373, 5, 110, 0, 0, 351, 352, 5, 102, 0, 0, 352,
		353, 5, 111, 0, 0, 353, 354, 5, 106, 0, 0, 354, 355, 5, 111, 0, 0, 355,
		356, 5, 105, 0, 0, 356, 373, 5, 110, 0, 0, 357, 358, 5, 99, 0, 0, 358,
		359, 5, 114, 0, 0, 359, 360, 5, 111, 0, 0, 360, 361, 5, 115, 0, 0, 361,
		362, 5, 115, 0, 0, 362, 363, 5, 95, 0, 0, 363, 364, 5, 106, 0, 0, 364,
		365, 5, 111, 0, 0, 365, 366, 5, 105, 0, 0, 366, 373, 5, 110, 0, 0, 367,
		368, 5, 120, 0, 0, 368, 369, 5, 106, 0, 0, 369, 370, 5, 111, 0, 0, 370,
		371, 5, 105, 0, 0, 371, 373, 5, 110, 0, 0, 372, 250, 1, 0, 0, 0, 372, 254,
		1, 0, 0, 0, 372, 264, 1, 0, 0, 0, 372, 273, 1, 0, 0, 0, 372, 278, 1, 0,
		0, 0, 372, 293, 1, 0, 0, 0, 372, 299, 1, 0, 0, 0, 372, 309, 1, 0, 0, 0,
		372, 314, 1, 0, 0, 0, 372, 330, 1, 0, 0, 0, 372, 336, 1, 0, 0, 0, 372,
		351, 1, 0, 0, 0, 372, 357, 1, 0, 0, 0, 372, 367, 1, 0, 0, 0, 373, 46, 1,
		0, 0, 0, 374, 375, 5, 119, 0, 0, 375, 376, 5, 104, 0, 0, 376, 377, 5, 101,
		0, 0, 377, 378, 5, 114, 0, 0, 378, 386, 5, 101, 0, 0, 379, 380, 5, 115,
		0, 0, 380, 381, 5, 101, 0, 0, 381, 382, 5, 108, 0, 0, 382, 383, 5, 101,
		0, 0, 383, 384, 5, 99, 0, 0, 384, 386, 5, 116, 0, 0, 385, 374, 1, 0, 0,
		0, 385, 379, 1, 0, 0, 0, 386, 48, 1, 0, 0, 0, 387, 388, 5, 103, 0, 0, 388,
		389, 5, 114, 0, 0, 389, 390, 5, 111, 0, 0, 390, 391, 5, 117, 0, 0, 391,
		392, 5, 112, 0, 0, 392, 393, 5, 95, 0, 0, 393, 394, 5, 98, 0, 0, 394, 395,
		5, 121, 0, 0, 395, 50, 1, 0, 0, 0, 396, 397, 5, 43, 0, 0, 397, 52, 1, 0,
		0, 0, 398, 399, 5, 45, 0, 0, 399, 54, 1, 0, 0, 0, 400, 401, 5, 111, 0,
		0, 401, 402, 5, 114, 0, 0, 402, 403, 5, 100, 0, 0, 403, 404, 5, 101, 0,
		0, 404, 405, 5, 114, 0, 0, 405, 406, 5, 95, 0, 0, 406, 407, 5, 98, 0, 0,
		407, 416, 5, 121, 0, 0, 408, 409, 5, 115, 0, 0, 409, 410, 5, 111, 0, 0,
		410, 411, 5, 114, 0, 0, 411, 412, 5, 116, 0, 0, 412, 413, 5, 95, 0, 0,
		413, 414, 5, 98, 0, 0, 414, 416, 5, 121, 0, 0, 415, 400, 1, 0, 0, 0, 415,
		408, 1, 0, 0, 0, 416, 56, 1, 0, 0, 0, 417, 418, 5, 58, 0, 0, 418, 419,
		5, 99, 0, 0, 419, 420, 5, 111, 0, 0, 420, 421, 5, 117, 0, 0, 421, 422,
		5, 110, 0, 0, 422, 474, 5, 116, 0, 0, 423, 424, 5, 58, 0, 0, 424, 425,
		5, 99, 0, 0, 425, 426, 5, 111, 0, 0, 426, 427, 5, 117, 0, 0, 427, 428,
		5, 110, 0, 0, 428, 429, 5, 116, 0, 0, 429, 430, 5, 95, 0, 0, 430, 431,
		5, 117, 0, 0, 431, 432, 5, 110, 0, 0, 432, 433, 5, 105, 0, 0, 433, 434,
		5, 113, 0, 0, 434, 435, 5, 117, 0, 0, 435, 474, 5, 101, 0, 0, 436, 437,
		5, 58, 0, 0, 437, 438, 5, 97, 0, 0, 438, 439, 5, 118, 0, 0, 439, 474, 5,
		103, 0, 0, 440, 441, 5, 58, 0, 0, 441, 442, 5, 103, 0, 0, 442, 443, 5,
		114, 0, 0, 443, 444, 5, 111, 0, 0, 444, 445, 5, 117, 0, 0, 445, 446, 5,
		112, 0, 0, 446, 447, 5, 95, 0, 0, 447, 448, 5, 98, 0, 0, 448, 474, 5, 121,
		0, 0, 449, 450, 5, 58, 0, 0, 450, 451, 5, 109, 0, 0, 451, 452, 5, 97, 0,
		0, 452, 474, 5, 120, 0, 0, 453, 454, 5, 58, 0, 0, 454, 455, 5, 109, 0,
		0, 455, 456, 5, 105, 0, 0, 456, 474, 5, 110, 0, 0, 457, 458, 5, 58, 0,
		0, 458, 459, 5, 111, 0, 0, 459, 460, 5, 114, 0, 0, 460, 461, 5, 100, 0,
		0, 461, 462, 5, 101, 0, 0, 462, 463, 5, 114, 0, 0, 463, 464, 5, 95, 0,
		0, 464, 465, 5, 98, 0, 0, 465, 474, 5, 121, 0, 0, 466, 467, 5, 58, 0, 0,
		467, 468, 5, 117, 0, 0, 468, 469, 5, 110, 0, 0, 469, 470, 5, 105, 0, 0,
		470, 471, 5, 113, 0, 0, 471, 472, 5, 117, 0, 0, 472, 474, 5, 101, 0, 0,
		473, 417, 1, 0, 0, 0, 473, 423, 1, 0, 0, 0, 473, 436, 1, 0, 0, 0, 473,
		440, 1, 0, 0, 0, 473, 449, 1, 0, 0, 0, 473, 453, 1, 0, 0, 0, 473, 457,
		1, 0, 0, 0, 473, 466, 1, 0, 0, 0, 474, 58, 1, 0, 0, 0, 475, 476, 5, 36,
		0, 0, 476, 477, 3, 63, 31, 0, 477, 60, 1, 0, 0, 0, 478, 479, 5, 110, 0,
		0, 479, 480, 5, 117, 0, 0, 480, 481, 5, 108, 0, 0, 481, 482, 5, 108, 0,
		0, 482, 62, 1, 0, 0, 0, 483, 487, 7, 0, 0, 0, 484, 486, 7, 1, 0, 0, 485,
		484, 1, 0, 0, 0, 486, 489, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488,
		1, 0, 0, 0, 488, 64, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 492, 7, 2,
		0, 0, 491, 490, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0,
		493, 494, 1, 0, 0, 0, 494, 495, 1, 0, 0, 0, 495, 496, 6, 32, 0, 0, 496,
		66, 1, 0, 0, 0, 497, 498, 5, 40, 0, 0, 498, 68, 1, 0, 0, 0, 499, 500, 5,
		41, 0, 0, 500, 70, 1, 0, 0, 0, 501, 502, 5, 91, 0, 0, 502, 72, 1, 0, 0,
		0, 503, 504, 5, 93, 0, 0, 504, 74, 1, 0, 0, 0, 505, 506, 5, 44, 0, 0, 506,
		76, 1, 0, 0, 0, 507, 508, 5, 124, 0, 0, 508, 78, 1, 0, 0, 0, 509, 510,
		5, 58, 0, 0, 510, 80, 1, 0, 0, 0, 511, 512, 3, 85, 42, 0, 512, 82, 1, 0,
		0, 0, 513, 538, 3, 81, 40, 0, 514, 516, 5, 45, 0, 0, 515, 514, 1, 0, 0,
		0, 515, 516, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 518, 3, 85, 42, 0,
		518, 520, 5, 46, 0, 0, 519, 521, 7, 3, 0, 0, 520, 519, 1, 0, 0, 0, 521,
		522, 1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 525,
		1, 0, 0, 0, 524, 526, 3, 87, 43, 0, 525, 524, 1, 0, 0, 0, 525, 526, 1,
		0, 0, 0, 526, 538, 1, 0, 0, 0, 527, 529, 5, 45, 0, 0, 528, 527, 1, 0, 0,
		0, 528, 529, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530, 531, 3, 85, 42, 0,
		531, 532, 3, 87, 43, 0, 532, 538, 1, 0, 0, 0, 533, 535, 5, 45, 0, 0, 534,
		533, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 538,
		3, 85, 42, 0, 537, 513, 1, 0, 0, 0, 537, 515, 1, 0, 0, 0, 537, 528, 1,
		0, 0, 0, 537, 534, 1, 0, 0, 0, 538, 84, 1, 0, 0, 0, 539, 548, 5, 48, 0,
		0, 540, 544, 7, 4, 0, 0, 541, 543, 7, 3, 0, 0, 542, 541, 1, 0, 0, 0, 543,
		546, 1, 0, 0, 0, 544, 542, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 548,
		1, 0, 0, 0, 546, 544, 1, 0, 0, 0, 547, 539, 1, 0, 0, 0, 547, 540, 1, 0,
		0, 0, 548, 86, 1, 0, 0, 0, 549, 551, 7, 5, 0, 0, 550, 552, 7, 6, 0, 0,
		551, 550, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0, 553,
		554, 3, 85, 42, 0, 554, 88, 1, 0, 0, 0, 555, 556, 5, 60, 0, 0, 556, 557,
		5, 61, 0, 0, 557, 90, 1, 0, 0, 0, 558, 559, 5, 60, 0, 0, 559, 92, 1, 0,
		0, 0, 560, 561, 5, 62, 0, 0, 561, 562, 5, 61, 0, 0, 562, 94, 1, 0, 0, 0,
		563, 564, 5, 62, 0, 0, 564, 96, 1, 0, 0, 0, 565, 566, 5, 33, 0, 0, 566,
		567, 5, 61, 0, 0, 567, 98, 1, 0, 0, 0, 568, 569, 5, 61, 0, 0, 569, 570,
		5, 61, 0, 0, 570, 100, 1, 0, 0, 0, 571, 575, 5, 46, 0, 0, 572, 576, 3,
		59, 29, 0, 573, 576, 3, 63, 31, 0, 574, 576, 3, 105, 52, 0, 575, 572, 1,
		0, 0, 0, 575, 573, 1, 0, 0, 0, 575, 574, 1, 0, 0, 0, 576, 102, 1, 0, 0,
		0, 577, 578, 5, 64, 0, 0, 578, 583, 3, 63, 31, 0, 579, 580, 5, 47, 0, 0,
		580, 582, 3, 63, 31, 0, 581, 579, 1, 0, 0, 0, 582, 585, 1, 0, 0, 0, 583,
		581, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0, 584, 104, 1, 0, 0, 0, 585, 583,
		1, 0, 0, 0, 586, 591, 5, 34, 0, 0, 587, 590, 3, 107, 53, 0, 588, 590, 8,
		7, 0, 0, 589, 587, 1, 0, 0, 0, 589, 588, 1, 0, 0, 0, 590, 593, 1, 0, 0,
		0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592, 594, 1, 0, 0, 0, 593,
		591, 1, 0, 0, 0, 594, 595, 5, 34, 0, 0, 595, 106, 1, 0, 0, 0, 596, 599,
		5, 92, 0, 0, 597, 600, 7, 8, 0, 0, 598, 600, 3, 109, 54, 0, 599, 597, 1,
		0, 0, 0, 599, 598, 1, 0, 0, 0, 600, 108, 1, 0, 0, 0, 601, 602, 5, 117,
		0, 0, 602, 603, 3, 111, 55, 0, 603, 604, 3, 111, 55, 0, 604, 605, 3, 111,
		55, 0, 605, 606, 3, 111, 55, 0, 606, 110, 1, 0, 0, 0, 607, 608, 7, 9, 0,
		0, 608, 112, 1, 0, 0, 0, 609, 610, 7, 3, 0, 0, 610, 114, 1, 0, 0, 0, 611,
		612, 7, 10, 0, 0, 612, 116, 1, 0, 0, 0, 613, 614, 7, 11, 0, 0, 614, 118,
		1, 0, 0, 0, 615, 616, 7, 12, 0, 0, 616, 120, 1, 0, 0, 0, 617, 618, 7, 13,
		0, 0, 618, 122, 1, 0, 0, 0, 619, 620, 7, 5, 0, 0, 620, 124, 1, 0, 0, 0,
		621, 622, 7, 14, 0, 0, 622, 126, 1, 0, 0, 0, 623, 624, 7, 15, 0, 0, 624,
		128, 1, 0, 0, 0, 625, 626, 7, 16, 0, 0, 626, 130, 1, 0, 0, 0, 627, 628,
		7, 17, 0, 0, 628, 132, 1, 0, 0, 0, 629, 630, 7, 18, 0, 0, 630, 134, 1,
		0, 0, 0, 631, 632, 7, 19, 0, 0, 632, 136, 1, 0, 0, 0, 633, 634, 7, 20,
		0, 0, 634, 138, 1, 0, 0, 0, 635, 636, 7, 21, 0, 0, 636, 140, 1, 0, 0, 0,
		637, 638, 7, 22, 0, 0, 638, 142, 1, 0, 0, 0, 639, 640, 7, 23, 0, 0, 640,
		144, 1, 0, 0, 0, 641, 642, 7, 24, 0, 0, 642, 146, 1, 0, 0, 0, 643, 644,
		7, 25, 0, 0, 644, 148, 1, 0, 0, 0, 645, 646, 7, 26, 0, 0, 646, 150, 1,
		0, 0, 0, 647, 648, 7, 27, 0, 0, 648, 152, 1, 0, 0, 0, 649, 650, 7, 28,
		0, 0, 650, 154, 1, 0, 0, 0, 651, 652, 7, 29, 0, 0, 652, 156, 1, 0, 0, 0,
		653, 654, 7, 30, 0, 0, 654, 158, 1, 0, 0, 0, 655, 656, 7, 31, 0, 0, 656,
		160, 1, 0, 0, 0, 657, 658, 7, 32, 0, 0, 658, 162, 1, 0, 0, 0, 659, 660,
		7, 33, 0, 0, 660, 164, 1, 0, 0, 0, 661, 662, 7, 34, 0, 0, 662, 166, 1,
		0, 0, 0, 663, 667, 5, 35, 0, 0, 664, 666, 9, 0, 0, 0, 665, 664, 1, 0, 0,
		0, 666, 669, 1, 0, 0, 0, 667, 668, 1, 0, 0, 0, 667, 665, 1, 0, 0, 0, 668,
		670, 1, 0, 0, 0, 669, 667, 1, 0, 0, 0, 670, 671, 5, 10, 0, 0, 671, 672,
		1, 0, 0, 0, 672, 673, 6, 83, 0, 0, 673, 168, 1, 0, 0, 0, 22, 0, 372, 385,
		415, 473, 487, 493, 515, 522, 525, 528, 534, 537, 544, 547, 551, 575, 583,
		589, 591, 599, 667, 1, 6, 0, 0,
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
	SLQLexerPROPRIETARY_FUNC_NAME = 22
	SLQLexerJOIN_TYPE             = 23
	SLQLexerWHERE                 = 24
	SLQLexerGROUP_BY              = 25
	SLQLexerORDER_ASC             = 26
	SLQLexerORDER_DESC            = 27
	SLQLexerORDER_BY              = 28
	SLQLexerALIAS_RESERVED        = 29
	SLQLexerARG                   = 30
	SLQLexerNULL                  = 31
	SLQLexerID                    = 32
	SLQLexerWS                    = 33
	SLQLexerLPAR                  = 34
	SLQLexerRPAR                  = 35
	SLQLexerLBRA                  = 36
	SLQLexerRBRA                  = 37
	SLQLexerCOMMA                 = 38
	SLQLexerPIPE                  = 39
	SLQLexerCOLON                 = 40
	SLQLexerNN                    = 41
	SLQLexerNUMBER                = 42
	SLQLexerLT_EQ                 = 43
	SLQLexerLT                    = 44
	SLQLexerGT_EQ                 = 45
	SLQLexerGT                    = 46
	SLQLexerNEQ                   = 47
	SLQLexerEQ                    = 48
	SLQLexerNAME                  = 49
	SLQLexerHANDLE                = 50
	SLQLexerSTRING                = 51
	SLQLexerLINECOMMENT           = 52
)
