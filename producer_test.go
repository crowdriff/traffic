package traffic_test

import (
	. "github.com/crowdriff/traffic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Producer", func() {
	Context("DefaultProducer", func() {
		It("should create a new default producer correctly", func() {
			d := NewDefaultProducer()
			Ω(d).ShouldNot(BeNil())
			i := d.Generate()
			Ω(i).ShouldNot(BeNil())
		})
	})
})
