package main

import (
	"testing"
)

func TestMinMaxSum(t *testing.T) {
	type args struct {
		origArr []int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "All positive numbers",
			args:    args{[]int{40, 19, 84, 84, 0}},
			want:    143,
			want1:   227,
			wantErr: false,
		},
		{
			name:    "All negative numbers",
			args:    args{[]int{-20, -19, -4, -8, -7}},
			want:    -54,
			want1:   -38,
			wantErr: false,
		},
		{
			name:    "All zeros",
			args:    args{[]int{0, 0, 0, 0, 0}},
			want:    0,
			want1:   0,
			wantErr: false,
		},
		{
			name:    "Array of 8 numbers",
			args:    args{[]int{12, 5, -1, 84, 84, 0, -98, 160}},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name:    "Array of 3 numbers",
			args:    args{[]int{12, 5, -1}},
			want:    0,
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := MinMaxSum(tt.args.origArr)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinMaxSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinMaxSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinMaxSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func BenchmarkMinMaxSum(b *testing.B) {
	arr := []int{40, 19, 84, 84, 0}
	for i := 0; i < b.N; i++ {
		MinMaxSum(arr)
	}

}

func BenchmarkMinMaxSumOld2(b *testing.B) {
	arr := []int{40, 19, 84, 84, 0}
	for i := 0; i < b.N; i++ {
		MinMaxSumOld2(arr)
	}

}

func BenchmarkMinMaxSumOld1(b *testing.B) {
	arr := [5]int{40, 19, 84, 84, 0}
	for i := 0; i < b.N; i++ {
		MinMaxSumOld1(arr)
	}

}
