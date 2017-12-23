package main

import "testing"

func Test_solveCaptcha(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Test cases from the problem description
		{
			name: "Simple match",
			args: args{"1122"},
			want: "3", // 1 + 2 = 3
		},
		{
			name: "All match",
			args: args{"1111"},
			want: "4", // each one matches next
		},
		{
			name: "No match",
			args: args{"1234"},
			want: "0",
		},
		{
			name: "Match at end",
			args: args{"91212129"},
			want: "9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveCaptcha(tt.args.s); got != tt.want {
				t.Errorf("solveCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
