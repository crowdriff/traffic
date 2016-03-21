package traffic_test

import (
	"time"

	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	It("should be able to properly run concurrently", func() {
		fn := func() {
			time.Sleep(1 * time.Millisecond)
		}

		g := NewGenerator(500, 1)
		Ω(g).ShouldNot(BeNil())
		g.AddPattern(NewPattern(100, fn))

		start := time.Now()
		g.Execute()
		diffSingle := time.Since(start)

		g2 := NewGenerator(500, 3)
		Ω(g2).ShouldNot(BeNil())
		g2.AddPattern(NewPattern(100, fn))
		start = time.Now()
		g2.Execute()
		diffMulti := time.Since(start)

		Ω(diffMulti * 2).Should(BeNumerically("<", diffSingle))
	})
})
