package textseach

const (
	d = 52
	p = 65713
)

var h = 1

func hInit(strLen int) int {
	h = 1
	for i := 1; i < strLen; i++ {
		h = (h * d) % p
	}

	return h
}

func ringHash(str []byte) int {
	strLen := len(str)
	if h == 0 {
		h = hInit(strLen)
	}
	hash := 0
	for i := 0; i < strLen; i++ {
		hash += (d*hash + int(str[i])) % p
	}

	if hash < 0 {
		hash += p
	}

	return hash
}

func ringReHash(str []byte, prev int) int {
	strLen := len(str)
	if h == 0 {
		h = hInit(strLen)
	}

	hash := (d * (prev - int(str[0])*h)) + int(str[strLen-1])%p
	if hash < 0 {
		hash += p
	}

	return hash
}

func RabinKarp(text, str string) int {
	strByteSlice, textByteSlice := []byte(str), []byte(text)
	strLen, textLen := len(strByteSlice), len(textByteSlice)
	strHash, textHash := ringHash(strByteSlice), ringHash(textByteSlice)

	for k := 0; k < textLen-strLen; k++ {
		if strHash == textHash {
			for i := 0; i < strLen; i++ {
				if i == strLen-1 {
					return k
				}
			}
		}
		textHash = ringReHash(textByteSlice[k:], textHash)
	}

	return -1

}
