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
	Default  float64
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

func (dt *DT) Majority(examples [][]float64) (result float64, majorityCount float64) {
    labelCounts := make(map[float64]float64)
	total := len(examples)
	if total == 0 {
        result = dt.Default
		return
	}
	length := 0
	if len(examples) > 0 {
		length = len(examples[0])
	}
	if length <= 0 {
		result = dt.Default
		return
	}
	for _, ex := range examples {
        label := ex[length-1]
        count, exists := labelCounts[label]
        if !exists {
            labelCounts[label] = 0
        }
        labelCounts[label] = count + 1
	}
    for label, count := range labelCounts {
        if count > majorityCount {
            majorityCount = count
            result = label
        }
    }
	return
}

func split(examples [][]float64, label int, value float64) (newExamples [][]float64) {
    for _, ex := range examples {
        if ex[label] == value {
            remaining := make([]float64, len(ex[:label]))
            copy(remaining, ex[:label])
            remaining = append(remaining, ex[label+1:]...)
            newExamples = append(newExamples, remaining)    
        }
    }
    return
}

func bestFeature(examples [][]float64) (best int) {
    var info float64 = 0.0
    tot := float64(len(examples))
    best = -1
    if tot == 0 {
        return
    }
    total := len(examples[0])
    baseEnt := entropy(examples)
    for i := 0; i < total; i++ {
        features := []float64{}
        for _, ex := range examples {
            features = append(features, ex[i])
        }
        var newEnt float64 = 0.0
        for _, val := range features {
            data := split(examples, i, val)
            prob := float64(len(data)) / tot
            newEnt += prob * entropy(data)
        }
        infoGain := baseEnt - newEnt
        if infoGain > info {
            info = infoGain
            best = i
        }
    }
    return
}
