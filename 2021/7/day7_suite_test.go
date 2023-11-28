package day7_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDay7(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day7 Suite")
}
