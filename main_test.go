package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	type args struct {
// 		s string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				s: "Hello, world",
// 			},
// 			want: "dlrow ,olleH",
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				s: " ",
// 			},
// 			want: " ",
// 		},
// 		{
// 			name: "3",
// 			args: args{
// 				s: "!12345",
// 			},
// 			want: "54321!",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Reverse(tt.args.s); got != tt.want {
// 				t.Errorf("Reverse() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
