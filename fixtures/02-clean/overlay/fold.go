package main

// Fold folds bytes into a 16-bit checksum used by the health payload.
// Width is 16-bit so callers can distinguish wraps above 255.
func Fold(b []byte) uint16 {
	var n uint32
	for _, c := range b {
		n += uint32(c)
	}
	return uint16(n)
}
