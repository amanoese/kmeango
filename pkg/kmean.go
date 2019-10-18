package pkg

import (
  "fmt"
  "math"
  "math/big"
  "math/rand"
  crand "crypto/rand"
)

type Point struct {
	values []int
}

func (point *Point) setValues(values []int) {
    newValues := make([]int,len(values))
    copy(newValues,values)
    point.values = newValues
}

type CategorizedPoint struct {
	values []int
	category []int
}

func Init_kpoint(points []Point,k int) []Point {

  _points := make([]Point,len(points))
  copy(_points,points)
  kpoints := make([]Point,k)

  seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
  rand.Seed(seed.Int64())

  for i :=0; i < k ; i++ {
    use_index := rand.Intn(len(_points) - 1)
    fmt.Printf("use_index:%d\n",use_index)

    var kpoint Point;
    kpoint.setValues(_points[use_index].values)
    kpoints[i] = kpoint

    _points = append(_points[:use_index],_points[use_index+1:]...)
  }
  return kpoints
}

func Calc(points []Point,k int) []CategorizedPoint {
	cpoints := []CategorizedPoint{}

	for _, point := range points {

		cpoint := CategorizedPoint{
			values:   point.values,
			category: point.values,
		}

		cpoints = append(cpoints,cpoint)
	}
	return cpoints
}
