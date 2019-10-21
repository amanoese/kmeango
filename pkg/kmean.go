package pkg

import (
  //"fmt"
  "math"
  "math/big"
  "math/rand"
  crand "crypto/rand"
)

type Point struct {
	Values []int
}

func (point *Point) setValues(values []int) {
    newValues := make([]int,len(values))
    copy(newValues,values)
    point.Values = newValues
}

//計算用のstruct
type CategoryPoint struct {
	Values []int
	category int
}

//結果用のstruct
type CategorizedPoint struct {
	Values []int
	CategoryNum int
	CategoryValue []int
}

//クラスタ中心の初期化
func Init_kpoint(points []Point,k int) []Point {

  _points := make([]Point,len(points))
  copy(_points,points)
  kpoints := make([]Point,k)

  seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
  rand.Seed(seed.Int64())

  //pointsにある要素からランダムにk個の要素を取り出す
  for i :=0; i < k ; i++ {
    use_index := rand.Intn(len(_points) - 1)
    //fmt.Printf("use_index:%d\n",use_index)

    var kpoint Point;
    kpoint.setValues(_points[use_index].Values)
    kpoints[i] = kpoint

    _points = append(_points[:use_index],_points[use_index+1:]...)
  }
  return kpoints
}

func Calc(points []Point,k int) []CategorizedPoint {
  //クラスタ中心
  clusters := Init_kpoint(points,k)

  //fmt.Println(clusters)
  //適当に初期化
  clusteryPoints := make([]CategoryPoint,cap(points))
	for i, point := range points {
    clusteryPoints[i] = CategoryPoint{ point.Values, 0 }
  }

  for loop := 0 ; loop < 100 ; loop = loop + 1 {
    for cp_i, point := range clusteryPoints {
      min := math.MaxFloat64

      for num, cluster_point := range clusters {
         diff := 0.0
         for ci, cvalue := range cluster_point.Values {
            diff = diff + math.Pow(float64(point.Values[ci] - cvalue),2)
         }
         if diff < min  {
            min = diff
            clusteryPoints[cp_i].category = num
         }
      }
    }
    for ci, cluster_point := range clusters {
      count := 0
      sump   := make([]int,cap(cluster_point.Values))
      for _, point := range clusteryPoints {
        if ci == point.category {
          count = count + 1
          for i, sum:= range sump {
            sump[i] = sum + point.Values[i]
          }
        }
      }
      for i, sum:= range sump {
        sump[i] = int(sum/count)
      }
      clusters[ci].Values = sump
    }
  }

	cpoints := []CategorizedPoint{}

	for _, point := range clusteryPoints {

		cpoint := CategorizedPoint{
			Values:   point.Values,
			CategoryNum: point.category,
			CategoryValue: clusters[point.category].Values,
		}

		cpoints = append(cpoints,cpoint)
	}
	return cpoints
}
