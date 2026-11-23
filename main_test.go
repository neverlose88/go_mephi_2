package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero size", 0},
		{"negative size", -5},
		{"one element", 1},
		{"regular size", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if data == nil {
				t.Fatal("expected non-nil slice")
			}
			wantLen := tt.size
			if wantLen < 0 {
				wantLen = 0
			}
			if len(data) != wantLen {
				t.Fatalf("len = %d, want %d", len(data), wantLen)
			}
			for i, v := range data {
				if v <= 0 {
					t.Fatalf("element %d = %d, want positive", i, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element", []int{7}, 7},
		{"max at start", []int{9, 2, 3}, 9},
		{"max in middle", []int{1, 8, 3}, 8},
		{"max at end", []int{1, 2, 10}, 10},
		{"all equal", []int{4, 4, 4, 4}, 4},
		{"unsorted", []int{5, 1, 99, 3, 42, 7}, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{"nil slice", nil},
		{"empty slice", []int{}},
		{"one element", []int{7}},
		{"fewer than chunks", []int{3, 9, 1}},
		{"exactly chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"not divisible, max in remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}},
		{"divisible", generateRandomElements(CHUNKS * 100)},
		{"large random", generateRandomElements(100_003)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := maximum(tt.data)
			if got := maxChunks(tt.data); got != want {
				t.Fatalf("maxChunks = %d, want %d", got, want)
			}
		})
	}
}
