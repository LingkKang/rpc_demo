package protocol

import "testing"

type ChecksumTestCase struct {
	name     string
	input    []byte
	checksum byte
}

var testCases = []ChecksumTestCase{
	{"single element", []byte{5}, 5},
	{"multiple elements", []byte{1, 2, 3, 4, 5}, 1 ^ 2 ^ 3 ^ 4 ^ 5},
	{"some empty elements", []byte{0, 0, 0}, 0},
}

func TestGenerateChecksum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run_result := generateChecksum(tc.input)
			if run_result != tc.checksum {
				t.Errorf(
					"expected %d, got %d",
					run_result, tc.checksum)
			}
		})
	}
}

func TestValidateChecksum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run_result := validateChecksum(tc.input, tc.checksum)
			if !run_result {
				t.Errorf(
					"validateChecksum(%v, %d) should be true",
					tc.input, tc.checksum)
			}
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_invalid", func(t *testing.T) {
			run_result := validateChecksum(tc.input, tc.checksum+1)
			if run_result {
				t.Errorf(
					"validateChecksum(%v, %d) should be false",
					tc.input, tc.checksum)
			}
		})
	}
}
