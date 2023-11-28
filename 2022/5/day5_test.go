package day5

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
		It("should equal CMZ", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal("CMZ"))
		})
	})

	Context("Part 1 Input", func() {
		It("should equal LJSVLTWQM", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal("LJSVLTWQM"))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal MCD", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal("MCD"))
		})
	})

	Context("Part 2 Input", func() {
		It("should equal BRQWDBBJM", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Equal("BRQWDBBJM"))
		})
	})
})
