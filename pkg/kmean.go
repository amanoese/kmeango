package pkg

type Point struct {
  values []int
}

type CategorizedPoint struct {
  values []int
  category []int
}

func Calc(points []Point) []CategorizedPoint {
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
