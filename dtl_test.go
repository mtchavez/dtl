package dtl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dtl", func() {

    Describe("information", func(){
        
        It("returns zero if x and y are 1.0", func(){
            Expect(information(1.0, 1.0)).To(Equal(0.0))
        })

        It("returns close to zero if x and y are 0.0", func(){
            var expected float64 = 4.650699332842307e-06
            Expect(information(0.0, 0.0)).To(Equal(expected))
        })
        
        It("returns expected calculation", func(){
            var expected float64 = 0.9643856189774724
            Expect(information(0.5, 0.2)).To(Equal(expected))
        })
    })


    Describe("gain", func(){


        It("returns 0 if total <= 0", func(){
            examples := [][]float64{}
            Expect(gain(-1.0, examples)).To(Equal(0.0))
            Expect(gain(0.0, examples)).To(Equal(0.0))
        })

        It("returns 0 if no examples", func(){
            examples := [][]float64{}
            Expect(gain(100.0, examples)).To(Equal(0.0))
        })

        It("caclulated correctly", func(){
            total := float64(3.0)
            examples := [][]float64{ []float64{1.0, 1.0}, []float64{2.0, 1.0}, []float64{3.0, 0.0} }
            expected := float64(0.9182958340544896)
            Expect(gain(total, examples)).To(Equal(expected))
        })
    })

    Describe("entropy", func(){

        var (
            examples [][]float64
            children []Node
            node Node
        )

        BeforeEach(func(){
            examples = [][]float64{ []float64{1.0, 1.0}, []float64{2.0, 1.0}, []float64{3.0, 0.0} }
            children = []Node { Node{ name: "test2", examples: examples } }
            node = Node{ name: "test", examples: examples, children: children}
 
        })

        
        It("calculates if remaining examples", func(){
            expected := float64(0.08170416594551044) 
            Expect(entropy(float64(3.0), node)).To(Equal(expected))
        })

        It("calculates if subset of total examples", func(){
            expected := float64(0.7245112497836532)
            Expect(entropy(float64(10.0), node)).To(Equal(expected))
        })
    })

    Describe("LoadExamples", func(){
    
        It("loads in examples from csv", func(){
            examples, labels := LoadExamples("./examples.csv")
            Expect(labels).To(HaveLen(4))
            Expect(examples).To(HaveLen(4))
        })
    })

    Describe("Majority", func(){
   
        var (
            dt *DT
            examples [][]float64
        )

        BeforeEach(func(){
            dt = &DT{Default: true}
            examples = [][]float64{[]float64{0.0, 1.0}, []float64{0.0, 0.0}, []float64{0.0, 0.0}}
        })

        It("returns defualt if no examples", func(){
            Expect(dt.Majority()).To(Equal(dt.Default))
        })

        It("returns clear majority", func(){
            dt.Examples = examples
            Expect(dt.Majority()).To(BeFalse())
        })

        It("returns default if results are even", func(){
            dt.Examples = append(examples, []float64{1.0, 1.0})
            Expect(dt.Majority()).To(Equal(dt.Default))
        })
    })

})
