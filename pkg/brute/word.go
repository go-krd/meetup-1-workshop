package brute

// NextWord returns next word after word.
func NextWord(word []byte) {
	for i := len(word) - 1; i >= 0; i-- {
		word[i] = NextByte(word[i])
		if word[i] != '0' {
			return
		}
	}
}
