package frame

// IsValidFrameName checks whether the given byte array contains
// a valid frame name. A valid frame name must be in uppercase
// or contains at least one digit between 1 and 4
func IsValidFrameName(frame []byte) bool {
	if len(frame) != 4 && len(frame) != 3 {
		return false
	}

	for _, b := range frame {
		if (b < 'A' || b > 'Z') && (b < '1' || b > '4') {
			return false
		}
	}

	return true
}
