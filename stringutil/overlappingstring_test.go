package stringutil

import (
	"testing"
)

func TestCommonStr1(t *testing.T) {
	got := CommonStr("HELLO", "HE")
	want := "HE"
	if got != want {
		t.Errorf("Wrong Answer! The Correct answer is %s and your output is %s", want, got)
	}
}

func TestCommonStr2(t *testing.T) {
	got := CommonStr("HELLOW", "HELL")
	want := "HELL"
	if got != want {
		t.Errorf("Wrong Answer! The Correct answer is %s and your output is %s", want, got)
	}
}

func TestCommonStr3(t *testing.T) {
	got := CommonStr("", "")
	want := ""
	if got != want {
		t.Errorf("Wrong Answer! The Correct answer is %s and your output is %s", want, got)
	}
}

func TestCommonStr4(t *testing.T) {
	got := CommonStr("abc", " ")
	want := ""
	if got != want {
		t.Errorf("Wrong Answer! The Correct answer is %s and your output is %s", want, got)
	}
}

func TestCommonStr5(t *testing.T) {
	got := CommonStr("abc", "a")
	want := "a"
	if got != want {
		t.Errorf("Wrong Answer! The Correct answer is %s and your output is %s", want, got)
	}
}

func TestCommonStr6(t *testing.T) {
	got := CommonStr("abcjshdsj", "cdjsjd")
	want := "c"
	if got != want {
		t.Errorf("Wrong Answer! The Correct answer is %s and your output is %s", want, got)
	}
}
