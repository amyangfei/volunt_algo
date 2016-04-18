package string

import (
	"testing"
)

func testManacherCase(t *testing.T, s string, expected string) {
	cals, start, length := Manacher(s)
	if cals != expected {
		t.Errorf("err for '%s'", s)
	}
	if s[start:start+length] != expected {
		t.Errorf("err for index or length")
	}
}

func TestManacher(t *testing.T) {
	s1 := "forgeeksskeegfor"
	e1 := "geeksskeeg"
	testManacherCase(t, s1, e1)

	s1 = "abaaba"
	e1 = "abaaba"
	testManacherCase(t, s1, e1)

	s1 = "00"
	e1 = "00"
	testManacherCase(t, s1, e1)

	s1 = "012134"
	e1 = "121"
	testManacherCase(t, s1, e1)

	s1 = "0"
	e1 = "0"
	testManacherCase(t, s1, e1)

	s1 = "11123321121"
	e1 = "11233211"
	testManacherCase(t, s1, e1)
}
