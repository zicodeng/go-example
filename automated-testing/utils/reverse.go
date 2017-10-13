package utils

// Reverse returns the reverse of the string passed as `s`
func Reverse(s string) string {
	// Convert string to a slice so we can manipulate it
	// since strings are immutable, this creates a copy of
	// the string so we won't be modifying the original.
	// Bug: this only works for characters within ASCII range
	// because Go uses UTF-8 encoding for strings,
	// and each ASCII character only takes up only one byte (8 bits),
	// but this will not work for characters beyond ASCII range such as 世界.
	// A non-ASCII character may consume anywhere from 1 to 4 bytes depending on its Unicode value.
	// The safest way is to convert it to a slice of runes, which is a complete Unicode character.
	chars := []rune(s)

	// Starting from each end, swap the values in the slice
	// elements, stopping when we get to the middle.
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	//return the reversed slice as a string
	return string(chars)
}
