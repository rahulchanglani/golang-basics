package main

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(15,5)
	if ans != 20 {
		t.Errorf("add was incorrect, got %d, want %d", ans, 20)
	}
}