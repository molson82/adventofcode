package day3_test

import (
	"github.com/molson82/adventofcode/internal/2021/day3"
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day3_input.txt"
)

var _ = Describe("Day3", func() {

	sampleInput := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	Context("Part 1 Sample", func() {
		It("should equal 150", func() {
			Expect(day3.Part1(sampleInput)).To(Equal(int64(198)))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 230", func() {
			Expect(day3.Part2(sampleInput)).To(Equal(int64(230)))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 3277364", func() {
				Expect(day3.Part1(utils.ReadInput(inputFile))).To(Equal(int64(3277364)))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal 5736383", func() {
				Expect(day3.Part2(utils.ReadInput(inputFile))).To(Equal(int64(5736383)))
			})
		})
	})
})
