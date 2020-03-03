// Code generated from antlr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // antlr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 61, 439,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	7, 2, 90, 10, 2, 12, 2, 14, 2, 93, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 104, 10, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 118, 10, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 127, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	132, 10, 7, 7, 7, 134, 10, 7, 12, 7, 14, 7, 137, 11, 7, 3, 8, 3, 8, 7,
	8, 141, 10, 8, 12, 8, 14, 8, 144, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 158, 10, 11, 12,
	11, 14, 11, 161, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 167, 10, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 7, 14, 182, 10, 14, 12, 14, 14, 14, 185, 11, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 197,
	10, 16, 12, 16, 14, 16, 200, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	5, 17, 207, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 214, 10,
	17, 12, 17, 14, 17, 217, 11, 17, 3, 17, 3, 17, 5, 17, 221, 10, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 7, 19, 235, 10, 19, 12, 19, 14, 19, 238, 11, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 7, 21, 246, 10, 21, 12, 21, 14, 21, 249, 11, 21, 3,
	21, 3, 21, 3, 22, 5, 22, 254, 10, 22, 3, 23, 3, 23, 7, 23, 258, 10, 23,
	12, 23, 14, 23, 261, 11, 23, 3, 23, 3, 23, 3, 24, 5, 24, 266, 10, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 273, 10, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 5, 25, 281, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 7, 25, 288, 10, 25, 12, 25, 14, 25, 291, 11, 25, 3, 25, 3, 25, 5, 25,
	295, 10, 25, 3, 26, 3, 26, 3, 26, 5, 26, 300, 10, 26, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 27, 7, 27, 307, 10, 27, 12, 27, 14, 27, 310, 11, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 5, 28, 316, 10, 28, 3, 29, 3, 29, 3, 29, 7, 29, 321,
	10, 29, 12, 29, 14, 29, 324, 11, 29, 3, 30, 3, 30, 5, 30, 328, 10, 30,
	3, 31, 5, 31, 331, 10, 31, 3, 31, 3, 31, 3, 31, 5, 31, 336, 10, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 344, 10, 31, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 32, 7, 32, 351, 10, 32, 12, 32, 14, 32, 354, 11, 32, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 361, 10, 33, 12, 33, 14, 33, 364,
	11, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 5, 34, 371, 10, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 379, 10, 34, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 5, 35, 396, 10, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 38, 7, 38, 407, 10, 38, 12, 38, 14, 38, 410, 11, 38, 3,
	39, 5, 39, 413, 10, 39, 3, 39, 3, 39, 5, 39, 417, 10, 39, 3, 39, 7, 39,
	420, 10, 39, 12, 39, 14, 39, 423, 11, 39, 3, 39, 3, 39, 3, 40, 3, 40, 5,
	40, 429, 10, 40, 3, 40, 3, 40, 5, 40, 433, 10, 40, 3, 40, 3, 40, 5, 40,
	437, 10, 40, 3, 40, 2, 2, 41, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 2, 9, 3, 2, 20, 21, 4, 2, 22, 22, 38,
	38, 8, 2, 4, 6, 9, 11, 13, 14, 28, 31, 33, 33, 36, 37, 8, 2, 4, 4, 9, 10,
	13, 14, 28, 31, 33, 33, 36, 37, 8, 2, 8, 8, 16, 16, 18, 19, 26, 27, 32,
	34, 38, 38, 3, 2, 50, 51, 3, 2, 57, 58, 2, 464, 2, 80, 3, 2, 2, 2, 4, 96,
	3, 2, 2, 2, 6, 101, 3, 2, 2, 2, 8, 108, 3, 2, 2, 2, 10, 112, 3, 2, 2, 2,
	12, 126, 3, 2, 2, 2, 14, 138, 3, 2, 2, 2, 16, 147, 3, 2, 2, 2, 18, 151,
	3, 2, 2, 2, 20, 155, 3, 2, 2, 2, 22, 166, 3, 2, 2, 2, 24, 168, 3, 2, 2,
	2, 26, 172, 3, 2, 2, 2, 28, 188, 3, 2, 2, 2, 30, 192, 3, 2, 2, 2, 32, 203,
	3, 2, 2, 2, 34, 224, 3, 2, 2, 2, 36, 228, 3, 2, 2, 2, 38, 241, 3, 2, 2,
	2, 40, 247, 3, 2, 2, 2, 42, 253, 3, 2, 2, 2, 44, 259, 3, 2, 2, 2, 46, 265,
	3, 2, 2, 2, 48, 267, 3, 2, 2, 2, 50, 296, 3, 2, 2, 2, 52, 303, 3, 2, 2,
	2, 54, 315, 3, 2, 2, 2, 56, 317, 3, 2, 2, 2, 58, 327, 3, 2, 2, 2, 60, 330,
	3, 2, 2, 2, 62, 347, 3, 2, 2, 2, 64, 355, 3, 2, 2, 2, 66, 367, 3, 2, 2,
	2, 68, 382, 3, 2, 2, 2, 70, 399, 3, 2, 2, 2, 72, 401, 3, 2, 2, 2, 74, 403,
	3, 2, 2, 2, 76, 412, 3, 2, 2, 2, 78, 436, 3, 2, 2, 2, 80, 91, 5, 4, 3,
	2, 81, 90, 5, 6, 4, 2, 82, 90, 5, 8, 5, 2, 83, 90, 5, 10, 6, 2, 84, 90,
	5, 24, 13, 2, 85, 90, 5, 28, 15, 2, 86, 90, 5, 36, 19, 2, 87, 90, 5, 18,
	10, 2, 88, 90, 7, 47, 2, 2, 89, 81, 3, 2, 2, 2, 89, 82, 3, 2, 2, 2, 89,
	83, 3, 2, 2, 2, 89, 84, 3, 2, 2, 2, 89, 85, 3, 2, 2, 2, 89, 86, 3, 2, 2,
	2, 89, 87, 3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89,
	3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2,
	94, 95, 7, 2, 2, 3, 95, 3, 3, 2, 2, 2, 96, 97, 7, 34, 2, 2, 97, 98, 7,
	53, 2, 2, 98, 99, 9, 2, 2, 2, 99, 100, 7, 47, 2, 2, 100, 5, 3, 2, 2, 2,
	101, 103, 7, 12, 2, 2, 102, 104, 9, 3, 2, 2, 103, 102, 3, 2, 2, 2, 103,
	104, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 7, 58, 2, 2, 106, 107,
	7, 47, 2, 2, 107, 7, 3, 2, 2, 2, 108, 109, 7, 19, 2, 2, 109, 110, 5, 74,
	38, 2, 110, 111, 7, 47, 2, 2, 111, 9, 3, 2, 2, 2, 112, 113, 7, 18, 2, 2,
	113, 114, 5, 12, 7, 2, 114, 117, 7, 53, 2, 2, 115, 118, 5, 78, 40, 2, 116,
	118, 5, 14, 8, 2, 117, 115, 3, 2, 2, 2, 117, 116, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 120, 7, 47, 2, 2, 120, 11, 3, 2, 2, 2, 121, 127, 7, 54,
	2, 2, 122, 123, 7, 39, 2, 2, 123, 124, 5, 74, 38, 2, 124, 125, 7, 40, 2,
	2, 125, 127, 3, 2, 2, 2, 126, 121, 3, 2, 2, 2, 126, 122, 3, 2, 2, 2, 127,
	135, 3, 2, 2, 2, 128, 131, 7, 49, 2, 2, 129, 132, 7, 54, 2, 2, 130, 132,
	5, 72, 37, 2, 131, 129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 134, 3,
	2, 2, 2, 133, 128, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2,
	2, 135, 136, 3, 2, 2, 2, 136, 13, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138,
	142, 7, 41, 2, 2, 139, 141, 5, 16, 9, 2, 140, 139, 3, 2, 2, 2, 141, 144,
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2,
	2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 42, 2, 2, 146, 15, 3, 2, 2, 2,
	147, 148, 5, 12, 7, 2, 148, 149, 7, 3, 2, 2, 149, 150, 5, 78, 40, 2, 150,
	17, 3, 2, 2, 2, 151, 152, 7, 8, 2, 2, 152, 153, 5, 22, 12, 2, 153, 154,
	5, 20, 11, 2, 154, 19, 3, 2, 2, 2, 155, 159, 7, 41, 2, 2, 156, 158, 5,
	60, 31, 2, 157, 156, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2,
	2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2,
	162, 163, 7, 42, 2, 2, 163, 21, 3, 2, 2, 2, 164, 167, 7, 54, 2, 2, 165,
	167, 5, 74, 38, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 23,
	3, 2, 2, 2, 168, 169, 7, 16, 2, 2, 169, 170, 7, 54, 2, 2, 170, 171, 5,
	26, 14, 2, 171, 25, 3, 2, 2, 2, 172, 183, 7, 41, 2, 2, 173, 182, 5, 60,
	31, 2, 174, 182, 5, 28, 15, 2, 175, 182, 5, 24, 13, 2, 176, 182, 5, 10,
	6, 2, 177, 182, 5, 64, 33, 2, 178, 182, 5, 68, 35, 2, 179, 182, 5, 50,
	26, 2, 180, 182, 7, 47, 2, 2, 181, 173, 3, 2, 2, 2, 181, 174, 3, 2, 2,
	2, 181, 175, 3, 2, 2, 2, 181, 176, 3, 2, 2, 2, 181, 177, 3, 2, 2, 2, 181,
	178, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 185,
	3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 186, 3, 2,
	2, 2, 185, 183, 3, 2, 2, 2, 186, 187, 7, 42, 2, 2, 187, 27, 3, 2, 2, 2,
	188, 189, 7, 7, 2, 2, 189, 190, 7, 54, 2, 2, 190, 191, 5, 30, 16, 2, 191,
	29, 3, 2, 2, 2, 192, 198, 7, 41, 2, 2, 193, 197, 5, 10, 6, 2, 194, 197,
	5, 32, 17, 2, 195, 197, 7, 47, 2, 2, 196, 193, 3, 2, 2, 2, 196, 194, 3,
	2, 2, 2, 196, 195, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2,
	2, 198, 199, 3, 2, 2, 2, 199, 201, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201,
	202, 7, 42, 2, 2, 202, 31, 3, 2, 2, 2, 203, 204, 7, 54, 2, 2, 204, 206,
	7, 53, 2, 2, 205, 207, 7, 50, 2, 2, 206, 205, 3, 2, 2, 2, 206, 207, 3,
	2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 220, 7, 55, 2, 2, 209, 210, 7, 43,
	2, 2, 210, 215, 5, 34, 18, 2, 211, 212, 7, 48, 2, 2, 212, 214, 5, 34, 18,
	2, 213, 211, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 218, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 219,
	7, 44, 2, 2, 219, 221, 3, 2, 2, 2, 220, 209, 3, 2, 2, 2, 220, 221, 3, 2,
	2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 7, 47, 2, 2, 223, 33, 3, 2, 2, 2,
	224, 225, 5, 12, 7, 2, 225, 226, 7, 53, 2, 2, 226, 227, 5, 78, 40, 2, 227,
	35, 3, 2, 2, 2, 228, 229, 7, 27, 2, 2, 229, 230, 7, 54, 2, 2, 230, 236,
	7, 41, 2, 2, 231, 235, 5, 10, 6, 2, 232, 235, 5, 48, 25, 2, 233, 235, 7,
	47, 2, 2, 234, 231, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 233, 3, 2, 2,
	2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237,
	239, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 240, 7, 42, 2, 2, 240, 37,
	3, 2, 2, 2, 241, 242, 7, 54, 2, 2, 242, 39, 3, 2, 2, 2, 243, 244, 7, 54,
	2, 2, 244, 246, 7, 49, 2, 2, 245, 243, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2,
	247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 250, 3, 2, 2, 2, 249,
	247, 3, 2, 2, 2, 250, 251, 7, 54, 2, 2, 251, 41, 3, 2, 2, 2, 252, 254,
	7, 32, 2, 2, 253, 252, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 43, 3, 2,
	2, 2, 255, 256, 7, 54, 2, 2, 256, 258, 7, 49, 2, 2, 257, 255, 3, 2, 2,
	2, 258, 261, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260,
	262, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 262, 263, 7, 54, 2, 2, 263, 45,
	3, 2, 2, 2, 264, 266, 7, 32, 2, 2, 265, 264, 3, 2, 2, 2, 265, 266, 3, 2,
	2, 2, 266, 47, 3, 2, 2, 2, 267, 268, 7, 26, 2, 2, 268, 269, 5, 38, 20,
	2, 269, 270, 7, 39, 2, 2, 270, 272, 5, 42, 22, 2, 271, 273, 7, 49, 2, 2,
	272, 271, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274,
	275, 5, 40, 21, 2, 275, 276, 7, 40, 2, 2, 276, 277, 7, 25, 2, 2, 277, 278,
	7, 39, 2, 2, 278, 280, 5, 46, 24, 2, 279, 281, 7, 49, 2, 2, 280, 279, 3,
	2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 5, 44, 23,
	2, 283, 294, 7, 40, 2, 2, 284, 289, 7, 41, 2, 2, 285, 288, 5, 10, 6, 2,
	286, 288, 7, 47, 2, 2, 287, 285, 3, 2, 2, 2, 287, 286, 3, 2, 2, 2, 288,
	291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 292,
	3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 295, 7, 42, 2, 2, 293, 295, 7, 47,
	2, 2, 294, 284, 3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295, 49, 3, 2, 2, 2,
	296, 299, 7, 24, 2, 2, 297, 300, 5, 52, 27, 2, 298, 300, 5, 56, 29, 2,
	299, 297, 3, 2, 2, 2, 299, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301,
	302, 7, 47, 2, 2, 302, 51, 3, 2, 2, 2, 303, 308, 5, 54, 28, 2, 304, 305,
	7, 48, 2, 2, 305, 307, 5, 54, 28, 2, 306, 304, 3, 2, 2, 2, 307, 310, 3,
	2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 53, 3, 2, 2,
	2, 310, 308, 3, 2, 2, 2, 311, 316, 7, 55, 2, 2, 312, 313, 7, 55, 2, 2,
	313, 314, 7, 35, 2, 2, 314, 316, 7, 55, 2, 2, 315, 311, 3, 2, 2, 2, 315,
	312, 3, 2, 2, 2, 316, 55, 3, 2, 2, 2, 317, 322, 7, 58, 2, 2, 318, 319,
	7, 48, 2, 2, 319, 321, 7, 58, 2, 2, 320, 318, 3, 2, 2, 2, 321, 324, 3,
	2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 57, 3, 2, 2,
	2, 324, 322, 3, 2, 2, 2, 325, 328, 9, 4, 2, 2, 326, 328, 5, 76, 39, 2,
	327, 325, 3, 2, 2, 2, 327, 326, 3, 2, 2, 2, 328, 59, 3, 2, 2, 2, 329, 331,
	7, 23, 2, 2, 330, 329, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 332, 3, 2,
	2, 2, 332, 335, 5, 58, 30, 2, 333, 336, 7, 54, 2, 2, 334, 336, 5, 72, 37,
	2, 335, 333, 3, 2, 2, 2, 335, 334, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337,
	338, 7, 53, 2, 2, 338, 343, 7, 55, 2, 2, 339, 340, 7, 43, 2, 2, 340, 341,
	5, 62, 32, 2, 341, 342, 7, 44, 2, 2, 342, 344, 3, 2, 2, 2, 343, 339, 3,
	2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 346, 7, 47, 2,
	2, 346, 61, 3, 2, 2, 2, 347, 352, 5, 34, 18, 2, 348, 349, 7, 48, 2, 2,
	349, 351, 5, 34, 18, 2, 350, 348, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352,
	350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 63, 3, 2, 2, 2, 354, 352, 3,
	2, 2, 2, 355, 356, 7, 17, 2, 2, 356, 357, 7, 54, 2, 2, 357, 362, 7, 41,
	2, 2, 358, 361, 5, 66, 34, 2, 359, 361, 7, 47, 2, 2, 360, 358, 3, 2, 2,
	2, 360, 359, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 362,
	363, 3, 2, 2, 2, 363, 365, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 365, 366,
	7, 42, 2, 2, 366, 65, 3, 2, 2, 2, 367, 370, 5, 58, 30, 2, 368, 371, 7,
	54, 2, 2, 369, 371, 5, 72, 37, 2, 370, 368, 3, 2, 2, 2, 370, 369, 3, 2,
	2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 7, 53, 2, 2, 373, 378, 7, 55, 2,
	2, 374, 375, 7, 43, 2, 2, 375, 376, 5, 62, 32, 2, 376, 377, 7, 44, 2, 2,
	377, 379, 3, 2, 2, 2, 378, 374, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379,
	380, 3, 2, 2, 2, 380, 381, 7, 47, 2, 2, 381, 67, 3, 2, 2, 2, 382, 383,
	7, 15, 2, 2, 383, 384, 7, 45, 2, 2, 384, 385, 5, 70, 36, 2, 385, 386, 7,
	48, 2, 2, 386, 387, 5, 58, 30, 2, 387, 388, 7, 46, 2, 2, 388, 389, 7, 54,
	2, 2, 389, 390, 7, 53, 2, 2, 390, 395, 7, 55, 2, 2, 391, 392, 7, 43, 2,
	2, 392, 393, 5, 62, 32, 2, 393, 394, 7, 44, 2, 2, 394, 396, 3, 2, 2, 2,
	395, 391, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397,
	398, 7, 47, 2, 2, 398, 69, 3, 2, 2, 2, 399, 400, 9, 5, 2, 2, 400, 71, 3,
	2, 2, 2, 401, 402, 9, 6, 2, 2, 402, 73, 3, 2, 2, 2, 403, 408, 7, 54, 2,
	2, 404, 405, 7, 49, 2, 2, 405, 407, 7, 54, 2, 2, 406, 404, 3, 2, 2, 2,
	407, 410, 3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409,
	75, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 411, 413, 7, 49, 2, 2, 412, 411,
	3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 421, 3, 2, 2, 2, 414, 417, 7, 54,
	2, 2, 415, 417, 5, 72, 37, 2, 416, 414, 3, 2, 2, 2, 416, 415, 3, 2, 2,
	2, 417, 418, 3, 2, 2, 2, 418, 420, 7, 49, 2, 2, 419, 416, 3, 2, 2, 2, 420,
	423, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 424,
	3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 424, 425, 7, 54, 2, 2, 425, 77, 3, 2,
	2, 2, 426, 437, 5, 74, 38, 2, 427, 429, 9, 7, 2, 2, 428, 427, 3, 2, 2,
	2, 428, 429, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 437, 7, 55, 2, 2, 431,
	433, 9, 7, 2, 2, 432, 431, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434,
	3, 2, 2, 2, 434, 437, 7, 56, 2, 2, 435, 437, 9, 8, 2, 2, 436, 426, 3, 2,
	2, 2, 436, 428, 3, 2, 2, 2, 436, 432, 3, 2, 2, 2, 436, 435, 3, 2, 2, 2,
	437, 79, 3, 2, 2, 2, 51, 89, 91, 103, 117, 126, 131, 135, 142, 159, 166,
	181, 183, 196, 198, 206, 215, 220, 234, 236, 247, 253, 259, 265, 272, 280,
	287, 289, 294, 299, 308, 315, 322, 327, 330, 335, 343, 352, 360, 362, 370,
	378, 395, 408, 412, 416, 421, 428, 432, 436,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'bool'", "'bytes'", "'double'", "'enum'", "'extend'", "'fixed32'",
	"'fixed64'", "'float'", "'import'", "'int32'", "'int64'", "'map'", "'message'",
	"'oneof'", "'option'", "'package'", "'\"proto3\"'", "''proto3''", "'public'",
	"'repeated'", "'reserved'", "'returns'", "'rpc'", "'service'", "'sfixed32'",
	"'sfixed64'", "'sint32'", "'sint64'", "'stream'", "'string'", "'syntax'",
	"'to'", "'uint32'", "'uint64'", "'weak'", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "'<'", "'>'", "';'", "','", "'.'", "'-'", "'+'", "'\\'", "'='",
}
var symbolicNames = []string{
	"", "", "BOOL", "BYTES", "DOUBLE", "ENUM", "EXTEND", "FIXED32", "FIXED64",
	"FLOAT", "IMPORT", "INT32", "INT64", "MAP", "MESSAGE", "ONEOF", "OPTION",
	"PACKAGE", "PROTO3_DOUBLE", "PROTO3_SINGLE", "PUBLIC", "REPEATED", "RESERVED",
	"RETURNS", "RPC", "SERVICE", "SFIXED32", "SFIXED64", "SINT32", "SINT64",
	"STREAM", "STRING", "SYNTAX", "TO", "UINT32", "UINT64", "WEAK", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "LCHEVR", "RCHEVR", "SEMI",
	"COMMA", "DOT", "MINUS", "PLUS", "SLASH", "EQ", "Ident", "IntLit", "FloatLit",
	"BoolLit", "StrLit", "WS", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"proto", "syntax", "imports", "packageName", "option", "optionName", "optionDecl",
	"optionDeclVar", "extend", "extendDecl", "extendName", "message", "messageDecl",
	"enum", "enumDecl", "enumField", "keyValue", "service", "rpcName", "rpcParams",
	"clientStreams", "rpcReturns", "serverStreams", "rpc", "reserved", "ranges",
	"rangeDef", "fieldNames", "fieldType", "field", "fieldOpts", "oneof", "oneofField",
	"mapField", "keyType", "reservedWord", "fullIdent", "messageOrEnumType",
	"constant",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type antlrParser struct {
	*antlr.BaseParser
}

func NewantlrParser(input antlr.TokenStream) *antlrParser {
	this := new(antlrParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "antlr.g4"

	return this
}

// antlrParser tokens.
const (
	antlrParserEOF           = antlr.TokenEOF
	antlrParserT__0          = 1
	antlrParserBOOL          = 2
	antlrParserBYTES         = 3
	antlrParserDOUBLE        = 4
	antlrParserENUM          = 5
	antlrParserEXTEND        = 6
	antlrParserFIXED32       = 7
	antlrParserFIXED64       = 8
	antlrParserFLOAT         = 9
	antlrParserIMPORT        = 10
	antlrParserINT32         = 11
	antlrParserINT64         = 12
	antlrParserMAP           = 13
	antlrParserMESSAGE       = 14
	antlrParserONEOF         = 15
	antlrParserOPTION        = 16
	antlrParserPACKAGE       = 17
	antlrParserPROTO3_DOUBLE = 18
	antlrParserPROTO3_SINGLE = 19
	antlrParserPUBLIC        = 20
	antlrParserREPEATED      = 21
	antlrParserRESERVED      = 22
	antlrParserRETURNS       = 23
	antlrParserRPC           = 24
	antlrParserSERVICE       = 25
	antlrParserSFIXED32      = 26
	antlrParserSFIXED64      = 27
	antlrParserSINT32        = 28
	antlrParserSINT64        = 29
	antlrParserSTREAM        = 30
	antlrParserSTRING        = 31
	antlrParserSYNTAX        = 32
	antlrParserTO            = 33
	antlrParserUINT32        = 34
	antlrParserUINT64        = 35
	antlrParserWEAK          = 36
	antlrParserLPAREN        = 37
	antlrParserRPAREN        = 38
	antlrParserLBRACE        = 39
	antlrParserRBRACE        = 40
	antlrParserLBRACK        = 41
	antlrParserRBRACK        = 42
	antlrParserLCHEVR        = 43
	antlrParserRCHEVR        = 44
	antlrParserSEMI          = 45
	antlrParserCOMMA         = 46
	antlrParserDOT           = 47
	antlrParserMINUS         = 48
	antlrParserPLUS          = 49
	antlrParserSLASH         = 50
	antlrParserEQ            = 51
	antlrParserIdent         = 52
	antlrParserIntLit        = 53
	antlrParserFloatLit      = 54
	antlrParserBoolLit       = 55
	antlrParserStrLit        = 56
	antlrParserWS            = 57
	antlrParserCOMMENT       = 58
	antlrParserLINE_COMMENT  = 59
)

// antlrParser rules.
const (
	antlrParserRULE_proto             = 0
	antlrParserRULE_syntax            = 1
	antlrParserRULE_imports           = 2
	antlrParserRULE_packageName       = 3
	antlrParserRULE_option            = 4
	antlrParserRULE_optionName        = 5
	antlrParserRULE_optionDecl        = 6
	antlrParserRULE_optionDeclVar     = 7
	antlrParserRULE_extend            = 8
	antlrParserRULE_extendDecl        = 9
	antlrParserRULE_extendName        = 10
	antlrParserRULE_message           = 11
	antlrParserRULE_messageDecl       = 12
	antlrParserRULE_enum              = 13
	antlrParserRULE_enumDecl          = 14
	antlrParserRULE_enumField         = 15
	antlrParserRULE_keyValue          = 16
	antlrParserRULE_service           = 17
	antlrParserRULE_rpcName           = 18
	antlrParserRULE_rpcParams         = 19
	antlrParserRULE_clientStreams     = 20
	antlrParserRULE_rpcReturns        = 21
	antlrParserRULE_serverStreams     = 22
	antlrParserRULE_rpc               = 23
	antlrParserRULE_reserved          = 24
	antlrParserRULE_ranges            = 25
	antlrParserRULE_rangeDef          = 26
	antlrParserRULE_fieldNames        = 27
	antlrParserRULE_fieldType         = 28
	antlrParserRULE_field             = 29
	antlrParserRULE_fieldOpts         = 30
	antlrParserRULE_oneof             = 31
	antlrParserRULE_oneofField        = 32
	antlrParserRULE_mapField          = 33
	antlrParserRULE_keyType           = 34
	antlrParserRULE_reservedWord      = 35
	antlrParserRULE_fullIdent         = 36
	antlrParserRULE_messageOrEnumType = 37
	antlrParserRULE_constant          = 38
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(antlrParserEOF, 0)
}

func (s *ProtoContext) AllImports() []IImportsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportsContext)(nil)).Elem())
	var tst = make([]IImportsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportsContext)
		}
	}

	return tst
}

func (s *ProtoContext) Imports(i int) IImportsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *ProtoContext) AllPackageName() []IPackageNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPackageNameContext)(nil)).Elem())
	var tst = make([]IPackageNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPackageNameContext)
		}
	}

	return tst
}

func (s *ProtoContext) PackageName(i int) IPackageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPackageNameContext)
}

func (s *ProtoContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *ProtoContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ProtoContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *ProtoContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *ProtoContext) AllEnum() []IEnumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumContext)(nil)).Elem())
	var tst = make([]IEnumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumContext)
		}
	}

	return tst
}

func (s *ProtoContext) Enum(i int) IEnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumContext)
}

func (s *ProtoContext) AllService() []IServiceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceContext)(nil)).Elem())
	var tst = make([]IServiceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceContext)
		}
	}

	return tst
}

func (s *ProtoContext) Service(i int) IServiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *ProtoContext) AllExtend() []IExtendContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtendContext)(nil)).Elem())
	var tst = make([]IExtendContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtendContext)
		}
	}

	return tst
}

func (s *ProtoContext) Extend(i int) IExtendContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtendContext)
}

func (s *ProtoContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(antlrParserSEMI)
}

func (s *ProtoContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, i)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *antlrParser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, antlrParserRULE_proto)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Syntax()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<antlrParserENUM)|(1<<antlrParserEXTEND)|(1<<antlrParserIMPORT)|(1<<antlrParserMESSAGE)|(1<<antlrParserOPTION)|(1<<antlrParserPACKAGE)|(1<<antlrParserSERVICE))) != 0) || _la == antlrParserSEMI {
		p.SetState(87)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case antlrParserIMPORT:
			{
				p.SetState(79)
				p.Imports()
			}

		case antlrParserPACKAGE:
			{
				p.SetState(80)
				p.PackageName()
			}

		case antlrParserOPTION:
			{
				p.SetState(81)
				p.Option()
			}

		case antlrParserMESSAGE:
			{
				p.SetState(82)
				p.Message()
			}

		case antlrParserENUM:
			{
				p.SetState(83)
				p.Enum()
			}

		case antlrParserSERVICE:
			{
				p.SetState(84)
				p.Service()
			}

		case antlrParserEXTEND:
			{
				p.SetState(85)
				p.Extend()
			}

		case antlrParserSEMI:
			{
				p.SetState(86)
				p.Match(antlrParserSEMI)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(antlrParserEOF)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(antlrParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *SyntaxContext) PROTO3_DOUBLE() antlr.TerminalNode {
	return s.GetToken(antlrParserPROTO3_DOUBLE, 0)
}

func (s *SyntaxContext) PROTO3_SINGLE() antlr.TerminalNode {
	return s.GetToken(antlrParserPROTO3_SINGLE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *antlrParser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, antlrParserRULE_syntax)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(antlrParserSYNTAX)
	}
	{
		p.SetState(95)
		p.Match(antlrParserEQ)
	}
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !(_la == antlrParserPROTO3_DOUBLE || _la == antlrParserPROTO3_SINGLE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(97)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IImportsContext is an interface to support dynamic dispatch.
type IImportsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportsContext differentiates from other interfaces.
	IsImportsContext()
}

type ImportsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportsContext() *ImportsContext {
	var p = new(ImportsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_imports
	return p
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	var p = new(ImportsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(antlrParserIMPORT, 0)
}

func (s *ImportsContext) StrLit() antlr.TerminalNode {
	return s.GetToken(antlrParserStrLit, 0)
}

func (s *ImportsContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *ImportsContext) WEAK() antlr.TerminalNode {
	return s.GetToken(antlrParserWEAK, 0)
}

func (s *ImportsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(antlrParserPUBLIC, 0)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterImports(s)
	}
}

func (s *ImportsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitImports(s)
	}
}

func (p *antlrParser) Imports() (localctx IImportsContext) {
	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, antlrParserRULE_imports)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(antlrParserIMPORT)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserPUBLIC || _la == antlrParserWEAK {
		{
			p.SetState(100)
			_la = p.GetTokenStream().LA(1)

			if !(_la == antlrParserPUBLIC || _la == antlrParserWEAK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(103)
		p.Match(antlrParserStrLit)
	}
	{
		p.SetState(104)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IPackageNameContext is an interface to support dynamic dispatch.
type IPackageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageNameContext differentiates from other interfaces.
	IsPackageNameContext()
}

type PackageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageNameContext() *PackageNameContext {
	var p = new(PackageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_packageName
	return p
}

func (*PackageNameContext) IsPackageNameContext() {}

func NewPackageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageNameContext {
	var p = new(PackageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_packageName

	return p
}

func (s *PackageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageNameContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(antlrParserPACKAGE, 0)
}

func (s *PackageNameContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageNameContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *PackageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterPackageName(s)
	}
}

func (s *PackageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitPackageName(s)
	}
}

func (p *antlrParser) PackageName() (localctx IPackageNameContext) {
	localctx = NewPackageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, antlrParserRULE_packageName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(antlrParserPACKAGE)
	}
	{
		p.SetState(107)
		p.FullIdent()
	}
	{
		p.SetState(108)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) OPTION() antlr.TerminalNode {
	return s.GetToken(antlrParserOPTION, 0)
}

func (s *OptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *OptionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *OptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionContext) OptionDecl() IOptionDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionDeclContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *antlrParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, antlrParserRULE_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(antlrParserOPTION)
	}
	{
		p.SetState(111)
		p.OptionName()
	}
	{
		p.SetState(112)
		p.Match(antlrParserEQ)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case antlrParserMINUS, antlrParserPLUS, antlrParserIdent, antlrParserIntLit, antlrParserFloatLit, antlrParserBoolLit, antlrParserStrLit:
		{
			p.SetState(113)
			p.Constant()
		}

	case antlrParserLBRACE:
		{
			p.SetState(114)
			p.OptionDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(117)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(antlrParserIdent)
}

func (s *OptionNameContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, i)
}

func (s *OptionNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(antlrParserLPAREN, 0)
}

func (s *OptionNameContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(antlrParserRPAREN, 0)
}

func (s *OptionNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(antlrParserDOT)
}

func (s *OptionNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserDOT, i)
}

func (s *OptionNameContext) AllReservedWord() []IReservedWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReservedWordContext)(nil)).Elem())
	var tst = make([]IReservedWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReservedWordContext)
		}
	}

	return tst
}

func (s *OptionNameContext) ReservedWord(i int) IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *antlrParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, antlrParserRULE_optionName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case antlrParserIdent:
		{
			p.SetState(119)
			p.Match(antlrParserIdent)
		}

	case antlrParserLPAREN:
		{
			p.SetState(120)
			p.Match(antlrParserLPAREN)
		}
		{
			p.SetState(121)
			p.FullIdent()
		}
		{
			p.SetState(122)
			p.Match(antlrParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserDOT {
		{
			p.SetState(126)
			p.Match(antlrParserDOT)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case antlrParserIdent:
			{
				p.SetState(127)
				p.Match(antlrParserIdent)
			}

		case antlrParserEXTEND, antlrParserMESSAGE, antlrParserOPTION, antlrParserPACKAGE, antlrParserRPC, antlrParserSERVICE, antlrParserSTREAM, antlrParserSTRING, antlrParserSYNTAX, antlrParserWEAK:
			{
				p.SetState(128)
				p.ReservedWord()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOptionDeclContext is an interface to support dynamic dispatch.
type IOptionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionDeclContext differentiates from other interfaces.
	IsOptionDeclContext()
}

type OptionDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionDeclContext() *OptionDeclContext {
	var p = new(OptionDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_optionDecl
	return p
}

func (*OptionDeclContext) IsOptionDeclContext() {}

func NewOptionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionDeclContext {
	var p = new(OptionDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_optionDecl

	return p
}

func (s *OptionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *OptionDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *OptionDeclContext) AllOptionDeclVar() []IOptionDeclVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionDeclVarContext)(nil)).Elem())
	var tst = make([]IOptionDeclVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionDeclVarContext)
		}
	}

	return tst
}

func (s *OptionDeclContext) OptionDeclVar(i int) IOptionDeclVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionDeclVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionDeclVarContext)
}

func (s *OptionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterOptionDecl(s)
	}
}

func (s *OptionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitOptionDecl(s)
	}
}

func (p *antlrParser) OptionDecl() (localctx IOptionDeclContext) {
	localctx = NewOptionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, antlrParserRULE_optionDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(antlrParserLBRACE)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserLPAREN || _la == antlrParserIdent {
		{
			p.SetState(137)
			p.OptionDeclVar()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(antlrParserRBRACE)
	}

	return localctx
}

// IOptionDeclVarContext is an interface to support dynamic dispatch.
type IOptionDeclVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionDeclVarContext differentiates from other interfaces.
	IsOptionDeclVarContext()
}

type OptionDeclVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionDeclVarContext() *OptionDeclVarContext {
	var p = new(OptionDeclVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_optionDeclVar
	return p
}

func (*OptionDeclVarContext) IsOptionDeclVarContext() {}

func NewOptionDeclVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionDeclVarContext {
	var p = new(OptionDeclVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_optionDeclVar

	return p
}

func (s *OptionDeclVarContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionDeclVarContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionDeclVarContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionDeclVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionDeclVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionDeclVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterOptionDeclVar(s)
	}
}

func (s *OptionDeclVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitOptionDeclVar(s)
	}
}

func (p *antlrParser) OptionDeclVar() (localctx IOptionDeclVarContext) {
	localctx = NewOptionDeclVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, antlrParserRULE_optionDeclVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.OptionName()
	}
	{
		p.SetState(146)
		p.Match(antlrParserT__0)
	}
	{
		p.SetState(147)
		p.Constant()
	}

	return localctx
}

// IExtendContext is an interface to support dynamic dispatch.
type IExtendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendContext differentiates from other interfaces.
	IsExtendContext()
}

type ExtendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendContext() *ExtendContext {
	var p = new(ExtendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_extend
	return p
}

func (*ExtendContext) IsExtendContext() {}

func NewExtendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendContext {
	var p = new(ExtendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_extend

	return p
}

func (s *ExtendContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(antlrParserEXTEND, 0)
}

func (s *ExtendContext) ExtendName() IExtendNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendNameContext)
}

func (s *ExtendContext) ExtendDecl() IExtendDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendDeclContext)
}

func (s *ExtendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterExtend(s)
	}
}

func (s *ExtendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitExtend(s)
	}
}

func (p *antlrParser) Extend() (localctx IExtendContext) {
	localctx = NewExtendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, antlrParserRULE_extend)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(antlrParserEXTEND)
	}
	{
		p.SetState(150)
		p.ExtendName()
	}
	{
		p.SetState(151)
		p.ExtendDecl()
	}

	return localctx
}

// IExtendDeclContext is an interface to support dynamic dispatch.
type IExtendDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendDeclContext differentiates from other interfaces.
	IsExtendDeclContext()
}

type ExtendDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendDeclContext() *ExtendDeclContext {
	var p = new(ExtendDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_extendDecl
	return p
}

func (*ExtendDeclContext) IsExtendDeclContext() {}

func NewExtendDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendDeclContext {
	var p = new(ExtendDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_extendDecl

	return p
}

func (s *ExtendDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *ExtendDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *ExtendDeclContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *ExtendDeclContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExtendDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterExtendDecl(s)
	}
}

func (s *ExtendDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitExtendDecl(s)
	}
}

func (p *antlrParser) ExtendDecl() (localctx IExtendDeclContext) {
	localctx = NewExtendDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, antlrParserRULE_extendDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(antlrParserLBRACE)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<antlrParserBOOL)|(1<<antlrParserBYTES)|(1<<antlrParserDOUBLE)|(1<<antlrParserEXTEND)|(1<<antlrParserFIXED32)|(1<<antlrParserFIXED64)|(1<<antlrParserFLOAT)|(1<<antlrParserINT32)|(1<<antlrParserINT64)|(1<<antlrParserMESSAGE)|(1<<antlrParserOPTION)|(1<<antlrParserPACKAGE)|(1<<antlrParserREPEATED)|(1<<antlrParserRPC)|(1<<antlrParserSERVICE)|(1<<antlrParserSFIXED32)|(1<<antlrParserSFIXED64)|(1<<antlrParserSINT32)|(1<<antlrParserSINT64)|(1<<antlrParserSTREAM)|(1<<antlrParserSTRING))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(antlrParserSYNTAX-32))|(1<<(antlrParserUINT32-32))|(1<<(antlrParserUINT64-32))|(1<<(antlrParserWEAK-32))|(1<<(antlrParserDOT-32))|(1<<(antlrParserIdent-32)))) != 0) {
		{
			p.SetState(154)
			p.Field()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
		p.Match(antlrParserRBRACE)
	}

	return localctx
}

// IExtendNameContext is an interface to support dynamic dispatch.
type IExtendNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendNameContext differentiates from other interfaces.
	IsExtendNameContext()
}

type ExtendNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendNameContext() *ExtendNameContext {
	var p = new(ExtendNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_extendName
	return p
}

func (*ExtendNameContext) IsExtendNameContext() {}

func NewExtendNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendNameContext {
	var p = new(ExtendNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_extendName

	return p
}

func (s *ExtendNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *ExtendNameContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ExtendNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterExtendName(s)
	}
}

func (s *ExtendNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitExtendName(s)
	}
}

func (p *antlrParser) ExtendName() (localctx IExtendNameContext) {
	localctx = NewExtendNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, antlrParserRULE_extendName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(162)
			p.Match(antlrParserIdent)
		}

	case 2:
		{
			p.SetState(163)
			p.FullIdent()
		}

	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(antlrParserMESSAGE, 0)
}

func (s *MessageContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *MessageContext) MessageDecl() IMessageDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageDeclContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *antlrParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, antlrParserRULE_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(antlrParserMESSAGE)
	}
	{
		p.SetState(167)
		p.Match(antlrParserIdent)
	}
	{
		p.SetState(168)
		p.MessageDecl()
	}

	return localctx
}

// IMessageDeclContext is an interface to support dynamic dispatch.
type IMessageDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageDeclContext differentiates from other interfaces.
	IsMessageDeclContext()
}

type MessageDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDeclContext() *MessageDeclContext {
	var p = new(MessageDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_messageDecl
	return p
}

func (*MessageDeclContext) IsMessageDeclContext() {}

func NewMessageDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDeclContext {
	var p = new(MessageDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_messageDecl

	return p
}

func (s *MessageDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *MessageDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *MessageDeclContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageDeclContext) AllEnum() []IEnumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumContext)(nil)).Elem())
	var tst = make([]IEnumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) Enum(i int) IEnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumContext)
}

func (s *MessageDeclContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageDeclContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *MessageDeclContext) AllOneof() []IOneofContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofContext)(nil)).Elem())
	var tst = make([]IOneofContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) Oneof(i int) IOneofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageDeclContext) AllMapField() []IMapFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapFieldContext)(nil)).Elem())
	var tst = make([]IMapFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapFieldContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) MapField(i int) IMapFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageDeclContext) AllReserved() []IReservedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReservedContext)(nil)).Elem())
	var tst = make([]IReservedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReservedContext)
		}
	}

	return tst
}

func (s *MessageDeclContext) Reserved(i int) IReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageDeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(antlrParserSEMI)
}

func (s *MessageDeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, i)
}

func (s *MessageDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterMessageDecl(s)
	}
}

func (s *MessageDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitMessageDecl(s)
	}
}

func (p *antlrParser) MessageDecl() (localctx IMessageDeclContext) {
	localctx = NewMessageDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, antlrParserRULE_messageDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(antlrParserLBRACE)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<antlrParserBOOL)|(1<<antlrParserBYTES)|(1<<antlrParserDOUBLE)|(1<<antlrParserENUM)|(1<<antlrParserEXTEND)|(1<<antlrParserFIXED32)|(1<<antlrParserFIXED64)|(1<<antlrParserFLOAT)|(1<<antlrParserINT32)|(1<<antlrParserINT64)|(1<<antlrParserMAP)|(1<<antlrParserMESSAGE)|(1<<antlrParserONEOF)|(1<<antlrParserOPTION)|(1<<antlrParserPACKAGE)|(1<<antlrParserREPEATED)|(1<<antlrParserRESERVED)|(1<<antlrParserRPC)|(1<<antlrParserSERVICE)|(1<<antlrParserSFIXED32)|(1<<antlrParserSFIXED64)|(1<<antlrParserSINT32)|(1<<antlrParserSINT64)|(1<<antlrParserSTREAM)|(1<<antlrParserSTRING))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(antlrParserSYNTAX-32))|(1<<(antlrParserUINT32-32))|(1<<(antlrParserUINT64-32))|(1<<(antlrParserWEAK-32))|(1<<(antlrParserSEMI-32))|(1<<(antlrParserDOT-32))|(1<<(antlrParserIdent-32)))) != 0) {
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(171)
				p.Field()
			}

		case 2:
			{
				p.SetState(172)
				p.Enum()
			}

		case 3:
			{
				p.SetState(173)
				p.Message()
			}

		case 4:
			{
				p.SetState(174)
				p.Option()
			}

		case 5:
			{
				p.SetState(175)
				p.Oneof()
			}

		case 6:
			{
				p.SetState(176)
				p.MapField()
			}

		case 7:
			{
				p.SetState(177)
				p.Reserved()
			}

		case 8:
			{
				p.SetState(178)
				p.Match(antlrParserSEMI)
			}

		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(antlrParserRBRACE)
	}

	return localctx
}

// IEnumContext is an interface to support dynamic dispatch.
type IEnumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumContext differentiates from other interfaces.
	IsEnumContext()
}

type EnumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumContext() *EnumContext {
	var p = new(EnumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_enum
	return p
}

func (*EnumContext) IsEnumContext() {}

func NewEnumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumContext {
	var p = new(EnumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_enum

	return p
}

func (s *EnumContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumContext) ENUM() antlr.TerminalNode {
	return s.GetToken(antlrParserENUM, 0)
}

func (s *EnumContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *EnumContext) EnumDecl() IEnumDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclContext)
}

func (s *EnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterEnum(s)
	}
}

func (s *EnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitEnum(s)
	}
}

func (p *antlrParser) Enum() (localctx IEnumContext) {
	localctx = NewEnumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, antlrParserRULE_enum)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(antlrParserENUM)
	}
	{
		p.SetState(187)
		p.Match(antlrParserIdent)
	}
	{
		p.SetState(188)
		p.EnumDecl()
	}

	return localctx
}

// IEnumDeclContext is an interface to support dynamic dispatch.
type IEnumDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDeclContext differentiates from other interfaces.
	IsEnumDeclContext()
}

type EnumDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDeclContext() *EnumDeclContext {
	var p = new(EnumDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_enumDecl
	return p
}

func (*EnumDeclContext) IsEnumDeclContext() {}

func NewEnumDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclContext {
	var p = new(EnumDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_enumDecl

	return p
}

func (s *EnumDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *EnumDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *EnumDeclContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *EnumDeclContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *EnumDeclContext) AllEnumField() []IEnumFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem())
	var tst = make([]IEnumFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumFieldContext)
		}
	}

	return tst
}

func (s *EnumDeclContext) EnumField(i int) IEnumFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumDeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(antlrParserSEMI)
}

func (s *EnumDeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, i)
}

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterEnumDecl(s)
	}
}

func (s *EnumDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitEnumDecl(s)
	}
}

func (p *antlrParser) EnumDecl() (localctx IEnumDeclContext) {
	localctx = NewEnumDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, antlrParserRULE_enumDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(antlrParserLBRACE)
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserOPTION || _la == antlrParserSEMI || _la == antlrParserIdent {
		p.SetState(194)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case antlrParserOPTION:
			{
				p.SetState(191)
				p.Option()
			}

		case antlrParserIdent:
			{
				p.SetState(192)
				p.EnumField()
			}

		case antlrParserSEMI:
			{
				p.SetState(193)
				p.Match(antlrParserSEMI)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
		p.Match(antlrParserRBRACE)
	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() antlr.TerminalNode {
	return s.GetToken(antlrParserIntLit, 0)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(antlrParserMINUS, 0)
}

func (s *EnumFieldContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACK, 0)
}

func (s *EnumFieldContext) AllKeyValue() []IKeyValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValueContext)(nil)).Elem())
	var tst = make([]IKeyValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValueContext)
		}
	}

	return tst
}

func (s *EnumFieldContext) KeyValue(i int) IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *EnumFieldContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACK, 0)
}

func (s *EnumFieldContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(antlrParserCOMMA)
}

func (s *EnumFieldContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserCOMMA, i)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *antlrParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, antlrParserRULE_enumField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(antlrParserIdent)
	}
	{
		p.SetState(202)
		p.Match(antlrParserEQ)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserMINUS {
		{
			p.SetState(203)
			p.Match(antlrParserMINUS)
		}

	}
	{
		p.SetState(206)
		p.Match(antlrParserIntLit)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserLBRACK {
		{
			p.SetState(207)
			p.Match(antlrParserLBRACK)
		}
		{
			p.SetState(208)
			p.KeyValue()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == antlrParserCOMMA {
			{
				p.SetState(209)
				p.Match(antlrParserCOMMA)
			}
			{
				p.SetState(210)
				p.KeyValue()
			}

			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(216)
			p.Match(antlrParserRBRACK)
		}

	}
	{
		p.SetState(220)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *KeyValueContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *KeyValueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (p *antlrParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, antlrParserRULE_keyValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.OptionName()
	}
	{
		p.SetState(223)
		p.Match(antlrParserEQ)
	}
	{
		p.SetState(224)
		p.Constant()
	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(antlrParserSERVICE, 0)
}

func (s *ServiceContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *ServiceContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *ServiceContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *ServiceContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *ServiceContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ServiceContext) AllRpc() []IRpcContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRpcContext)(nil)).Elem())
	var tst = make([]IRpcContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRpcContext)
		}
	}

	return tst
}

func (s *ServiceContext) Rpc(i int) IRpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(antlrParserSEMI)
}

func (s *ServiceContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, i)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *antlrParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, antlrParserRULE_service)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(antlrParserSERVICE)
	}
	{
		p.SetState(227)
		p.Match(antlrParserIdent)
	}
	{
		p.SetState(228)
		p.Match(antlrParserLBRACE)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(antlrParserOPTION-16))|(1<<(antlrParserRPC-16))|(1<<(antlrParserSEMI-16)))) != 0 {
		p.SetState(232)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case antlrParserOPTION:
			{
				p.SetState(229)
				p.Option()
			}

		case antlrParserRPC:
			{
				p.SetState(230)
				p.Rpc()
			}

		case antlrParserSEMI:
			{
				p.SetState(231)
				p.Match(antlrParserSEMI)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(237)
		p.Match(antlrParserRBRACE)
	}

	return localctx
}

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_rpcName
	return p
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (p *antlrParser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, antlrParserRULE_rpcName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(antlrParserIdent)
	}

	return localctx
}

// IRpcParamsContext is an interface to support dynamic dispatch.
type IRpcParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcParamsContext differentiates from other interfaces.
	IsRpcParamsContext()
}

type RpcParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcParamsContext() *RpcParamsContext {
	var p = new(RpcParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_rpcParams
	return p
}

func (*RpcParamsContext) IsRpcParamsContext() {}

func NewRpcParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcParamsContext {
	var p = new(RpcParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_rpcParams

	return p
}

func (s *RpcParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcParamsContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(antlrParserIdent)
}

func (s *RpcParamsContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, i)
}

func (s *RpcParamsContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(antlrParserDOT)
}

func (s *RpcParamsContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserDOT, i)
}

func (s *RpcParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterRpcParams(s)
	}
}

func (s *RpcParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitRpcParams(s)
	}
}

func (p *antlrParser) RpcParams() (localctx IRpcParamsContext) {
	localctx = NewRpcParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, antlrParserRULE_rpcParams)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(241)
				p.Match(antlrParserIdent)
			}
			{
				p.SetState(242)
				p.Match(antlrParserDOT)
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	{
		p.SetState(248)
		p.Match(antlrParserIdent)
	}

	return localctx
}

// IClientStreamsContext is an interface to support dynamic dispatch.
type IClientStreamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClientStreamsContext differentiates from other interfaces.
	IsClientStreamsContext()
}

type ClientStreamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClientStreamsContext() *ClientStreamsContext {
	var p = new(ClientStreamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_clientStreams
	return p
}

func (*ClientStreamsContext) IsClientStreamsContext() {}

func NewClientStreamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClientStreamsContext {
	var p = new(ClientStreamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_clientStreams

	return p
}

func (s *ClientStreamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ClientStreamsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(antlrParserSTREAM, 0)
}

func (s *ClientStreamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClientStreamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClientStreamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterClientStreams(s)
	}
}

func (s *ClientStreamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitClientStreams(s)
	}
}

func (p *antlrParser) ClientStreams() (localctx IClientStreamsContext) {
	localctx = NewClientStreamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, antlrParserRULE_clientStreams)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserSTREAM {
		{
			p.SetState(250)
			p.Match(antlrParserSTREAM)
		}

	}

	return localctx
}

// IRpcReturnsContext is an interface to support dynamic dispatch.
type IRpcReturnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcReturnsContext differentiates from other interfaces.
	IsRpcReturnsContext()
}

type RpcReturnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcReturnsContext() *RpcReturnsContext {
	var p = new(RpcReturnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_rpcReturns
	return p
}

func (*RpcReturnsContext) IsRpcReturnsContext() {}

func NewRpcReturnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcReturnsContext {
	var p = new(RpcReturnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_rpcReturns

	return p
}

func (s *RpcReturnsContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcReturnsContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(antlrParserIdent)
}

func (s *RpcReturnsContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, i)
}

func (s *RpcReturnsContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(antlrParserDOT)
}

func (s *RpcReturnsContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserDOT, i)
}

func (s *RpcReturnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcReturnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcReturnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterRpcReturns(s)
	}
}

func (s *RpcReturnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitRpcReturns(s)
	}
}

func (p *antlrParser) RpcReturns() (localctx IRpcReturnsContext) {
	localctx = NewRpcReturnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, antlrParserRULE_rpcReturns)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(253)
				p.Match(antlrParserIdent)
			}
			{
				p.SetState(254)
				p.Match(antlrParserDOT)
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	{
		p.SetState(260)
		p.Match(antlrParserIdent)
	}

	return localctx
}

// IServerStreamsContext is an interface to support dynamic dispatch.
type IServerStreamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServerStreamsContext differentiates from other interfaces.
	IsServerStreamsContext()
}

type ServerStreamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServerStreamsContext() *ServerStreamsContext {
	var p = new(ServerStreamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_serverStreams
	return p
}

func (*ServerStreamsContext) IsServerStreamsContext() {}

func NewServerStreamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServerStreamsContext {
	var p = new(ServerStreamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_serverStreams

	return p
}

func (s *ServerStreamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ServerStreamsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(antlrParserSTREAM, 0)
}

func (s *ServerStreamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServerStreamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServerStreamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterServerStreams(s)
	}
}

func (s *ServerStreamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitServerStreams(s)
	}
}

func (p *antlrParser) ServerStreams() (localctx IServerStreamsContext) {
	localctx = NewServerStreamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, antlrParserRULE_serverStreams)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserSTREAM {
		{
			p.SetState(262)
			p.Match(antlrParserSTREAM)
		}

	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(antlrParserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(antlrParserLPAREN)
}

func (s *RpcContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserLPAREN, i)
}

func (s *RpcContext) ClientStreams() IClientStreamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClientStreamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClientStreamsContext)
}

func (s *RpcContext) RpcParams() IRpcParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcParamsContext)
}

func (s *RpcContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(antlrParserRPAREN)
}

func (s *RpcContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserRPAREN, i)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(antlrParserRETURNS, 0)
}

func (s *RpcContext) ServerStreams() IServerStreamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServerStreamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServerStreamsContext)
}

func (s *RpcContext) RpcReturns() IRpcReturnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcReturnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcReturnsContext)
}

func (s *RpcContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(antlrParserSEMI)
}

func (s *RpcContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, i)
}

func (s *RpcContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(antlrParserDOT)
}

func (s *RpcContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserDOT, i)
}

func (s *RpcContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *RpcContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *RpcContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *RpcContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitRpc(s)
	}
}

func (p *antlrParser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, antlrParserRULE_rpc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(antlrParserRPC)
	}
	{
		p.SetState(266)
		p.RpcName()
	}
	{
		p.SetState(267)
		p.Match(antlrParserLPAREN)
	}
	{
		p.SetState(268)
		p.ClientStreams()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserDOT {
		{
			p.SetState(269)
			p.Match(antlrParserDOT)
		}

	}
	{
		p.SetState(272)
		p.RpcParams()
	}
	{
		p.SetState(273)
		p.Match(antlrParserRPAREN)
	}
	{
		p.SetState(274)
		p.Match(antlrParserRETURNS)
	}
	{
		p.SetState(275)
		p.Match(antlrParserLPAREN)
	}
	{
		p.SetState(276)
		p.ServerStreams()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserDOT {
		{
			p.SetState(277)
			p.Match(antlrParserDOT)
		}

	}
	{
		p.SetState(280)
		p.RpcReturns()
	}
	{
		p.SetState(281)
		p.Match(antlrParserRPAREN)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case antlrParserLBRACE:
		{
			p.SetState(282)
			p.Match(antlrParserLBRACE)
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == antlrParserOPTION || _la == antlrParserSEMI {
			p.SetState(285)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case antlrParserOPTION:
				{
					p.SetState(283)
					p.Option()
				}

			case antlrParserSEMI:
				{
					p.SetState(284)
					p.Match(antlrParserSEMI)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(289)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(290)
			p.Match(antlrParserRBRACE)
		}

	case antlrParserSEMI:
		{
			p.SetState(291)
			p.Match(antlrParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(antlrParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) FieldNames() IFieldNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitReserved(s)
	}
}

func (p *antlrParser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, antlrParserRULE_reserved)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(antlrParserRESERVED)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case antlrParserIntLit:
		{
			p.SetState(295)
			p.Ranges()
		}

	case antlrParserStrLit:
		{
			p.SetState(296)
			p.FieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(299)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRangeDef() []IRangeDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRangeDefContext)(nil)).Elem())
	var tst = make([]IRangeDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRangeDefContext)
		}
	}

	return tst
}

func (s *RangesContext) RangeDef(i int) IRangeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRangeDefContext)
}

func (s *RangesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(antlrParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *antlrParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, antlrParserRULE_ranges)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.RangeDef()
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserCOMMA {
		{
			p.SetState(302)
			p.Match(antlrParserCOMMA)
		}
		{
			p.SetState(303)
			p.RangeDef()
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRangeDefContext is an interface to support dynamic dispatch.
type IRangeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeDefContext differentiates from other interfaces.
	IsRangeDefContext()
}

type RangeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeDefContext() *RangeDefContext {
	var p = new(RangeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_rangeDef
	return p
}

func (*RangeDefContext) IsRangeDefContext() {}

func NewRangeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeDefContext {
	var p = new(RangeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_rangeDef

	return p
}

func (s *RangeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeDefContext) AllIntLit() []antlr.TerminalNode {
	return s.GetTokens(antlrParserIntLit)
}

func (s *RangeDefContext) IntLit(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserIntLit, i)
}

func (s *RangeDefContext) TO() antlr.TerminalNode {
	return s.GetToken(antlrParserTO, 0)
}

func (s *RangeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterRangeDef(s)
	}
}

func (s *RangeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitRangeDef(s)
	}
}

func (p *antlrParser) RangeDef() (localctx IRangeDefContext) {
	localctx = NewRangeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, antlrParserRULE_rangeDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(antlrParserIntLit)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Match(antlrParserIntLit)
		}
		{
			p.SetState(311)
			p.Match(antlrParserTO)
		}
		{
			p.SetState(312)
			p.Match(antlrParserIntLit)
		}

	}

	return localctx
}

// IFieldNamesContext is an interface to support dynamic dispatch.
type IFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNamesContext differentiates from other interfaces.
	IsFieldNamesContext()
}

type FieldNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNamesContext() *FieldNamesContext {
	var p = new(FieldNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_fieldNames
	return p
}

func (*FieldNamesContext) IsFieldNamesContext() {}

func NewFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNamesContext {
	var p = new(FieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_fieldNames

	return p
}

func (s *FieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNamesContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(antlrParserStrLit)
}

func (s *FieldNamesContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserStrLit, i)
}

func (s *FieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(antlrParserCOMMA)
}

func (s *FieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserCOMMA, i)
}

func (s *FieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterFieldNames(s)
	}
}

func (s *FieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitFieldNames(s)
	}
}

func (p *antlrParser) FieldNames() (localctx IFieldNamesContext) {
	localctx = NewFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, antlrParserRULE_fieldNames)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(antlrParserStrLit)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserCOMMA {
		{
			p.SetState(316)
			p.Match(antlrParserCOMMA)
		}
		{
			p.SetState(317)
			p.Match(antlrParserStrLit)
		}

		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(antlrParserDOUBLE, 0)
}

func (s *FieldTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(antlrParserFLOAT, 0)
}

func (s *FieldTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(antlrParserINT32, 0)
}

func (s *FieldTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(antlrParserINT64, 0)
}

func (s *FieldTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(antlrParserUINT32, 0)
}

func (s *FieldTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(antlrParserUINT64, 0)
}

func (s *FieldTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(antlrParserSINT32, 0)
}

func (s *FieldTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(antlrParserSINT64, 0)
}

func (s *FieldTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(antlrParserFIXED32, 0)
}

func (s *FieldTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(antlrParserFIXED64, 0)
}

func (s *FieldTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(antlrParserSFIXED32, 0)
}

func (s *FieldTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(antlrParserSFIXED64, 0)
}

func (s *FieldTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(antlrParserBOOL, 0)
}

func (s *FieldTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(antlrParserSTRING, 0)
}

func (s *FieldTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(antlrParserBYTES, 0)
}

func (s *FieldTypeContext) MessageOrEnumType() IMessageOrEnumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageOrEnumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumTypeContext)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *antlrParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, antlrParserRULE_fieldType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<antlrParserBOOL)|(1<<antlrParserBYTES)|(1<<antlrParserDOUBLE)|(1<<antlrParserFIXED32)|(1<<antlrParserFIXED64)|(1<<antlrParserFLOAT)|(1<<antlrParserINT32)|(1<<antlrParserINT64)|(1<<antlrParserSFIXED32)|(1<<antlrParserSFIXED64)|(1<<antlrParserSINT32)|(1<<antlrParserSINT64)|(1<<antlrParserSTRING))) != 0) || _la == antlrParserUINT32 || _la == antlrParserUINT64) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.MessageOrEnumType()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *FieldContext) IntLit() antlr.TerminalNode {
	return s.GetToken(antlrParserIntLit, 0)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *FieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *FieldContext) ReservedWord() IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(antlrParserREPEATED, 0)
}

func (s *FieldContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACK, 0)
}

func (s *FieldContext) FieldOpts() IFieldOptsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptsContext)
}

func (s *FieldContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACK, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *antlrParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, antlrParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserREPEATED {
		{
			p.SetState(327)
			p.Match(antlrParserREPEATED)
		}

	}
	{
		p.SetState(330)
		p.FieldType()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case antlrParserIdent:
		{
			p.SetState(331)
			p.Match(antlrParserIdent)
		}

	case antlrParserEXTEND, antlrParserMESSAGE, antlrParserOPTION, antlrParserPACKAGE, antlrParserRPC, antlrParserSERVICE, antlrParserSTREAM, antlrParserSTRING, antlrParserSYNTAX, antlrParserWEAK:
		{
			p.SetState(332)
			p.ReservedWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(335)
		p.Match(antlrParserEQ)
	}
	{
		p.SetState(336)
		p.Match(antlrParserIntLit)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserLBRACK {
		{
			p.SetState(337)
			p.Match(antlrParserLBRACK)
		}
		{
			p.SetState(338)
			p.FieldOpts()
		}
		{
			p.SetState(339)
			p.Match(antlrParserRBRACK)
		}

	}
	{
		p.SetState(343)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IFieldOptsContext is an interface to support dynamic dispatch.
type IFieldOptsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptsContext differentiates from other interfaces.
	IsFieldOptsContext()
}

type FieldOptsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptsContext() *FieldOptsContext {
	var p = new(FieldOptsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_fieldOpts
	return p
}

func (*FieldOptsContext) IsFieldOptsContext() {}

func NewFieldOptsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptsContext {
	var p = new(FieldOptsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_fieldOpts

	return p
}

func (s *FieldOptsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptsContext) AllKeyValue() []IKeyValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValueContext)(nil)).Elem())
	var tst = make([]IKeyValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValueContext)
		}
	}

	return tst
}

func (s *FieldOptsContext) KeyValue(i int) IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *FieldOptsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(antlrParserCOMMA)
}

func (s *FieldOptsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserCOMMA, i)
}

func (s *FieldOptsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterFieldOpts(s)
	}
}

func (s *FieldOptsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitFieldOpts(s)
	}
}

func (p *antlrParser) FieldOpts() (localctx IFieldOptsContext) {
	localctx = NewFieldOptsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, antlrParserRULE_fieldOpts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.KeyValue()
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserCOMMA {
		{
			p.SetState(346)
			p.Match(antlrParserCOMMA)
		}
		{
			p.SetState(347)
			p.KeyValue()
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(antlrParserONEOF, 0)
}

func (s *OneofContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *OneofContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACE, 0)
}

func (s *OneofContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACE, 0)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem())
	var tst = make([]IOneofFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofFieldContext)
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(antlrParserSEMI)
}

func (s *OneofContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, i)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitOneof(s)
	}
}

func (p *antlrParser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, antlrParserRULE_oneof)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(antlrParserONEOF)
	}
	{
		p.SetState(354)
		p.Match(antlrParserIdent)
	}
	{
		p.SetState(355)
		p.Match(antlrParserLBRACE)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<antlrParserBOOL)|(1<<antlrParserBYTES)|(1<<antlrParserDOUBLE)|(1<<antlrParserEXTEND)|(1<<antlrParserFIXED32)|(1<<antlrParserFIXED64)|(1<<antlrParserFLOAT)|(1<<antlrParserINT32)|(1<<antlrParserINT64)|(1<<antlrParserMESSAGE)|(1<<antlrParserOPTION)|(1<<antlrParserPACKAGE)|(1<<antlrParserRPC)|(1<<antlrParserSERVICE)|(1<<antlrParserSFIXED32)|(1<<antlrParserSFIXED64)|(1<<antlrParserSINT32)|(1<<antlrParserSINT64)|(1<<antlrParserSTREAM)|(1<<antlrParserSTRING))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(antlrParserSYNTAX-32))|(1<<(antlrParserUINT32-32))|(1<<(antlrParserUINT64-32))|(1<<(antlrParserWEAK-32))|(1<<(antlrParserSEMI-32))|(1<<(antlrParserDOT-32))|(1<<(antlrParserIdent-32)))) != 0) {
		p.SetState(358)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case antlrParserBOOL, antlrParserBYTES, antlrParserDOUBLE, antlrParserEXTEND, antlrParserFIXED32, antlrParserFIXED64, antlrParserFLOAT, antlrParserINT32, antlrParserINT64, antlrParserMESSAGE, antlrParserOPTION, antlrParserPACKAGE, antlrParserRPC, antlrParserSERVICE, antlrParserSFIXED32, antlrParserSFIXED64, antlrParserSINT32, antlrParserSINT64, antlrParserSTREAM, antlrParserSTRING, antlrParserSYNTAX, antlrParserUINT32, antlrParserUINT64, antlrParserWEAK, antlrParserDOT, antlrParserIdent:
			{
				p.SetState(356)
				p.OneofField()
			}

		case antlrParserSEMI:
			{
				p.SetState(357)
				p.Match(antlrParserSEMI)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(363)
		p.Match(antlrParserRBRACE)
	}

	return localctx
}

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *OneofFieldContext) IntLit() antlr.TerminalNode {
	return s.GetToken(antlrParserIntLit, 0)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *OneofFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *OneofFieldContext) ReservedWord() IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *OneofFieldContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACK, 0)
}

func (s *OneofFieldContext) FieldOpts() IFieldOptsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptsContext)
}

func (s *OneofFieldContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACK, 0)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (p *antlrParser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, antlrParserRULE_oneofField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.FieldType()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case antlrParserIdent:
		{
			p.SetState(366)
			p.Match(antlrParserIdent)
		}

	case antlrParserEXTEND, antlrParserMESSAGE, antlrParserOPTION, antlrParserPACKAGE, antlrParserRPC, antlrParserSERVICE, antlrParserSTREAM, antlrParserSTRING, antlrParserSYNTAX, antlrParserWEAK:
		{
			p.SetState(367)
			p.ReservedWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(370)
		p.Match(antlrParserEQ)
	}
	{
		p.SetState(371)
		p.Match(antlrParserIntLit)
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserLBRACK {
		{
			p.SetState(372)
			p.Match(antlrParserLBRACK)
		}
		{
			p.SetState(373)
			p.FieldOpts()
		}
		{
			p.SetState(374)
			p.Match(antlrParserRBRACK)
		}

	}
	{
		p.SetState(378)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(antlrParserMAP, 0)
}

func (s *MapFieldContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(antlrParserLCHEVR, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(antlrParserCOMMA, 0)
}

func (s *MapFieldContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *MapFieldContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(antlrParserRCHEVR, 0)
}

func (s *MapFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, 0)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(antlrParserEQ, 0)
}

func (s *MapFieldContext) IntLit() antlr.TerminalNode {
	return s.GetToken(antlrParserIntLit, 0)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(antlrParserSEMI, 0)
}

func (s *MapFieldContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserLBRACK, 0)
}

func (s *MapFieldContext) FieldOpts() IFieldOptsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptsContext)
}

func (s *MapFieldContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(antlrParserRBRACK, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *antlrParser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, antlrParserRULE_mapField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(antlrParserMAP)
	}
	{
		p.SetState(381)
		p.Match(antlrParserLCHEVR)
	}
	{
		p.SetState(382)
		p.KeyType()
	}
	{
		p.SetState(383)
		p.Match(antlrParserCOMMA)
	}
	{
		p.SetState(384)
		p.FieldType()
	}
	{
		p.SetState(385)
		p.Match(antlrParserRCHEVR)
	}
	{
		p.SetState(386)
		p.Match(antlrParserIdent)
	}
	{
		p.SetState(387)
		p.Match(antlrParserEQ)
	}
	{
		p.SetState(388)
		p.Match(antlrParserIntLit)
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserLBRACK {
		{
			p.SetState(389)
			p.Match(antlrParserLBRACK)
		}
		{
			p.SetState(390)
			p.FieldOpts()
		}
		{
			p.SetState(391)
			p.Match(antlrParserRBRACK)
		}

	}
	{
		p.SetState(395)
		p.Match(antlrParserSEMI)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(antlrParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(antlrParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(antlrParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(antlrParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(antlrParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(antlrParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(antlrParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(antlrParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(antlrParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(antlrParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(antlrParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(antlrParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *antlrParser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, antlrParserRULE_keyType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<antlrParserBOOL)|(1<<antlrParserFIXED32)|(1<<antlrParserFIXED64)|(1<<antlrParserINT32)|(1<<antlrParserINT64)|(1<<antlrParserSFIXED32)|(1<<antlrParserSFIXED64)|(1<<antlrParserSINT32)|(1<<antlrParserSINT64)|(1<<antlrParserSTRING))) != 0) || _la == antlrParserUINT32 || _la == antlrParserUINT64) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReservedWordContext is an interface to support dynamic dispatch.
type IReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedWordContext differentiates from other interfaces.
	IsReservedWordContext()
}

type ReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedWordContext() *ReservedWordContext {
	var p = new(ReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_reservedWord
	return p
}

func (*ReservedWordContext) IsReservedWordContext() {}

func NewReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedWordContext {
	var p = new(ReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_reservedWord

	return p
}

func (s *ReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedWordContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(antlrParserMESSAGE, 0)
}

func (s *ReservedWordContext) OPTION() antlr.TerminalNode {
	return s.GetToken(antlrParserOPTION, 0)
}

func (s *ReservedWordContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(antlrParserPACKAGE, 0)
}

func (s *ReservedWordContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(antlrParserSERVICE, 0)
}

func (s *ReservedWordContext) STREAM() antlr.TerminalNode {
	return s.GetToken(antlrParserSTREAM, 0)
}

func (s *ReservedWordContext) STRING() antlr.TerminalNode {
	return s.GetToken(antlrParserSTRING, 0)
}

func (s *ReservedWordContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(antlrParserSYNTAX, 0)
}

func (s *ReservedWordContext) WEAK() antlr.TerminalNode {
	return s.GetToken(antlrParserWEAK, 0)
}

func (s *ReservedWordContext) RPC() antlr.TerminalNode {
	return s.GetToken(antlrParserRPC, 0)
}

func (s *ReservedWordContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(antlrParserEXTEND, 0)
}

func (s *ReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterReservedWord(s)
	}
}

func (s *ReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitReservedWord(s)
	}
}

func (p *antlrParser) ReservedWord() (localctx IReservedWordContext) {
	localctx = NewReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, antlrParserRULE_reservedWord)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(antlrParserEXTEND-6))|(1<<(antlrParserMESSAGE-6))|(1<<(antlrParserOPTION-6))|(1<<(antlrParserPACKAGE-6))|(1<<(antlrParserRPC-6))|(1<<(antlrParserSERVICE-6))|(1<<(antlrParserSTREAM-6))|(1<<(antlrParserSTRING-6))|(1<<(antlrParserSYNTAX-6))|(1<<(antlrParserWEAK-6)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(antlrParserIdent)
}

func (s *FullIdentContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, i)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(antlrParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *antlrParser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, antlrParserRULE_fullIdent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Match(antlrParserIdent)
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == antlrParserDOT {
		{
			p.SetState(402)
			p.Match(antlrParserDOT)
		}
		{
			p.SetState(403)
			p.Match(antlrParserIdent)
		}

		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageOrEnumTypeContext is an interface to support dynamic dispatch.
type IMessageOrEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageOrEnumTypeContext differentiates from other interfaces.
	IsMessageOrEnumTypeContext()
}

type MessageOrEnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumTypeContext() *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_messageOrEnumType
	return p
}

func (*MessageOrEnumTypeContext) IsMessageOrEnumTypeContext() {}

func NewMessageOrEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_messageOrEnumType

	return p
}

func (s *MessageOrEnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(antlrParserIdent)
}

func (s *MessageOrEnumTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserIdent, i)
}

func (s *MessageOrEnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(antlrParserDOT)
}

func (s *MessageOrEnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(antlrParserDOT, i)
}

func (s *MessageOrEnumTypeContext) AllReservedWord() []IReservedWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReservedWordContext)(nil)).Elem())
	var tst = make([]IReservedWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReservedWordContext)
		}
	}

	return tst
}

func (s *MessageOrEnumTypeContext) ReservedWord(i int) IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *MessageOrEnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterMessageOrEnumType(s)
	}
}

func (s *MessageOrEnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitMessageOrEnumType(s)
	}
}

func (p *antlrParser) MessageOrEnumType() (localctx IMessageOrEnumTypeContext) {
	localctx = NewMessageOrEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, antlrParserRULE_messageOrEnumType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == antlrParserDOT {
		{
			p.SetState(409)
			p.Match(antlrParserDOT)
		}

	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(414)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case antlrParserIdent:
				{
					p.SetState(412)
					p.Match(antlrParserIdent)
				}

			case antlrParserEXTEND, antlrParserMESSAGE, antlrParserOPTION, antlrParserPACKAGE, antlrParserRPC, antlrParserSERVICE, antlrParserSTREAM, antlrParserSTRING, antlrParserSYNTAX, antlrParserWEAK:
				{
					p.SetState(413)
					p.ReservedWord()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(416)
				p.Match(antlrParserDOT)
			}

		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}
	{
		p.SetState(422)
		p.Match(antlrParserIdent)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = antlrParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = antlrParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() antlr.TerminalNode {
	return s.GetToken(antlrParserIntLit, 0)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(antlrParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(antlrParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() antlr.TerminalNode {
	return s.GetToken(antlrParserFloatLit, 0)
}

func (s *ConstantContext) StrLit() antlr.TerminalNode {
	return s.GetToken(antlrParserStrLit, 0)
}

func (s *ConstantContext) BoolLit() antlr.TerminalNode {
	return s.GetToken(antlrParserBoolLit, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(antlrListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *antlrParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, antlrParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == antlrParserMINUS || _la == antlrParserPLUS {
			{
				p.SetState(425)
				_la = p.GetTokenStream().LA(1)

				if !(_la == antlrParserMINUS || _la == antlrParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(428)
			p.Match(antlrParserIntLit)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == antlrParserMINUS || _la == antlrParserPLUS {
			{
				p.SetState(429)
				_la = p.GetTokenStream().LA(1)

				if !(_la == antlrParserMINUS || _la == antlrParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(432)
			p.Match(antlrParserFloatLit)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(433)
			_la = p.GetTokenStream().LA(1)

			if !(_la == antlrParserBoolLit || _la == antlrParserStrLit) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}
