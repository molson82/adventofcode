package day7_test

import (
	"github.com/molson82/adventofcode/internal/2021/day7"
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day7_input.txt"
)

var _ = Describe("Day7", func() {

	Context("Part 1 Sample", func() {
		It("should equal 37", func() {
			Expect(day7.Part1(utils.ReadInput("sample_input.txt"))).To(Equal(37))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 168", func() {
			Expect(day7.Part2(utils.ReadInput("sample_input.txt"))).To(Equal(168))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 331067", func() {
				Expect(day7.Part1(utils.ReadInput(inputFile))).To(Equal(331067))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal 92881128", func() {
				Expect(day7.Part2(utils.ReadInput(inputFile))).To(Equal(92881128))
			})
		})
	})

})
