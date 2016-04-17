package string

import (
	"testing"
)

func TestManacher(t *testing.T) {
	test1 := "forgeeksskeegfor"
	if Manacher(test1) != "geeksskeeg" {
		t.Errorf("err for '%s'", test1)
	}

	test2 := "abaaba"
	if Manacher(test2) != "abaaba" {
		t.Errorf("err for '%s'", test2)
	}

	test3 := "12212321"
	if Manacher(test3) != "12321" {
		t.Errorf("err for '%s'", test3)
	}

	test4 := "00"
	if Manacher(test4) != "00" {
		t.Errorf("err for '%s'", test4)
	}

	test5 := "0"
	if Manacher(test5) != "0" {
		t.Errorf("err for '%s'", test5)
	}

	test6 := "11123321121"
	if Manacher(test6) != "11233211" {
		t.Errorf("err for '%s'", test6)
	}
}
