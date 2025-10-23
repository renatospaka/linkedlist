package linkedlist_test

import (
	"reflect"
	"testing"

	"github.com/renatospaka/linkedlist"
)

func TestNewListNode(t *testing.T) {
	node := linkedlist.NewListNode()
	if node == nil {
		t.Fatal("NewListNode returned nil")
	}
	if node.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", node.Next)
	}
	if node.Val != 0 {
		t.Errorf("Expected Val to be 0, got %v", node.Val)
	}
}

func TestPush(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{0}},
		{"single value", []int{5}, []int{5}},
		{"multiple values", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"negative values", []int{-1, -2}, []int{-1, -2}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := linkedlist.NewListNode()
			node.Push(tc.input)
			// Cover List() method for code coverage
			node.List()
			got := node.ToArray()
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Push: expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		adds     []int
		expected []int
	}{
		{"add to empty node", []int{10}, []int{0, 10}},
		{"add multiple", []int{10, 20}, []int{0, 10, 20}},
		{"add negative", []int{-5}, []int{0, -5}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := linkedlist.NewListNode()
			for _, v := range tc.adds {
				node.Add(v)
			}
			got := node.ToArray()
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Add: expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestToArray(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{0}},
		{"single node", []int{7}, []int{7}},
		{"multi node", []int{5, 6, 7}, []int{5, 6, 7}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := linkedlist.NewListNode()
			node.Push(tc.input)
			got := node.ToArray()
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ToArray: expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestToInverse(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"invert empty list", []int{}, []int{0}},
		{"invert single node", []int{42}, []int{42}},
		{"invert multi node", []int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := linkedlist.NewListNode()
			node.Push(tc.input)
			inv := node.ToInverse()
			got := inv.ToArray()
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ToInverse: expected %v, got %v", tc.expected, got)
			}
		})
	}
}
