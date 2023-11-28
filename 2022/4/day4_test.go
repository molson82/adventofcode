package day4

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
		It("should equal 2", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(2))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 548: too high", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(548)))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 529: too high", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(529)))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 533: too high", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(533)))
		})
	})

	Context("Part 1 Input", func() {
		It("should equal 528", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal(528))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 4", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(4))
		})
	})

	Context("Part 2 Input", func() {
		It("should equal 881", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Equal(881))
		})
	})
})
