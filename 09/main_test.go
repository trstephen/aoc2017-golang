package main

import "testing"

func Test_scoreGroups(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single group",
			args: args{"{}"},
			want: 1,
		}, {
			name: "simple nesting",
			args: args{"{{{}}}"},
			want: 6,
		}, {
			name: "horizontal nesting",
			args: args{"{{},{}}"},
			want: 5,
		}, {
			name: "uneven horizontal nesting",
			args: args{"{{{},{},{{}}}}"},
			want: 16,
		}, {
			name: "simple garbage",
			args: args{"{<a>,<a>,<a>,<a>}"},
			want: 1,
		}, {
			name: "nested groups and garbage",
			args: args{"{{<ab>},{<ab>},{<ab>},{<ab>}}`"},
			want: 9,
		}, {
			name: "cancel the cancelation",
			args: args{"{{<!!>},{<!!>},{<!!>},{<!!>}}"},
			want: 9,
		}, {
			name: "garbage cancelation",
			args: args{"{{<a!>},{<a!>},{<a!>},{<ab>}}`"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreGroups(tt.args.s); got != tt.want {
				t.Errorf("scoreGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
