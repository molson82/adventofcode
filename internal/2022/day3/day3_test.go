package day3

import (
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile  = "./input.txt"
	sampleFile = "./sample.txt"
)

var _ = Describe("Day1", func() {

	Context("Part 1 Sample", func() {
		It("should equal 0", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(0))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 0", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal(0))
		})
	})
})
