// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	//"fmt"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate2(s string) string {
	//words := regexp.MustCompile(`[ \-_]`).Split(s, -1)

	buf := bytes.Buffer{}

	//fmt.Printf("words = %s\n", words)

	//words := re.Split(s, -1).FindAllString(`\w`)

	//strs2 := regexp.MustCompile(`\a([[:alpha:]])|\b([[:alpha:]])|`).FindAllStringSubmatch(s, -1)
	words := regexp.MustCompile(`^([[:alpha:]])|[ \-_^]([[:alpha:]])`).FindAllStringSubmatch(s, -1)
	//fmt.Printf("s = %s\n", s)
	//fmt.Printf("words = %s\n", words)


        for _, word := range words {
		//fmt.Printf("word = %s\n", word)
		//fmt.Printf("word join = %s\n", strings.Join(word, `", "`))
		//for i, w := range word {
		//	fmt.Printf("w[%d] = %s\n", i, w)
		//}
		//fmt.Printf("word[0] = %s\n", word[1])
                if len(word) > 0 {
			buf.WriteString(strings.Join(word[1:], ``))
		}
        }

        return strings.ToUpper(buf.String())
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	words := regexp.MustCompile(`[ \-_]`).Split(s, -1)

	buf := bytes.Buffer{}

	//fmt.Printf("words = %s\n", words)

	//words := re.Split(s, -1).FindAllString(`\w`)

	//strs2 := regexp.MustCompile(`\a([[:alpha:]])|\b([[:alpha:]])|`).FindAllStringSubmatch(s, -1)
	//strs2 := regexp.MustCompile(`^([[:alpha:]])|[ \-_]([[:alpha:]])`).FindAllStringSubmatch(s, -1)
	//fmt.Printf("s = %s\n", s)
	//fmt.Printf("strs2 = %s\n", strs2)


        for _, word := range words {
		//fmt.Printf("word = %s\n", word)
                if len(word) > 0 {
			buf.WriteByte(word[0])
		}
        }

        return strings.ToUpper(buf.String())
}
