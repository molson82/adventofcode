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
		It("should equal 157", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(157))
		})
	})

	Context("Part 1 Input", func() {
		It("should equal 8018", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal(8018))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 70", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(70))
		})
	})

	Context("Part 2 Input", func() {
		It("should equal 2518", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Equal(2518))
		})
	})
})
