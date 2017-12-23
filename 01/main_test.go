package main

import "testing"

func Test_solveCaptchaPart1(t *testing.T) {
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
			if got := solveCaptchaPart1(tt.args.s); got != tt.want {
				t.Errorf("solveCaptchaPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveCaptchaPart2(t *testing.T) {
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
			name: "All match 1",
			args: args{"1212"},
			want: "6", // (1 + 1) + (2 + 2)
		},
		{
			name: "All match 2",
			args: args{"123123"},
			want: "12", // (1 + 1) + (2 + 2) + (3 + 3)
		},
		{
			name: "No match",
			args: args{"1221"},
			want: "0",
		},
		{
			name: "Single match",
			args: args{"123425"},
			want: "4", // (2 + 2)
		},
		{
			name: "Double match",
			args: args{"12131415"},
			want: "4", // (1 + 1) + (1 + 1)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveCaptchaPart2(tt.args.s); got != tt.want {
				t.Errorf("solveCaptchaPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
