package archive

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCreateFromBuffer(t *testing.T) {
	result, err := createFromBuffer([]byte("okok"), "a.csv", 7000)
	expected := []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 74, 212, 75, 46, 46, 99, 160, 45, 48, 48, 48, 52, 55, 52, 49,
		100, 48, 128, 0, 116, 218, 192, 192, 192, 4, 137, 109, 192, 96, 96, 104, 96, 104, 108, 194, 160, 96, 64, 99,
		119, 129, 65, 105, 113, 73, 98, 17, 131, 1, 197, 118, 161, 123, 110, 136, 128, 252, 236, 252, 236, 129, 118,
		195, 40, 24, 5, 163, 96, 20, 140, 2, 250, 3, 64, 0, 0, 0, 255, 255, 236, 190, 228, 172, 0, 8, 0, 0}

	if err != nil {
		t.Error("Error not expected")
	}
	resultB := result.Bytes()
	fmt.Println(resultB)
	if bytes.Compare(resultB, expected) != 0 {
		t.Errorf("Expected slice %q, got %q", expected, resultB)
	}
}
