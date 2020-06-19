package leetcode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "[1,2,3,null,null,4,5]"},
		{name: "[3,2,4,1]"},
		{name: "[1,2]"},
		{name: "[]"},
	}

	obj := Constructor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := obj.serialize(obj.deserialize(tt.name))
			if got != tt.name {
				t.Errorf("%v, want %v", got, tt.name)
			}
		})
	}
}
