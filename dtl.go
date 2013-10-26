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

func entropy(examples [][]float64) (ent float64) {
    labelCounts := make(map[float64]float64)
    total := float64(len(examples))
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
		label := ex[length-1]
        count, exists := labelCounts[label]
        if !exists {
            labelCounts[label] = 0
        }
        labelCounts[label] = count + 1
	}

    for _, count := range labelCounts {
	    prob := count / total
        if prob <= 0.0 {
		    prob = 0.0000001
	    }
        ent -= prob * math.Log2(prob)
    }

	return 
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

func split(examples [][]float64, label int, value float64) (newExamples [][]float64) {
    for _, ex := range examples {
        if ex[label] == value {
            remaining := ex[:label]
            remaining = append(remaining, ex[label+1:]...)
            newExamples = append(newExamples, remaining)    
        }
    }
    return
}

//func (dt *DT) BestFeature() (featidx int) {
//    max := -1.0
//    total := len(dt.Examples)
//    var info float64
//    for _, label := range dt.Labels {
//        
//    }
//    return
//}
