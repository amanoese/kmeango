package pkg

import (
	"fmt"
  "strconv"
	"testing"
)

func TestInitKpoint(t *testing.T) {
  v1 := Point{ Values: []int{1} }
  v2 := Point{ Values: []int{2} }
  v3 := Point{ Values: []int{3} }

	points := []Point{v1,v2,v3}
  kpoints := Init_kpoint(points,1)

  t.Log(kpoints)

	if len(kpoints) != 1 {
		t.Fatal("failed test")
  }
}

func TestExampleSuccess(t *testing.T) {
  v1 := Point{ Values: []int{1,2} }
	v2 := Point{ Values: []int{1,2} }

	points := []Point{v1,v2}
	fmt.Print(points)
	result := Calc(points,1)

	if len(result) == 0 {
		t.Fatal("failed test")
	}
}

func Test4meansSuccess(t *testing.T) {
  points := make([]Point,4*20)

  for i,_ := range points {
    points[i] = Point{ Values: []int{i,i} }
  }

	fmt.Print(points)
	result := Calc(points,4)

  for _, point := range result {
    category := ""
    values := ""
    for _,str := range point.CategoryValue {
      category = category + " " + strconv.Itoa(str)
    }
    for _,str := range point.Values {
      values = values + " " + strconv.Itoa(str)
    }
    //fmt.Println(category,":",values)
  }

	if len(result) == 0 {
		t.Fatal("failed test")
	}
}
