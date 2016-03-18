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
			Ω(g.Next()).ShouldNot(BeNil())
		})
	})
})
