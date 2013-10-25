package dtl

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type DT struct {
	Labels   []string
	Examples [][]float64
	Default  bool
	T        *Tree
}

type Tree struct {
	root Node
}

type Node struct {
	name     string
	children []Node
	examples [][]float64
}

func information(x float64, y float64) float64 {
	if x <= 0.0 {
		x = 0.0000001
	}
	if y <= 0.0 {
		y = 0.0000001
	}
	return math.Abs(-x*math.Log2(x) - y*math.Log2(y))
}

func gain(total float64, examples [][]float64) float64 {
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
		last := ex[length-1]
		if last == 1.0 {
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

func entropy(totalExamples float64, node Node) float64 {
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

func LoadExamples(filepath string) ([][]float64, []string) {
	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	var labels []string
	examples := [][]float64{}
	for scanner.Scan() {
		if len(labels) == 0 {
			labels = strings.Split(strings.Replace(scanner.Text(), "\"", "", -1), ",")
			labels = labels[:len(labels)-1]
			continue
		}
		ex := strings.Split(scanner.Text(), ",")
		converted := []float64{}
		for _, val := range ex {
			f, err := strconv.ParseFloat(val, 64)
			if err != nil {
				panic(err)
			}
			converted = append(converted, f)
		}
		examples = append(examples, converted)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return examples, labels
}

func (dt *DT) Majority() (result bool) {
	total := len(dt.Examples)
	if total == 0 {
		return dt.Default
	}
	good := 0
	bad := 0
	for _, ex := range dt.Examples {
		if ex[len(ex)-1] == 1.0 {
			good++
		} else {
			bad++
		}
	}
	if good > (total / 2) {
		result = true
	} else if bad > (total / 2) {
		result = false
	} else {
		result = dt.Default
	}
	return
}
