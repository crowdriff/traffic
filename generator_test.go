package traffic_test

import (
	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	Context("NewGenerator", func() {
		It("should generate a new generator correctly", func() {
			g := NewGenerator(1)
			Ω(g).ShouldNot(BeNil())
			i, err := g.Next()
			Ω(i).Should(BeNil())
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("AddPattern", func() {
		It("should be able to add patterns correctly", func() {
			g := NewGenerator(1)
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
		It("should be callable at most `interations` # of times (simple)",
			func() {
				g := NewGenerator(1)
				Ω(g).ShouldNot(BeNil())
				g.AddPattern(NewPattern(100, func() (interface{}, error) {
					return struct{}{}, nil
				}))

				_, e := g.Next()
				Ω(e).Should(BeNil())
				_, e = g.Next()
				Ω(e).ShouldNot(BeNil())
			})

		It("should be callable at most `interations` # of times (complex)",
			func() {
				g := NewGenerator(100)
				Ω(g).ShouldNot(BeNil())
				g.AddPattern(NewPattern(100, func() (interface{}, error) {
					return struct{}{}, nil
				}))
				var e error
				for i := 0; i < 100; i++ {
					_, e = g.Next()
					Ω(e).Should(BeNil())
				}
				_, e = g.Next()
				Ω(e).ShouldNot(BeNil())
			})
	})
})
