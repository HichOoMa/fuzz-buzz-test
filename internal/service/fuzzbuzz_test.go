package service

import (
	"reflect"
	"testing"
)

func TestFuzzBuzz(t *testing.T) {
	tests := []struct {
		name  string
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
		want  []string
	}{
		{
			name:  "classic fizzbuzz to 15",
			int1:  3,
			int2:  5,
			limit: 15,
			str1:  "fizz",
			str2:  "buzz",
			want: []string{
				"1", "2", "fizz", "4", "buzz",
				"fizz", "7", "8", "fizz", "buzz",
				"11", "fizz", "13", "14", "fizzbuzz",
			},
		},
		{
			name:  "custom strings",
			int1:  2,
			int2:  3,
			limit: 6,
			str1:  "foo",
			str2:  "bar",
			want:  []string{"1", "foo", "bar", "foo", "5", "foobar"},
		},
		{
			name:  "limit smaller than both divisors",
			int1:  3,
			int2:  5,
			limit: 2,
			str1:  "fizz",
			str2:  "buzz",
			want:  []string{"1", "2"},
		},
		{
			name:  "limit of zero returns empty slice",
			int1:  3,
			int2:  5,
			limit: 0,
			str1:  "fizz",
			str2:  "buzz",
			want:  []string{},
		},
		{
			name:  "int1 equal to int2",
			int1:  3,
			int2:  3,
			limit: 6,
			str1:  "fizz",
			str2:  "buzz",
			want:  []string{"1", "2", "fizzbuzz", "4", "5", "fizzbuzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FuzzBuzz(tt.int1, tt.int2, tt.limit, tt.str1, tt.str2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FuzzBuzz(%d, %d, %d, %q, %q) = %v, want %v",
					tt.int1, tt.int2, tt.limit, tt.str1, tt.str2, got, tt.want)
			}
		})
	}
}
