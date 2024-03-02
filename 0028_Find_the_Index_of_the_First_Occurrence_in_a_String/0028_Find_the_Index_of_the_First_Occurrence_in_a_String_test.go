package leetcode_0028_Find_the_Index_of_the_First_Occurrence_in_a_String

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				haystack: "xsadbutsad",
				needle:   "sad",
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "Test 3",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "Test 4",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "Test 5",
			args: args{
				haystack: "mississi",
				needle:   "issip",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
