package day1_test

import (
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day1_input.txt"
)

var _ = Describe("Day1", func() {

	sampleInput := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}

	Context("Part 1 Sample", func() {
		It("should equal 7", func() {
			Expect(Part1(sampleInput)).To(Equal(7))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 5", func() {
			Expect(Part2(sampleInput)).To(Equal(5))
		})
	})

	Describe("Part 1 solution", func() {
		Context("Input file", func() {
			It("should equal len 2,000", func() {
				input := utils.ReadInput(inputFile)
				Expect(len(input)).To(Equal(2000))
				Expect(input[len(input)-1]).To(Equal("8895"))
			})
		})
		Context("Answer", func() {
			It("should equal 1616", func() {
				Expect(Part1(utils.ReadInput(inputFile))).To(Equal(1616))
			})
		})
	})

	Describe("Part 2 Solution", func() {
		Context("Answer", func() {
			It("should equal 1645", func() {
				Expect(Part2(utils.ReadInput(inputFile))).To(Equal(1645))
			})
		})
	})
})
