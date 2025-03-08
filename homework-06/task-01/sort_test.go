package main

import (
	"testing"
	"sort"
	"reflect"
)

// Простой тест для функции sort.Ints
func TestInts(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	want := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}
}