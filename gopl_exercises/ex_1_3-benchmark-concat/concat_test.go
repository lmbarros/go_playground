package concat

import "testing"

func TestSimple(t *testing.T) {
	args := []string{"dummy", "bla", "ble", "bli"}

	if concatenatingConcat(args) != "bla ble bli" {
		t.Error("concatenatingConcat failed!")
	}

	if joiningConcat(args) != "bla ble bli" {
		t.Error("joiningConcat failed!")
	}
}

var benchArgs = []string{"dummy", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func BenchmarkConcatenating(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatenatingConcat(benchArgs)
	}
}

func BenchmarkJoining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joiningConcat(benchArgs)
	}
}
