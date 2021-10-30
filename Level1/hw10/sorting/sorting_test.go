package sorting

import (
	"reflect"
	"testing"
)

var test1 = []int{100, 50, 11, 21}
var expectTest1 = []int{11, 21, 50, 100}
var test2 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
var expectTest2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestSorting(t *testing.T) {
	a := sortingByInsert(test1)
	b := expectTest1
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Test fail")
	}
}

func TestSorting2(t *testing.T) {
	a := sortingByInsert(test2)
	b := expectTest2
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Test fail")
	}
}
