package protocol

func generateChecksum(data []byte) byte {
	var checksum byte = 0
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

func validateChecksum(data []byte, checksum byte) bool {
	return generateChecksum(data) == checksum
}
