package protodecode

import (
	"testing"
)

func TestIPv4(t *testing.T) {
	b := []byte{
		69, 0, 0, 126, 186, 180, 64, 0, 64, 6, 151, 204,
		192, 168, 0, 103,
		173, 194, 121, 39,
		187, 33, 0, 80, 130, 127, 178, 159, 110, 68, 148, 175, 128, 24, 0,
		229, 76, 61, 0, 0, 1, 1, 8, 10, 0, 63, 8, 48, 14, 52, 9, 77, 71, 69,
		84, 32, 47, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 85, 115, 101,
		114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55,
		46, 51, 53, 46, 48, 13, 10, 72, 111, 115, 116, 58, 32, 103, 111, 111,
		103, 108, 101, 46, 99, 111, 109, 13, 10, 65, 99, 99, 101, 112, 116,
		58, 32, 42, 47, 42, 13, 10, 13, 10,
	}

	packet := DecodeIPv4(b)
	if packet.Source.String() != "192.168.0.103" {
		t.Errorf("expected source IPv4 address %v, got %v", "192.168.0.103", packet.Source)
	}
}

func BenchmarkIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := []byte{
			69, 0, 0, 126, 186, 180, 64, 0, 64, 6, 151, 204,
			192, 168, 0, 103,
			173, 194, 121, 39,
			187, 33, 0, 80, 130, 127, 178, 159, 110, 68, 148, 175, 128, 24, 0,
			229, 76, 61, 0, 0, 1, 1, 8, 10, 0, 63, 8, 48, 14, 52, 9, 77, 71, 69,
			84, 32, 47, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 85, 115, 101,
			114, 45, 65, 103, 101, 110, 116, 58, 32, 99, 117, 114, 108, 47, 55,
			46, 51, 53, 46, 48, 13, 10, 72, 111, 115, 116, 58, 32, 103, 111, 111,
			103, 108, 101, 46, 99, 111, 109, 13, 10, 65, 99, 99, 101, 112, 116,
			58, 32, 42, 47, 42, 13, 10, 13, 10,
		}

		DecodeIPv4(b)
	}
}
