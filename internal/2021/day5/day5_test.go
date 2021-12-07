package day5_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/molson82/adventofcode/internal/2021/day5"
	"github.com/molson82/adventofcode/pkg/utils"
)

const (
	inputFile = "../../../inputs/day5_input.txt"
)

var _ = Describe("Day5", func() {

	Context("Part 1 Sample", func() {
		It("should equal 4512", func() {
			Expect(day5.Part1(utils.ReadInput("sample_input.txt"))).To(Equal(5))
		})
	})

	Context("Part 2 sample", func() {
		It("should equal 12", func() {
			Expect(day5.Part2(utils.ReadInput("sample_input.txt"))).To(Equal(12))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 6311", func() {
				Expect(day5.Part1(utils.ReadInput(inputFile))).To(Equal(6311))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal 19929", func() {
				Expect(day5.Part2(utils.ReadInput(inputFile))).To(Equal(19929))
			})
		})
	})
})
