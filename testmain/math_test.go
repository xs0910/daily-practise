package testmain

import (
	"fmt"
	"math"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("do some setup")
	m.Run()
	fmt.Println("do some cleanup")
}

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func TestMax(t *testing.T) {
	got := math.Max(1, 2)
	if got != 2 {
		t.Errorf("Max(1,2) = %f; want 2", got)
	}
}
