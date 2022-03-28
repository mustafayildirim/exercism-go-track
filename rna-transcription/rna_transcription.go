package strand

import (
	"unicode/utf8"
)

const AdenineUnicode = "\u0041"
const CytosineUnicode = "\u0043"
const GuanineUnicode = "\u0047"
const ThymineUnicode = "\u0054"

func ToRNA(dna string) string {

	AdenineRune, _ := utf8.DecodeRuneInString(AdenineUnicode)
	CytosineRune, _ := utf8.DecodeRuneInString(CytosineUnicode)
	GuanineRune, _ := utf8.DecodeRuneInString(GuanineUnicode)
	ThymineRune, _ := utf8.DecodeRuneInString(ThymineUnicode)

	rna := ""

	for i, w := 0, 0; i < len(dna); i += w {
		runeValue, width := utf8.DecodeRuneInString(dna[i:])
		w = width

		switch runeValue {
		case AdenineRune:
			rna += "U"
		case CytosineRune:
			rna += "G"
		case GuanineRune:
			rna += "C"
		case ThymineRune:
			rna += "A"
		default:
			break

		}
	}
	return rna
}
