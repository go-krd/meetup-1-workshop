package brute

// NextByte returns next byte after b.
func NextByte(b byte) byte {
	switch b {
	case 'z':
		return '0'
	case '9':
		return 'a'
	default:
		return b + 1
	}
}
