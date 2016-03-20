package traffic_test

import (
	"sync/atomic"

	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	It("should be able to add 2 patterns with 50/50% prob", func() {
		g := NewGenerator(1000)
		Ω(g).ShouldNot(BeNil())

		var first, second uint32
		g.AddPattern(NewPattern(50, func() {
			atomic.AddUint32(&first, 1)
		}))
		g.AddPattern(NewPattern(50, func() {
			atomic.AddUint32(&second, 1)
		}))
		g.Execute()

		var lower uint32 = 450
		var upper uint32 = 550
		Ω(first).Should(BeNumerically(">=", lower))
		Ω(first).Should(BeNumerically("<=", upper))
		Ω(second).Should(BeNumerically(">=", lower))
		Ω(second).Should(BeNumerically("<=", upper))
	})

	It("should be able to add 4 patterns with 25% prob each", func() {
		// Prep
		g := NewGenerator(10000)
		Ω(g).ShouldNot(BeNil())

		var first, second, third, fourth uint32
		g.AddPattern(NewPattern(25, func() {
			atomic.AddUint32(&first, 1)
		}))
		g.AddPattern(NewPattern(25, func() {
			atomic.AddUint32(&second, 1)
		}))
		g.AddPattern(NewPattern(25, func() {
			atomic.AddUint32(&third, 1)
		}))
		g.AddPattern(NewPattern(25, func() {
			atomic.AddUint32(&fourth, 1)
		}))

		// Run the generator
		g.Execute()

		// Assert
		var lower uint32 = 2000
		var upper uint32 = 3000
		Ω(first).Should(BeNumerically(">=", lower))
		Ω(first).Should(BeNumerically("<=", upper))
		Ω(second).Should(BeNumerically(">=", lower))
		Ω(second).Should(BeNumerically("<=", upper))
		Ω(third).Should(BeNumerically(">=", lower))
		Ω(third).Should(BeNumerically("<=", upper))
		Ω(fourth).Should(BeNumerically(">=", lower))
		Ω(fourth).Should(BeNumerically("<=", upper))
	})

	It("should be able to add 2 patterns with 25/75% prob", func() {
		g := NewGenerator(1000)
		Ω(g).ShouldNot(BeNil())

		var first, second uint32
		g.AddPattern(NewPattern(25, func() {
			atomic.AddUint32(&first, 1)
		}))
		g.AddPattern(NewPattern(75, func() {
			atomic.AddUint32(&second, 1)
		}))

		g.Execute()

		Ω(first).Should(BeNumerically(">=", uint32(200)))
		Ω(first).Should(BeNumerically("<=", uint32(300)))
		Ω(second).Should(BeNumerically(">=", uint32(700)))
		Ω(second).Should(BeNumerically("<=", uint32(800)))
	})
})
