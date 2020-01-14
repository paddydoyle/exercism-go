package reverse

/*
import (
	"fmt"
	"strings"
)
*/
import (
	"fmt"
)

// Reverse returns a reversal of the input string
func Reverse(input string) string {
	// solution from: https://yourbasic.org/golang/reverse-utf8-string/
	output := make([]byte, len(input))
	prevPos, outputPos := 0, len(input)

	fmt.Println("input: ", input)
	fmt.Println("input bytes: ", []byte(input))
	fmt.Println("prevPos: ", prevPos)
	fmt.Println("outputPos: ", outputPos)

	for pos := range input {
		fmt.Println("START pos: ", pos, "prevPos: ", prevPos)
		outputPos -= pos - prevPos
		fmt.Println("outputPos now: ", outputPos)
		fmt.Println("copying into output[outputPos:], from input[prevPos:pos]: ", input[prevPos:pos])
		copy(output[outputPos:], input[prevPos:pos])
		fmt.Println("output: ", output)
		prevPos = pos
		fmt.Println("prevPos now: ", prevPos)
		fmt.Println("")
	}
	copy(output[0:], input[prevPos:])
	fmt.Println("output: ", output)
	fmt.Println("output string: ", string(output))
	fmt.Println("")
	return string(output)

	/*
		// won't work because of utf8 runes, which do not equal bytes
		// 'range' gives you runes, but copying single chars doesn't work
			inputLen := len(input)
			output := make([]rune, len(input))

			fmt.Println("input: ", input)
			fmt.Println("len: ", inputLen)

			for i, v := range input {
				//output[i], output[inputLen-i] = input[i], input[inputLen-i]
				fmt.Println("index: ", i)
				fmt.Println("value: ", v)
			}

			return string(output)
	*/

	/*
		// doesn't work; the FieldsFunc doesn't do what I thought it would
			fmt.Println(input)

			output := new(strings.Builder)

			// Convert to slice of single-rune strings
			runes := strings.FieldsFunc(input, func(c rune) bool { return true })
			lenRunes := len(runes)

			fmt.Println("Runes = ", runes)
			fmt.Println("len runes = ", lenRunes)
			fmt.Println("---------------------")

			for i := range runes {
				output.WriteString(runes[lenRunes-i-1])
				fmt.Println(output.String())
				fmt.Println("------------0000000000000000000000000000000000---------")
			}

			return output.String()
	*/
}
