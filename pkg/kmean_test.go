package pkg

import (
	"fmt"
	"testing"
)

func TestInitKpoint(t *testing.T) {
  v1 := Point{ values: []int{1} }
  v2 := Point{ values: []int{2} }
  v3 := Point{ values: []int{3} }

	points := []Point{v1,v2,v3}
  kpoints := Init_kpoint(points,1)

  t.Log(kpoints)

	if len(kpoints) != 1 {
		t.Fatal("failed test")
  }
}

func TestExampleSuccess(t *testing.T) {
  v1 := Point{ values: []int{1,2} }
	v2 := Point{ values: []int{1,2} }

	points := []Point{v1,v2}
	fmt.Print(points)
	result := Calc(points,1)

	if len(result) == 0 {
		t.Fatal("failed test")
	}
}
