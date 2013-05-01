package dtl

import (
  "math"
)

type Tree struct {
  root Node
}

type Node struct {
  name string
  children []Node
  examples [][]string
}

func DTL(examples [][]string) {

}

func information(x float64, y float64) float64 {
  if x <= 0.0 {
    x = 0.0000001
  }
  if y <= 0.0 {
    y = 0.0000001
  }
  return math.Abs(-x * math.Log2(x) - y * math.Log2(y))
}

func gain(total float64, examples [][]string) float64 {
  var good, bad int = 0, 0
  for _, ex := range examples {
    last := ex[len(ex) - 1]
    if last == "true" {
      good++
    } else {
      bad++
    }
  }
  count := float64(good + bad)
  factor := count / float64(total)
  x := float64(good) / count
  y := float64(bad) / count

  return factor * information(x, y)
}

func entropy(totalExamples float64, node Node) float64{
  var gains []float64
  for _, child := range node.children {
    gains = append(gains, gain(totalExamples, child.examples))
  }
  sum := float64(0.0)
  for _, val := range gains {
      sum += val
  }
  return float64(1.0 - sum)
}
