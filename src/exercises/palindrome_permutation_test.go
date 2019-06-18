package exercises

import (
	"testing"
)

func Test_PalindromePermutation_CorrectInput_GetsTrue(t *testing.T) {
	str := "affai"
	v := IsPermutationOfPalindrome(str)

	if !v {
		t.Errorf("Result is incorrect.")
	}
}

func Test_PalindromePermutation_WrongInput_GetsFalse(t *testing.T) {
	str := "affaui"
	v := IsPermutationOfPalindrome(str)

	if v {
		t.Errorf("Result is incorrect.")
	}
}
