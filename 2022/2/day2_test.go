package day2

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
		It("should equal 15", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(15))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 10359", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(10359)))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 12", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(12))
		})
	})

	Context("Part 2 Input", func() {
		It("should equal 11258", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Equal(11258))
		})
	})
})
