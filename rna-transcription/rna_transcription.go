package strand

import "bytes"

func ToRNA(dna string) string {
	sidesMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	buf := bytes.Buffer{}

	for _, ch := range dna {
		buf.WriteRune(sidesMap[ch])
	}

	return buf.String()
}
