package day10

import (
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile  = "./input.txt"
	sampleFile = "./sample.txt"
)

var _ = Describe("Day0", func() {

	Context("Part 1 Sample", func() {
		It("should equal 0", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(0))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 0", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(0))
		})
	})
})
