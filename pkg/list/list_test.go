package list

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	var act []int
	var exp []int
	
	act = []int{0, 1, 2, 3, 4, 5}
	act = remove(act, 1)
	exp = []int{0, 2, 3, 4, 5}
	if ! reflect.DeepEqual(act, exp) {
		t.Fatalf("failed test %v", act)
	}
	
	act = []int{0, 1, 2, 3, 4, 5}
	act = remove(act, 0)
	exp = []int{1, 2, 3, 4, 5}
	if ! reflect.DeepEqual(act, exp) {
		t.Fatalf("failed test %#v", act)
	}
	
	act = []int{0, 1, 2, 3, 4, 5}
	act = remove(act, 5)
	exp = []int{0, 1, 2, 3, 4}
	if ! reflect.DeepEqual(act, exp) {
		t.Fatalf("failed test %#v", act)
	}
	
	act = []int{0, 1, 2, 3, 4, 5}
	act = remove(act, 0)
	act = remove(act, 0)
	act = remove(act, 3)
	exp = []int{2, 3, 4}
	if ! reflect.DeepEqual(act, exp) {
		t.Fatalf("failed test %#v", act)
	}
}

func TestMax(t *testing.T) {
	if i, n := max([]int{0, 1, 2, 3, 4, 5}); i != 5 || n != 5 {
		t.Fatalf("failed test")
	}
	if i, n := max([]int{5, 4, 3, 2, 1, 0}); i != 0 || n != 5 {
		t.Fatalf("failed test")
	}
	if i, n := max([]int{0, 2, 4, 3, 5, 1}); i != 4 || n != 5 {
		t.Fatalf("failed test")
	}
	if i, n := max([]int{0, 50, 10, 30, 40, 20}); i != 1 || n != 50 {
		t.Fatalf("failed test")
	}
	if i, n := max([]int{50, 50, 50, 50, 50, 50}); i != 0 || n != 50 {
		t.Fatalf("failed test")
	}
}
