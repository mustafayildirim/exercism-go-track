package protein

import "errors"

var (
	ErrStop        error = errors.New("STOP")
	ErrInvalidBase error = errors.New("Invalid codon")
)

var proteinList = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	codon := ""
	list := []string{}
	for i, v := range rna {
		codon += string(v)
		if i%3 == 2 {
			c := proteinList[codon]
			if c == "STOP" {
				break
			}
			list = append(list, c)
			codon = ""
		}
	}
	return list, nil
}

func FromCodon(codon string) (string, error) {
	c, present := proteinList[codon]
	if !present {
		return "", ErrInvalidBase
	}
	if c == "STOP" {
		return "", ErrStop
	}

	return c, nil
}
