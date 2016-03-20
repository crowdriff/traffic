package traffic_test

import (
	"sync/atomic"

	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	Context("NewGenerator", func() {
		It("should generate a new generator correctly", func() {
			g := NewGenerator(1)
			Ω(g).ShouldNot(BeNil())
			g.Execute()
		})
	})

	Context("AddPattern", func() {
		It("should be able to add patterns correctly", func() {
			g := NewGenerator(1)
			Ω(g).ShouldNot(BeNil())

			var first uint32
			g.AddPattern(NewPattern(25, func() {
				atomic.AddUint32(&first, 1)
			}))

			g.Execute()
		})
	})
})
