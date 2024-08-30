package word_test

import (
	"testing"

	word "github.com/hungqd/go-learning/chap11/word1"
)

func TestIsPalindrome(t *testing.T) {
	if !word.IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !word.IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonIsPalindrome(t *testing.T) {
	if word.IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
