package protocol

// Use bit-wise XOR to generate checksum.
func generateChecksum(data []byte) byte {
	var checksum byte = 0
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

// Verify the checksum over a slice of bytes.
func validateChecksum(data []byte, checksum byte) bool {
	return generateChecksum(data) == checksum
}
