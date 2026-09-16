package main

// Fold folds bytes into a 16-bit checksum used by the health payload.
func Fold(b []byte) uint16 {
	var n uint32
	for _, c := range b {
		n += uint32(c)
	}
	return uint16(n)
}
