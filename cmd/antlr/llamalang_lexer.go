// Code generated from ../../grammar/LlamaLang.g4 by ANTLR 4.9.1. DO NOT EDIT.

package antlr

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 52, 377,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 7, 4, 135, 10, 4, 12, 4, 14, 4, 138, 11, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 7, 41, 229, 10, 41, 12, 41, 14, 41, 232, 11, 41, 3, 42, 3,
	42, 7, 42, 236, 10, 42, 12, 42, 14, 42, 239, 11, 42, 3, 43, 3, 43, 3, 43,
	6, 43, 244, 10, 43, 13, 43, 14, 43, 245, 3, 44, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 5, 45, 254, 10, 45, 3, 45, 5, 45, 257, 10, 45, 3, 45, 5, 45,
	260, 10, 45, 3, 45, 3, 45, 3, 45, 5, 45, 265, 10, 45, 5, 45, 267, 10, 45,
	3, 46, 3, 46, 3, 46, 5, 46, 272, 10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 7, 47, 279, 10, 47, 12, 47, 14, 47, 282, 11, 47, 3, 47, 3, 47, 3, 48,
	6, 48, 287, 10, 48, 13, 48, 14, 48, 288, 3, 48, 3, 48, 3, 49, 3, 49, 3,
	49, 3, 49, 7, 49, 297, 10, 49, 12, 49, 14, 49, 300, 11, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 50, 6, 50, 308, 10, 50, 13, 50, 14, 50, 309, 3,
	50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 7, 51, 318, 10, 51, 12, 51, 14,
	51, 321, 11, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 5, 52,
	351, 10, 52, 3, 53, 6, 53, 354, 10, 53, 13, 53, 14, 53, 355, 3, 54, 3,
	54, 3, 55, 3, 55, 3, 56, 3, 56, 5, 56, 364, 10, 56, 3, 56, 3, 56, 3, 57,
	3, 57, 5, 57, 370, 10, 57, 3, 58, 5, 58, 373, 10, 58, 3, 59, 5, 59, 376,
	10, 59, 3, 298, 2, 60, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71,
	37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89,
	46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 2, 105, 2, 107,
	2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 3, 2, 17, 3, 2, 51, 59, 3, 2,
	50, 59, 4, 2, 90, 90, 122, 122, 4, 2, 72, 72, 104, 104, 4, 2, 12, 12, 94,
	94, 4, 2, 36, 36, 94, 94, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 11,
	2, 36, 36, 41, 41, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116, 118,
	118, 120, 120, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 71, 71,
	103, 103, 4, 2, 45, 45, 47, 47, 22, 2, 50, 59, 1634, 1643, 1778, 1787,
	2408, 2417, 2536, 2545, 2664, 2673, 2792, 2801, 2920, 2929, 3049, 3057,
	3176, 3185, 3304, 3313, 3432, 3441, 3666, 3675, 3794, 3803, 3874, 3883,
	4162, 4171, 4971, 4979, 6114, 6123, 6162, 6171, 65298, 65307, 260, 2, 67,
	92, 99, 124, 172, 172, 183, 183, 188, 188, 194, 216, 218, 248, 250, 545,
	548, 565, 594, 687, 690, 698, 701, 707, 722, 723, 738, 742, 752, 752, 892,
	892, 904, 904, 906, 908, 910, 910, 912, 931, 933, 976, 978, 985, 988, 1013,
	1026, 1155, 1166, 1222, 1225, 1226, 1229, 1230, 1234, 1271, 1274, 1275,
	1331, 1368, 1371, 1371, 1379, 1417, 1490, 1516, 1522, 1524, 1571, 1596,
	1602, 1612, 1651, 1749, 1751, 1751, 1767, 1768, 1788, 1790, 1810, 1810,
	1812, 1838, 1922, 1959, 2311, 2363, 2367, 2367, 2386, 2386, 2394, 2403,
	2439, 2446, 2449, 2450, 2453, 2474, 2476, 2482, 2484, 2484, 2488, 2491,
	2526, 2527, 2529, 2531, 2546, 2547, 2567, 2572, 2577, 2578, 2581, 2602,
	2604, 2610, 2612, 2613, 2615, 2616, 2618, 2619, 2651, 2654, 2656, 2656,
	2676, 2678, 2695, 2701, 2703, 2703, 2705, 2707, 2709, 2730, 2732, 2738,
	2740, 2741, 2743, 2747, 2751, 2751, 2770, 2770, 2786, 2786, 2823, 2830,
	2833, 2834, 2837, 2858, 2860, 2866, 2868, 2869, 2872, 2875, 2879, 2879,
	2910, 2911, 2913, 2915, 2951, 2956, 2960, 2962, 2964, 2967, 2971, 2972,
	2974, 2974, 2976, 2977, 2981, 2982, 2986, 2988, 2992, 2999, 3001, 3003,
	3079, 3086, 3088, 3090, 3092, 3114, 3116, 3125, 3127, 3131, 3170, 3171,
	3207, 3214, 3216, 3218, 3220, 3242, 3244, 3253, 3255, 3259, 3296, 3296,
	3298, 3299, 3335, 3342, 3344, 3346, 3348, 3370, 3372, 3387, 3426, 3427,
	3463, 3480, 3484, 3507, 3509, 3517, 3519, 3519, 3522, 3528, 3587, 3634,
	3636, 3637, 3650, 3656, 3715, 3716, 3718, 3718, 3721, 3722, 3724, 3724,
	3727, 3727, 3734, 3737, 3739, 3745, 3747, 3749, 3751, 3751, 3753, 3753,
	3756, 3757, 3759, 3762, 3764, 3765, 3775, 3782, 3784, 3784, 3806, 3807,
	3842, 3842, 3906, 3948, 3978, 3981, 4098, 4131, 4133, 4137, 4139, 4140,
	4178, 4183, 4258, 4295, 4306, 4344, 4354, 4443, 4449, 4516, 4522, 4603,
	4610, 4616, 4618, 4680, 4682, 4682, 4684, 4687, 4690, 4696, 4698, 4698,
	4700, 4703, 4706, 4744, 4746, 4746, 4748, 4751, 4754, 4784, 4786, 4786,
	4788, 4791, 4794, 4800, 4802, 4802, 4804, 4807, 4810, 4816, 4818, 4824,
	4826, 4848, 4850, 4880, 4882, 4882, 4884, 4887, 4890, 4896, 4898, 4936,
	4938, 4956, 5026, 5110, 5123, 5752, 5763, 5788, 5794, 5868, 6018, 6069,
	6178, 6265, 6274, 6314, 7682, 7837, 7842, 7931, 7938, 7959, 7962, 7967,
	7970, 8007, 8010, 8015, 8018, 8025, 8027, 8027, 8029, 8029, 8031, 8031,
	8033, 8063, 8066, 8118, 8120, 8126, 8128, 8128, 8132, 8134, 8136, 8142,
	8146, 8149, 8152, 8157, 8162, 8174, 8180, 8182, 8184, 8190, 8321, 8321,
	8452, 8452, 8457, 8457, 8460, 8469, 8471, 8471, 8475, 8479, 8486, 8486,
	8488, 8488, 8490, 8490, 8492, 8495, 8497, 8499, 8501, 8507, 8546, 8581,
	12295, 12297, 12323, 12331, 12339, 12343, 12346, 12348, 12355, 12438, 12447,
	12448, 12451, 12540, 12542, 12544, 12551, 12590, 12595, 12688, 12706, 12729,
	13314, 13314, 19895, 19895, 19970, 19970, 40871, 40871, 40962, 42126, 44034,
	44034, 55205, 55205, 63746, 64047, 64258, 64264, 64277, 64281, 64287, 64287,
	64289, 64298, 64300, 64312, 64314, 64318, 64320, 64320, 64322, 64323, 64325,
	64326, 64328, 64435, 64469, 64831, 64850, 64913, 64916, 64969, 65010, 65021,
	65138, 65140, 65142, 65142, 65144, 65278, 65315, 65340, 65347, 65372, 65384,
	65472, 65476, 65481, 65484, 65489, 65492, 65497, 65500, 65502, 2, 392,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2,
	49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2,
	2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2,
	2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2,
	2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3,
	2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87,
	3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2,
	95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2,
	2, 3, 119, 3, 2, 2, 2, 5, 124, 3, 2, 2, 2, 7, 131, 3, 2, 2, 2, 9, 139,
	3, 2, 2, 2, 11, 141, 3, 2, 2, 2, 13, 143, 3, 2, 2, 2, 15, 145, 3, 2, 2,
	2, 17, 147, 3, 2, 2, 2, 19, 149, 3, 2, 2, 2, 21, 151, 3, 2, 2, 2, 23, 153,
	3, 2, 2, 2, 25, 155, 3, 2, 2, 2, 27, 157, 3, 2, 2, 2, 29, 159, 3, 2, 2,
	2, 31, 161, 3, 2, 2, 2, 33, 164, 3, 2, 2, 2, 35, 167, 3, 2, 2, 2, 37, 170,
	3, 2, 2, 2, 39, 174, 3, 2, 2, 2, 41, 177, 3, 2, 2, 2, 43, 180, 3, 2, 2,
	2, 45, 183, 3, 2, 2, 2, 47, 186, 3, 2, 2, 2, 49, 188, 3, 2, 2, 2, 51, 191,
	3, 2, 2, 2, 53, 193, 3, 2, 2, 2, 55, 196, 3, 2, 2, 2, 57, 198, 3, 2, 2,
	2, 59, 200, 3, 2, 2, 2, 61, 202, 3, 2, 2, 2, 63, 205, 3, 2, 2, 2, 65, 208,
	3, 2, 2, 2, 67, 211, 3, 2, 2, 2, 69, 213, 3, 2, 2, 2, 71, 215, 3, 2, 2,
	2, 73, 217, 3, 2, 2, 2, 75, 219, 3, 2, 2, 2, 77, 221, 3, 2, 2, 2, 79, 223,
	3, 2, 2, 2, 81, 226, 3, 2, 2, 2, 83, 233, 3, 2, 2, 2, 85, 240, 3, 2, 2,
	2, 87, 247, 3, 2, 2, 2, 89, 266, 3, 2, 2, 2, 91, 268, 3, 2, 2, 2, 93, 275,
	3, 2, 2, 2, 95, 286, 3, 2, 2, 2, 97, 292, 3, 2, 2, 2, 99, 307, 3, 2, 2,
	2, 101, 313, 3, 2, 2, 2, 103, 324, 3, 2, 2, 2, 105, 353, 3, 2, 2, 2, 107,
	357, 3, 2, 2, 2, 109, 359, 3, 2, 2, 2, 111, 361, 3, 2, 2, 2, 113, 369,
	3, 2, 2, 2, 115, 372, 3, 2, 2, 2, 117, 375, 3, 2, 2, 2, 119, 120, 7, 104,
	2, 2, 120, 121, 7, 119, 2, 2, 121, 122, 7, 112, 2, 2, 122, 123, 7, 101,
	2, 2, 123, 4, 3, 2, 2, 2, 124, 125, 7, 116, 2, 2, 125, 126, 7, 103, 2,
	2, 126, 127, 7, 118, 2, 2, 127, 128, 7, 119, 2, 2, 128, 129, 7, 116, 2,
	2, 129, 130, 7, 112, 2, 2, 130, 6, 3, 2, 2, 2, 131, 136, 5, 113, 57, 2,
	132, 135, 5, 113, 57, 2, 133, 135, 5, 115, 58, 2, 134, 132, 3, 2, 2, 2,
	134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 8, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 140, 7,
	42, 2, 2, 140, 10, 3, 2, 2, 2, 141, 142, 7, 43, 2, 2, 142, 12, 3, 2, 2,
	2, 143, 144, 7, 125, 2, 2, 144, 14, 3, 2, 2, 2, 145, 146, 7, 127, 2, 2,
	146, 16, 3, 2, 2, 2, 147, 148, 7, 93, 2, 2, 148, 18, 3, 2, 2, 2, 149, 150,
	7, 95, 2, 2, 150, 20, 3, 2, 2, 2, 151, 152, 7, 63, 2, 2, 152, 22, 3, 2,
	2, 2, 153, 154, 7, 46, 2, 2, 154, 24, 3, 2, 2, 2, 155, 156, 7, 61, 2, 2,
	156, 26, 3, 2, 2, 2, 157, 158, 7, 60, 2, 2, 158, 28, 3, 2, 2, 2, 159, 160,
	7, 48, 2, 2, 160, 30, 3, 2, 2, 2, 161, 162, 7, 45, 2, 2, 162, 163, 7, 45,
	2, 2, 163, 32, 3, 2, 2, 2, 164, 165, 7, 47, 2, 2, 165, 166, 7, 47, 2, 2,
	166, 34, 3, 2, 2, 2, 167, 168, 7, 60, 2, 2, 168, 169, 7, 63, 2, 2, 169,
	36, 3, 2, 2, 2, 170, 171, 7, 48, 2, 2, 171, 172, 7, 48, 2, 2, 172, 173,
	7, 48, 2, 2, 173, 38, 3, 2, 2, 2, 174, 175, 7, 126, 2, 2, 175, 176, 7,
	126, 2, 2, 176, 40, 3, 2, 2, 2, 177, 178, 7, 40, 2, 2, 178, 179, 7, 40,
	2, 2, 179, 42, 3, 2, 2, 2, 180, 181, 7, 63, 2, 2, 181, 182, 7, 63, 2, 2,
	182, 44, 3, 2, 2, 2, 183, 184, 7, 35, 2, 2, 184, 185, 7, 63, 2, 2, 185,
	46, 3, 2, 2, 2, 186, 187, 7, 62, 2, 2, 187, 48, 3, 2, 2, 2, 188, 189, 7,
	62, 2, 2, 189, 190, 7, 63, 2, 2, 190, 50, 3, 2, 2, 2, 191, 192, 7, 64,
	2, 2, 192, 52, 3, 2, 2, 2, 193, 194, 7, 64, 2, 2, 194, 195, 7, 63, 2, 2,
	195, 54, 3, 2, 2, 2, 196, 197, 7, 126, 2, 2, 197, 56, 3, 2, 2, 2, 198,
	199, 7, 49, 2, 2, 199, 58, 3, 2, 2, 2, 200, 201, 7, 39, 2, 2, 201, 60,
	3, 2, 2, 2, 202, 203, 7, 62, 2, 2, 203, 204, 7, 62, 2, 2, 204, 62, 3, 2,
	2, 2, 205, 206, 7, 64, 2, 2, 206, 207, 7, 64, 2, 2, 207, 64, 3, 2, 2, 2,
	208, 209, 7, 40, 2, 2, 209, 210, 7, 96, 2, 2, 210, 66, 3, 2, 2, 2, 211,
	212, 7, 35, 2, 2, 212, 68, 3, 2, 2, 2, 213, 214, 7, 45, 2, 2, 214, 70,
	3, 2, 2, 2, 215, 216, 7, 47, 2, 2, 216, 72, 3, 2, 2, 2, 217, 218, 7, 96,
	2, 2, 218, 74, 3, 2, 2, 2, 219, 220, 7, 44, 2, 2, 220, 76, 3, 2, 2, 2,
	221, 222, 7, 40, 2, 2, 222, 78, 3, 2, 2, 2, 223, 224, 7, 47, 2, 2, 224,
	225, 7, 64, 2, 2, 225, 80, 3, 2, 2, 2, 226, 230, 9, 2, 2, 2, 227, 229,
	9, 3, 2, 2, 228, 227, 3, 2, 2, 2, 229, 232, 3, 2, 2, 2, 230, 228, 3, 2,
	2, 2, 230, 231, 3, 2, 2, 2, 231, 82, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2,
	233, 237, 7, 50, 2, 2, 234, 236, 5, 107, 54, 2, 235, 234, 3, 2, 2, 2, 236,
	239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 84, 3,
	2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 7, 50, 2, 2, 241, 243, 9, 4, 2,
	2, 242, 244, 5, 109, 55, 2, 243, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2,
	245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 86, 3, 2, 2, 2, 247, 248,
	5, 89, 45, 2, 248, 249, 9, 5, 2, 2, 249, 88, 3, 2, 2, 2, 250, 259, 5, 105,
	53, 2, 251, 253, 7, 48, 2, 2, 252, 254, 5, 105, 53, 2, 253, 252, 3, 2,
	2, 2, 253, 254, 3, 2, 2, 2, 254, 256, 3, 2, 2, 2, 255, 257, 5, 111, 56,
	2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258,
	260, 5, 111, 56, 2, 259, 251, 3, 2, 2, 2, 259, 258, 3, 2, 2, 2, 260, 267,
	3, 2, 2, 2, 261, 262, 7, 48, 2, 2, 262, 264, 5, 105, 53, 2, 263, 265, 5,
	111, 56, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 267, 3, 2,
	2, 2, 266, 250, 3, 2, 2, 2, 266, 261, 3, 2, 2, 2, 267, 90, 3, 2, 2, 2,
	268, 271, 7, 41, 2, 2, 269, 272, 10, 6, 2, 2, 270, 272, 5, 103, 52, 2,
	271, 269, 3, 2, 2, 2, 271, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	274, 7, 41, 2, 2, 274, 92, 3, 2, 2, 2, 275, 280, 7, 36, 2, 2, 276, 279,
	10, 7, 2, 2, 277, 279, 5, 103, 52, 2, 278, 276, 3, 2, 2, 2, 278, 277, 3,
	2, 2, 2, 279, 282, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2,
	2, 281, 283, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 283, 284, 7, 36, 2, 2, 284,
	94, 3, 2, 2, 2, 285, 287, 9, 8, 2, 2, 286, 285, 3, 2, 2, 2, 287, 288, 3,
	2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 3, 2, 2,
	2, 290, 291, 8, 48, 2, 2, 291, 96, 3, 2, 2, 2, 292, 293, 7, 49, 2, 2, 293,
	294, 7, 44, 2, 2, 294, 298, 3, 2, 2, 2, 295, 297, 11, 2, 2, 2, 296, 295,
	3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 298, 296, 3, 2,
	2, 2, 299, 301, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 301, 302, 7, 44, 2, 2,
	302, 303, 7, 49, 2, 2, 303, 304, 3, 2, 2, 2, 304, 305, 8, 49, 2, 2, 305,
	98, 3, 2, 2, 2, 306, 308, 9, 9, 2, 2, 307, 306, 3, 2, 2, 2, 308, 309, 3,
	2, 2, 2, 309, 307, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 3, 2, 2,
	2, 311, 312, 8, 50, 2, 2, 312, 100, 3, 2, 2, 2, 313, 314, 7, 49, 2, 2,
	314, 315, 7, 49, 2, 2, 315, 319, 3, 2, 2, 2, 316, 318, 10, 9, 2, 2, 317,
	316, 3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319, 320,
	3, 2, 2, 2, 320, 322, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 322, 323, 8, 51,
	2, 2, 323, 102, 3, 2, 2, 2, 324, 350, 7, 94, 2, 2, 325, 326, 7, 119, 2,
	2, 326, 327, 5, 109, 55, 2, 327, 328, 5, 109, 55, 2, 328, 329, 5, 109,
	55, 2, 329, 330, 5, 109, 55, 2, 330, 351, 3, 2, 2, 2, 331, 332, 7, 87,
	2, 2, 332, 333, 5, 109, 55, 2, 333, 334, 5, 109, 55, 2, 334, 335, 5, 109,
	55, 2, 335, 336, 5, 109, 55, 2, 336, 337, 5, 109, 55, 2, 337, 338, 5, 109,
	55, 2, 338, 339, 5, 109, 55, 2, 339, 340, 5, 109, 55, 2, 340, 351, 3, 2,
	2, 2, 341, 351, 9, 10, 2, 2, 342, 343, 5, 107, 54, 2, 343, 344, 5, 107,
	54, 2, 344, 345, 5, 107, 54, 2, 345, 351, 3, 2, 2, 2, 346, 347, 7, 122,
	2, 2, 347, 348, 5, 109, 55, 2, 348, 349, 5, 109, 55, 2, 349, 351, 3, 2,
	2, 2, 350, 325, 3, 2, 2, 2, 350, 331, 3, 2, 2, 2, 350, 341, 3, 2, 2, 2,
	350, 342, 3, 2, 2, 2, 350, 346, 3, 2, 2, 2, 351, 104, 3, 2, 2, 2, 352,
	354, 9, 3, 2, 2, 353, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 353,
	3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 106, 3, 2, 2, 2, 357, 358, 9, 11,
	2, 2, 358, 108, 3, 2, 2, 2, 359, 360, 9, 12, 2, 2, 360, 110, 3, 2, 2, 2,
	361, 363, 9, 13, 2, 2, 362, 364, 9, 14, 2, 2, 363, 362, 3, 2, 2, 2, 363,
	364, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 366, 5, 105, 53, 2, 366, 112,
	3, 2, 2, 2, 367, 370, 5, 117, 59, 2, 368, 370, 7, 97, 2, 2, 369, 367, 3,
	2, 2, 2, 369, 368, 3, 2, 2, 2, 370, 114, 3, 2, 2, 2, 371, 373, 9, 15, 2,
	2, 372, 371, 3, 2, 2, 2, 373, 116, 3, 2, 2, 2, 374, 376, 9, 16, 2, 2, 375,
	374, 3, 2, 2, 2, 376, 118, 3, 2, 2, 2, 26, 2, 134, 136, 230, 237, 245,
	253, 256, 259, 264, 266, 271, 278, 280, 288, 298, 309, 319, 350, 355, 363,
	369, 372, 375, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'func'", "'return'", "", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'='", "','", "';'", "':'", "'.'", "'++'", "'--'", "':='", "'...'", "'||'",
	"'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'",
	"'<<'", "'>>'", "'&^'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'", "'->'",
}

var lexerSymbolicNames = []string{
	"", "FUNC", "RETURN", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS",
	"MINUS_MINUS", "DECLARE_ASSIGN", "ELLIPSIS", "LOGICAL_OR", "LOGICAL_AND",
	"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
	"OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "EXCLAMATION", "PLUS",
	"MINUS", "CARET", "STAR", "AMPERSAND", "ARROW", "DECIMAL_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "DOUBLE_LIT", "RUNE_LIT", "INTERPRETED_STRING_LIT",
	"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"FUNC", "RETURN", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS",
	"MINUS_MINUS", "DECLARE_ASSIGN", "ELLIPSIS", "LOGICAL_OR", "LOGICAL_AND",
	"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
	"OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "EXCLAMATION", "PLUS",
	"MINUS", "CARET", "STAR", "AMPERSAND", "ARROW", "DECIMAL_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "DOUBLE_LIT", "RUNE_LIT", "INTERPRETED_STRING_LIT",
	"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", "ESCAPED_VALUE", "DECIMALS",
	"OCTAL_DIGIT", "HEX_DIGIT", "EXPONENT", "LETTER", "UNICODE_DIGIT", "UNICODE_LETTER",
}

type LlamaLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewLlamaLangLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *LlamaLangLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLlamaLangLexer(input antlr.CharStream) *LlamaLangLexer {
	l := new(LlamaLangLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "LlamaLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LlamaLangLexer tokens.
const (
	LlamaLangLexerFUNC                   = 1
	LlamaLangLexerRETURN                 = 2
	LlamaLangLexerIDENTIFIER             = 3
	LlamaLangLexerL_PAREN                = 4
	LlamaLangLexerR_PAREN                = 5
	LlamaLangLexerL_CURLY                = 6
	LlamaLangLexerR_CURLY                = 7
	LlamaLangLexerL_BRACKET              = 8
	LlamaLangLexerR_BRACKET              = 9
	LlamaLangLexerASSIGN                 = 10
	LlamaLangLexerCOMMA                  = 11
	LlamaLangLexerSEMI                   = 12
	LlamaLangLexerCOLON                  = 13
	LlamaLangLexerDOT                    = 14
	LlamaLangLexerPLUS_PLUS              = 15
	LlamaLangLexerMINUS_MINUS            = 16
	LlamaLangLexerDECLARE_ASSIGN         = 17
	LlamaLangLexerELLIPSIS               = 18
	LlamaLangLexerLOGICAL_OR             = 19
	LlamaLangLexerLOGICAL_AND            = 20
	LlamaLangLexerEQUALS                 = 21
	LlamaLangLexerNOT_EQUALS             = 22
	LlamaLangLexerLESS                   = 23
	LlamaLangLexerLESS_OR_EQUALS         = 24
	LlamaLangLexerGREATER                = 25
	LlamaLangLexerGREATER_OR_EQUALS      = 26
	LlamaLangLexerOR                     = 27
	LlamaLangLexerDIV                    = 28
	LlamaLangLexerMOD                    = 29
	LlamaLangLexerLSHIFT                 = 30
	LlamaLangLexerRSHIFT                 = 31
	LlamaLangLexerBIT_CLEAR              = 32
	LlamaLangLexerEXCLAMATION            = 33
	LlamaLangLexerPLUS                   = 34
	LlamaLangLexerMINUS                  = 35
	LlamaLangLexerCARET                  = 36
	LlamaLangLexerSTAR                   = 37
	LlamaLangLexerAMPERSAND              = 38
	LlamaLangLexerARROW                  = 39
	LlamaLangLexerDECIMAL_LIT            = 40
	LlamaLangLexerOCTAL_LIT              = 41
	LlamaLangLexerHEX_LIT                = 42
	LlamaLangLexerFLOAT_LIT              = 43
	LlamaLangLexerDOUBLE_LIT             = 44
	LlamaLangLexerRUNE_LIT               = 45
	LlamaLangLexerINTERPRETED_STRING_LIT = 46
	LlamaLangLexerWS                     = 47
	LlamaLangLexerCOMMENT                = 48
	LlamaLangLexerTERMINATOR             = 49
	LlamaLangLexerLINE_COMMENT           = 50
)