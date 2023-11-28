package day1_test

import (
	day1 "github.com/molson82/adventofcode/2022/1"
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
		It("should equal 24000", func() {
			input := utils.ReadInput(sampleFile)
			Expect(day1.Part1(input)).To(Equal(24000))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 45000", func() {
			input := utils.ReadInput(sampleFile)
			Expect(day1.Part2(input)).To(Equal(45000))
		})
	})

	Context("Part 1 Real Input", func() {
		It("should equal 69528", func() {
			input := utils.ReadInput(inputFile)
			Expect(day1.Part1(input)).To(Equal(69528))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 206152", func() {
			input := utils.ReadInput(inputFile)
			Expect(day1.Part2(input)).To(Equal(206152))
		})
	})
})
