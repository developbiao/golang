package main

import "testing"

// Unit test
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// normal cases
		{"hello", 3},
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"你好这个世界", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d",
				actual, tt.s, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))
	ans := 8
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;  expected %d", actual, s, ans)
		}
	}

}
