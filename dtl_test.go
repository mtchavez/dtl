package dtl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dtl", func() {

    Describe("split", func(){
    
        It("returns no examples if none left", func(){
            examples := [][]float64{}
            Expect(split(examples, 1, 1.0)).To(HaveLen(0))
        })

        It("returns no examples if not found for value", func(){
            examples := [][]float64{[]float64{3.0, 4.0, 5.0, 1.0}, []float64{9.0, 8.0, 7.0, 0.0}}
            Expect(split(examples, 1, 6.0)).To(HaveLen(0))
        })

        It("returns examples split on value", func(){
            examples := [][]float64{[]float64{3.0, 4.0, 5.0, 1.0}, []float64{9.0, 8.0, 7.0, 0.0}}
            expected := [][]float64{[]float64{4.0, 5.0, 1.0}}
            Expect(split(examples, 0, 3.0)).To(Equal(expected))
        })
    })

	Describe("entropy", func() {

		var (
			examples [][]float64
		)

		BeforeEach(func() {
			examples = [][]float64{{1.0, 1.0}, {2.0, 1.0}, {3.0, 0.0}}

		})

        It("returns zero if no examples", func() {
            Expect(entropy([][]float64{})).To(Equal(0.0))
        })

        It("returns zero if no data for example", func() {
            Expect(entropy([][]float64{[]float64{}})).To(Equal(0.0))
        })

		It("calculates information from passed in examples", func() {
			expected := float64(0.9182958340544896)
			Expect(entropy(examples)).To(Equal(expected))
		})
	})

	Describe("LoadExamples", func() {

		It("loads in examples from csv", func() {
			examples, labels := LoadExamples("./examples.csv")
			Expect(labels).To(HaveLen(4))
			Expect(examples).To(HaveLen(4))
		})
	})

	Describe("Majority", func() {

		var (
			dt       *DT
			examples [][]float64
		)

		BeforeEach(func() {
			dt = &DT{Default: true}
			examples = [][]float64{{0.0, 1.0}, {0.0, 0.0}, {0.0, 0.0}}
		})

		It("returns defualt if no examples", func() {
			Expect(dt.Majority()).To(Equal(dt.Default))
		})

		It("returns clear majority", func() {
			dt.Examples = examples
			Expect(dt.Majority()).To(BeFalse())
		})

		It("returns default if results are even", func() {
			dt.Examples = append(examples, []float64{1.0, 1.0})
			Expect(dt.Majority()).To(Equal(dt.Default))
		})
	})

    Describe("bestFeature", func(){
    
        It("returns -1 if no examples", func(){
            Expect(bestFeature([][]float64{})).To(Equal(-1))   
        })

        It("returns -1 if no features", func(){
            Expect(bestFeature([][]float64{[]float64{}})).To(Equal(-1))   
        })

        It("returns index of next best feature", func(){
            examples := [][]float64{
                []float64{0.1, 0.4, 0.2, 1.0},
                []float64{0.0, 0.3, 1.1, 0.0},
                []float64{1.1, 0.4, 3.2, 1.0},
                []float64{1.1, 0.4, 3.2, 1.0},
            }
            Expect(bestFeature(examples)).To(Equal(0))
        })
    })

})
