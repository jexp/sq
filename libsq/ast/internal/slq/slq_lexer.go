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
		"", "';'", "'*'", "'sum'", "'avg'", "'max'", "'min'", "'unique'", "'count'",
		"'.['", "'||'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'&&'", "'~'",
		"'!'", "", "", "", "'group_by'", "'+'", "'-'", "", "", "", "'null'",
		"", "", "'('", "')'", "'['", "']'", "','", "'|'", "':'", "", "", "'<='",
		"'<'", "'>='", "'>'", "'!='", "'=='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "PROPRIETARY_FUNC_NAME", "JOIN_TYPE", "WHERE", "GROUP_BY", "ORDER_ASC",
		"ORDER_DESC", "ORDER_BY", "ALIAS_RESERVED", "ARG", "NULL", "ID", "WS",
		"LPAR", "RPAR", "LBRA", "RBRA", "COMMA", "PIPE", "COLON", "NN", "NUMBER",
		"LT_EQ", "LT", "GT_EQ", "GT", "NEQ", "EQ", "NAME", "HANDLE", "STRING",
		"LINECOMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "PROPRIETARY_FUNC_NAME", "JOIN_TYPE", "WHERE", "GROUP_BY",
		"ORDER_ASC", "ORDER_DESC", "ORDER_BY", "ALIAS_RESERVED", "ARG", "NULL",
		"ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA", "COMMA", "PIPE", "COLON",
		"NN", "NUMBER", "INTF", "EXP", "LT_EQ", "LT", "GT_EQ", "GT", "NEQ",
		"EQ", "NAME", "HANDLE", "STRING", "ESC", "UNICODE", "HEX", "DIGIT",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "LINECOMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 49, 648, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 347, 8, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 360,
		8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 390, 8,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 448, 8, 25, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 5, 28, 460, 8, 28, 10, 28,
		12, 28, 463, 9, 28, 1, 29, 4, 29, 466, 8, 29, 11, 29, 12, 29, 467, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 3, 38, 490,
		8, 38, 1, 38, 1, 38, 1, 38, 4, 38, 495, 8, 38, 11, 38, 12, 38, 496, 1,
		38, 3, 38, 500, 8, 38, 1, 38, 3, 38, 503, 8, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 3, 38, 509, 8, 38, 1, 38, 3, 38, 512, 8, 38, 1, 39, 1, 39, 1, 39, 5,
		39, 517, 8, 39, 10, 39, 12, 39, 520, 9, 39, 3, 39, 522, 8, 39, 1, 40, 1,
		40, 3, 40, 526, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 550, 8, 47, 1, 48, 1, 48, 1, 48,
		1, 48, 5, 48, 556, 8, 48, 10, 48, 12, 48, 559, 9, 48, 1, 49, 1, 49, 1,
		49, 5, 49, 564, 8, 49, 10, 49, 12, 49, 567, 9, 49, 1, 49, 1, 49, 1, 50,
		1, 50, 1, 50, 3, 50, 574, 8, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56,
		1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1,
		62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67,
		1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 1, 70, 1, 71, 1, 71, 1, 72, 1,
		72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77,
		1, 78, 1, 78, 1, 79, 1, 79, 1, 80, 1, 80, 5, 80, 640, 8, 80, 10, 80, 12,
		80, 643, 9, 80, 1, 80, 1, 80, 1, 80, 1, 80, 1, 641, 0, 81, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		0, 81, 0, 83, 40, 85, 41, 87, 42, 89, 43, 91, 44, 93, 45, 95, 46, 97, 47,
		99, 48, 101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115, 0,
		117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131, 0, 133, 0,
		135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149, 0, 151, 0,
		153, 0, 155, 0, 157, 0, 159, 0, 161, 49, 1, 0, 35, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45,
		45, 2, 0, 34, 34, 92, 92, 8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102,
		110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 65,
		65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100,
		100, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104,
		104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107,
		107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110,
		110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113,
		113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116,
		116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119,
		119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122,
		122, 657, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 161, 1, 0, 0, 0, 1, 163,
		1, 0, 0, 0, 3, 165, 1, 0, 0, 0, 5, 167, 1, 0, 0, 0, 7, 171, 1, 0, 0, 0,
		9, 175, 1, 0, 0, 0, 11, 179, 1, 0, 0, 0, 13, 183, 1, 0, 0, 0, 15, 190,
		1, 0, 0, 0, 17, 196, 1, 0, 0, 0, 19, 199, 1, 0, 0, 0, 21, 202, 1, 0, 0,
		0, 23, 204, 1, 0, 0, 0, 25, 206, 1, 0, 0, 0, 27, 209, 1, 0, 0, 0, 29, 212,
		1, 0, 0, 0, 31, 214, 1, 0, 0, 0, 33, 217, 1, 0, 0, 0, 35, 219, 1, 0, 0,
		0, 37, 221, 1, 0, 0, 0, 39, 346, 1, 0, 0, 0, 41, 359, 1, 0, 0, 0, 43, 361,
		1, 0, 0, 0, 45, 370, 1, 0, 0, 0, 47, 372, 1, 0, 0, 0, 49, 389, 1, 0, 0,
		0, 51, 447, 1, 0, 0, 0, 53, 449, 1, 0, 0, 0, 55, 452, 1, 0, 0, 0, 57, 457,
		1, 0, 0, 0, 59, 465, 1, 0, 0, 0, 61, 471, 1, 0, 0, 0, 63, 473, 1, 0, 0,
		0, 65, 475, 1, 0, 0, 0, 67, 477, 1, 0, 0, 0, 69, 479, 1, 0, 0, 0, 71, 481,
		1, 0, 0, 0, 73, 483, 1, 0, 0, 0, 75, 485, 1, 0, 0, 0, 77, 511, 1, 0, 0,
		0, 79, 521, 1, 0, 0, 0, 81, 523, 1, 0, 0, 0, 83, 529, 1, 0, 0, 0, 85, 532,
		1, 0, 0, 0, 87, 534, 1, 0, 0, 0, 89, 537, 1, 0, 0, 0, 91, 539, 1, 0, 0,
		0, 93, 542, 1, 0, 0, 0, 95, 545, 1, 0, 0, 0, 97, 551, 1, 0, 0, 0, 99, 560,
		1, 0, 0, 0, 101, 570, 1, 0, 0, 0, 103, 575, 1, 0, 0, 0, 105, 581, 1, 0,
		0, 0, 107, 583, 1, 0, 0, 0, 109, 585, 1, 0, 0, 0, 111, 587, 1, 0, 0, 0,
		113, 589, 1, 0, 0, 0, 115, 591, 1, 0, 0, 0, 117, 593, 1, 0, 0, 0, 119,
		595, 1, 0, 0, 0, 121, 597, 1, 0, 0, 0, 123, 599, 1, 0, 0, 0, 125, 601,
		1, 0, 0, 0, 127, 603, 1, 0, 0, 0, 129, 605, 1, 0, 0, 0, 131, 607, 1, 0,
		0, 0, 133, 609, 1, 0, 0, 0, 135, 611, 1, 0, 0, 0, 137, 613, 1, 0, 0, 0,
		139, 615, 1, 0, 0, 0, 141, 617, 1, 0, 0, 0, 143, 619, 1, 0, 0, 0, 145,
		621, 1, 0, 0, 0, 147, 623, 1, 0, 0, 0, 149, 625, 1, 0, 0, 0, 151, 627,
		1, 0, 0, 0, 153, 629, 1, 0, 0, 0, 155, 631, 1, 0, 0, 0, 157, 633, 1, 0,
		0, 0, 159, 635, 1, 0, 0, 0, 161, 637, 1, 0, 0, 0, 163, 164, 5, 59, 0, 0,
		164, 2, 1, 0, 0, 0, 165, 166, 5, 42, 0, 0, 166, 4, 1, 0, 0, 0, 167, 168,
		5, 115, 0, 0, 168, 169, 5, 117, 0, 0, 169, 170, 5, 109, 0, 0, 170, 6, 1,
		0, 0, 0, 171, 172, 5, 97, 0, 0, 172, 173, 5, 118, 0, 0, 173, 174, 5, 103,
		0, 0, 174, 8, 1, 0, 0, 0, 175, 176, 5, 109, 0, 0, 176, 177, 5, 97, 0, 0,
		177, 178, 5, 120, 0, 0, 178, 10, 1, 0, 0, 0, 179, 180, 5, 109, 0, 0, 180,
		181, 5, 105, 0, 0, 181, 182, 5, 110, 0, 0, 182, 12, 1, 0, 0, 0, 183, 184,
		5, 117, 0, 0, 184, 185, 5, 110, 0, 0, 185, 186, 5, 105, 0, 0, 186, 187,
		5, 113, 0, 0, 187, 188, 5, 117, 0, 0, 188, 189, 5, 101, 0, 0, 189, 14,
		1, 0, 0, 0, 190, 191, 5, 99, 0, 0, 191, 192, 5, 111, 0, 0, 192, 193, 5,
		117, 0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 116, 0, 0, 195, 16, 1,
		0, 0, 0, 196, 197, 5, 46, 0, 0, 197, 198, 5, 91, 0, 0, 198, 18, 1, 0, 0,
		0, 199, 200, 5, 124, 0, 0, 200, 201, 5, 124, 0, 0, 201, 20, 1, 0, 0, 0,
		202, 203, 5, 47, 0, 0, 203, 22, 1, 0, 0, 0, 204, 205, 5, 37, 0, 0, 205,
		24, 1, 0, 0, 0, 206, 207, 5, 60, 0, 0, 207, 208, 5, 60, 0, 0, 208, 26,
		1, 0, 0, 0, 209, 210, 5, 62, 0, 0, 210, 211, 5, 62, 0, 0, 211, 28, 1, 0,
		0, 0, 212, 213, 5, 38, 0, 0, 213, 30, 1, 0, 0, 0, 214, 215, 5, 38, 0, 0,
		215, 216, 5, 38, 0, 0, 216, 32, 1, 0, 0, 0, 217, 218, 5, 126, 0, 0, 218,
		34, 1, 0, 0, 0, 219, 220, 5, 33, 0, 0, 220, 36, 1, 0, 0, 0, 221, 222, 5,
		95, 0, 0, 222, 223, 3, 57, 28, 0, 223, 38, 1, 0, 0, 0, 224, 225, 5, 106,
		0, 0, 225, 226, 5, 111, 0, 0, 226, 227, 5, 105, 0, 0, 227, 347, 5, 110,
		0, 0, 228, 229, 5, 105, 0, 0, 229, 230, 5, 110, 0, 0, 230, 231, 5, 110,
		0, 0, 231, 232, 5, 101, 0, 0, 232, 233, 5, 114, 0, 0, 233, 234, 5, 95,
		0, 0, 234, 235, 5, 106, 0, 0, 235, 236, 5, 111, 0, 0, 236, 237, 5, 105,
		0, 0, 237, 347, 5, 110, 0, 0, 238, 239, 5, 108, 0, 0, 239, 240, 5, 101,
		0, 0, 240, 241, 5, 102, 0, 0, 241, 242, 5, 116, 0, 0, 242, 243, 5, 95,
		0, 0, 243, 244, 5, 106, 0, 0, 244, 245, 5, 111, 0, 0, 245, 246, 5, 105,
		0, 0, 246, 347, 5, 110, 0, 0, 247, 248, 5, 108, 0, 0, 248, 249, 5, 106,
		0, 0, 249, 250, 5, 111, 0, 0, 250, 251, 5, 105, 0, 0, 251, 347, 5, 110,
		0, 0, 252, 253, 5, 108, 0, 0, 253, 254, 5, 101, 0, 0, 254, 255, 5, 102,
		0, 0, 255, 256, 5, 116, 0, 0, 256, 257, 5, 95, 0, 0, 257, 258, 5, 111,
		0, 0, 258, 259, 5, 117, 0, 0, 259, 260, 5, 116, 0, 0, 260, 261, 5, 101,
		0, 0, 261, 262, 5, 114, 0, 0, 262, 263, 5, 95, 0, 0, 263, 264, 5, 106,
		0, 0, 264, 265, 5, 111, 0, 0, 265, 266, 5, 105, 0, 0, 266, 347, 5, 110,
		0, 0, 267, 268, 5, 108, 0, 0, 268, 269, 5, 111, 0, 0, 269, 270, 5, 106,
		0, 0, 270, 271, 5, 111, 0, 0, 271, 272, 5, 105, 0, 0, 272, 347, 5, 110,
		0, 0, 273, 274, 5, 114, 0, 0, 274, 275, 5, 105, 0, 0, 275, 276, 5, 103,
		0, 0, 276, 277, 5, 104, 0, 0, 277, 278, 5, 116, 0, 0, 278, 279, 5, 95,
		0, 0, 279, 280, 5, 106, 0, 0, 280, 281, 5, 111, 0, 0, 281, 282, 5, 105,
		0, 0, 282, 347, 5, 110, 0, 0, 283, 284, 5, 114, 0, 0, 284, 285, 5, 106,
		0, 0, 285, 286, 5, 111, 0, 0, 286, 287, 5, 105, 0, 0, 287, 347, 5, 110,
		0, 0, 288, 289, 5, 114, 0, 0, 289, 290, 5, 105, 0, 0, 290, 291, 5, 103,
		0, 0, 291, 292, 5, 104, 0, 0, 292, 293, 5, 116, 0, 0, 293, 294, 5, 95,
		0, 0, 294, 295, 5, 111, 0, 0, 295, 296, 5, 117, 0, 0, 296, 297, 5, 116,
		0, 0, 297, 298, 5, 101, 0, 0, 298, 299, 5, 114, 0, 0, 299, 300, 5, 95,
		0, 0, 300, 301, 5, 106, 0, 0, 301, 302, 5, 111, 0, 0, 302, 303, 5, 105,
		0, 0, 303, 347, 5, 110, 0, 0, 304, 305, 5, 114, 0, 0, 305, 306, 5, 111,
		0, 0, 306, 307, 5, 106, 0, 0, 307, 308, 5, 111, 0, 0, 308, 309, 5, 105,
		0, 0, 309, 347, 5, 110, 0, 0, 310, 311, 5, 102, 0, 0, 311, 312, 5, 117,
		0, 0, 312, 313, 5, 108, 0, 0, 313, 314, 5, 108, 0, 0, 314, 315, 5, 95,
		0, 0, 315, 316, 5, 111, 0, 0, 316, 317, 5, 117, 0, 0, 317, 318, 5, 116,
		0, 0, 318, 319, 5, 101, 0, 0, 319, 320, 5, 114, 0, 0, 320, 321, 5, 95,
		0, 0, 321, 322, 5, 106, 0, 0, 322, 323, 5, 111, 0, 0, 323, 324, 5, 105,
		0, 0, 324, 347, 5, 110, 0, 0, 325, 326, 5, 102, 0, 0, 326, 327, 5, 111,
		0, 0, 327, 328, 5, 106, 0, 0, 328, 329, 5, 111, 0, 0, 329, 330, 5, 105,
		0, 0, 330, 347, 5, 110, 0, 0, 331, 332, 5, 99, 0, 0, 332, 333, 5, 114,
		0, 0, 333, 334, 5, 111, 0, 0, 334, 335, 5, 115, 0, 0, 335, 336, 5, 115,
		0, 0, 336, 337, 5, 95, 0, 0, 337, 338, 5, 106, 0, 0, 338, 339, 5, 111,
		0, 0, 339, 340, 5, 105, 0, 0, 340, 347, 5, 110, 0, 0, 341, 342, 5, 120,
		0, 0, 342, 343, 5, 106, 0, 0, 343, 344, 5, 111, 0, 0, 344, 345, 5, 105,
		0, 0, 345, 347, 5, 110, 0, 0, 346, 224, 1, 0, 0, 0, 346, 228, 1, 0, 0,
		0, 346, 238, 1, 0, 0, 0, 346, 247, 1, 0, 0, 0, 346, 252, 1, 0, 0, 0, 346,
		267, 1, 0, 0, 0, 346, 273, 1, 0, 0, 0, 346, 283, 1, 0, 0, 0, 346, 288,
		1, 0, 0, 0, 346, 304, 1, 0, 0, 0, 346, 310, 1, 0, 0, 0, 346, 325, 1, 0,
		0, 0, 346, 331, 1, 0, 0, 0, 346, 341, 1, 0, 0, 0, 347, 40, 1, 0, 0, 0,
		348, 349, 5, 119, 0, 0, 349, 350, 5, 104, 0, 0, 350, 351, 5, 101, 0, 0,
		351, 352, 5, 114, 0, 0, 352, 360, 5, 101, 0, 0, 353, 354, 5, 115, 0, 0,
		354, 355, 5, 101, 0, 0, 355, 356, 5, 108, 0, 0, 356, 357, 5, 101, 0, 0,
		357, 358, 5, 99, 0, 0, 358, 360, 5, 116, 0, 0, 359, 348, 1, 0, 0, 0, 359,
		353, 1, 0, 0, 0, 360, 42, 1, 0, 0, 0, 361, 362, 5, 103, 0, 0, 362, 363,
		5, 114, 0, 0, 363, 364, 5, 111, 0, 0, 364, 365, 5, 117, 0, 0, 365, 366,
		5, 112, 0, 0, 366, 367, 5, 95, 0, 0, 367, 368, 5, 98, 0, 0, 368, 369, 5,
		121, 0, 0, 369, 44, 1, 0, 0, 0, 370, 371, 5, 43, 0, 0, 371, 46, 1, 0, 0,
		0, 372, 373, 5, 45, 0, 0, 373, 48, 1, 0, 0, 0, 374, 375, 5, 111, 0, 0,
		375, 376, 5, 114, 0, 0, 376, 377, 5, 100, 0, 0, 377, 378, 5, 101, 0, 0,
		378, 379, 5, 114, 0, 0, 379, 380, 5, 95, 0, 0, 380, 381, 5, 98, 0, 0, 381,
		390, 5, 121, 0, 0, 382, 383, 5, 115, 0, 0, 383, 384, 5, 111, 0, 0, 384,
		385, 5, 114, 0, 0, 385, 386, 5, 116, 0, 0, 386, 387, 5, 95, 0, 0, 387,
		388, 5, 98, 0, 0, 388, 390, 5, 121, 0, 0, 389, 374, 1, 0, 0, 0, 389, 382,
		1, 0, 0, 0, 390, 50, 1, 0, 0, 0, 391, 392, 5, 58, 0, 0, 392, 393, 5, 99,
		0, 0, 393, 394, 5, 111, 0, 0, 394, 395, 5, 117, 0, 0, 395, 396, 5, 110,
		0, 0, 396, 448, 5, 116, 0, 0, 397, 398, 5, 58, 0, 0, 398, 399, 5, 99, 0,
		0, 399, 400, 5, 111, 0, 0, 400, 401, 5, 117, 0, 0, 401, 402, 5, 110, 0,
		0, 402, 403, 5, 116, 0, 0, 403, 404, 5, 95, 0, 0, 404, 405, 5, 117, 0,
		0, 405, 406, 5, 110, 0, 0, 406, 407, 5, 105, 0, 0, 407, 408, 5, 113, 0,
		0, 408, 409, 5, 117, 0, 0, 409, 448, 5, 101, 0, 0, 410, 411, 5, 58, 0,
		0, 411, 412, 5, 97, 0, 0, 412, 413, 5, 118, 0, 0, 413, 448, 5, 103, 0,
		0, 414, 415, 5, 58, 0, 0, 415, 416, 5, 103, 0, 0, 416, 417, 5, 114, 0,
		0, 417, 418, 5, 111, 0, 0, 418, 419, 5, 117, 0, 0, 419, 420, 5, 112, 0,
		0, 420, 421, 5, 95, 0, 0, 421, 422, 5, 98, 0, 0, 422, 448, 5, 121, 0, 0,
		423, 424, 5, 58, 0, 0, 424, 425, 5, 109, 0, 0, 425, 426, 5, 97, 0, 0, 426,
		448, 5, 120, 0, 0, 427, 428, 5, 58, 0, 0, 428, 429, 5, 109, 0, 0, 429,
		430, 5, 105, 0, 0, 430, 448, 5, 110, 0, 0, 431, 432, 5, 58, 0, 0, 432,
		433, 5, 111, 0, 0, 433, 434, 5, 114, 0, 0, 434, 435, 5, 100, 0, 0, 435,
		436, 5, 101, 0, 0, 436, 437, 5, 114, 0, 0, 437, 438, 5, 95, 0, 0, 438,
		439, 5, 98, 0, 0, 439, 448, 5, 121, 0, 0, 440, 441, 5, 58, 0, 0, 441, 442,
		5, 117, 0, 0, 442, 443, 5, 110, 0, 0, 443, 444, 5, 105, 0, 0, 444, 445,
		5, 113, 0, 0, 445, 446, 5, 117, 0, 0, 446, 448, 5, 101, 0, 0, 447, 391,
		1, 0, 0, 0, 447, 397, 1, 0, 0, 0, 447, 410, 1, 0, 0, 0, 447, 414, 1, 0,
		0, 0, 447, 423, 1, 0, 0, 0, 447, 427, 1, 0, 0, 0, 447, 431, 1, 0, 0, 0,
		447, 440, 1, 0, 0, 0, 448, 52, 1, 0, 0, 0, 449, 450, 5, 36, 0, 0, 450,
		451, 3, 57, 28, 0, 451, 54, 1, 0, 0, 0, 452, 453, 5, 110, 0, 0, 453, 454,
		5, 117, 0, 0, 454, 455, 5, 108, 0, 0, 455, 456, 5, 108, 0, 0, 456, 56,
		1, 0, 0, 0, 457, 461, 7, 0, 0, 0, 458, 460, 7, 1, 0, 0, 459, 458, 1, 0,
		0, 0, 460, 463, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0,
		462, 58, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 464, 466, 7, 2, 0, 0, 465, 464,
		1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 467, 468, 1, 0,
		0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 6, 29, 0, 0, 470, 60, 1, 0, 0, 0,
		471, 472, 5, 40, 0, 0, 472, 62, 1, 0, 0, 0, 473, 474, 5, 41, 0, 0, 474,
		64, 1, 0, 0, 0, 475, 476, 5, 91, 0, 0, 476, 66, 1, 0, 0, 0, 477, 478, 5,
		93, 0, 0, 478, 68, 1, 0, 0, 0, 479, 480, 5, 44, 0, 0, 480, 70, 1, 0, 0,
		0, 481, 482, 5, 124, 0, 0, 482, 72, 1, 0, 0, 0, 483, 484, 5, 58, 0, 0,
		484, 74, 1, 0, 0, 0, 485, 486, 3, 79, 39, 0, 486, 76, 1, 0, 0, 0, 487,
		512, 3, 75, 37, 0, 488, 490, 5, 45, 0, 0, 489, 488, 1, 0, 0, 0, 489, 490,
		1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492, 3, 79, 39, 0, 492, 494, 5,
		46, 0, 0, 493, 495, 7, 3, 0, 0, 494, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0,
		0, 496, 494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 499, 1, 0, 0, 0, 498,
		500, 3, 81, 40, 0, 499, 498, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 512,
		1, 0, 0, 0, 501, 503, 5, 45, 0, 0, 502, 501, 1, 0, 0, 0, 502, 503, 1, 0,
		0, 0, 503, 504, 1, 0, 0, 0, 504, 505, 3, 79, 39, 0, 505, 506, 3, 81, 40,
		0, 506, 512, 1, 0, 0, 0, 507, 509, 5, 45, 0, 0, 508, 507, 1, 0, 0, 0, 508,
		509, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 512, 3, 79, 39, 0, 511, 487,
		1, 0, 0, 0, 511, 489, 1, 0, 0, 0, 511, 502, 1, 0, 0, 0, 511, 508, 1, 0,
		0, 0, 512, 78, 1, 0, 0, 0, 513, 522, 5, 48, 0, 0, 514, 518, 7, 4, 0, 0,
		515, 517, 7, 3, 0, 0, 516, 515, 1, 0, 0, 0, 517, 520, 1, 0, 0, 0, 518,
		516, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 522, 1, 0, 0, 0, 520, 518,
		1, 0, 0, 0, 521, 513, 1, 0, 0, 0, 521, 514, 1, 0, 0, 0, 522, 80, 1, 0,
		0, 0, 523, 525, 7, 5, 0, 0, 524, 526, 7, 6, 0, 0, 525, 524, 1, 0, 0, 0,
		525, 526, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 528, 3, 79, 39, 0, 528,
		82, 1, 0, 0, 0, 529, 530, 5, 60, 0, 0, 530, 531, 5, 61, 0, 0, 531, 84,
		1, 0, 0, 0, 532, 533, 5, 60, 0, 0, 533, 86, 1, 0, 0, 0, 534, 535, 5, 62,
		0, 0, 535, 536, 5, 61, 0, 0, 536, 88, 1, 0, 0, 0, 537, 538, 5, 62, 0, 0,
		538, 90, 1, 0, 0, 0, 539, 540, 5, 33, 0, 0, 540, 541, 5, 61, 0, 0, 541,
		92, 1, 0, 0, 0, 542, 543, 5, 61, 0, 0, 543, 544, 5, 61, 0, 0, 544, 94,
		1, 0, 0, 0, 545, 549, 5, 46, 0, 0, 546, 550, 3, 53, 26, 0, 547, 550, 3,
		57, 28, 0, 548, 550, 3, 99, 49, 0, 549, 546, 1, 0, 0, 0, 549, 547, 1, 0,
		0, 0, 549, 548, 1, 0, 0, 0, 550, 96, 1, 0, 0, 0, 551, 552, 5, 64, 0, 0,
		552, 557, 3, 57, 28, 0, 553, 554, 5, 47, 0, 0, 554, 556, 3, 57, 28, 0,
		555, 553, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 557,
		558, 1, 0, 0, 0, 558, 98, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 560, 565, 5,
		34, 0, 0, 561, 564, 3, 101, 50, 0, 562, 564, 8, 7, 0, 0, 563, 561, 1, 0,
		0, 0, 563, 562, 1, 0, 0, 0, 564, 567, 1, 0, 0, 0, 565, 563, 1, 0, 0, 0,
		565, 566, 1, 0, 0, 0, 566, 568, 1, 0, 0, 0, 567, 565, 1, 0, 0, 0, 568,
		569, 5, 34, 0, 0, 569, 100, 1, 0, 0, 0, 570, 573, 5, 92, 0, 0, 571, 574,
		7, 8, 0, 0, 572, 574, 3, 103, 51, 0, 573, 571, 1, 0, 0, 0, 573, 572, 1,
		0, 0, 0, 574, 102, 1, 0, 0, 0, 575, 576, 5, 117, 0, 0, 576, 577, 3, 105,
		52, 0, 577, 578, 3, 105, 52, 0, 578, 579, 3, 105, 52, 0, 579, 580, 3, 105,
		52, 0, 580, 104, 1, 0, 0, 0, 581, 582, 7, 9, 0, 0, 582, 106, 1, 0, 0, 0,
		583, 584, 7, 3, 0, 0, 584, 108, 1, 0, 0, 0, 585, 586, 7, 10, 0, 0, 586,
		110, 1, 0, 0, 0, 587, 588, 7, 11, 0, 0, 588, 112, 1, 0, 0, 0, 589, 590,
		7, 12, 0, 0, 590, 114, 1, 0, 0, 0, 591, 592, 7, 13, 0, 0, 592, 116, 1,
		0, 0, 0, 593, 594, 7, 5, 0, 0, 594, 118, 1, 0, 0, 0, 595, 596, 7, 14, 0,
		0, 596, 120, 1, 0, 0, 0, 597, 598, 7, 15, 0, 0, 598, 122, 1, 0, 0, 0, 599,
		600, 7, 16, 0, 0, 600, 124, 1, 0, 0, 0, 601, 602, 7, 17, 0, 0, 602, 126,
		1, 0, 0, 0, 603, 604, 7, 18, 0, 0, 604, 128, 1, 0, 0, 0, 605, 606, 7, 19,
		0, 0, 606, 130, 1, 0, 0, 0, 607, 608, 7, 20, 0, 0, 608, 132, 1, 0, 0, 0,
		609, 610, 7, 21, 0, 0, 610, 134, 1, 0, 0, 0, 611, 612, 7, 22, 0, 0, 612,
		136, 1, 0, 0, 0, 613, 614, 7, 23, 0, 0, 614, 138, 1, 0, 0, 0, 615, 616,
		7, 24, 0, 0, 616, 140, 1, 0, 0, 0, 617, 618, 7, 25, 0, 0, 618, 142, 1,
		0, 0, 0, 619, 620, 7, 26, 0, 0, 620, 144, 1, 0, 0, 0, 621, 622, 7, 27,
		0, 0, 622, 146, 1, 0, 0, 0, 623, 624, 7, 28, 0, 0, 624, 148, 1, 0, 0, 0,
		625, 626, 7, 29, 0, 0, 626, 150, 1, 0, 0, 0, 627, 628, 7, 30, 0, 0, 628,
		152, 1, 0, 0, 0, 629, 630, 7, 31, 0, 0, 630, 154, 1, 0, 0, 0, 631, 632,
		7, 32, 0, 0, 632, 156, 1, 0, 0, 0, 633, 634, 7, 33, 0, 0, 634, 158, 1,
		0, 0, 0, 635, 636, 7, 34, 0, 0, 636, 160, 1, 0, 0, 0, 637, 641, 5, 35,
		0, 0, 638, 640, 9, 0, 0, 0, 639, 638, 1, 0, 0, 0, 640, 643, 1, 0, 0, 0,
		641, 642, 1, 0, 0, 0, 641, 639, 1, 0, 0, 0, 642, 644, 1, 0, 0, 0, 643,
		641, 1, 0, 0, 0, 644, 645, 5, 10, 0, 0, 645, 646, 1, 0, 0, 0, 646, 647,
		6, 80, 0, 0, 647, 162, 1, 0, 0, 0, 22, 0, 346, 359, 389, 447, 461, 467,
		489, 496, 499, 502, 508, 511, 518, 521, 525, 549, 557, 563, 565, 573, 641,
		1, 6, 0, 0,
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
	SLQLexerPROPRIETARY_FUNC_NAME = 19
	SLQLexerJOIN_TYPE             = 20
	SLQLexerWHERE                 = 21
	SLQLexerGROUP_BY              = 22
	SLQLexerORDER_ASC             = 23
	SLQLexerORDER_DESC            = 24
	SLQLexerORDER_BY              = 25
	SLQLexerALIAS_RESERVED        = 26
	SLQLexerARG                   = 27
	SLQLexerNULL                  = 28
	SLQLexerID                    = 29
	SLQLexerWS                    = 30
	SLQLexerLPAR                  = 31
	SLQLexerRPAR                  = 32
	SLQLexerLBRA                  = 33
	SLQLexerRBRA                  = 34
	SLQLexerCOMMA                 = 35
	SLQLexerPIPE                  = 36
	SLQLexerCOLON                 = 37
	SLQLexerNN                    = 38
	SLQLexerNUMBER                = 39
	SLQLexerLT_EQ                 = 40
	SLQLexerLT                    = 41
	SLQLexerGT_EQ                 = 42
	SLQLexerGT                    = 43
	SLQLexerNEQ                   = 44
	SLQLexerEQ                    = 45
	SLQLexerNAME                  = 46
	SLQLexerHANDLE                = 47
	SLQLexerSTRING                = 48
	SLQLexerLINECOMMENT           = 49
)
