// Code generated from ../../grammar/LlamaLang.g4 by ANTLR 4.9.1. DO NOT EDIT.

package antlr // LlamaLang
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 337,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 2, 3, 2, 7,
	2, 89, 10, 2, 12, 2, 14, 2, 92, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7,
	3, 99, 10, 3, 12, 3, 14, 3, 102, 11, 3, 3, 4, 3, 4, 3, 4, 7, 4, 107, 10,
	4, 12, 4, 14, 4, 110, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 124, 10, 6, 3, 7, 3, 7, 5, 7, 128,
	10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 6, 8, 135, 10, 8, 13, 8, 14, 8, 136,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 143, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 148,
	10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 5, 13, 157, 10,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 165, 10, 15, 3, 16,
	3, 16, 3, 16, 5, 16, 170, 10, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 5, 19, 181, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21,
	5, 21, 187, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 193, 10, 22, 12,
	22, 14, 22, 196, 11, 22, 5, 22, 198, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 209, 10, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 7, 24, 226, 10, 24, 12, 24, 14, 24, 229, 11, 24, 3, 25, 3,
	25, 3, 25, 5, 25, 234, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 240,
	10, 25, 7, 25, 242, 10, 25, 12, 25, 14, 25, 245, 11, 25, 3, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 251, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 257,
	10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	5, 28, 268, 10, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3,
	31, 5, 31, 278, 10, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 5, 34,
	286, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	37, 5, 37, 297, 10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	5, 39, 306, 10, 39, 5, 39, 308, 10, 39, 3, 39, 5, 39, 311, 10, 39, 3, 39,
	5, 39, 314, 10, 39, 5, 39, 316, 10, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 329, 10, 41, 3, 41,
	3, 41, 5, 41, 333, 10, 41, 3, 42, 3, 42, 3, 42, 2, 4, 46, 48, 43, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 2, 10, 4, 2, 29, 34, 36, 40, 4, 2, 30, 34, 39, 40, 4, 2, 29, 29,
	36, 38, 3, 2, 23, 28, 5, 2, 17, 18, 35, 37, 39, 40, 3, 2, 42, 44, 3, 2,
	45, 46, 3, 3, 14, 14, 2, 341, 2, 90, 3, 2, 2, 2, 4, 95, 3, 2, 2, 2, 6,
	103, 3, 2, 2, 2, 8, 111, 3, 2, 2, 2, 10, 118, 3, 2, 2, 2, 12, 125, 3, 2,
	2, 2, 14, 134, 3, 2, 2, 2, 16, 142, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2, 20,
	149, 3, 2, 2, 2, 22, 151, 3, 2, 2, 2, 24, 156, 3, 2, 2, 2, 26, 160, 3,
	2, 2, 2, 28, 162, 3, 2, 2, 2, 30, 169, 3, 2, 2, 2, 32, 171, 3, 2, 2, 2,
	34, 174, 3, 2, 2, 2, 36, 180, 3, 2, 2, 2, 38, 182, 3, 2, 2, 2, 40, 186,
	3, 2, 2, 2, 42, 188, 3, 2, 2, 2, 44, 201, 3, 2, 2, 2, 46, 208, 3, 2, 2,
	2, 48, 233, 3, 2, 2, 2, 50, 250, 3, 2, 2, 2, 52, 252, 3, 2, 2, 2, 54, 267,
	3, 2, 2, 2, 56, 269, 3, 2, 2, 2, 58, 271, 3, 2, 2, 2, 60, 277, 3, 2, 2,
	2, 62, 279, 3, 2, 2, 2, 64, 281, 3, 2, 2, 2, 66, 285, 3, 2, 2, 2, 68, 287,
	3, 2, 2, 2, 70, 291, 3, 2, 2, 2, 72, 293, 3, 2, 2, 2, 74, 298, 3, 2, 2,
	2, 76, 300, 3, 2, 2, 2, 78, 319, 3, 2, 2, 2, 80, 332, 3, 2, 2, 2, 82, 334,
	3, 2, 2, 2, 84, 89, 5, 8, 5, 2, 85, 86, 5, 10, 6, 2, 86, 87, 5, 82, 42,
	2, 87, 89, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 92,
	3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 93, 3, 2, 2, 2,
	92, 90, 3, 2, 2, 2, 93, 94, 5, 82, 42, 2, 94, 3, 3, 2, 2, 2, 95, 100, 7,
	5, 2, 2, 96, 97, 7, 13, 2, 2, 97, 99, 7, 5, 2, 2, 98, 96, 3, 2, 2, 2, 99,
	102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 5, 3,
	2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 108, 5, 46, 24, 2, 104, 105, 7, 13,
	2, 2, 105, 107, 5, 46, 24, 2, 106, 104, 3, 2, 2, 2, 107, 110, 3, 2, 2,
	2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 7, 3, 2, 2, 2, 110,
	108, 3, 2, 2, 2, 111, 112, 7, 3, 2, 2, 112, 113, 7, 5, 2, 2, 113, 114,
	5, 38, 20, 2, 114, 115, 7, 41, 2, 2, 115, 116, 5, 30, 16, 2, 116, 117,
	5, 12, 7, 2, 117, 9, 3, 2, 2, 2, 118, 119, 7, 5, 2, 2, 119, 120, 7, 15,
	2, 2, 120, 123, 5, 30, 16, 2, 121, 122, 7, 12, 2, 2, 122, 124, 5, 6, 4,
	2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 11, 3, 2, 2, 2, 125,
	127, 7, 8, 2, 2, 126, 128, 5, 14, 8, 2, 127, 126, 3, 2, 2, 2, 127, 128,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 7, 9, 2, 2, 130, 13, 3, 2,
	2, 2, 131, 132, 5, 16, 9, 2, 132, 133, 5, 82, 42, 2, 133, 135, 3, 2, 2,
	2, 134, 131, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 15, 3, 2, 2, 2, 138, 143, 5, 10, 6, 2, 139, 143,
	5, 18, 10, 2, 140, 143, 5, 28, 15, 2, 141, 143, 5, 12, 7, 2, 142, 138,
	3, 2, 2, 2, 142, 139, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2,
	2, 2, 143, 17, 3, 2, 2, 2, 144, 148, 5, 20, 11, 2, 145, 148, 5, 22, 12,
	2, 146, 148, 5, 26, 14, 2, 147, 144, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2,
	147, 146, 3, 2, 2, 2, 148, 19, 3, 2, 2, 2, 149, 150, 5, 46, 24, 2, 150,
	21, 3, 2, 2, 2, 151, 152, 7, 5, 2, 2, 152, 153, 5, 24, 13, 2, 153, 154,
	5, 6, 4, 2, 154, 23, 3, 2, 2, 2, 155, 157, 9, 2, 2, 2, 156, 155, 3, 2,
	2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 7, 12, 2, 2,
	159, 25, 3, 2, 2, 2, 160, 161, 7, 14, 2, 2, 161, 27, 3, 2, 2, 2, 162, 164,
	7, 4, 2, 2, 163, 165, 5, 46, 24, 2, 164, 163, 3, 2, 2, 2, 164, 165, 3,
	2, 2, 2, 165, 29, 3, 2, 2, 2, 166, 170, 5, 36, 19, 2, 167, 170, 5, 32,
	17, 2, 168, 170, 5, 34, 18, 2, 169, 166, 3, 2, 2, 2, 169, 167, 3, 2, 2,
	2, 169, 168, 3, 2, 2, 2, 170, 31, 3, 2, 2, 2, 171, 172, 7, 39, 2, 2, 172,
	173, 5, 30, 16, 2, 173, 33, 3, 2, 2, 2, 174, 175, 7, 10, 2, 2, 175, 176,
	7, 11, 2, 2, 176, 177, 5, 36, 19, 2, 177, 35, 3, 2, 2, 2, 178, 181, 7,
	5, 2, 2, 179, 181, 5, 68, 35, 2, 180, 178, 3, 2, 2, 2, 180, 179, 3, 2,
	2, 2, 181, 37, 3, 2, 2, 2, 182, 183, 5, 42, 22, 2, 183, 39, 3, 2, 2, 2,
	184, 187, 5, 42, 22, 2, 185, 187, 5, 30, 16, 2, 186, 184, 3, 2, 2, 2, 186,
	185, 3, 2, 2, 2, 187, 41, 3, 2, 2, 2, 188, 197, 7, 6, 2, 2, 189, 194, 5,
	44, 23, 2, 190, 191, 7, 13, 2, 2, 191, 193, 5, 44, 23, 2, 192, 190, 3,
	2, 2, 2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2,
	2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 189, 3, 2, 2, 2, 197,
	198, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 7, 7, 2, 2, 200, 43, 3,
	2, 2, 2, 201, 202, 7, 5, 2, 2, 202, 203, 7, 15, 2, 2, 203, 204, 5, 30,
	16, 2, 204, 45, 3, 2, 2, 2, 205, 206, 8, 24, 1, 2, 206, 209, 5, 48, 25,
	2, 207, 209, 5, 50, 26, 2, 208, 205, 3, 2, 2, 2, 208, 207, 3, 2, 2, 2,
	209, 227, 3, 2, 2, 2, 210, 211, 12, 7, 2, 2, 211, 212, 9, 3, 2, 2, 212,
	226, 5, 46, 24, 8, 213, 214, 12, 6, 2, 2, 214, 215, 9, 4, 2, 2, 215, 226,
	5, 46, 24, 7, 216, 217, 12, 5, 2, 2, 217, 218, 9, 5, 2, 2, 218, 226, 5,
	46, 24, 6, 219, 220, 12, 4, 2, 2, 220, 221, 7, 22, 2, 2, 221, 226, 5, 46,
	24, 5, 222, 223, 12, 3, 2, 2, 223, 224, 7, 21, 2, 2, 224, 226, 5, 46, 24,
	4, 225, 210, 3, 2, 2, 2, 225, 213, 3, 2, 2, 2, 225, 216, 3, 2, 2, 2, 225,
	219, 3, 2, 2, 2, 225, 222, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 47, 3, 2, 2, 2, 229, 227, 3, 2,
	2, 2, 230, 231, 8, 25, 1, 2, 231, 234, 5, 54, 28, 2, 232, 234, 5, 52, 27,
	2, 233, 230, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234, 243, 3, 2, 2, 2, 235,
	239, 12, 3, 2, 2, 236, 237, 7, 16, 2, 2, 237, 240, 7, 5, 2, 2, 238, 240,
	5, 76, 39, 2, 239, 236, 3, 2, 2, 2, 239, 238, 3, 2, 2, 2, 240, 242, 3,
	2, 2, 2, 241, 235, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 241, 3, 2, 2,
	2, 243, 244, 3, 2, 2, 2, 244, 49, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246,
	251, 5, 48, 25, 2, 247, 248, 5, 56, 29, 2, 248, 249, 5, 46, 24, 2, 249,
	251, 3, 2, 2, 2, 250, 246, 3, 2, 2, 2, 250, 247, 3, 2, 2, 2, 251, 51, 3,
	2, 2, 2, 252, 253, 5, 30, 16, 2, 253, 254, 7, 6, 2, 2, 254, 256, 5, 46,
	24, 2, 255, 257, 7, 13, 2, 2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2,
	2, 257, 258, 3, 2, 2, 2, 258, 259, 7, 7, 2, 2, 259, 53, 3, 2, 2, 2, 260,
	268, 5, 58, 30, 2, 261, 268, 5, 66, 34, 2, 262, 268, 5, 78, 40, 2, 263,
	264, 7, 6, 2, 2, 264, 265, 5, 46, 24, 2, 265, 266, 7, 7, 2, 2, 266, 268,
	3, 2, 2, 2, 267, 260, 3, 2, 2, 2, 267, 261, 3, 2, 2, 2, 267, 262, 3, 2,
	2, 2, 267, 263, 3, 2, 2, 2, 268, 55, 3, 2, 2, 2, 269, 270, 9, 6, 2, 2,
	270, 57, 3, 2, 2, 2, 271, 272, 5, 60, 31, 2, 272, 59, 3, 2, 2, 2, 273,
	278, 5, 62, 32, 2, 274, 278, 5, 74, 38, 2, 275, 278, 5, 64, 33, 2, 276,
	278, 7, 47, 2, 2, 277, 273, 3, 2, 2, 2, 277, 274, 3, 2, 2, 2, 277, 275,
	3, 2, 2, 2, 277, 276, 3, 2, 2, 2, 278, 61, 3, 2, 2, 2, 279, 280, 9, 7,
	2, 2, 280, 63, 3, 2, 2, 2, 281, 282, 9, 8, 2, 2, 282, 65, 3, 2, 2, 2, 283,
	286, 7, 5, 2, 2, 284, 286, 5, 68, 35, 2, 285, 283, 3, 2, 2, 2, 285, 284,
	3, 2, 2, 2, 286, 67, 3, 2, 2, 2, 287, 288, 7, 5, 2, 2, 288, 289, 7, 16,
	2, 2, 289, 290, 7, 5, 2, 2, 290, 69, 3, 2, 2, 2, 291, 292, 5, 36, 19, 2,
	292, 71, 3, 2, 2, 2, 293, 294, 5, 4, 3, 2, 294, 296, 5, 30, 16, 2, 295,
	297, 5, 74, 38, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 73,
	3, 2, 2, 2, 298, 299, 7, 48, 2, 2, 299, 75, 3, 2, 2, 2, 300, 315, 7, 6,
	2, 2, 301, 308, 5, 6, 4, 2, 302, 305, 5, 30, 16, 2, 303, 304, 7, 13, 2,
	2, 304, 306, 5, 6, 4, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306,
	308, 3, 2, 2, 2, 307, 301, 3, 2, 2, 2, 307, 302, 3, 2, 2, 2, 308, 310,
	3, 2, 2, 2, 309, 311, 7, 20, 2, 2, 310, 309, 3, 2, 2, 2, 310, 311, 3, 2,
	2, 2, 311, 313, 3, 2, 2, 2, 312, 314, 7, 13, 2, 2, 313, 312, 3, 2, 2, 2,
	313, 314, 3, 2, 2, 2, 314, 316, 3, 2, 2, 2, 315, 307, 3, 2, 2, 2, 315,
	316, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318, 7, 7, 2, 2, 318, 77, 3,
	2, 2, 2, 319, 320, 5, 80, 41, 2, 320, 321, 7, 16, 2, 2, 321, 322, 7, 5,
	2, 2, 322, 79, 3, 2, 2, 2, 323, 333, 5, 36, 19, 2, 324, 328, 7, 6, 2, 2,
	325, 326, 7, 39, 2, 2, 326, 329, 5, 36, 19, 2, 327, 329, 5, 80, 41, 2,
	328, 325, 3, 2, 2, 2, 328, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330,
	331, 7, 7, 2, 2, 331, 333, 3, 2, 2, 2, 332, 323, 3, 2, 2, 2, 332, 324,
	3, 2, 2, 2, 333, 81, 3, 2, 2, 2, 334, 335, 9, 9, 2, 2, 335, 83, 3, 2, 2,
	2, 37, 88, 90, 100, 108, 123, 127, 136, 142, 147, 156, 164, 169, 180, 186,
	194, 197, 208, 225, 227, 233, 239, 243, 250, 256, 267, 277, 285, 296, 305,
	307, 310, 313, 315, 328, 332,
}
var literalNames = []string{
	"", "'func'", "'return'", "", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'='", "','", "';'", "':'", "'.'", "'++'", "'--'", "':='", "'...'", "'||'",
	"'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'",
	"'<<'", "'>>'", "'&^'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'", "'->'",
}
var symbolicNames = []string{
	"", "FUNC", "RETURN", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS",
	"MINUS_MINUS", "DECLARE_ASSIGN", "ELLIPSIS", "LOGICAL_OR", "LOGICAL_AND",
	"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
	"OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "EXCLAMATION", "PLUS",
	"MINUS", "CARET", "STAR", "AMPERSAND", "ARROW", "DECIMAL_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "DOUBLE_LIT", "RUNE_LIT", "INTERPRETED_STRING_LIT",
	"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
}

var ruleNames = []string{
	"sourceFile", "identifierList", "expressionList", "functionDef", "varDef",
	"block", "statementList", "statement", "simpleStmt", "expressionStmt",
	"assignment", "assign_op", "emptyStmt", "returnStmt", "type_", "pointer",
	"array", "typeName", "signature", "result", "parameters", "parameterDecl",
	"expression", "primaryExpr", "unaryExpr", "conversion", "operand", "unaryOp",
	"literal", "basicLit", "integer", "floatingPoint", "operandName", "qualifiedIdent",
	"literalType", "fieldDecl", "string_", "arguments", "methodExpr", "receiverType",
	"eos",
}

type LlamaLangParser struct {
	*antlr.BaseParser
}

// NewLlamaLangParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LlamaLangParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLlamaLangParser(input antlr.TokenStream) *LlamaLangParser {
	this := new(LlamaLangParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LlamaLang.g4"

	return this
}

// LlamaLangParser tokens.
const (
	LlamaLangParserEOF                    = antlr.TokenEOF
	LlamaLangParserFUNC                   = 1
	LlamaLangParserRETURN                 = 2
	LlamaLangParserIDENTIFIER             = 3
	LlamaLangParserL_PAREN                = 4
	LlamaLangParserR_PAREN                = 5
	LlamaLangParserL_CURLY                = 6
	LlamaLangParserR_CURLY                = 7
	LlamaLangParserL_BRACKET              = 8
	LlamaLangParserR_BRACKET              = 9
	LlamaLangParserASSIGN                 = 10
	LlamaLangParserCOMMA                  = 11
	LlamaLangParserSEMI                   = 12
	LlamaLangParserCOLON                  = 13
	LlamaLangParserDOT                    = 14
	LlamaLangParserPLUS_PLUS              = 15
	LlamaLangParserMINUS_MINUS            = 16
	LlamaLangParserDECLARE_ASSIGN         = 17
	LlamaLangParserELLIPSIS               = 18
	LlamaLangParserLOGICAL_OR             = 19
	LlamaLangParserLOGICAL_AND            = 20
	LlamaLangParserEQUALS                 = 21
	LlamaLangParserNOT_EQUALS             = 22
	LlamaLangParserLESS                   = 23
	LlamaLangParserLESS_OR_EQUALS         = 24
	LlamaLangParserGREATER                = 25
	LlamaLangParserGREATER_OR_EQUALS      = 26
	LlamaLangParserOR                     = 27
	LlamaLangParserDIV                    = 28
	LlamaLangParserMOD                    = 29
	LlamaLangParserLSHIFT                 = 30
	LlamaLangParserRSHIFT                 = 31
	LlamaLangParserBIT_CLEAR              = 32
	LlamaLangParserEXCLAMATION            = 33
	LlamaLangParserPLUS                   = 34
	LlamaLangParserMINUS                  = 35
	LlamaLangParserCARET                  = 36
	LlamaLangParserSTAR                   = 37
	LlamaLangParserAMPERSAND              = 38
	LlamaLangParserARROW                  = 39
	LlamaLangParserDECIMAL_LIT            = 40
	LlamaLangParserOCTAL_LIT              = 41
	LlamaLangParserHEX_LIT                = 42
	LlamaLangParserFLOAT_LIT              = 43
	LlamaLangParserDOUBLE_LIT             = 44
	LlamaLangParserRUNE_LIT               = 45
	LlamaLangParserINTERPRETED_STRING_LIT = 46
	LlamaLangParserWS                     = 47
	LlamaLangParserCOMMENT                = 48
	LlamaLangParserTERMINATOR             = 49
	LlamaLangParserLINE_COMMENT           = 50
)

// LlamaLangParser rules.
const (
	LlamaLangParserRULE_sourceFile     = 0
	LlamaLangParserRULE_identifierList = 1
	LlamaLangParserRULE_expressionList = 2
	LlamaLangParserRULE_functionDef    = 3
	LlamaLangParserRULE_varDef         = 4
	LlamaLangParserRULE_block          = 5
	LlamaLangParserRULE_statementList  = 6
	LlamaLangParserRULE_statement      = 7
	LlamaLangParserRULE_simpleStmt     = 8
	LlamaLangParserRULE_expressionStmt = 9
	LlamaLangParserRULE_assignment     = 10
	LlamaLangParserRULE_assign_op      = 11
	LlamaLangParserRULE_emptyStmt      = 12
	LlamaLangParserRULE_returnStmt     = 13
	LlamaLangParserRULE_type_          = 14
	LlamaLangParserRULE_pointer        = 15
	LlamaLangParserRULE_array          = 16
	LlamaLangParserRULE_typeName       = 17
	LlamaLangParserRULE_signature      = 18
	LlamaLangParserRULE_result         = 19
	LlamaLangParserRULE_parameters     = 20
	LlamaLangParserRULE_parameterDecl  = 21
	LlamaLangParserRULE_expression     = 22
	LlamaLangParserRULE_primaryExpr    = 23
	LlamaLangParserRULE_unaryExpr      = 24
	LlamaLangParserRULE_conversion     = 25
	LlamaLangParserRULE_operand        = 26
	LlamaLangParserRULE_unaryOp        = 27
	LlamaLangParserRULE_literal        = 28
	LlamaLangParserRULE_basicLit       = 29
	LlamaLangParserRULE_integer        = 30
	LlamaLangParserRULE_floatingPoint  = 31
	LlamaLangParserRULE_operandName    = 32
	LlamaLangParserRULE_qualifiedIdent = 33
	LlamaLangParserRULE_literalType    = 34
	LlamaLangParserRULE_fieldDecl      = 35
	LlamaLangParserRULE_string_        = 36
	LlamaLangParserRULE_arguments      = 37
	LlamaLangParserRULE_methodExpr     = 38
	LlamaLangParserRULE_receiverType   = 39
	LlamaLangParserRULE_eos            = 40
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SourceFileContext) AllFunctionDef() []IFunctionDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDefContext)(nil)).Elem())
	var tst = make([]IFunctionDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDefContext)
		}
	}

	return tst
}

func (s *SourceFileContext) FunctionDef(i int) IFunctionDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *SourceFileContext) AllVarDef() []IVarDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDefContext)(nil)).Elem())
	var tst = make([]IVarDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDefContext)
		}
	}

	return tst
}

func (s *SourceFileContext) VarDef(i int) IVarDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDefContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LlamaLangParserRULE_sourceFile)
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
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LlamaLangParserFUNC || _la == LlamaLangParserIDENTIFIER {
		p.SetState(86)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LlamaLangParserFUNC:
			{
				p.SetState(82)
				p.FunctionDef()
			}

		case LlamaLangParserIDENTIFIER:
			{
				p.SetState(83)
				p.VarDef()
			}
			{
				p.SetState(84)
				p.Eos()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Eos()
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(LlamaLangParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LlamaLangParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LlamaLangParserRULE_identifierList)
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
		p.SetState(93)
		p.Match(LlamaLangParserIDENTIFIER)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LlamaLangParserCOMMA {
		{
			p.SetState(94)
			p.Match(LlamaLangParserCOMMA)
		}
		{
			p.SetState(95)
			p.Match(LlamaLangParserIDENTIFIER)
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LlamaLangParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LlamaLangParserRULE_expressionList)

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
	{
		p.SetState(101)
		p.expression(0)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(102)
				p.Match(LlamaLangParserCOMMA)
			}
			{
				p.SetState(103)
				p.expression(0)
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_functionDef
	return p
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserFUNC, 0)
}

func (s *FunctionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *FunctionDefContext) Signature() ISignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignatureContext)
}

func (s *FunctionDefContext) ARROW() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserARROW, 0)
}

func (s *FunctionDefContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FunctionDefContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitFunctionDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LlamaLangParserRULE_functionDef)

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
		p.SetState(109)
		p.Match(LlamaLangParserFUNC)
	}
	{
		p.SetState(110)
		p.Match(LlamaLangParserIDENTIFIER)
	}
	{
		p.SetState(111)
		p.Signature()
	}
	{
		p.SetState(112)
		p.Match(LlamaLangParserARROW)
	}
	{
		p.SetState(113)
		p.Type_()
	}
	{
		p.SetState(114)
		p.Block()
	}

	return localctx
}

// IVarDefContext is an interface to support dynamic dispatch.
type IVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDefContext differentiates from other interfaces.
	IsVarDefContext()
}

type VarDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDefContext() *VarDefContext {
	var p = new(VarDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_varDef
	return p
}

func (*VarDefContext) IsVarDefContext() {}

func NewVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefContext {
	var p = new(VarDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_varDef

	return p
}

func (s *VarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *VarDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOLON, 0)
}

func (s *VarDefContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarDefContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserASSIGN, 0)
}

func (s *VarDefContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitVarDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) VarDef() (localctx IVarDefContext) {
	localctx = NewVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LlamaLangParserRULE_varDef)
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
		p.SetState(116)
		p.Match(LlamaLangParserIDENTIFIER)
	}
	{
		p.SetState(117)
		p.Match(LlamaLangParserCOLON)
	}
	{
		p.SetState(118)
		p.Type_()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LlamaLangParserASSIGN {
		{
			p.SetState(119)
			p.Match(LlamaLangParserASSIGN)
		}
		{
			p.SetState(120)
			p.ExpressionList()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_CURLY, 0)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_CURLY, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LlamaLangParserRULE_block)
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
		p.SetState(123)
		p.Match(LlamaLangParserL_CURLY)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LlamaLangParserRETURN)|(1<<LlamaLangParserIDENTIFIER)|(1<<LlamaLangParserL_PAREN)|(1<<LlamaLangParserL_CURLY)|(1<<LlamaLangParserL_BRACKET)|(1<<LlamaLangParserSEMI)|(1<<LlamaLangParserPLUS_PLUS)|(1<<LlamaLangParserMINUS_MINUS))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LlamaLangParserEXCLAMATION-33))|(1<<(LlamaLangParserPLUS-33))|(1<<(LlamaLangParserMINUS-33))|(1<<(LlamaLangParserSTAR-33))|(1<<(LlamaLangParserAMPERSAND-33))|(1<<(LlamaLangParserDECIMAL_LIT-33))|(1<<(LlamaLangParserOCTAL_LIT-33))|(1<<(LlamaLangParserHEX_LIT-33))|(1<<(LlamaLangParserFLOAT_LIT-33))|(1<<(LlamaLangParserDOUBLE_LIT-33))|(1<<(LlamaLangParserRUNE_LIT-33))|(1<<(LlamaLangParserINTERPRETED_STRING_LIT-33)))) != 0) {
		{
			p.SetState(124)
			p.StatementList()
		}

	}
	{
		p.SetState(127)
		p.Match(LlamaLangParserR_CURLY)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementListContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LlamaLangParserRULE_statementList)
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LlamaLangParserRETURN)|(1<<LlamaLangParserIDENTIFIER)|(1<<LlamaLangParserL_PAREN)|(1<<LlamaLangParserL_CURLY)|(1<<LlamaLangParserL_BRACKET)|(1<<LlamaLangParserSEMI)|(1<<LlamaLangParserPLUS_PLUS)|(1<<LlamaLangParserMINUS_MINUS))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LlamaLangParserEXCLAMATION-33))|(1<<(LlamaLangParserPLUS-33))|(1<<(LlamaLangParserMINUS-33))|(1<<(LlamaLangParserSTAR-33))|(1<<(LlamaLangParserAMPERSAND-33))|(1<<(LlamaLangParserDECIMAL_LIT-33))|(1<<(LlamaLangParserOCTAL_LIT-33))|(1<<(LlamaLangParserHEX_LIT-33))|(1<<(LlamaLangParserFLOAT_LIT-33))|(1<<(LlamaLangParserDOUBLE_LIT-33))|(1<<(LlamaLangParserRUNE_LIT-33))|(1<<(LlamaLangParserINTERPRETED_STRING_LIT-33)))) != 0) {
		{
			p.SetState(129)
			p.Statement()
		}
		{
			p.SetState(130)
			p.Eos()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarDef() IVarDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDefContext)
}

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LlamaLangParserRULE_statement)

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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.VarDef()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.SimpleStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.ReturnStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(139)
			p.Block()
		}

	}

	return localctx
}

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_simpleStmt
	return p
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) ExpressionStmt() IExpressionStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStmtContext)
}

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleStmtContext) EmptyStmt() IEmptyStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStmtContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LlamaLangParserRULE_simpleStmt)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.ExpressionStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.EmptyStmt()
		}

	}

	return localctx
}

// IExpressionStmtContext is an interface to support dynamic dispatch.
type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}

type ExpressionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_expressionStmt
	return p
}

func (*ExpressionStmtContext) IsExpressionStmtContext() {}

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_expressionStmt

	return p
}

func (s *ExpressionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitExpressionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) ExpressionStmt() (localctx IExpressionStmtContext) {
	localctx = NewExpressionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LlamaLangParserRULE_expressionStmt)

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
		p.SetState(147)
		p.expression(0)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *AssignmentContext) Assign_op() IAssign_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_opContext)
}

func (s *AssignmentContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LlamaLangParserRULE_assignment)

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
		p.Match(LlamaLangParserIDENTIFIER)
	}
	{
		p.SetState(150)
		p.Assign_op()
	}
	{
		p.SetState(151)
		p.ExpressionList()
	}

	return localctx
}

// IAssign_opContext is an interface to support dynamic dispatch.
type IAssign_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssign_opContext differentiates from other interfaces.
	IsAssign_opContext()
}

type Assign_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_opContext() *Assign_opContext {
	var p = new(Assign_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_assign_op
	return p
}

func (*Assign_opContext) IsAssign_opContext() {}

func NewAssign_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_opContext {
	var p = new(Assign_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_assign_op

	return p
}

func (s *Assign_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_opContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserASSIGN, 0)
}

func (s *Assign_opContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserPLUS, 0)
}

func (s *Assign_opContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserMINUS, 0)
}

func (s *Assign_opContext) OR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserOR, 0)
}

func (s *Assign_opContext) CARET() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCARET, 0)
}

func (s *Assign_opContext) STAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSTAR, 0)
}

func (s *Assign_opContext) DIV() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDIV, 0)
}

func (s *Assign_opContext) MOD() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserMOD, 0)
}

func (s *Assign_opContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserLSHIFT, 0)
}

func (s *Assign_opContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserRSHIFT, 0)
}

func (s *Assign_opContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserAMPERSAND, 0)
}

func (s *Assign_opContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserBIT_CLEAR, 0)
}

func (s *Assign_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitAssign_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Assign_op() (localctx IAssign_opContext) {
	localctx = NewAssign_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LlamaLangParserRULE_assign_op)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(LlamaLangParserOR-27))|(1<<(LlamaLangParserDIV-27))|(1<<(LlamaLangParserMOD-27))|(1<<(LlamaLangParserLSHIFT-27))|(1<<(LlamaLangParserRSHIFT-27))|(1<<(LlamaLangParserBIT_CLEAR-27))|(1<<(LlamaLangParserPLUS-27))|(1<<(LlamaLangParserMINUS-27))|(1<<(LlamaLangParserCARET-27))|(1<<(LlamaLangParserSTAR-27))|(1<<(LlamaLangParserAMPERSAND-27)))) != 0 {
		{
			p.SetState(153)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(LlamaLangParserOR-27))|(1<<(LlamaLangParserDIV-27))|(1<<(LlamaLangParserMOD-27))|(1<<(LlamaLangParserLSHIFT-27))|(1<<(LlamaLangParserRSHIFT-27))|(1<<(LlamaLangParserBIT_CLEAR-27))|(1<<(LlamaLangParserPLUS-27))|(1<<(LlamaLangParserMINUS-27))|(1<<(LlamaLangParserCARET-27))|(1<<(LlamaLangParserSTAR-27))|(1<<(LlamaLangParserAMPERSAND-27)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(156)
		p.Match(LlamaLangParserASSIGN)
	}

	return localctx
}

// IEmptyStmtContext is an interface to support dynamic dispatch.
type IEmptyStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStmtContext differentiates from other interfaces.
	IsEmptyStmtContext()
}

type EmptyStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStmtContext() *EmptyStmtContext {
	var p = new(EmptyStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_emptyStmt
	return p
}

func (*EmptyStmtContext) IsEmptyStmtContext() {}

func NewEmptyStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStmtContext {
	var p = new(EmptyStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_emptyStmt

	return p
}

func (s *EmptyStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSEMI, 0)
}

func (s *EmptyStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitEmptyStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) EmptyStmt() (localctx IEmptyStmtContext) {
	localctx = NewEmptyStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LlamaLangParserRULE_emptyStmt)

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
		p.SetState(158)
		p.Match(LlamaLangParserSEMI)
	}

	return localctx
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_returnStmt
	return p
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LlamaLangParserRULE_returnStmt)
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
		p.SetState(160)
		p.Match(LlamaLangParserRETURN)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LlamaLangParserIDENTIFIER)|(1<<LlamaLangParserL_PAREN)|(1<<LlamaLangParserL_BRACKET)|(1<<LlamaLangParserPLUS_PLUS)|(1<<LlamaLangParserMINUS_MINUS))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LlamaLangParserEXCLAMATION-33))|(1<<(LlamaLangParserPLUS-33))|(1<<(LlamaLangParserMINUS-33))|(1<<(LlamaLangParserSTAR-33))|(1<<(LlamaLangParserAMPERSAND-33))|(1<<(LlamaLangParserDECIMAL_LIT-33))|(1<<(LlamaLangParserOCTAL_LIT-33))|(1<<(LlamaLangParserHEX_LIT-33))|(1<<(LlamaLangParserFLOAT_LIT-33))|(1<<(LlamaLangParserDOUBLE_LIT-33))|(1<<(LlamaLangParserRUNE_LIT-33))|(1<<(LlamaLangParserINTERPRETED_STRING_LIT-33)))) != 0) {
		{
			p.SetState(161)
			p.expression(0)
		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *Type_Context) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *Type_Context) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LlamaLangParserRULE_type_)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LlamaLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.TypeName()
		}

	case LlamaLangParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Pointer()
		}

	case LlamaLangParserL_BRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointerContext is an interface to support dynamic dispatch.
type IPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerContext differentiates from other interfaces.
	IsPointerContext()
}

type PointerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerContext() *PointerContext {
	var p = new(PointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_pointer
	return p
}

func (*PointerContext) IsPointerContext() {}

func NewPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerContext {
	var p = new(PointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_pointer

	return p
}

func (s *PointerContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerContext) STAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSTAR, 0)
}

func (s *PointerContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *PointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitPointer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Pointer() (localctx IPointerContext) {
	localctx = NewPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LlamaLangParserRULE_pointer)

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
		p.SetState(169)
		p.Match(LlamaLangParserSTAR)
	}
	{
		p.SetState(170)
		p.Type_()
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_BRACKET, 0)
}

func (s *ArrayContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_BRACKET, 0)
}

func (s *ArrayContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LlamaLangParserRULE_array)

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
		p.SetState(172)
		p.Match(LlamaLangParserL_BRACKET)
	}
	{
		p.SetState(173)
		p.Match(LlamaLangParserR_BRACKET)
	}
	{
		p.SetState(174)
		p.TypeName()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *TypeNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LlamaLangParserRULE_typeName)

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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(LlamaLangParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.QualifiedIdent()
		}

	}

	return localctx
}

// ISignatureContext is an interface to support dynamic dispatch.
type ISignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignatureContext differentiates from other interfaces.
	IsSignatureContext()
}

type SignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignatureContext() *SignatureContext {
	var p = new(SignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_signature
	return p
}

func (*SignatureContext) IsSignatureContext() {}

func NewSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignatureContext {
	var p = new(SignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_signature

	return p
}

func (s *SignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *SignatureContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *SignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitSignature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Signature() (localctx ISignatureContext) {
	localctx = NewSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LlamaLangParserRULE_signature)

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
		p.SetState(180)
		p.Parameters()
	}

	return localctx
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_result
	return p
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ResultContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitResult(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Result() (localctx IResultContext) {
	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LlamaLangParserRULE_result)

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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LlamaLangParserL_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Parameters()
		}

	case LlamaLangParserIDENTIFIER, LlamaLangParserL_BRACKET, LlamaLangParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Type_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_PAREN, 0)
}

func (s *ParametersContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_PAREN, 0)
}

func (s *ParametersContext) AllParameterDecl() []IParameterDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterDeclContext)(nil)).Elem())
	var tst = make([]IParameterDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterDeclContext)
		}
	}

	return tst
}

func (s *ParametersContext) ParameterDecl(i int) IParameterDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterDeclContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LlamaLangParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LlamaLangParserRULE_parameters)
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
		p.SetState(186)
		p.Match(LlamaLangParserL_PAREN)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LlamaLangParserIDENTIFIER {
		{
			p.SetState(187)
			p.ParameterDecl()
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LlamaLangParserCOMMA {
			{
				p.SetState(188)
				p.Match(LlamaLangParserCOMMA)
			}
			{
				p.SetState(189)
				p.ParameterDecl()
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(197)
		p.Match(LlamaLangParserR_PAREN)
	}

	return localctx
}

// IParameterDeclContext is an interface to support dynamic dispatch.
type IParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterDeclContext differentiates from other interfaces.
	IsParameterDeclContext()
}

type ParameterDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclContext() *ParameterDeclContext {
	var p = new(ParameterDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_parameterDecl
	return p
}

func (*ParameterDeclContext) IsParameterDeclContext() {}

func NewParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclContext {
	var p = new(ParameterDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_parameterDecl

	return p
}

func (s *ParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *ParameterDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOLON, 0)
}

func (s *ParameterDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitParameterDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) ParameterDecl() (localctx IParameterDeclContext) {
	localctx = NewParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LlamaLangParserRULE_parameterDecl)

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
		p.SetState(199)
		p.Match(LlamaLangParserIDENTIFIER)
	}
	{
		p.SetState(200)
		p.Match(LlamaLangParserCOLON)
	}
	{
		p.SetState(201)
		p.Type_()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpressionContext
	right  IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserMOD, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserRSHIFT, 0)
}

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserAMPERSAND, 0)
}

func (s *ExpressionContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserBIT_CLEAR, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserMINUS, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserOR, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCARET, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LlamaLangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, LlamaLangParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(204)
			p.primaryExpr(0)
		}

	case 2:
		{
			p.SetState(205)
			p.UnaryExpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LlamaLangParserRULE_expression)
				p.SetState(208)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(209)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(LlamaLangParserDIV-28))|(1<<(LlamaLangParserMOD-28))|(1<<(LlamaLangParserLSHIFT-28))|(1<<(LlamaLangParserRSHIFT-28))|(1<<(LlamaLangParserBIT_CLEAR-28))|(1<<(LlamaLangParserSTAR-28))|(1<<(LlamaLangParserAMPERSAND-28)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(210)

					var _x = p.expression(6)

					localctx.(*ExpressionContext).right = _x
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LlamaLangParserRULE_expression)
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(212)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(LlamaLangParserOR-27))|(1<<(LlamaLangParserPLUS-27))|(1<<(LlamaLangParserMINUS-27))|(1<<(LlamaLangParserCARET-27)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(213)

					var _x = p.expression(5)

					localctx.(*ExpressionContext).right = _x
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LlamaLangParserRULE_expression)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(215)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LlamaLangParserEQUALS)|(1<<LlamaLangParserNOT_EQUALS)|(1<<LlamaLangParserLESS)|(1<<LlamaLangParserLESS_OR_EQUALS)|(1<<LlamaLangParserGREATER)|(1<<LlamaLangParserGREATER_OR_EQUALS))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(216)

					var _x = p.expression(4)

					localctx.(*ExpressionContext).right = _x
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LlamaLangParserRULE_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(218)
					p.Match(LlamaLangParserLOGICAL_AND)
				}
				{
					p.SetState(219)

					var _x = p.expression(3)

					localctx.(*ExpressionContext).right = _x
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LlamaLangParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(221)
					p.Match(LlamaLangParserLOGICAL_OR)
				}
				{
					p.SetState(222)

					var _x = p.expression(2)

					localctx.(*ExpressionContext).right = _x
				}

			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) Conversion() IConversionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConversionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConversionContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDOT, 0)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *LlamaLangParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, LlamaLangParserRULE_primaryExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(229)
			p.Operand()
		}

	case 2:
		{
			p.SetState(230)
			p.Conversion()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, LlamaLangParserRULE_primaryExpr)
			p.SetState(233)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(237)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case LlamaLangParserDOT:
				{
					p.SetState(234)
					p.Match(LlamaLangParserDOT)
				}
				{
					p.SetState(235)
					p.Match(LlamaLangParserIDENTIFIER)
				}

			case LlamaLangParserL_PAREN:
				{
					p.SetState(236)
					p.Arguments()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) UnaryOp() IUnaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LlamaLangParserRULE_unaryExpr)

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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			p.primaryExpr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)
			p.UnaryOp()
		}
		{
			p.SetState(246)
			p.expression(0)
		}

	}

	return localctx
}

// IConversionContext is an interface to support dynamic dispatch.
type IConversionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConversionContext differentiates from other interfaces.
	IsConversionContext()
}

type ConversionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConversionContext() *ConversionContext {
	var p = new(ConversionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_conversion
	return p
}

func (*ConversionContext) IsConversionContext() {}

func NewConversionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConversionContext {
	var p = new(ConversionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_conversion

	return p
}

func (s *ConversionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConversionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ConversionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_PAREN, 0)
}

func (s *ConversionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConversionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_PAREN, 0)
}

func (s *ConversionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOMMA, 0)
}

func (s *ConversionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConversionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConversionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitConversion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Conversion() (localctx IConversionContext) {
	localctx = NewConversionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LlamaLangParserRULE_conversion)
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
		p.SetState(250)
		p.Type_()
	}
	{
		p.SetState(251)
		p.Match(LlamaLangParserL_PAREN)
	}
	{
		p.SetState(252)
		p.expression(0)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LlamaLangParserCOMMA {
		{
			p.SetState(253)
			p.Match(LlamaLangParserCOMMA)
		}

	}
	{
		p.SetState(256)
		p.Match(LlamaLangParserR_PAREN)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) MethodExpr() IMethodExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodExprContext)
}

func (s *OperandContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_PAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_PAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LlamaLangParserRULE_operand)

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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.OperandName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.MethodExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(261)
			p.Match(LlamaLangParserL_PAREN)
		}
		{
			p.SetState(262)
			p.expression(0)
		}
		{
			p.SetState(263)
			p.Match(LlamaLangParserR_PAREN)
		}

	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserPLUS, 0)
}

func (s *UnaryOpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserMINUS, 0)
}

func (s *UnaryOpContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserEXCLAMATION, 0)
}

func (s *UnaryOpContext) STAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSTAR, 0)
}

func (s *UnaryOpContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserAMPERSAND, 0)
}

func (s *UnaryOpContext) MINUS_MINUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserMINUS_MINUS, 0)
}

func (s *UnaryOpContext) PLUS_PLUS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserPLUS_PLUS, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LlamaLangParserRULE_unaryOp)
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
		p.SetState(267)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-15)&-(0x1f+1)) == 0 && ((1<<uint((_la-15)))&((1<<(LlamaLangParserPLUS_PLUS-15))|(1<<(LlamaLangParserMINUS_MINUS-15))|(1<<(LlamaLangParserEXCLAMATION-15))|(1<<(LlamaLangParserPLUS-15))|(1<<(LlamaLangParserMINUS-15))|(1<<(LlamaLangParserSTAR-15))|(1<<(LlamaLangParserAMPERSAND-15)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LlamaLangParserRULE_literal)

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
		p.SetState(269)
		p.BasicLit()
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *BasicLitContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *BasicLitContext) FloatingPoint() IFloatingPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatingPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatingPointContext)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserRUNE_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LlamaLangParserRULE_basicLit)

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

	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LlamaLangParserDECIMAL_LIT, LlamaLangParserOCTAL_LIT, LlamaLangParserHEX_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Integer()
		}

	case LlamaLangParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.String_()
		}

	case LlamaLangParserFLOAT_LIT, LlamaLangParserDOUBLE_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.FloatingPoint()
		}

	case LlamaLangParserRUNE_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(274)
			p.Match(LlamaLangParserRUNE_LIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserHEX_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LlamaLangParserRULE_integer)
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
		p.SetState(277)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(LlamaLangParserDECIMAL_LIT-40))|(1<<(LlamaLangParserOCTAL_LIT-40))|(1<<(LlamaLangParserHEX_LIT-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloatingPointContext is an interface to support dynamic dispatch.
type IFloatingPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatingPointContext differentiates from other interfaces.
	IsFloatingPointContext()
}

type FloatingPointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatingPointContext() *FloatingPointContext {
	var p = new(FloatingPointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_floatingPoint
	return p
}

func (*FloatingPointContext) IsFloatingPointContext() {}

func NewFloatingPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatingPointContext {
	var p = new(FloatingPointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_floatingPoint

	return p
}

func (s *FloatingPointContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatingPointContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserFLOAT_LIT, 0)
}

func (s *FloatingPointContext) DOUBLE_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDOUBLE_LIT, 0)
}

func (s *FloatingPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatingPointContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitFloatingPoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) FloatingPoint() (localctx IFloatingPointContext) {
	localctx = NewFloatingPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LlamaLangParserRULE_floatingPoint)
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
		p.SetState(279)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LlamaLangParserFLOAT_LIT || _la == LlamaLangParserDOUBLE_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *OperandNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LlamaLangParserRULE_operandName)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(LlamaLangParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.QualifiedIdent()
		}

	}

	return localctx
}

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_qualifiedIdent
	return p
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(LlamaLangParserIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, i)
}

func (s *QualifiedIdentContext) DOT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDOT, 0)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitQualifiedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) QualifiedIdent() (localctx IQualifiedIdentContext) {
	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LlamaLangParserRULE_qualifiedIdent)

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
		p.SetState(285)
		p.Match(LlamaLangParserIDENTIFIER)
	}
	{
		p.SetState(286)
		p.Match(LlamaLangParserDOT)
	}
	{
		p.SetState(287)
		p.Match(LlamaLangParserIDENTIFIER)
	}

	return localctx
}

// ILiteralTypeContext is an interface to support dynamic dispatch.
type ILiteralTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralTypeContext differentiates from other interfaces.
	IsLiteralTypeContext()
}

type LiteralTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralTypeContext() *LiteralTypeContext {
	var p = new(LiteralTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_literalType
	return p
}

func (*LiteralTypeContext) IsLiteralTypeContext() {}

func NewLiteralTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralTypeContext {
	var p = new(LiteralTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_literalType

	return p
}

func (s *LiteralTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *LiteralTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitLiteralType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) LiteralType() (localctx ILiteralTypeContext) {
	localctx = NewLiteralTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LlamaLangParserRULE_literalType)

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
		p.SetState(289)
		p.TypeName()
	}

	return localctx
}

// IFieldDeclContext is an interface to support dynamic dispatch.
type IFieldDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_fieldDecl
	return p
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *FieldDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDeclContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitFieldDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) FieldDecl() (localctx IFieldDeclContext) {
	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LlamaLangParserRULE_fieldDecl)
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
		p.SetState(291)
		p.IdentifierList()
	}
	{
		p.SetState(292)
		p.Type_()
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LlamaLangParserINTERPRETED_STRING_LIT {
		{
			p.SetState(293)
			p.String_()
		}

	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, LlamaLangParserRULE_string_)

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
		p.SetState(296)
		p.Match(LlamaLangParserINTERPRETED_STRING_LIT)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_PAREN, 0)
}

func (s *ArgumentsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_PAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArgumentsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserELLIPSIS, 0)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LlamaLangParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LlamaLangParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, LlamaLangParserRULE_arguments)
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
		p.SetState(298)
		p.Match(LlamaLangParserL_PAREN)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LlamaLangParserIDENTIFIER)|(1<<LlamaLangParserL_PAREN)|(1<<LlamaLangParserL_BRACKET)|(1<<LlamaLangParserPLUS_PLUS)|(1<<LlamaLangParserMINUS_MINUS))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LlamaLangParserEXCLAMATION-33))|(1<<(LlamaLangParserPLUS-33))|(1<<(LlamaLangParserMINUS-33))|(1<<(LlamaLangParserSTAR-33))|(1<<(LlamaLangParserAMPERSAND-33))|(1<<(LlamaLangParserDECIMAL_LIT-33))|(1<<(LlamaLangParserOCTAL_LIT-33))|(1<<(LlamaLangParserHEX_LIT-33))|(1<<(LlamaLangParserFLOAT_LIT-33))|(1<<(LlamaLangParserDOUBLE_LIT-33))|(1<<(LlamaLangParserRUNE_LIT-33))|(1<<(LlamaLangParserINTERPRETED_STRING_LIT-33)))) != 0) {
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(299)
				p.ExpressionList()
			}

		case 2:
			{
				p.SetState(300)
				p.Type_()
			}
			p.SetState(303)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(301)
					p.Match(LlamaLangParserCOMMA)
				}
				{
					p.SetState(302)
					p.ExpressionList()
				}

			}

		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LlamaLangParserELLIPSIS {
			{
				p.SetState(307)
				p.Match(LlamaLangParserELLIPSIS)
			}

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LlamaLangParserCOMMA {
			{
				p.SetState(310)
				p.Match(LlamaLangParserCOMMA)
			}

		}

	}
	{
		p.SetState(315)
		p.Match(LlamaLangParserR_PAREN)
	}

	return localctx
}

// IMethodExprContext is an interface to support dynamic dispatch.
type IMethodExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodExprContext differentiates from other interfaces.
	IsMethodExprContext()
}

type MethodExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodExprContext() *MethodExprContext {
	var p = new(MethodExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_methodExpr
	return p
}

func (*MethodExprContext) IsMethodExprContext() {}

func NewMethodExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodExprContext {
	var p = new(MethodExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_methodExpr

	return p
}

func (s *MethodExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodExprContext) ReceiverType() IReceiverTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReceiverTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReceiverTypeContext)
}

func (s *MethodExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserDOT, 0)
}

func (s *MethodExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserIDENTIFIER, 0)
}

func (s *MethodExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitMethodExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) MethodExpr() (localctx IMethodExprContext) {
	localctx = NewMethodExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, LlamaLangParserRULE_methodExpr)

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
		p.SetState(317)
		p.ReceiverType()
	}
	{
		p.SetState(318)
		p.Match(LlamaLangParserDOT)
	}
	{
		p.SetState(319)
		p.Match(LlamaLangParserIDENTIFIER)
	}

	return localctx
}

// IReceiverTypeContext is an interface to support dynamic dispatch.
type IReceiverTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReceiverTypeContext differentiates from other interfaces.
	IsReceiverTypeContext()
}

type ReceiverTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReceiverTypeContext() *ReceiverTypeContext {
	var p = new(ReceiverTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_receiverType
	return p
}

func (*ReceiverTypeContext) IsReceiverTypeContext() {}

func NewReceiverTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiverTypeContext {
	var p = new(ReceiverTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_receiverType

	return p
}

func (s *ReceiverTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiverTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ReceiverTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserL_PAREN, 0)
}

func (s *ReceiverTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserR_PAREN, 0)
}

func (s *ReceiverTypeContext) STAR() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSTAR, 0)
}

func (s *ReceiverTypeContext) ReceiverType() IReceiverTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReceiverTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReceiverTypeContext)
}

func (s *ReceiverTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiverTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiverTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitReceiverType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) ReceiverType() (localctx IReceiverTypeContext) {
	localctx = NewReceiverTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, LlamaLangParserRULE_receiverType)

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

	p.SetState(330)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LlamaLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.TypeName()
		}

	case LlamaLangParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(LlamaLangParserL_PAREN)
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LlamaLangParserSTAR:
			{
				p.SetState(323)
				p.Match(LlamaLangParserSTAR)
			}
			{
				p.SetState(324)
				p.TypeName()
			}

		case LlamaLangParserIDENTIFIER, LlamaLangParserL_PAREN:
			{
				p.SetState(325)
				p.ReceiverType()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(328)
			p.Match(LlamaLangParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LlamaLangParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LlamaLangParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserSEMI, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(LlamaLangParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LlamaLangVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LlamaLangParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, LlamaLangParserRULE_eos)
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
		p.SetState(332)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LlamaLangParserEOF || _la == LlamaLangParserSEMI) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *LlamaLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 23:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LlamaLangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LlamaLangParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
