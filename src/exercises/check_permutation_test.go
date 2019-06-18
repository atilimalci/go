package exercises

import (
	"testing"
)

func Test_PermutationBySorting_CorrectInput_GetsTrue(t *testing.T) {
	s1, s2 := "atilim", "milita"

	result := PermutationBySorting(s1, s2)

	if !result {
		t.Errorf("Result is incorrect.")
	}

}

func Test_PermutationBySorting_BadInput_GetsFalse(t *testing.T) {
	s1, s2 := "atilim", "alci"

	result := PermutationBySorting(s1, s2)

	if result {
		t.Errorf("Result is incorrect.")
	}

}
