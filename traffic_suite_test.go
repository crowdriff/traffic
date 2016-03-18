package traffic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTraffic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Traffic Suite")
}
