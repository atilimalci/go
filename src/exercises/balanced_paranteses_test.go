package exercises

import "testing"

/*
{name: "TestExpression_isParanthesesBalanced true",
			e:    "[()]{}{[()()]()}",
			want: true},
		{name: "TestExpression_isParanthesesBalanced false",
			e:    "[(])",
			want: false}}
*/

func TestIsParanthesesBalanced(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestExpression_isParanthesesBalanced true",
			args: args{s: "[()]{}{[()()]()}"},
			want: true},
		{name: "TestExpression_isParanthesesBalanced false",
			args: args{s: "[(])"},
			want: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParanthesesBalanced(tt.args.s); got != tt.want {
				t.Errorf("IsParanthesesBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
