package dtl


import (
  "testing"
)

func TestInformation(t *testing.T) {
  var x float64 = 1.0
  var y float64 = 1.0

  result := information(x, y)
  if result != 0 {
    t.Errorf("Information should be 0")
  }
  x = 0.5
  y = 0.2
  result = information(x, y)
  var expected float64 = 0.9643856189774724
  if result != expected {
    t.Errorf("Information should be %f", expected)
  }
}

func TestGain(t *testing.T) {
  total := float64(3.0)
  examples := [][]string{ []string{"1", "true"}, []string{"2", "true"}, []string{"3", "false"} }
  result := gain(total, examples)
  expected := float64(0.9182958340544896)
  if result != expected {
    t.Errorf("Gain should be %f", expected)
  }
}

func TestEntropy(t *testing.T) {
  examples := [][]string{ []string{"1", "true"}, []string{"2", "true"}, []string{"3", "false"} }
  children := []Node { Node{ name: "test2", examples: examples } }
  node := Node{ name: "test", examples: examples, children: children}

  result := entropy(float64(3.0), node)
  expected := float64(0.08170416594551044)
  if result != expected {
    t.Errorf("Entropy should be %f", expected)
  }

  result = entropy(float64(10.0), node)
  expected = float64(0.7245112497836532)
  if result != expected {
    t.Errorf("Entropy should be %f", expected)
  }
}
