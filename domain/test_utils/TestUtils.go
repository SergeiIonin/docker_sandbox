package test_utils

import (
	"slices"
	"testing"
)

type TestUtils struct {
}

func NewTestUtils() *TestUtils {
	return &TestUtils{}
}

func (tu *TestUtils) CompareSlices(left []string, right []string, expectedEqual bool, t *testing.T) {
	res := slices.Equal(left, right)
	if expectedEqual && !res {
		t.Fatalf("slices %v and %v aren't equal as expected", left, right)
	}
	if !expectedEqual && res {
		t.Fatalf("slices %v and %v are equal, but different slices expected", left, right)
	}
	return
}
