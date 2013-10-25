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
    var good, bad float64 = 0.0, 0.0
    if total <= 0.0 {
        return 0.0
    }
    length := 0
    if len(examples) > 0 {
        length = len(examples[0])
    }
    if length <= 0 {
        return 0.0
    }
    for _, ex := range examples {
        last := ex[length - 1]
        if last == "true" {
            good++
        } else {
            bad++
        }
    }
    count := good + bad
    factor := count / total
    x := good / count
    y := bad / count

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
