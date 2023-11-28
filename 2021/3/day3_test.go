package test

import (
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/input.txt"
)

var _ = Describe(", func() {

	sampleInput := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	Context("Part 1 Sample", func() {
		It("should equal 150", func() {
			Expect(Part1(sampleInput)).To(Equal(int64(198)))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 230", func() {
			Expect(Part2(sampleInput)).To(Equal(int64(230)))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 3277364", func() {
				Expect(Part1(utils.ReadInput(inputFile))).To(Equal(int64(3277364)))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal 5736383", func() {
				Expect(Part2(utils.ReadInput(inputFile))).To(Equal(int64(5736383)))
			})
		})
	})
})
