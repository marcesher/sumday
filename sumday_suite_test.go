package sumday_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSumday(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumDay Suite")
}
