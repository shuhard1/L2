package main

import "testing"

func TestStrUnpacking(t *testing.T) {
	cases := []string{
		"a4bc2d5e",
		"abcd",
		"456",
		"",
		"a1b2c1",
	}

	results := []string{
		"aaaabccddddde",
		"abcd",
		"",
		"",
		"abbc",
	}

	for i := range cases {
		s, _ := StrUnpacking(cases[i])
		if s != results[i] {
			t.Errorf("got %s, want %s", results[i], s)
		}
	}
}

func TestStrUnpackingWithSlash(t *testing.T) {
	cases := []string{
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
		`\qwe\\`,
		`qwe45\5`,
		`qwe\44`,
		`\`,
		`\\`,
		`\9\`,
		`4\\\5`,
		`\\\\\4`,
	}

	results := []string{
		"qwe45",
		"qwe44444",
		`qwe\\\\\`,
		`qwe\`,
		"qweeee44445",
		"qwe4444",
		"",
		`\`,
		"9",
		"",
		`\\4`,
	}

	for i := range cases {
		s, _ := StrUnpacking(cases[i])
		if s != results[i] {
			t.Errorf("got %s, want %s", results[i], s)
		}
	}
}
