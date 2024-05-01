// Package vec4 contains a 4 float64 components vector type T and functions.
package vec4

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/float64/generic"
	"github.com/ungerik/go3d/float64/vec3"
)

var (
	// Zero holds a zero vector.
	Zero = T{}

	// UnitXW holds a vector with X and W set to one.
	UnitXW = T{1, 0, 0, 1}
	// UnitYW holds a vector with Y and W set to one.
	UnitYW = T{0, 1, 0, 1}
	// UnitZW holds a vector with Z and W set to one.
	UnitZW = T{0, 0, 1, 1}
	// UnitW holds a vector with W set to one.
	UnitW = T{0, 0, 0, 1}
	// UnitXYZW holds a vector with X, Y, Z, W set to one.
	UnitXYZW = T{1, 1, 1, 1}

	// Red holds the color red.
	Red = T{1, 0, 0, 1}
	// Green holds the color green.
	Green = T{0, 1, 0, 1}
	// Blue holds the color blue.
	Blue = T{0, 0, 1, 1}
	// Black holds the color black.
	Black = T{0, 0, 0, 1}
	// White holds the color white.
	White = T{1, 1, 1, 1}

	// MinVal holds a vector with the smallest possible component values.
	MinVal = T{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, 1}
	// MaxVal holds a vector with the highest possible component values.
	MaxVal = T{+math.MaxFloat64, +math.MaxFloat64, +math.MaxFloat64, 1}
)

type ShuffleMask int

const (
	XXXX ShuffleMask = 0x0
	XXXY ShuffleMask = 0x40
	XXXZ ShuffleMask = 0x80
	XXXW ShuffleMask = 0xC0
	XXYX ShuffleMask = 0x10
	XXYY ShuffleMask = 0x50
	XXYZ ShuffleMask = 0x90
	XXYW ShuffleMask = 0xD0
	XXZX ShuffleMask = 0x20
	XXZY ShuffleMask = 0x60
	XXZZ ShuffleMask = 0xA0
	XXZW ShuffleMask = 0xE0
	XXWX ShuffleMask = 0x30
	XXWY ShuffleMask = 0x70
	XXWZ ShuffleMask = 0xB0
	XXWW ShuffleMask = 0xF0
	XYXX ShuffleMask = 0x4
	XYXY ShuffleMask = 0x44
	XYXZ ShuffleMask = 0x84
	XYXW ShuffleMask = 0xC4
	XYYX ShuffleMask = 0x14
	XYYY ShuffleMask = 0x54
	XYYZ ShuffleMask = 0x94
	XYYW ShuffleMask = 0xD4
	XYZX ShuffleMask = 0x24
	XYZY ShuffleMask = 0x64
	XYZZ ShuffleMask = 0xA4
	XYZW ShuffleMask = 0xE4
	XYWX ShuffleMask = 0x34
	XYWY ShuffleMask = 0x74
	XYWZ ShuffleMask = 0xB4
	XYWW ShuffleMask = 0xF4
	XZXX ShuffleMask = 0x8
	XZXY ShuffleMask = 0x48
	XZXZ ShuffleMask = 0x88
	XZXW ShuffleMask = 0xC8
	XZYX ShuffleMask = 0x18
	XZYY ShuffleMask = 0x58
	XZYZ ShuffleMask = 0x98
	XZYW ShuffleMask = 0xD8
	XZZX ShuffleMask = 0x28
	XZZY ShuffleMask = 0x68
	XZZZ ShuffleMask = 0xA8
	XZZW ShuffleMask = 0xE8
	XZWX ShuffleMask = 0x38
	XZWY ShuffleMask = 0x78
	XZWZ ShuffleMask = 0xB8
	XZWW ShuffleMask = 0xF8
	XWXX ShuffleMask = 0xC
	XWXY ShuffleMask = 0x4C
	XWXZ ShuffleMask = 0x8C
	XWXW ShuffleMask = 0xCC
	XWYX ShuffleMask = 0x1C
	XWYY ShuffleMask = 0x5C
	XWYZ ShuffleMask = 0x9C
	XWYW ShuffleMask = 0xDC
	XWZX ShuffleMask = 0x2C
	XWZY ShuffleMask = 0x6C
	XWZZ ShuffleMask = 0xAC
	XWZW ShuffleMask = 0xEC
	XWWX ShuffleMask = 0x3C
	XWWY ShuffleMask = 0x7C
	XWWZ ShuffleMask = 0xBC
	XWWW ShuffleMask = 0xFC
	YXXX ShuffleMask = 0x1
	YXXY ShuffleMask = 0x41
	YXXZ ShuffleMask = 0x81
	YXXW ShuffleMask = 0xC1
	YXYX ShuffleMask = 0x11
	YXYY ShuffleMask = 0x51
	YXYZ ShuffleMask = 0x91
	YXYW ShuffleMask = 0xD1
	YXZX ShuffleMask = 0x21
	YXZY ShuffleMask = 0x61
	YXZZ ShuffleMask = 0xA1
	YXZW ShuffleMask = 0xE1
	YXWX ShuffleMask = 0x31
	YXWY ShuffleMask = 0x71
	YXWZ ShuffleMask = 0xB1
	YXWW ShuffleMask = 0xF1
	YYXX ShuffleMask = 0x5
	YYXY ShuffleMask = 0x45
	YYXZ ShuffleMask = 0x85
	YYXW ShuffleMask = 0xC5
	YYYX ShuffleMask = 0x15
	YYYY ShuffleMask = 0x55
	YYYZ ShuffleMask = 0x95
	YYYW ShuffleMask = 0xD5
	YYZX ShuffleMask = 0x25
	YYZY ShuffleMask = 0x65
	YYZZ ShuffleMask = 0xA5
	YYZW ShuffleMask = 0xE5
	YYWX ShuffleMask = 0x35
	YYWY ShuffleMask = 0x75
	YYWZ ShuffleMask = 0xB5
	YYWW ShuffleMask = 0xF5
	YZXX ShuffleMask = 0x9
	YZXY ShuffleMask = 0x49
	YZXZ ShuffleMask = 0x89
	YZXW ShuffleMask = 0xC9
	YZYX ShuffleMask = 0x19
	YZYY ShuffleMask = 0x59
	YZYZ ShuffleMask = 0x99
	YZYW ShuffleMask = 0xD9
	YZZX ShuffleMask = 0x29
	YZZY ShuffleMask = 0x69
	YZZZ ShuffleMask = 0xA9
	YZZW ShuffleMask = 0xE9
	YZWX ShuffleMask = 0x39
	YZWY ShuffleMask = 0x79
	YZWZ ShuffleMask = 0xB9
	YZWW ShuffleMask = 0xF9
	YWXX ShuffleMask = 0xD
	YWXY ShuffleMask = 0x4D
	YWXZ ShuffleMask = 0x8D
	YWXW ShuffleMask = 0xCD
	YWYX ShuffleMask = 0x1D
	YWYY ShuffleMask = 0x5D
	YWYZ ShuffleMask = 0x9D
	YWYW ShuffleMask = 0xDD
	YWZX ShuffleMask = 0x2D
	YWZY ShuffleMask = 0x6D
	YWZZ ShuffleMask = 0xAD
	YWZW ShuffleMask = 0xED
	YWWX ShuffleMask = 0x3D
	YWWY ShuffleMask = 0x7D
	YWWZ ShuffleMask = 0xBD
	YWWW ShuffleMask = 0xFD
	ZXXX ShuffleMask = 0x2
	ZXXY ShuffleMask = 0x42
	ZXXZ ShuffleMask = 0x82
	ZXXW ShuffleMask = 0xC2
	ZXYX ShuffleMask = 0x12
	ZXYY ShuffleMask = 0x52
	ZXYZ ShuffleMask = 0x92
	ZXYW ShuffleMask = 0xD2
	ZXZX ShuffleMask = 0x22
	ZXZY ShuffleMask = 0x62
	ZXZZ ShuffleMask = 0xA2
	ZXZW ShuffleMask = 0xE2
	ZXWX ShuffleMask = 0x32
	ZXWY ShuffleMask = 0x72
	ZXWZ ShuffleMask = 0xB2
	ZXWW ShuffleMask = 0xF2
	ZYXX ShuffleMask = 0x6
	ZYXY ShuffleMask = 0x46
	ZYXZ ShuffleMask = 0x86
	ZYXW ShuffleMask = 0xC6
	ZYYX ShuffleMask = 0x16
	ZYYY ShuffleMask = 0x56
	ZYYZ ShuffleMask = 0x96
	ZYYW ShuffleMask = 0xD6
	ZYZX ShuffleMask = 0x26
	ZYZY ShuffleMask = 0x66
	ZYZZ ShuffleMask = 0xA6
	ZYZW ShuffleMask = 0xE6
	ZYWX ShuffleMask = 0x36
	ZYWY ShuffleMask = 0x76
	ZYWZ ShuffleMask = 0xB6
	ZYWW ShuffleMask = 0xF6
	ZZXX ShuffleMask = 0xA
	ZZXY ShuffleMask = 0x4A
	ZZXZ ShuffleMask = 0x8A
	ZZXW ShuffleMask = 0xCA
	ZZYX ShuffleMask = 0x1A
	ZZYY ShuffleMask = 0x5A
	ZZYZ ShuffleMask = 0x9A
	ZZYW ShuffleMask = 0xDA
	ZZZX ShuffleMask = 0x2A
	ZZZY ShuffleMask = 0x6A
	ZZZZ ShuffleMask = 0xAA
	ZZZW ShuffleMask = 0xEA
	ZZWX ShuffleMask = 0x3A
	ZZWY ShuffleMask = 0x7A
	ZZWZ ShuffleMask = 0xBA
	ZZWW ShuffleMask = 0xFA
	ZWXX ShuffleMask = 0xE
	ZWXY ShuffleMask = 0x4E
	ZWXZ ShuffleMask = 0x8E
	ZWXW ShuffleMask = 0xCE
	ZWYX ShuffleMask = 0x1E
	ZWYY ShuffleMask = 0x5E
	ZWYZ ShuffleMask = 0x9E
	ZWYW ShuffleMask = 0xDE
	ZWZX ShuffleMask = 0x2E
	ZWZY ShuffleMask = 0x6E
	ZWZZ ShuffleMask = 0xAE
	ZWZW ShuffleMask = 0xEE
	ZWWX ShuffleMask = 0x3E
	ZWWY ShuffleMask = 0x7E
	ZWWZ ShuffleMask = 0xBE
	ZWWW ShuffleMask = 0xFE
	WXXX ShuffleMask = 0x3
	WXXY ShuffleMask = 0x43
	WXXZ ShuffleMask = 0x83
	WXXW ShuffleMask = 0xC3
	WXYX ShuffleMask = 0x13
	WXYY ShuffleMask = 0x53
	WXYZ ShuffleMask = 0x93
	WXYW ShuffleMask = 0xD3
	WXZX ShuffleMask = 0x23
	WXZY ShuffleMask = 0x63
	WXZZ ShuffleMask = 0xA3
	WXZW ShuffleMask = 0xE3
	WXWX ShuffleMask = 0x33
	WXWY ShuffleMask = 0x73
	WXWZ ShuffleMask = 0xB3
	WXWW ShuffleMask = 0xF3
	WYXX ShuffleMask = 0x7
	WYXY ShuffleMask = 0x47
	WYXZ ShuffleMask = 0x87
	WYXW ShuffleMask = 0xC7
	WYYX ShuffleMask = 0x17
	WYYY ShuffleMask = 0x57
	WYYZ ShuffleMask = 0x97
	WYYW ShuffleMask = 0xD7
	WYZX ShuffleMask = 0x27
	WYZY ShuffleMask = 0x67
	WYZZ ShuffleMask = 0xA7
	WYZW ShuffleMask = 0xE7
	WYWX ShuffleMask = 0x37
	WYWY ShuffleMask = 0x77
	WYWZ ShuffleMask = 0xB7
	WYWW ShuffleMask = 0xF7
	WZXX ShuffleMask = 0xB
	WZXY ShuffleMask = 0x4B
	WZXZ ShuffleMask = 0x8B
	WZXW ShuffleMask = 0xCB
	WZYX ShuffleMask = 0x1B
	WZYY ShuffleMask = 0x5B
	WZYZ ShuffleMask = 0x9B
	WZYW ShuffleMask = 0xDB
	WZZX ShuffleMask = 0x2B
	WZZY ShuffleMask = 0x6B
	WZZZ ShuffleMask = 0xAB
	WZZW ShuffleMask = 0xEB
	WZWX ShuffleMask = 0x3B
	WZWY ShuffleMask = 0x7B
	WZWZ ShuffleMask = 0xBB
	WZWW ShuffleMask = 0xFB
	WWXX ShuffleMask = 0xF
	WWXY ShuffleMask = 0x4F
	WWXZ ShuffleMask = 0x8F
	WWXW ShuffleMask = 0xCF
	WWYX ShuffleMask = 0x1F
	WWYY ShuffleMask = 0x5F
	WWYZ ShuffleMask = 0x9F
	WWYW ShuffleMask = 0xDF
	WWZX ShuffleMask = 0x2F
	WWZY ShuffleMask = 0x6F
	WWZZ ShuffleMask = 0xAF
	WWZW ShuffleMask = 0xEF
	WWWX ShuffleMask = 0x3F
	WWWY ShuffleMask = 0x7F
	WWWZ ShuffleMask = 0xBF
	WWWW ShuffleMask = 0xFF
)

const (
	RRRR ShuffleMask = 0x0
	RRRG ShuffleMask = 0x40
	RRRB ShuffleMask = 0x80
	RRRA ShuffleMask = 0xC0
	RRGR ShuffleMask = 0x10
	RRGG ShuffleMask = 0x50
	RRGB ShuffleMask = 0x90
	RRGA ShuffleMask = 0xD0
	RRBR ShuffleMask = 0x20
	RRBG ShuffleMask = 0x60
	RRBB ShuffleMask = 0xA0
	RRBA ShuffleMask = 0xE0
	RRAR ShuffleMask = 0x30
	RRAG ShuffleMask = 0x70
	RRAB ShuffleMask = 0xB0
	RRAA ShuffleMask = 0xF0
	RGRR ShuffleMask = 0x4
	RGRG ShuffleMask = 0x44
	RGRB ShuffleMask = 0x84
	RGRA ShuffleMask = 0xC4
	RGGR ShuffleMask = 0x14
	RGGG ShuffleMask = 0x54
	RGGB ShuffleMask = 0x94
	RGGA ShuffleMask = 0xD4
	RGBR ShuffleMask = 0x24
	RGBG ShuffleMask = 0x64
	RGBB ShuffleMask = 0xA4
	RGBA ShuffleMask = 0xE4
	RGAR ShuffleMask = 0x34
	RGAG ShuffleMask = 0x74
	RGAB ShuffleMask = 0xB4
	RGAA ShuffleMask = 0xF4
	RBRR ShuffleMask = 0x8
	RBRG ShuffleMask = 0x48
	RBRB ShuffleMask = 0x88
	RBRA ShuffleMask = 0xC8
	RBGR ShuffleMask = 0x18
	RBGG ShuffleMask = 0x58
	RBGB ShuffleMask = 0x98
	RBGA ShuffleMask = 0xD8
	RBBR ShuffleMask = 0x28
	RBBG ShuffleMask = 0x68
	RBBB ShuffleMask = 0xA8
	RBBA ShuffleMask = 0xE8
	RBAR ShuffleMask = 0x38
	RBAG ShuffleMask = 0x78
	RBAB ShuffleMask = 0xB8
	RBAA ShuffleMask = 0xF8
	RARR ShuffleMask = 0xC
	RARG ShuffleMask = 0x4C
	RARB ShuffleMask = 0x8C
	RARA ShuffleMask = 0xCC
	RAGR ShuffleMask = 0x1C
	RAGG ShuffleMask = 0x5C
	RAGB ShuffleMask = 0x9C
	RAGA ShuffleMask = 0xDC
	RABR ShuffleMask = 0x2C
	RABG ShuffleMask = 0x6C
	RABB ShuffleMask = 0xAC
	RABA ShuffleMask = 0xEC
	RAAR ShuffleMask = 0x3C
	RAAG ShuffleMask = 0x7C
	RAAB ShuffleMask = 0xBC
	RAAA ShuffleMask = 0xFC
	GRRR ShuffleMask = 0x1
	GRRG ShuffleMask = 0x41
	GRRB ShuffleMask = 0x81
	GRRA ShuffleMask = 0xC1
	GRGR ShuffleMask = 0x11
	GRGG ShuffleMask = 0x51
	GRGB ShuffleMask = 0x91
	GRGA ShuffleMask = 0xD1
	GRBR ShuffleMask = 0x21
	GRBG ShuffleMask = 0x61
	GRBB ShuffleMask = 0xA1
	GRBA ShuffleMask = 0xE1
	GRAR ShuffleMask = 0x31
	GRAG ShuffleMask = 0x71
	GRAB ShuffleMask = 0xB1
	GRAA ShuffleMask = 0xF1
	GGRR ShuffleMask = 0x5
	GGRG ShuffleMask = 0x45
	GGRB ShuffleMask = 0x85
	GGRA ShuffleMask = 0xC5
	GGGR ShuffleMask = 0x15
	GGGG ShuffleMask = 0x55
	GGGB ShuffleMask = 0x95
	GGGA ShuffleMask = 0xD5
	GGBR ShuffleMask = 0x25
	GGBG ShuffleMask = 0x65
	GGBB ShuffleMask = 0xA5
	GGBA ShuffleMask = 0xE5
	GGAR ShuffleMask = 0x35
	GGAG ShuffleMask = 0x75
	GGAB ShuffleMask = 0xB5
	GGAA ShuffleMask = 0xF5
	GBRR ShuffleMask = 0x9
	GBRG ShuffleMask = 0x49
	GBRB ShuffleMask = 0x89
	GBRA ShuffleMask = 0xC9
	GBGR ShuffleMask = 0x19
	GBGG ShuffleMask = 0x59
	GBGB ShuffleMask = 0x99
	GBGA ShuffleMask = 0xD9
	GBBR ShuffleMask = 0x29
	GBBG ShuffleMask = 0x69
	GBBB ShuffleMask = 0xA9
	GBBA ShuffleMask = 0xE9
	GBAR ShuffleMask = 0x39
	GBAG ShuffleMask = 0x79
	GBAB ShuffleMask = 0xB9
	GBAA ShuffleMask = 0xF9
	GARR ShuffleMask = 0xD
	GARG ShuffleMask = 0x4D
	GARB ShuffleMask = 0x8D
	GARA ShuffleMask = 0xCD
	GAGR ShuffleMask = 0x1D
	GAGG ShuffleMask = 0x5D
	GAGB ShuffleMask = 0x9D
	GAGA ShuffleMask = 0xDD
	GABR ShuffleMask = 0x2D
	GABG ShuffleMask = 0x6D
	GABB ShuffleMask = 0xAD
	GABA ShuffleMask = 0xED
	GAAR ShuffleMask = 0x3D
	GAAG ShuffleMask = 0x7D
	GAAB ShuffleMask = 0xBD
	GAAA ShuffleMask = 0xFD
	BRRR ShuffleMask = 0x2
	BRRG ShuffleMask = 0x42
	BRRB ShuffleMask = 0x82
	BRRA ShuffleMask = 0xC2
	BRGR ShuffleMask = 0x12
	BRGG ShuffleMask = 0x52
	BRGB ShuffleMask = 0x92
	BRGA ShuffleMask = 0xD2
	BRBR ShuffleMask = 0x22
	BRBG ShuffleMask = 0x62
	BRBB ShuffleMask = 0xA2
	BRBA ShuffleMask = 0xE2
	BRAR ShuffleMask = 0x32
	BRAG ShuffleMask = 0x72
	BRAB ShuffleMask = 0xB2
	BRAA ShuffleMask = 0xF2
	BGRR ShuffleMask = 0x6
	BGRG ShuffleMask = 0x46
	BGRB ShuffleMask = 0x86
	BGRA ShuffleMask = 0xC6
	BGGR ShuffleMask = 0x16
	BGGG ShuffleMask = 0x56
	BGGB ShuffleMask = 0x96
	BGGA ShuffleMask = 0xD6
	BGBR ShuffleMask = 0x26
	BGBG ShuffleMask = 0x66
	BGBB ShuffleMask = 0xA6
	BGBA ShuffleMask = 0xE6
	BGAR ShuffleMask = 0x36
	BGAG ShuffleMask = 0x76
	BGAB ShuffleMask = 0xB6
	BGAA ShuffleMask = 0xF6
	BBRR ShuffleMask = 0xA
	BBRG ShuffleMask = 0x4A
	BBRB ShuffleMask = 0x8A
	BBRA ShuffleMask = 0xCA
	BBGR ShuffleMask = 0x1A
	BBGG ShuffleMask = 0x5A
	BBGB ShuffleMask = 0x9A
	BBGA ShuffleMask = 0xDA
	BBBR ShuffleMask = 0x2A
	BBBG ShuffleMask = 0x6A
	BBBB ShuffleMask = 0xAA
	BBBA ShuffleMask = 0xEA
	BBAR ShuffleMask = 0x3A
	BBAG ShuffleMask = 0x7A
	BBAB ShuffleMask = 0xBA
	BBAA ShuffleMask = 0xFA
	BARR ShuffleMask = 0xE
	BARG ShuffleMask = 0x4E
	BARB ShuffleMask = 0x8E
	BARA ShuffleMask = 0xCE
	BAGR ShuffleMask = 0x1E
	BAGG ShuffleMask = 0x5E
	BAGB ShuffleMask = 0x9E
	BAGA ShuffleMask = 0xDE
	BABR ShuffleMask = 0x2E
	BABG ShuffleMask = 0x6E
	BABB ShuffleMask = 0xAE
	BABA ShuffleMask = 0xEE
	BAAR ShuffleMask = 0x3E
	BAAG ShuffleMask = 0x7E
	BAAB ShuffleMask = 0xBE
	BAAA ShuffleMask = 0xFE
	ARRR ShuffleMask = 0x3
	ARRG ShuffleMask = 0x43
	ARRB ShuffleMask = 0x83
	ARRA ShuffleMask = 0xC3
	ARGR ShuffleMask = 0x13
	ARGG ShuffleMask = 0x53
	ARGB ShuffleMask = 0x93
	ARGA ShuffleMask = 0xD3
	ARBR ShuffleMask = 0x23
	ARBG ShuffleMask = 0x63
	ARBB ShuffleMask = 0xA3
	ARBA ShuffleMask = 0xE3
	ARAR ShuffleMask = 0x33
	ARAG ShuffleMask = 0x73
	ARAB ShuffleMask = 0xB3
	ARAA ShuffleMask = 0xF3
	AGRR ShuffleMask = 0x7
	AGRG ShuffleMask = 0x47
	AGRB ShuffleMask = 0x87
	AGRA ShuffleMask = 0xC7
	AGGR ShuffleMask = 0x17
	AGGG ShuffleMask = 0x57
	AGGB ShuffleMask = 0x97
	AGGA ShuffleMask = 0xD7
	AGBR ShuffleMask = 0x27
	AGBG ShuffleMask = 0x67
	AGBB ShuffleMask = 0xA7
	AGBA ShuffleMask = 0xE7
	AGAR ShuffleMask = 0x37
	AGAG ShuffleMask = 0x77
	AGAB ShuffleMask = 0xB7
	AGAA ShuffleMask = 0xF7
	ABRR ShuffleMask = 0xB
	ABRG ShuffleMask = 0x4B
	ABRB ShuffleMask = 0x8B
	ABRA ShuffleMask = 0xCB
	ABGR ShuffleMask = 0x1B
	ABGG ShuffleMask = 0x5B
	ABGB ShuffleMask = 0x9B
	ABGA ShuffleMask = 0xDB
	ABBR ShuffleMask = 0x2B
	ABBG ShuffleMask = 0x6B
	ABBB ShuffleMask = 0xAB
	ABBA ShuffleMask = 0xEB
	ABAR ShuffleMask = 0x3B
	ABAG ShuffleMask = 0x7B
	ABAB ShuffleMask = 0xBB
	ABAA ShuffleMask = 0xFB
	AARR ShuffleMask = 0xF
	AARG ShuffleMask = 0x4F
	AARB ShuffleMask = 0x8F
	AARA ShuffleMask = 0xCF
	AAGR ShuffleMask = 0x1F
	AAGG ShuffleMask = 0x5F
	AAGB ShuffleMask = 0x9F
	AAGA ShuffleMask = 0xDF
	AABR ShuffleMask = 0x2F
	AABG ShuffleMask = 0x6F
	AABB ShuffleMask = 0xAF
	AABA ShuffleMask = 0xEF
	AAAR ShuffleMask = 0x3F
	AAAG ShuffleMask = 0x7F
	AAAB ShuffleMask = 0xBF
	AAAA ShuffleMask = 0xFF
)

// T represents a 4 component vector.
type T [4]float64

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
	switch other.Size() {
	case 2:
		return T{other.Get(0, 0), other.Get(0, 1), 0, 1}
	case 3:
		return T{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2), 1}
	case 4:
		return T{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2), other.Get(0, 3)}
	default:
		panic("Unsupported type")
	}
}

// FromVec3 returns a vector with the first 3 components copied from a vec3.T.
func FromVec3(other *vec3.T) T {
	return T{other[0], other[1], other[2], 1}
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscan(s, &r[0], &r[1], &r[2], &r[3])
	return r, err
}

// String formats T as string. See also Parse().
func (vec *T) String() string {
	return fmt.Sprint(vec[0], vec[1], vec[2], vec[3])
}

// Rows returns the number of rows of the vector.
func (vec *T) Rows() int {
	return 4
}

// Cols returns the number of columns of the vector.
func (vec *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *T) Size() int {
	return 4
}

// Slice returns the elements of the vector as slice.
func (vec *T) Slice() []float64 {
	return vec[:]
}

// Get returns one element of the vector.
func (vec *T) Get(col, row int) float64 {
	return vec[row]
}

// IsZero checks if all elements of the vector are zero.
func (vec *T) IsZero() bool {
	return vec[0] == 0 && vec[1] == 0 && vec[2] == 0 && vec[3] == 0
}

// Shuffle returns the vector with its components shuffled in the order according to mask.
//
// Example:
// 		vec := vec4.T{1, 2, 3, 4}
//		vec.Shuffle(vec4.WZYX)
//		check := vec == T{4, 3, 2, 1} // true
//
func (vec *T) Shuffle(mask ShuffleMask) *T {
	*vec = vec.Shuffled(mask)
	return vec
}

// Shuffled returns a copy of the vector with its components shuffled in the order according to mask.
//
// Example:
// 		vec := vec4.T{1, 2, 3, 4}
//		check := vec.Shuffled(vec4.WZYX) == T{4, 3, 2, 1} // true
//
func (vec *T) Shuffled(mask ShuffleMask) (result T) {
	result[0] = vec[mask&3]
	result[1] = vec[(mask>>2)&3]
	result[2] = vec[(mask>>4)&3]
	result[3] = vec[(mask>>6)&3]
	return result
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *T) Length() float64 {
	v3 := vec.Vec3DividedByW()
	return v3.Length()
}

// LengthSqr returns the squared length of the vector.
// See also Length and Normalize.
func (vec *T) LengthSqr() float64 {
	v3 := vec.Vec3DividedByW()
	return v3.LengthSqr()
}

// Scale multiplies the first 3 element of the vector by f and returns vec.
func (vec *T) Scale(f float64) *T {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	return vec
}

// Scaled returns a copy of vec with the first 3 elements multiplies by f.
func (vec *T) Scaled(f float64) T {
	return T{vec[0] * f, vec[1] * f, vec[2] * f, vec[3]}
}

// Invert inverts the vector.
func (vec *T) Invert() {
	vec[0] = -vec[0]
	vec[1] = -vec[1]
	vec[2] = -vec[2]
}

// Inverted returns an inverted copy of the vector.
func (vec *T) Inverted() T {
	return T{-vec[0], -vec[1], -vec[2], vec[3]}
}

// Normalize normalizes the vector to unit length.
func (vec *T) Normalize() *T {
	v3 := vec.Vec3DividedByW()
	v3.Normalize()
	vec[0] = v3[0]
	vec[1] = v3[1]
	vec[2] = v3[2]
	vec[3] = 1
	return vec
}

// Normalized returns a unit length normalized copy of the vector.
func (vec *T) Normalized() T {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns an orthogonal vector.
func (vec *T) Normal() T {
	v3 := vec.Vec3()
	n3 := v3.Normal()
	return T{n3[0], n3[1], n3[2], 1}
}

// DivideByW divides the first three components (XYZ) by the last one (W).
func (vec *T) DivideByW() *T {
	oow := 1 / vec[3]
	vec[0] *= oow
	vec[1] *= oow
	vec[2] *= oow
	vec[3] = 1
	return vec
}

// DividedByW returns a copy of the vector with the first three components (XYZ) divided by the last one (W).
func (vec *T) DividedByW() T {
	oow := 1 / vec[3]
	return T{vec[0] * oow, vec[1] * oow, vec[2] * oow, 1}
}

// Vec3DividedByW returns a vec3.T version of the vector by dividing the first three vector components (XYZ) by the last one (W).
func (vec *T) Vec3DividedByW() vec3.T {
	oow := 1 / vec[3]
	return vec3.T{vec[0] * oow, vec[1] * oow, vec[2] * oow}
}

// Vec3 returns a vec3.T with the first three components of the vector.
// See also Vec3DividedByW
func (vec *T) Vec3() vec3.T {
	return vec3.T{vec[0], vec[1], vec[2]}
}

// AssignVec3 assigns v to the first three components and sets the fourth to 1.
func (vec *T) AssignVec3(v *vec3.T) *T {
	vec[0] = v[0]
	vec[1] = v[1]
	vec[2] = v[2]
	vec[3] = 1
	return vec
}

// Add adds another vector to vec.
func (vec *T) Add(v *T) *T {
	if v[3] == vec[3] {
		vec[0] += v[0]
		vec[1] += v[1]
		vec[2] += v[2]
	} else {
		vec.DividedByW()
		v3 := v.Vec3DividedByW()
		vec[0] += v3[0]
		vec[1] += v3[1]
		vec[2] += v3[2]
	}
	return vec
}

// Sub subtracts another vector from vec.
func (vec *T) Sub(v *T) *T {
	if v[3] == vec[3] {
		vec[0] -= v[0]
		vec[1] -= v[1]
		vec[2] -= v[2]
	} else {
		vec.DividedByW()
		v3 := v.Vec3DividedByW()
		vec[0] -= v3[0]
		vec[1] -= v3[1]
		vec[2] -= v3[2]
	}
	return vec
}

// Add returns the sum of two vectors.
func Add(a, b *T) T {
	if a[3] == b[3] {
		return T{a[0] + b[0], a[1] + b[1], a[2] + b[2], 1}
	} else {
		a3 := a.Vec3DividedByW()
		b3 := b.Vec3DividedByW()
		return T{a3[0] + b3[0], a3[1] + b3[1], a3[2] + b3[2], 1}
	}
}

// Sub returns the difference of two vectors.
func Sub(a, b *T) T {
	if a[3] == b[3] {
		return T{a[0] - b[0], a[1] - b[1], a[2] - b[2], 1}
	} else {
		a3 := a.Vec3DividedByW()
		b3 := b.Vec3DividedByW()
		return T{a3[0] - b3[0], a3[1] - b3[1], a3[2] - b3[2], 1}
	}
}

// Dot returns the dot product of two (dived by w) vectors.
func Dot(a, b *T) float64 {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	return vec3.Dot(&a3, &b3)
}

// Dot4 returns the 4 element vdot product of two vectors.
func Dot4(a, b *T) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// Cross returns the cross product of two vectors.
func Cross(a, b *T) T {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	c3 := vec3.Cross(&a3, &b3)
	return T{c3[0], c3[1], c3[2], 1}
}

// Angle returns the angle between two vectors.
func Angle(a, b *T) float64 {
	v := Dot(a, b) / (a.Length() * b.Length())
	// prevent NaN
	if v > 1. {
		v = v - 2
	} else if v < -1. {
		v = v + 2
	}
	return math.Acos(v)
}

// Interpolate interpolates between a and b at t (0,1).
func Interpolate(a, b *T, t float64) T {
	t1 := 1 - t
	return T{
		a[0]*t1 + b[0]*t,
		a[1]*t1 + b[1]*t,
		a[2]*t1 + b[2]*t,
		a[3]*t1 + b[3]*t,
	}
}

// Clamp clamps the vector's components to be in the range of min to max.
func (vec *T) Clamp(min, max *T) *T {
	for i := range vec {
		if vec[i] < min[i] {
			vec[i] = min[i]
		} else if vec[i] > max[i] {
			vec[i] = max[i]
		}
	}
	return vec
}

// Clamped returns a copy of the vector with the components clamped to be in the range of min to max.
func (vec *T) Clamped(min, max *T) T {
	result := *vec
	result.Clamp(min, max)
	return result
}

// Clamp01 clamps the vector's components to be in the range of 0 to 1.
func (vec *T) Clamp01() *T {
	return vec.Clamp(&Zero, &UnitXYZW)
}

// Clamped01 returns a copy of the vector with the components clamped to be in the range of 0 to 1.
func (vec *T) Clamped01() T {
	result := *vec
	result.Clamp01()
	return result
}
