package reverseWords

import (
	"testing"

	"github.com/abhishekvaid/leetcode-golang/reverseString/util"
)

func TestReverseWords(t *testing.T) {
	t.Run("reverse words of a single word sentence", func(t *testing.T) {
		got := ReverseWords("Hello")
		want := "olleH"
		util.AssumeSame(t, got, want)
	})

	t.Run("reverse words of multi word sentence", func(t *testing.T) {
		got := ReverseWords("Let's take LeetCode contest")
		want := "s'teL ekat edoCteeL tsetnoc"
		util.AssumeSame(t, got, want)
	})

	t.Run("reverse words of an empty sentence", func(t *testing.T) {
		got := ReverseWords("")
		want := ""
		util.AssumeSame(t, got, want)
	})

	t.Run("reverse words of multi word sentence with trailing spaces", func(t *testing.T) {
		got := ReverseWords("Let's take LeetCode contest    ")
		want := "s'teL ekat edoCteeL tsetnoc"
		util.AssumeSame(t, got, want)
	})

	t.Run("reverse words of multi word sentence with leading spaces", func(t *testing.T) {
		got := ReverseWords("       Let's take LeetCode contest    ")
		want := "s'teL ekat edoCteeL tsetnoc"
		util.AssumeSame(t, got, want)

	})
}
