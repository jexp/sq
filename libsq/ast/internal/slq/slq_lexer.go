// Code generated from SLQ.g4 by ANTLR 4.12.0. DO NOT EDIT.

package slq

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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

var slqlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func slqlexerLexerInit() {
	staticData := &slqlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "'*'", "'join'", "'unique'", "'count'", "'.['", "'||'", "'/'",
		"'%'", "'<<'", "'>>'", "'&'", "'&&'", "'~'", "'!'", "'group_by'", "'+'",
		"'-'", "", "", "", "'null'", "", "", "'('", "')'", "'['", "']'", "','",
		"'|'", "':'", "", "", "'<='", "'<'", "'>='", "'>'", "'!='", "'=='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "GROUP_BY",
		"ORDER_ASC", "ORDER_DESC", "ORDER_BY", "ALIAS_RESERVED", "ARG", "NULL",
		"ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA", "COMMA", "PIPE", "COLON",
		"NN", "NUMBER", "LT_EQ", "LT", "GT_EQ", "GT", "NEQ", "EQ", "NAME", "HANDLE",
		"STRING", "LINECOMMENT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "GROUP_BY", "ORDER_ASC",
		"ORDER_DESC", "ORDER_BY", "ALIAS_RESERVED", "ARG", "NULL", "ID", "WS",
		"LPAR", "RPAR", "LBRA", "RBRA", "COMMA", "PIPE", "COLON", "NN", "NUMBER",
		"INTF", "EXP", "LT_EQ", "LT", "GT_EQ", "GT", "NEQ", "EQ", "NAME", "HANDLE",
		"STRING", "ESC", "UNICODE", "HEX", "DIGIT", "A", "B", "C", "D", "E",
		"F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z", "LINECOMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 485, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		73, 7, 73, 2, 74, 7, 74, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 227, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 285, 8, 19, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22,
		297, 8, 22, 10, 22, 12, 22, 300, 9, 22, 1, 23, 4, 23, 303, 8, 23, 11, 23,
		12, 23, 304, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 3, 32, 327, 8, 32, 1, 32, 1, 32, 1, 32, 4, 32, 332, 8, 32, 11, 32,
		12, 32, 333, 1, 32, 3, 32, 337, 8, 32, 1, 32, 3, 32, 340, 8, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 3, 32, 346, 8, 32, 1, 32, 3, 32, 349, 8, 32, 1, 33,
		1, 33, 1, 33, 5, 33, 354, 8, 33, 10, 33, 12, 33, 357, 9, 33, 3, 33, 359,
		8, 33, 1, 34, 1, 34, 3, 34, 363, 8, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 387, 8, 41, 1,
		42, 1, 42, 1, 42, 1, 42, 5, 42, 393, 8, 42, 10, 42, 12, 42, 396, 9, 42,
		1, 43, 1, 43, 1, 43, 5, 43, 401, 8, 43, 10, 43, 12, 43, 404, 9, 43, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 44, 3, 44, 411, 8, 44, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65,
		1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 1,
		70, 1, 71, 1, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 5, 74, 477,
		8, 74, 10, 74, 12, 74, 480, 9, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 478,
		0, 75, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 0, 69, 0, 71, 34, 73, 35, 75,
		36, 77, 37, 79, 38, 81, 39, 83, 40, 85, 41, 87, 42, 89, 0, 91, 0, 93, 0,
		95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113,
		0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131,
		0, 133, 0, 135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149,
		43, 1, 0, 35, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 1, 0, 49, 57, 2,
		0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 34, 34, 92, 92, 8, 0,
		34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116,
		3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98,
		98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100, 2, 0, 70, 70, 102, 102,
		2, 0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105,
		2, 0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108,
		2, 0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111,
		2, 0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114,
		2, 0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117,
		2, 0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120,
		2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 480, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0,
		0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1,
		0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 1, 151,
		1, 0, 0, 0, 3, 153, 1, 0, 0, 0, 5, 155, 1, 0, 0, 0, 7, 160, 1, 0, 0, 0,
		9, 167, 1, 0, 0, 0, 11, 173, 1, 0, 0, 0, 13, 176, 1, 0, 0, 0, 15, 179,
		1, 0, 0, 0, 17, 181, 1, 0, 0, 0, 19, 183, 1, 0, 0, 0, 21, 186, 1, 0, 0,
		0, 23, 189, 1, 0, 0, 0, 25, 191, 1, 0, 0, 0, 27, 194, 1, 0, 0, 0, 29, 196,
		1, 0, 0, 0, 31, 198, 1, 0, 0, 0, 33, 207, 1, 0, 0, 0, 35, 209, 1, 0, 0,
		0, 37, 226, 1, 0, 0, 0, 39, 284, 1, 0, 0, 0, 41, 286, 1, 0, 0, 0, 43, 289,
		1, 0, 0, 0, 45, 294, 1, 0, 0, 0, 47, 302, 1, 0, 0, 0, 49, 308, 1, 0, 0,
		0, 51, 310, 1, 0, 0, 0, 53, 312, 1, 0, 0, 0, 55, 314, 1, 0, 0, 0, 57, 316,
		1, 0, 0, 0, 59, 318, 1, 0, 0, 0, 61, 320, 1, 0, 0, 0, 63, 322, 1, 0, 0,
		0, 65, 348, 1, 0, 0, 0, 67, 358, 1, 0, 0, 0, 69, 360, 1, 0, 0, 0, 71, 366,
		1, 0, 0, 0, 73, 369, 1, 0, 0, 0, 75, 371, 1, 0, 0, 0, 77, 374, 1, 0, 0,
		0, 79, 376, 1, 0, 0, 0, 81, 379, 1, 0, 0, 0, 83, 382, 1, 0, 0, 0, 85, 388,
		1, 0, 0, 0, 87, 397, 1, 0, 0, 0, 89, 407, 1, 0, 0, 0, 91, 412, 1, 0, 0,
		0, 93, 418, 1, 0, 0, 0, 95, 420, 1, 0, 0, 0, 97, 422, 1, 0, 0, 0, 99, 424,
		1, 0, 0, 0, 101, 426, 1, 0, 0, 0, 103, 428, 1, 0, 0, 0, 105, 430, 1, 0,
		0, 0, 107, 432, 1, 0, 0, 0, 109, 434, 1, 0, 0, 0, 111, 436, 1, 0, 0, 0,
		113, 438, 1, 0, 0, 0, 115, 440, 1, 0, 0, 0, 117, 442, 1, 0, 0, 0, 119,
		444, 1, 0, 0, 0, 121, 446, 1, 0, 0, 0, 123, 448, 1, 0, 0, 0, 125, 450,
		1, 0, 0, 0, 127, 452, 1, 0, 0, 0, 129, 454, 1, 0, 0, 0, 131, 456, 1, 0,
		0, 0, 133, 458, 1, 0, 0, 0, 135, 460, 1, 0, 0, 0, 137, 462, 1, 0, 0, 0,
		139, 464, 1, 0, 0, 0, 141, 466, 1, 0, 0, 0, 143, 468, 1, 0, 0, 0, 145,
		470, 1, 0, 0, 0, 147, 472, 1, 0, 0, 0, 149, 474, 1, 0, 0, 0, 151, 152,
		5, 59, 0, 0, 152, 2, 1, 0, 0, 0, 153, 154, 5, 42, 0, 0, 154, 4, 1, 0, 0,
		0, 155, 156, 5, 106, 0, 0, 156, 157, 5, 111, 0, 0, 157, 158, 5, 105, 0,
		0, 158, 159, 5, 110, 0, 0, 159, 6, 1, 0, 0, 0, 160, 161, 5, 117, 0, 0,
		161, 162, 5, 110, 0, 0, 162, 163, 5, 105, 0, 0, 163, 164, 5, 113, 0, 0,
		164, 165, 5, 117, 0, 0, 165, 166, 5, 101, 0, 0, 166, 8, 1, 0, 0, 0, 167,
		168, 5, 99, 0, 0, 168, 169, 5, 111, 0, 0, 169, 170, 5, 117, 0, 0, 170,
		171, 5, 110, 0, 0, 171, 172, 5, 116, 0, 0, 172, 10, 1, 0, 0, 0, 173, 174,
		5, 46, 0, 0, 174, 175, 5, 91, 0, 0, 175, 12, 1, 0, 0, 0, 176, 177, 5, 124,
		0, 0, 177, 178, 5, 124, 0, 0, 178, 14, 1, 0, 0, 0, 179, 180, 5, 47, 0,
		0, 180, 16, 1, 0, 0, 0, 181, 182, 5, 37, 0, 0, 182, 18, 1, 0, 0, 0, 183,
		184, 5, 60, 0, 0, 184, 185, 5, 60, 0, 0, 185, 20, 1, 0, 0, 0, 186, 187,
		5, 62, 0, 0, 187, 188, 5, 62, 0, 0, 188, 22, 1, 0, 0, 0, 189, 190, 5, 38,
		0, 0, 190, 24, 1, 0, 0, 0, 191, 192, 5, 38, 0, 0, 192, 193, 5, 38, 0, 0,
		193, 26, 1, 0, 0, 0, 194, 195, 5, 126, 0, 0, 195, 28, 1, 0, 0, 0, 196,
		197, 5, 33, 0, 0, 197, 30, 1, 0, 0, 0, 198, 199, 5, 103, 0, 0, 199, 200,
		5, 114, 0, 0, 200, 201, 5, 111, 0, 0, 201, 202, 5, 117, 0, 0, 202, 203,
		5, 112, 0, 0, 203, 204, 5, 95, 0, 0, 204, 205, 5, 98, 0, 0, 205, 206, 5,
		121, 0, 0, 206, 32, 1, 0, 0, 0, 207, 208, 5, 43, 0, 0, 208, 34, 1, 0, 0,
		0, 209, 210, 5, 45, 0, 0, 210, 36, 1, 0, 0, 0, 211, 212, 5, 111, 0, 0,
		212, 213, 5, 114, 0, 0, 213, 214, 5, 100, 0, 0, 214, 215, 5, 101, 0, 0,
		215, 216, 5, 114, 0, 0, 216, 217, 5, 95, 0, 0, 217, 218, 5, 98, 0, 0, 218,
		227, 5, 121, 0, 0, 219, 220, 5, 115, 0, 0, 220, 221, 5, 111, 0, 0, 221,
		222, 5, 114, 0, 0, 222, 223, 5, 116, 0, 0, 223, 224, 5, 95, 0, 0, 224,
		225, 5, 98, 0, 0, 225, 227, 5, 121, 0, 0, 226, 211, 1, 0, 0, 0, 226, 219,
		1, 0, 0, 0, 227, 38, 1, 0, 0, 0, 228, 229, 5, 58, 0, 0, 229, 230, 5, 99,
		0, 0, 230, 231, 5, 111, 0, 0, 231, 232, 5, 117, 0, 0, 232, 233, 5, 110,
		0, 0, 233, 285, 5, 116, 0, 0, 234, 235, 5, 58, 0, 0, 235, 236, 5, 99, 0,
		0, 236, 237, 5, 111, 0, 0, 237, 238, 5, 117, 0, 0, 238, 239, 5, 110, 0,
		0, 239, 240, 5, 116, 0, 0, 240, 241, 5, 95, 0, 0, 241, 242, 5, 117, 0,
		0, 242, 243, 5, 110, 0, 0, 243, 244, 5, 105, 0, 0, 244, 245, 5, 113, 0,
		0, 245, 246, 5, 117, 0, 0, 246, 285, 5, 101, 0, 0, 247, 248, 5, 58, 0,
		0, 248, 249, 5, 97, 0, 0, 249, 250, 5, 118, 0, 0, 250, 285, 5, 103, 0,
		0, 251, 252, 5, 58, 0, 0, 252, 253, 5, 103, 0, 0, 253, 254, 5, 114, 0,
		0, 254, 255, 5, 111, 0, 0, 255, 256, 5, 117, 0, 0, 256, 257, 5, 112, 0,
		0, 257, 258, 5, 95, 0, 0, 258, 259, 5, 98, 0, 0, 259, 285, 5, 121, 0, 0,
		260, 261, 5, 58, 0, 0, 261, 262, 5, 109, 0, 0, 262, 263, 5, 97, 0, 0, 263,
		285, 5, 120, 0, 0, 264, 265, 5, 58, 0, 0, 265, 266, 5, 109, 0, 0, 266,
		267, 5, 105, 0, 0, 267, 285, 5, 110, 0, 0, 268, 269, 5, 58, 0, 0, 269,
		270, 5, 111, 0, 0, 270, 271, 5, 114, 0, 0, 271, 272, 5, 100, 0, 0, 272,
		273, 5, 101, 0, 0, 273, 274, 5, 114, 0, 0, 274, 275, 5, 95, 0, 0, 275,
		276, 5, 98, 0, 0, 276, 285, 5, 121, 0, 0, 277, 278, 5, 58, 0, 0, 278, 279,
		5, 117, 0, 0, 279, 280, 5, 110, 0, 0, 280, 281, 5, 105, 0, 0, 281, 282,
		5, 113, 0, 0, 282, 283, 5, 117, 0, 0, 283, 285, 5, 101, 0, 0, 284, 228,
		1, 0, 0, 0, 284, 234, 1, 0, 0, 0, 284, 247, 1, 0, 0, 0, 284, 251, 1, 0,
		0, 0, 284, 260, 1, 0, 0, 0, 284, 264, 1, 0, 0, 0, 284, 268, 1, 0, 0, 0,
		284, 277, 1, 0, 0, 0, 285, 40, 1, 0, 0, 0, 286, 287, 5, 36, 0, 0, 287,
		288, 3, 45, 22, 0, 288, 42, 1, 0, 0, 0, 289, 290, 5, 110, 0, 0, 290, 291,
		5, 117, 0, 0, 291, 292, 5, 108, 0, 0, 292, 293, 5, 108, 0, 0, 293, 44,
		1, 0, 0, 0, 294, 298, 7, 0, 0, 0, 295, 297, 7, 1, 0, 0, 296, 295, 1, 0,
		0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0,
		299, 46, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 303, 7, 2, 0, 0, 302, 301,
		1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0,
		0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 6, 23, 0, 0, 307, 48, 1, 0, 0, 0,
		308, 309, 5, 40, 0, 0, 309, 50, 1, 0, 0, 0, 310, 311, 5, 41, 0, 0, 311,
		52, 1, 0, 0, 0, 312, 313, 5, 91, 0, 0, 313, 54, 1, 0, 0, 0, 314, 315, 5,
		93, 0, 0, 315, 56, 1, 0, 0, 0, 316, 317, 5, 44, 0, 0, 317, 58, 1, 0, 0,
		0, 318, 319, 5, 124, 0, 0, 319, 60, 1, 0, 0, 0, 320, 321, 5, 58, 0, 0,
		321, 62, 1, 0, 0, 0, 322, 323, 3, 67, 33, 0, 323, 64, 1, 0, 0, 0, 324,
		349, 3, 63, 31, 0, 325, 327, 5, 45, 0, 0, 326, 325, 1, 0, 0, 0, 326, 327,
		1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329, 3, 67, 33, 0, 329, 331, 5,
		46, 0, 0, 330, 332, 7, 3, 0, 0, 331, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0,
		0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335,
		337, 3, 69, 34, 0, 336, 335, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 349,
		1, 0, 0, 0, 338, 340, 5, 45, 0, 0, 339, 338, 1, 0, 0, 0, 339, 340, 1, 0,
		0, 0, 340, 341, 1, 0, 0, 0, 341, 342, 3, 67, 33, 0, 342, 343, 3, 69, 34,
		0, 343, 349, 1, 0, 0, 0, 344, 346, 5, 45, 0, 0, 345, 344, 1, 0, 0, 0, 345,
		346, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349, 3, 67, 33, 0, 348, 324,
		1, 0, 0, 0, 348, 326, 1, 0, 0, 0, 348, 339, 1, 0, 0, 0, 348, 345, 1, 0,
		0, 0, 349, 66, 1, 0, 0, 0, 350, 359, 5, 48, 0, 0, 351, 355, 7, 4, 0, 0,
		352, 354, 7, 3, 0, 0, 353, 352, 1, 0, 0, 0, 354, 357, 1, 0, 0, 0, 355,
		353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357, 355,
		1, 0, 0, 0, 358, 350, 1, 0, 0, 0, 358, 351, 1, 0, 0, 0, 359, 68, 1, 0,
		0, 0, 360, 362, 7, 5, 0, 0, 361, 363, 7, 6, 0, 0, 362, 361, 1, 0, 0, 0,
		362, 363, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 365, 3, 67, 33, 0, 365,
		70, 1, 0, 0, 0, 366, 367, 5, 60, 0, 0, 367, 368, 5, 61, 0, 0, 368, 72,
		1, 0, 0, 0, 369, 370, 5, 60, 0, 0, 370, 74, 1, 0, 0, 0, 371, 372, 5, 62,
		0, 0, 372, 373, 5, 61, 0, 0, 373, 76, 1, 0, 0, 0, 374, 375, 5, 62, 0, 0,
		375, 78, 1, 0, 0, 0, 376, 377, 5, 33, 0, 0, 377, 378, 5, 61, 0, 0, 378,
		80, 1, 0, 0, 0, 379, 380, 5, 61, 0, 0, 380, 381, 5, 61, 0, 0, 381, 82,
		1, 0, 0, 0, 382, 386, 5, 46, 0, 0, 383, 387, 3, 41, 20, 0, 384, 387, 3,
		45, 22, 0, 385, 387, 3, 87, 43, 0, 386, 383, 1, 0, 0, 0, 386, 384, 1, 0,
		0, 0, 386, 385, 1, 0, 0, 0, 387, 84, 1, 0, 0, 0, 388, 389, 5, 64, 0, 0,
		389, 394, 3, 45, 22, 0, 390, 391, 5, 47, 0, 0, 391, 393, 3, 45, 22, 0,
		392, 390, 1, 0, 0, 0, 393, 396, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394,
		395, 1, 0, 0, 0, 395, 86, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 397, 402, 5,
		34, 0, 0, 398, 401, 3, 89, 44, 0, 399, 401, 8, 7, 0, 0, 400, 398, 1, 0,
		0, 0, 400, 399, 1, 0, 0, 0, 401, 404, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0,
		402, 403, 1, 0, 0, 0, 403, 405, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 405,
		406, 5, 34, 0, 0, 406, 88, 1, 0, 0, 0, 407, 410, 5, 92, 0, 0, 408, 411,
		7, 8, 0, 0, 409, 411, 3, 91, 45, 0, 410, 408, 1, 0, 0, 0, 410, 409, 1,
		0, 0, 0, 411, 90, 1, 0, 0, 0, 412, 413, 5, 117, 0, 0, 413, 414, 3, 93,
		46, 0, 414, 415, 3, 93, 46, 0, 415, 416, 3, 93, 46, 0, 416, 417, 3, 93,
		46, 0, 417, 92, 1, 0, 0, 0, 418, 419, 7, 9, 0, 0, 419, 94, 1, 0, 0, 0,
		420, 421, 7, 3, 0, 0, 421, 96, 1, 0, 0, 0, 422, 423, 7, 10, 0, 0, 423,
		98, 1, 0, 0, 0, 424, 425, 7, 11, 0, 0, 425, 100, 1, 0, 0, 0, 426, 427,
		7, 12, 0, 0, 427, 102, 1, 0, 0, 0, 428, 429, 7, 13, 0, 0, 429, 104, 1,
		0, 0, 0, 430, 431, 7, 5, 0, 0, 431, 106, 1, 0, 0, 0, 432, 433, 7, 14, 0,
		0, 433, 108, 1, 0, 0, 0, 434, 435, 7, 15, 0, 0, 435, 110, 1, 0, 0, 0, 436,
		437, 7, 16, 0, 0, 437, 112, 1, 0, 0, 0, 438, 439, 7, 17, 0, 0, 439, 114,
		1, 0, 0, 0, 440, 441, 7, 18, 0, 0, 441, 116, 1, 0, 0, 0, 442, 443, 7, 19,
		0, 0, 443, 118, 1, 0, 0, 0, 444, 445, 7, 20, 0, 0, 445, 120, 1, 0, 0, 0,
		446, 447, 7, 21, 0, 0, 447, 122, 1, 0, 0, 0, 448, 449, 7, 22, 0, 0, 449,
		124, 1, 0, 0, 0, 450, 451, 7, 23, 0, 0, 451, 126, 1, 0, 0, 0, 452, 453,
		7, 24, 0, 0, 453, 128, 1, 0, 0, 0, 454, 455, 7, 25, 0, 0, 455, 130, 1,
		0, 0, 0, 456, 457, 7, 26, 0, 0, 457, 132, 1, 0, 0, 0, 458, 459, 7, 27,
		0, 0, 459, 134, 1, 0, 0, 0, 460, 461, 7, 28, 0, 0, 461, 136, 1, 0, 0, 0,
		462, 463, 7, 29, 0, 0, 463, 138, 1, 0, 0, 0, 464, 465, 7, 30, 0, 0, 465,
		140, 1, 0, 0, 0, 466, 467, 7, 31, 0, 0, 467, 142, 1, 0, 0, 0, 468, 469,
		7, 32, 0, 0, 469, 144, 1, 0, 0, 0, 470, 471, 7, 33, 0, 0, 471, 146, 1,
		0, 0, 0, 472, 473, 7, 34, 0, 0, 473, 148, 1, 0, 0, 0, 474, 478, 5, 35,
		0, 0, 475, 477, 9, 0, 0, 0, 476, 475, 1, 0, 0, 0, 477, 480, 1, 0, 0, 0,
		478, 479, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 479, 481, 1, 0, 0, 0, 480,
		478, 1, 0, 0, 0, 481, 482, 5, 10, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484,
		6, 74, 0, 0, 484, 150, 1, 0, 0, 0, 20, 0, 226, 284, 298, 304, 326, 333,
		336, 339, 345, 348, 355, 358, 362, 386, 394, 400, 402, 410, 478, 1, 6,
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
	staticData := &slqlexerLexerStaticData
	staticData.once.Do(slqlexerLexerInit)
}

// NewSLQLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSLQLexer(input antlr.CharStream) *SLQLexer {
	SLQLexerInit()
	l := new(SLQLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &slqlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SLQ.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SLQLexer tokens.
const (
	SLQLexerT__0           = 1
	SLQLexerT__1           = 2
	SLQLexerT__2           = 3
	SLQLexerT__3           = 4
	SLQLexerT__4           = 5
	SLQLexerT__5           = 6
	SLQLexerT__6           = 7
	SLQLexerT__7           = 8
	SLQLexerT__8           = 9
	SLQLexerT__9           = 10
	SLQLexerT__10          = 11
	SLQLexerT__11          = 12
	SLQLexerT__12          = 13
	SLQLexerT__13          = 14
	SLQLexerT__14          = 15
	SLQLexerGROUP_BY       = 16
	SLQLexerORDER_ASC      = 17
	SLQLexerORDER_DESC     = 18
	SLQLexerORDER_BY       = 19
	SLQLexerALIAS_RESERVED = 20
	SLQLexerARG            = 21
	SLQLexerNULL           = 22
	SLQLexerID             = 23
	SLQLexerWS             = 24
	SLQLexerLPAR           = 25
	SLQLexerRPAR           = 26
	SLQLexerLBRA           = 27
	SLQLexerRBRA           = 28
	SLQLexerCOMMA          = 29
	SLQLexerPIPE           = 30
	SLQLexerCOLON          = 31
	SLQLexerNN             = 32
	SLQLexerNUMBER         = 33
	SLQLexerLT_EQ          = 34
	SLQLexerLT             = 35
	SLQLexerGT_EQ          = 36
	SLQLexerGT             = 37
	SLQLexerNEQ            = 38
	SLQLexerEQ             = 39
	SLQLexerNAME           = 40
	SLQLexerHANDLE         = 41
	SLQLexerSTRING         = 42
	SLQLexerLINECOMMENT    = 43
)
