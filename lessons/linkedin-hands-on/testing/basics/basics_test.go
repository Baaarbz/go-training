package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := sum(1, 2, 3)
	want := 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("=======  Set-up  =======")
	m.Run()
	fmt.Println("======= Teardown =======")
}
