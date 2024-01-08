package protocol

import (
	"reflect"
	"testing"
)

// Test case struct for converting message payload.
type PayloadConversionTestCase struct {
	name    string
	payload []byte
	floats  []float64
}

func TestParsePayloadToFloat64s_normal(t *testing.T) {
	testCases := []PayloadConversionTestCase{
		{
			name: "single flaot",
			// 00111111 11110000 00000000 00000000 00000000 00000000 00000000 00000000
			payload: []byte{63, 240, 0, 0, 0, 0, 0, 0},
			floats:  []float64{1.0},
		}, {
			name: "impercise float",
			// 00111111 10111001 10011001 10011001 10011001 10011001 10011001 10011010
			payload: []byte{63, 185, 153, 153, 153, 153, 153, 154},
			floats:  []float64{0.1000000000000000055511151231257827021181583404541015625},
		}, {
			name: "two floats",
			payload: []byte{
				// 01000000 00101000 00000000 00000000 00000000 00000000 00000000 00000000
				64, 40, 0, 0, 0, 0, 0, 0,
				// 01000000 00101010 00000000 00000000 00000000 00000000 00000000 00000000
				64, 42, 0, 0, 0, 0, 0, 0,
			},
			floats: []float64{12.0, 13.0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results, err := ParsePayloadToFloat64s(tc.payload)
			if err != nil {
				t.Error("unexpected fail ", err.Error())
			}
			for i, result := range results {
				float := tc.floats[i]
				if !reflect.DeepEqual(result, float) {
					t.Errorf("expected %f, got %f", float, result)
				}
			}
		})
	}
}

func TestParsePayloadToFloat64s_error(t *testing.T) {
	testCases := []PayloadConversionTestCase{
		{
			name: "not a number",
			// 01111111 11111000 00000000 00000000 00000000 00000000 00000000 00000000
			payload: []byte{127, 248, 0, 0, 0, 0, 0, 0},
			floats:  nil,
		}, {
			name:    "invalid payload length",
			payload: []byte{1, 3, 5, 7, 9},
			floats:  nil,
		}, {
			name: "infinity",
			// 01111111 11110000 00000000 00000000 00000000 00000000 00000000 00000000
			payload: []byte{127, 240, 0, 0, 0, 0, 0, 0},
			floats:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParsePayloadToFloat64s(tc.payload)
			if err == nil {
				t.Errorf("expected an error, got nil")
			}
		})
	}
}
