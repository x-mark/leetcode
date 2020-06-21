package leetcode

import (
	"reflect"
	"testing"
)

func Test_recoverFromPreorder(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{S: "1-2--3--4-5--6--7"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverFromPreorder(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
