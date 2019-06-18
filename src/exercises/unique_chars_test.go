package exercises

import (
	"testing"
)

func Test_IsUniqueChars_UniqueString_ReturnsTrue(t *testing.T) {
	str := "abcdefgh"
	result := IsUniqueChars(str)

	if !result {
		t.Errorf("Result is incorrect.")
	}

}

func Test_IsUniqueChars_NonUniqueString_ReturnsFalse(t *testing.T) {
	str := "abcdefgha"
	result := IsUniqueChars(str)

	if result {
		t.Error("Result is incorrect.")
	}
}
