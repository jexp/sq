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
		"", "';'", "'*'", "'join'", "'JOIN'", "'j'", "'group'", "'GROUP'", "'g'",
		"'.['", "'sum'", "'SUM'", "'avg'", "'AVG'", "'count'", "'COUNT'", "'where'",
		"'WHERE'", "'||'", "'/'", "'%'", "'+'", "'-'", "'<<'", "'>>'", "'&'",
		"'&&'", "'~'", "'!'", "", "", "'('", "')'", "'['", "']'", "','", "'|'",
		"':'", "", "", "", "'<='", "'<'", "'>='", "'>'", "'!='", "'=='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "ID", "WS", "LPAR",
		"RPAR", "LBRA", "RBRA", "COMMA", "PIPE", "COLON", "NULL", "NN", "NUMBER",
		"LT_EQ", "LT", "GT_EQ", "GT", "NEQ", "EQ", "NAME", "HANDLE", "STRING",
		"LINECOMMENT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "ID", "WS", "LPAR", "RPAR", "LBRA", "RBRA",
		"COMMA", "PIPE", "COLON", "NULL", "NN", "NUMBER", "INTF", "EXP", "LT_EQ",
		"LT", "GT_EQ", "GT", "NEQ", "EQ", "NAME", "HANDLE", "STRING", "ESC",
		"UNICODE", "HEX", "DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "LINECOMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 460, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 5, 28, 267,
		8, 28, 10, 28, 12, 28, 270, 9, 28, 1, 29, 4, 29, 273, 8, 29, 11, 29, 12,
		29, 274, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 301, 8, 37, 1, 38, 1, 38, 1, 39,
		1, 39, 3, 39, 307, 8, 39, 1, 39, 1, 39, 1, 39, 4, 39, 312, 8, 39, 11, 39,
		12, 39, 313, 1, 39, 3, 39, 317, 8, 39, 1, 39, 3, 39, 320, 8, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 3, 39, 326, 8, 39, 1, 39, 3, 39, 329, 8, 39, 1, 40,
		1, 40, 1, 40, 5, 40, 334, 8, 40, 10, 40, 12, 40, 337, 9, 40, 3, 40, 339,
		8, 40, 1, 41, 1, 41, 3, 41, 343, 8, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46,
		1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 3, 48, 366, 8, 48, 1, 49, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 50, 5, 50, 374, 8, 50, 10, 50, 12, 50, 377,
		9, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 3, 51, 384, 8, 51, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1,
		61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66,
		1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 1, 70, 1, 71, 1,
		71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76,
		1, 77, 1, 77, 1, 78, 1, 78, 1, 79, 1, 79, 1, 80, 1, 80, 1, 81, 1, 81, 1,
		81, 1, 81, 5, 81, 452, 8, 81, 10, 81, 12, 81, 455, 9, 81, 1, 81, 1, 81,
		1, 81, 1, 81, 1, 453, 0, 82, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69,
		35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 0, 83, 0, 85, 41, 87, 42,
		89, 43, 91, 44, 93, 45, 95, 46, 97, 47, 99, 48, 101, 49, 103, 0, 105, 0,
		107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 0,
		125, 0, 127, 0, 129, 0, 131, 0, 133, 0, 135, 0, 137, 0, 139, 0, 141, 0,
		143, 0, 145, 0, 147, 0, 149, 0, 151, 0, 153, 0, 155, 0, 157, 0, 159, 0,
		161, 0, 163, 50, 1, 0, 35, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
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
		88, 88, 120, 120, 2, 0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 446,
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
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89,
		1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0,
		97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 163, 1, 0, 0,
		0, 1, 165, 1, 0, 0, 0, 3, 167, 1, 0, 0, 0, 5, 169, 1, 0, 0, 0, 7, 174,
		1, 0, 0, 0, 9, 179, 1, 0, 0, 0, 11, 181, 1, 0, 0, 0, 13, 187, 1, 0, 0,
		0, 15, 193, 1, 0, 0, 0, 17, 195, 1, 0, 0, 0, 19, 198, 1, 0, 0, 0, 21, 202,
		1, 0, 0, 0, 23, 206, 1, 0, 0, 0, 25, 210, 1, 0, 0, 0, 27, 214, 1, 0, 0,
		0, 29, 220, 1, 0, 0, 0, 31, 226, 1, 0, 0, 0, 33, 232, 1, 0, 0, 0, 35, 238,
		1, 0, 0, 0, 37, 241, 1, 0, 0, 0, 39, 243, 1, 0, 0, 0, 41, 245, 1, 0, 0,
		0, 43, 247, 1, 0, 0, 0, 45, 249, 1, 0, 0, 0, 47, 252, 1, 0, 0, 0, 49, 255,
		1, 0, 0, 0, 51, 257, 1, 0, 0, 0, 53, 260, 1, 0, 0, 0, 55, 262, 1, 0, 0,
		0, 57, 264, 1, 0, 0, 0, 59, 272, 1, 0, 0, 0, 61, 278, 1, 0, 0, 0, 63, 280,
		1, 0, 0, 0, 65, 282, 1, 0, 0, 0, 67, 284, 1, 0, 0, 0, 69, 286, 1, 0, 0,
		0, 71, 288, 1, 0, 0, 0, 73, 290, 1, 0, 0, 0, 75, 300, 1, 0, 0, 0, 77, 302,
		1, 0, 0, 0, 79, 328, 1, 0, 0, 0, 81, 338, 1, 0, 0, 0, 83, 340, 1, 0, 0,
		0, 85, 346, 1, 0, 0, 0, 87, 349, 1, 0, 0, 0, 89, 351, 1, 0, 0, 0, 91, 354,
		1, 0, 0, 0, 93, 356, 1, 0, 0, 0, 95, 359, 1, 0, 0, 0, 97, 362, 1, 0, 0,
		0, 99, 367, 1, 0, 0, 0, 101, 370, 1, 0, 0, 0, 103, 380, 1, 0, 0, 0, 105,
		385, 1, 0, 0, 0, 107, 391, 1, 0, 0, 0, 109, 393, 1, 0, 0, 0, 111, 395,
		1, 0, 0, 0, 113, 397, 1, 0, 0, 0, 115, 399, 1, 0, 0, 0, 117, 401, 1, 0,
		0, 0, 119, 403, 1, 0, 0, 0, 121, 405, 1, 0, 0, 0, 123, 407, 1, 0, 0, 0,
		125, 409, 1, 0, 0, 0, 127, 411, 1, 0, 0, 0, 129, 413, 1, 0, 0, 0, 131,
		415, 1, 0, 0, 0, 133, 417, 1, 0, 0, 0, 135, 419, 1, 0, 0, 0, 137, 421,
		1, 0, 0, 0, 139, 423, 1, 0, 0, 0, 141, 425, 1, 0, 0, 0, 143, 427, 1, 0,
		0, 0, 145, 429, 1, 0, 0, 0, 147, 431, 1, 0, 0, 0, 149, 433, 1, 0, 0, 0,
		151, 435, 1, 0, 0, 0, 153, 437, 1, 0, 0, 0, 155, 439, 1, 0, 0, 0, 157,
		441, 1, 0, 0, 0, 159, 443, 1, 0, 0, 0, 161, 445, 1, 0, 0, 0, 163, 447,
		1, 0, 0, 0, 165, 166, 5, 59, 0, 0, 166, 2, 1, 0, 0, 0, 167, 168, 5, 42,
		0, 0, 168, 4, 1, 0, 0, 0, 169, 170, 5, 106, 0, 0, 170, 171, 5, 111, 0,
		0, 171, 172, 5, 105, 0, 0, 172, 173, 5, 110, 0, 0, 173, 6, 1, 0, 0, 0,
		174, 175, 5, 74, 0, 0, 175, 176, 5, 79, 0, 0, 176, 177, 5, 73, 0, 0, 177,
		178, 5, 78, 0, 0, 178, 8, 1, 0, 0, 0, 179, 180, 5, 106, 0, 0, 180, 10,
		1, 0, 0, 0, 181, 182, 5, 103, 0, 0, 182, 183, 5, 114, 0, 0, 183, 184, 5,
		111, 0, 0, 184, 185, 5, 117, 0, 0, 185, 186, 5, 112, 0, 0, 186, 12, 1,
		0, 0, 0, 187, 188, 5, 71, 0, 0, 188, 189, 5, 82, 0, 0, 189, 190, 5, 79,
		0, 0, 190, 191, 5, 85, 0, 0, 191, 192, 5, 80, 0, 0, 192, 14, 1, 0, 0, 0,
		193, 194, 5, 103, 0, 0, 194, 16, 1, 0, 0, 0, 195, 196, 5, 46, 0, 0, 196,
		197, 5, 91, 0, 0, 197, 18, 1, 0, 0, 0, 198, 199, 5, 115, 0, 0, 199, 200,
		5, 117, 0, 0, 200, 201, 5, 109, 0, 0, 201, 20, 1, 0, 0, 0, 202, 203, 5,
		83, 0, 0, 203, 204, 5, 85, 0, 0, 204, 205, 5, 77, 0, 0, 205, 22, 1, 0,
		0, 0, 206, 207, 5, 97, 0, 0, 207, 208, 5, 118, 0, 0, 208, 209, 5, 103,
		0, 0, 209, 24, 1, 0, 0, 0, 210, 211, 5, 65, 0, 0, 211, 212, 5, 86, 0, 0,
		212, 213, 5, 71, 0, 0, 213, 26, 1, 0, 0, 0, 214, 215, 5, 99, 0, 0, 215,
		216, 5, 111, 0, 0, 216, 217, 5, 117, 0, 0, 217, 218, 5, 110, 0, 0, 218,
		219, 5, 116, 0, 0, 219, 28, 1, 0, 0, 0, 220, 221, 5, 67, 0, 0, 221, 222,
		5, 79, 0, 0, 222, 223, 5, 85, 0, 0, 223, 224, 5, 78, 0, 0, 224, 225, 5,
		84, 0, 0, 225, 30, 1, 0, 0, 0, 226, 227, 5, 119, 0, 0, 227, 228, 5, 104,
		0, 0, 228, 229, 5, 101, 0, 0, 229, 230, 5, 114, 0, 0, 230, 231, 5, 101,
		0, 0, 231, 32, 1, 0, 0, 0, 232, 233, 5, 87, 0, 0, 233, 234, 5, 72, 0, 0,
		234, 235, 5, 69, 0, 0, 235, 236, 5, 82, 0, 0, 236, 237, 5, 69, 0, 0, 237,
		34, 1, 0, 0, 0, 238, 239, 5, 124, 0, 0, 239, 240, 5, 124, 0, 0, 240, 36,
		1, 0, 0, 0, 241, 242, 5, 47, 0, 0, 242, 38, 1, 0, 0, 0, 243, 244, 5, 37,
		0, 0, 244, 40, 1, 0, 0, 0, 245, 246, 5, 43, 0, 0, 246, 42, 1, 0, 0, 0,
		247, 248, 5, 45, 0, 0, 248, 44, 1, 0, 0, 0, 249, 250, 5, 60, 0, 0, 250,
		251, 5, 60, 0, 0, 251, 46, 1, 0, 0, 0, 252, 253, 5, 62, 0, 0, 253, 254,
		5, 62, 0, 0, 254, 48, 1, 0, 0, 0, 255, 256, 5, 38, 0, 0, 256, 50, 1, 0,
		0, 0, 257, 258, 5, 38, 0, 0, 258, 259, 5, 38, 0, 0, 259, 52, 1, 0, 0, 0,
		260, 261, 5, 126, 0, 0, 261, 54, 1, 0, 0, 0, 262, 263, 5, 33, 0, 0, 263,
		56, 1, 0, 0, 0, 264, 268, 7, 0, 0, 0, 265, 267, 7, 1, 0, 0, 266, 265, 1,
		0, 0, 0, 267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0,
		0, 269, 58, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 273, 7, 2, 0, 0, 272,
		271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 6, 29, 0, 0, 277, 60, 1, 0,
		0, 0, 278, 279, 5, 40, 0, 0, 279, 62, 1, 0, 0, 0, 280, 281, 5, 41, 0, 0,
		281, 64, 1, 0, 0, 0, 282, 283, 5, 91, 0, 0, 283, 66, 1, 0, 0, 0, 284, 285,
		5, 93, 0, 0, 285, 68, 1, 0, 0, 0, 286, 287, 5, 44, 0, 0, 287, 70, 1, 0,
		0, 0, 288, 289, 5, 124, 0, 0, 289, 72, 1, 0, 0, 0, 290, 291, 5, 58, 0,
		0, 291, 74, 1, 0, 0, 0, 292, 293, 5, 110, 0, 0, 293, 294, 5, 117, 0, 0,
		294, 295, 5, 108, 0, 0, 295, 301, 5, 108, 0, 0, 296, 297, 5, 78, 0, 0,
		297, 298, 5, 85, 0, 0, 298, 299, 5, 76, 0, 0, 299, 301, 5, 76, 0, 0, 300,
		292, 1, 0, 0, 0, 300, 296, 1, 0, 0, 0, 301, 76, 1, 0, 0, 0, 302, 303, 3,
		81, 40, 0, 303, 78, 1, 0, 0, 0, 304, 329, 3, 77, 38, 0, 305, 307, 5, 45,
		0, 0, 306, 305, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0,
		308, 309, 3, 81, 40, 0, 309, 311, 5, 46, 0, 0, 310, 312, 7, 3, 0, 0, 311,
		310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314,
		1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 317, 3, 83, 41, 0, 316, 315, 1,
		0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 329, 1, 0, 0, 0, 318, 320, 5, 45, 0,
		0, 319, 318, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		322, 3, 81, 40, 0, 322, 323, 3, 83, 41, 0, 323, 329, 1, 0, 0, 0, 324, 326,
		5, 45, 0, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 1, 0,
		0, 0, 327, 329, 3, 81, 40, 0, 328, 304, 1, 0, 0, 0, 328, 306, 1, 0, 0,
		0, 328, 319, 1, 0, 0, 0, 328, 325, 1, 0, 0, 0, 329, 80, 1, 0, 0, 0, 330,
		339, 5, 48, 0, 0, 331, 335, 7, 4, 0, 0, 332, 334, 7, 3, 0, 0, 333, 332,
		1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0,
		0, 0, 336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338, 330, 1, 0, 0, 0,
		338, 331, 1, 0, 0, 0, 339, 82, 1, 0, 0, 0, 340, 342, 7, 5, 0, 0, 341, 343,
		7, 6, 0, 0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 1, 0,
		0, 0, 344, 345, 3, 81, 40, 0, 345, 84, 1, 0, 0, 0, 346, 347, 5, 60, 0,
		0, 347, 348, 5, 61, 0, 0, 348, 86, 1, 0, 0, 0, 349, 350, 5, 60, 0, 0, 350,
		88, 1, 0, 0, 0, 351, 352, 5, 62, 0, 0, 352, 353, 5, 61, 0, 0, 353, 90,
		1, 0, 0, 0, 354, 355, 5, 62, 0, 0, 355, 92, 1, 0, 0, 0, 356, 357, 5, 33,
		0, 0, 357, 358, 5, 61, 0, 0, 358, 94, 1, 0, 0, 0, 359, 360, 5, 61, 0, 0,
		360, 361, 5, 61, 0, 0, 361, 96, 1, 0, 0, 0, 362, 365, 5, 46, 0, 0, 363,
		366, 3, 57, 28, 0, 364, 366, 3, 101, 50, 0, 365, 363, 1, 0, 0, 0, 365,
		364, 1, 0, 0, 0, 366, 98, 1, 0, 0, 0, 367, 368, 5, 64, 0, 0, 368, 369,
		3, 57, 28, 0, 369, 100, 1, 0, 0, 0, 370, 375, 5, 34, 0, 0, 371, 374, 3,
		103, 51, 0, 372, 374, 8, 7, 0, 0, 373, 371, 1, 0, 0, 0, 373, 372, 1, 0,
		0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0,
		376, 378, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 379, 5, 34, 0, 0, 379,
		102, 1, 0, 0, 0, 380, 383, 5, 92, 0, 0, 381, 384, 7, 8, 0, 0, 382, 384,
		3, 105, 52, 0, 383, 381, 1, 0, 0, 0, 383, 382, 1, 0, 0, 0, 384, 104, 1,
		0, 0, 0, 385, 386, 5, 117, 0, 0, 386, 387, 3, 107, 53, 0, 387, 388, 3,
		107, 53, 0, 388, 389, 3, 107, 53, 0, 389, 390, 3, 107, 53, 0, 390, 106,
		1, 0, 0, 0, 391, 392, 7, 9, 0, 0, 392, 108, 1, 0, 0, 0, 393, 394, 7, 3,
		0, 0, 394, 110, 1, 0, 0, 0, 395, 396, 7, 10, 0, 0, 396, 112, 1, 0, 0, 0,
		397, 398, 7, 11, 0, 0, 398, 114, 1, 0, 0, 0, 399, 400, 7, 12, 0, 0, 400,
		116, 1, 0, 0, 0, 401, 402, 7, 13, 0, 0, 402, 118, 1, 0, 0, 0, 403, 404,
		7, 5, 0, 0, 404, 120, 1, 0, 0, 0, 405, 406, 7, 14, 0, 0, 406, 122, 1, 0,
		0, 0, 407, 408, 7, 15, 0, 0, 408, 124, 1, 0, 0, 0, 409, 410, 7, 16, 0,
		0, 410, 126, 1, 0, 0, 0, 411, 412, 7, 17, 0, 0, 412, 128, 1, 0, 0, 0, 413,
		414, 7, 18, 0, 0, 414, 130, 1, 0, 0, 0, 415, 416, 7, 19, 0, 0, 416, 132,
		1, 0, 0, 0, 417, 418, 7, 20, 0, 0, 418, 134, 1, 0, 0, 0, 419, 420, 7, 21,
		0, 0, 420, 136, 1, 0, 0, 0, 421, 422, 7, 22, 0, 0, 422, 138, 1, 0, 0, 0,
		423, 424, 7, 23, 0, 0, 424, 140, 1, 0, 0, 0, 425, 426, 7, 24, 0, 0, 426,
		142, 1, 0, 0, 0, 427, 428, 7, 25, 0, 0, 428, 144, 1, 0, 0, 0, 429, 430,
		7, 26, 0, 0, 430, 146, 1, 0, 0, 0, 431, 432, 7, 27, 0, 0, 432, 148, 1,
		0, 0, 0, 433, 434, 7, 28, 0, 0, 434, 150, 1, 0, 0, 0, 435, 436, 7, 29,
		0, 0, 436, 152, 1, 0, 0, 0, 437, 438, 7, 30, 0, 0, 438, 154, 1, 0, 0, 0,
		439, 440, 7, 31, 0, 0, 440, 156, 1, 0, 0, 0, 441, 442, 7, 32, 0, 0, 442,
		158, 1, 0, 0, 0, 443, 444, 7, 33, 0, 0, 444, 160, 1, 0, 0, 0, 445, 446,
		7, 34, 0, 0, 446, 162, 1, 0, 0, 0, 447, 448, 5, 47, 0, 0, 448, 449, 5,
		47, 0, 0, 449, 453, 1, 0, 0, 0, 450, 452, 9, 0, 0, 0, 451, 450, 1, 0, 0,
		0, 452, 455, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454,
		456, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 456, 457, 5, 10, 0, 0, 457, 458,
		1, 0, 0, 0, 458, 459, 6, 81, 0, 0, 459, 164, 1, 0, 0, 0, 18, 0, 268, 274,
		300, 306, 313, 316, 319, 325, 328, 335, 338, 342, 365, 373, 375, 383, 453,
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
	SLQLexerT__0        = 1
	SLQLexerT__1        = 2
	SLQLexerT__2        = 3
	SLQLexerT__3        = 4
	SLQLexerT__4        = 5
	SLQLexerT__5        = 6
	SLQLexerT__6        = 7
	SLQLexerT__7        = 8
	SLQLexerT__8        = 9
	SLQLexerT__9        = 10
	SLQLexerT__10       = 11
	SLQLexerT__11       = 12
	SLQLexerT__12       = 13
	SLQLexerT__13       = 14
	SLQLexerT__14       = 15
	SLQLexerT__15       = 16
	SLQLexerT__16       = 17
	SLQLexerT__17       = 18
	SLQLexerT__18       = 19
	SLQLexerT__19       = 20
	SLQLexerT__20       = 21
	SLQLexerT__21       = 22
	SLQLexerT__22       = 23
	SLQLexerT__23       = 24
	SLQLexerT__24       = 25
	SLQLexerT__25       = 26
	SLQLexerT__26       = 27
	SLQLexerT__27       = 28
	SLQLexerID          = 29
	SLQLexerWS          = 30
	SLQLexerLPAR        = 31
	SLQLexerRPAR        = 32
	SLQLexerLBRA        = 33
	SLQLexerRBRA        = 34
	SLQLexerCOMMA       = 35
	SLQLexerPIPE        = 36
	SLQLexerCOLON       = 37
	SLQLexerNULL        = 38
	SLQLexerNN          = 39
	SLQLexerNUMBER      = 40
	SLQLexerLT_EQ       = 41
	SLQLexerLT          = 42
	SLQLexerGT_EQ       = 43
	SLQLexerGT          = 44
	SLQLexerNEQ         = 45
	SLQLexerEQ          = 46
	SLQLexerNAME        = 47
	SLQLexerHANDLE      = 48
	SLQLexerSTRING      = 49
	SLQLexerLINECOMMENT = 50
)
