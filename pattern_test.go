package traffic_test

import (
	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pattern", func() {
	Context("NewPattern", func() {
		It("should create a new pattern successfully", func() {
			p := NewPattern(50, func() (interface{}, error) {
				return struct{}{}, nil
			})
			Ω(p).ShouldNot(BeNil())
			Ω(p.Fn()).ShouldNot(BeNil())
		})
	})
})
