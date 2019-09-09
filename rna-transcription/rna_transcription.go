package strand

import "bytes"

var sidesMap = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

func ToRNA(dna string) string {

	buf := bytes.Buffer{}

	for _, ch := range dna {
		buf.WriteRune(sidesMap[ch])
	}

	return buf.String()
}
