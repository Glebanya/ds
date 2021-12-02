package textseach

import "fmt"

const charSize int = 256

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func BoyerMooreTextSeach(text, pattern string) {
	badChars := [charSize]int{}
	patternLen := len(pattern)

	for i := 0; i < charSize; i++ {
		badChars[i] = -1
	}

	textLen := len(text)
	for i := 0; i < textLen; i++ {
		badChars[text[i]] = i
	}

	for s := 0; textLen-patternLen >= s; {
		var j int
		for j = patternLen - 1; j > 0 && pattern[j] == text[s+j]; j-- {
		}
		if j < 0 {
			fmt.Printf("find pattern at %d", s)
			if textLen > s+patternLen {
				s += 1
			} else {
				s += patternLen - badChars[text[s+patternLen]]
			}
		} else {
			s = max(1, j-badChars[text[s+j]])
		}
	}

}
