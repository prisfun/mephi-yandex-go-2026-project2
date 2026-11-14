package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, 7, 8, 9, 1000} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			data := generateRandomElements(size)
			if len(data) != max(size, 0) {
				t.Fatalf("length = %d, want %d", len(data), max(size, 0))
			}
			for index, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", index, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	testMaximum(t, maximum)
}

func TestMaxChunks(t *testing.T) {
	testMaximum(t, maxChunks)
}

func testMaximum(t *testing.T, findMaximum func([]int) int) {
	t.Helper()
	maxInt := int(^uint(0) >> 1)
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"first", []int{99, 3, 2, 1}, 99},
		{"middle", []int{1, 99, 2}, 99},
		{"last", []int{1, 2, 3, 99}, 99},
		{"duplicates", []int{7, 7, 7, 7, 7, 7, 7, 7}, 7},
		{"zeros", []int{0, 0, 0}, 0},
		{"negative singleton", []int{-5}, -5},
		{"negative", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
		{"mixed", []int{-7, 0, 5, -1, 2}, 5},
		{"integer limits", []int{-maxInt - 1, maxInt}, maxInt},
		{"minimum integer", []int{-maxInt - 1}, -maxInt - 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before := slices.Clone(test.data)
			if got := findMaximum(test.data); got != test.want {
				t.Errorf("maximum = %d, want %d", got, test.want)
			}
			if !slices.Equal(test.data, before) {
				t.Error("input slice was modified")
			}
		})
	}
}

func TestMaxChunksCoversEveryElement(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 15, 16, 17, 65} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			for position := range size {
				data := make([]int, size)
				for index := range data {
					data[index] = -2
				}
				data[position] = -1
				if got := maxChunks(data); got != -1 {
					t.Fatalf("maximum at index %d: got %d, want -1", position, got)
				}
			}
		})
	}
}
