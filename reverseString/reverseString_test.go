package reverseString

import (
	"testing"

	"github.com/abhishekvaid/leetcode-golang/reverseString/util"
)

func TestReverseString(t *testing.T) {

	t.Run("reverse a non empty string", func(t *testing.T) {
		got := ReverseString("Abhishek")
		want := "kehsihbA"
		util.AssumeSame(t, got, want)
	})

}
