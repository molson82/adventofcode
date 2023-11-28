package day1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDay1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day1 Suite")
}
