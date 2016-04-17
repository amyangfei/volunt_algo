/*
1.1 单词翻转
*/
package string

import (
	"testing"
)

func TestReverseWord(t *testing.T) {
	test1 := "I am a student."
	test2 := "I am   a good     student."
	test3 := "   I am   a good     student.   "
	if ReverseWord(test1) != "student. a am I" {
		t.Errorf("err for '%s'", test1)
	}
	if ReverseWord(test2) != "student.     good a   am I" {
		t.Errorf("err for '%s'", test2)
	}
	if ReverseWord(test3) != "   student.     good a   am I   " {
		t.Errorf("err for '%s'", test3)
	}
}
