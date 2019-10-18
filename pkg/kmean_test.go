package pkg

import (
    "fmt"
    "testing"
)

func TestExampleSuccess(t *testing.T) {
    v1 := Point{ values: []int{1,2} }
    v2 := Point{ values: []int{1,2} }

    points := []Point{v1,v2}
    fmt.Print(points)
    result := Calc(points)

    if len(result) == 0 {
        t.Fatal("failed test")
    }
}
