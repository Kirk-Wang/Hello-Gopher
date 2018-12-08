package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinses support
		{"这里是中国", 5},
		{"一二三三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8
	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer() // 上面我们准备数据的时间不算

	// 既然是benchmark，做一遍肯定不够，得做很多遍，具体多少遍，是有具体的算法的。
	// 在这里，不用我们操心，系统会告诉我们多少遍
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
