package traffic_test

import (
	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	Context("NewGenerator", func() {
		It("should generate a new generator correctly", func() {
			g := NewGenerator()
			Ω(g).ShouldNot(BeNil())
			i, err := g.Next()
			Ω(i).Should(BeNil())
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("AddPattern", func() {
		It("should be able to add patterns correctly", func() {
			g := NewGenerator()
			Ω(g).ShouldNot(BeNil())
			g.AddPattern(NewPattern(25, func() (interface{}, error) {
				return []int{25}, nil
			}))

			val, err := g.Next()
			Ω(err).Should(BeNil())
			Ω(val).ShouldNot(BeNil())
		})
	})

	Context("Next", func() {
		It("should be able to add 2 patterns with 50 per cent prob", func() {
			g := NewGenerator()
			Ω(g).ShouldNot(BeNil())
			g.AddPattern(NewPattern(50, func() (interface{}, error) {
				return []int{1}, nil
			}))
			g.AddPattern(NewPattern(50, func() (interface{}, error) {
				return []int{2}, nil
			}))

			first := 0
			second := 0
			for i := 0; i < 1000; i++ {
				res, err := g.Next()
				Ω(err).Should(BeNil())
				Ω(res).ShouldNot(BeNil())

				r := res.([]int)
				Ω(len(r)).Should(Equal(1))
				if r[0] == 1 {
					first++
				} else {
					second++
				}
			}

			Ω(first).Should(BeNumerically(">=", 450))
			Ω(first).Should(BeNumerically("<=", 550))
			Ω(second).Should(BeNumerically(">=", 450))
			Ω(second).Should(BeNumerically("<=", 550))
		})

		It("should be able to add 4 patterns with 25 per cent prob", func() {
			g := NewGenerator()
			Ω(g).ShouldNot(BeNil())
			g.AddPattern(NewPattern(25, func() (interface{}, error) {
				return []int{1}, nil
			}))
			g.AddPattern(NewPattern(25, func() (interface{}, error) {
				return []int{2}, nil
			}))
			g.AddPattern(NewPattern(25, func() (interface{}, error) {
				return []int{3}, nil
			}))
			g.AddPattern(NewPattern(25, func() (interface{}, error) {
				return []int{4}, nil
			}))

			responses := map[int]int{1: 0, 2: 0, 3: 0, 4: 0}
			for i := 0; i < 10000; i++ {
				res, err := g.Next()
				Ω(err).Should(BeNil())
				Ω(res).ShouldNot(BeNil())

				r := res.([]int)
				Ω(len(r)).Should(Equal(1))
				responses[r[0]] = responses[r[0]] + 1
			}

			for i := 1; i < 5; i++ {
				Ω(responses[i]).Should(BeNumerically(">=", 2000))
				Ω(responses[i]).Should(BeNumerically("<=", 3000))
			}
		})
	})
})
