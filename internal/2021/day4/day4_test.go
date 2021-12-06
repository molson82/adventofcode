package day4_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/molson82/adventofcode/internal/2021/day4"
	"github.com/molson82/adventofcode/pkg/utils"
)

const (
	inputFile = "../../../inputs/day4_input.txt"
)

var _ = Describe("Day4", func() {

	Context("Part 1 Sample", func() {
		It("should equal 4512", func() {
			Expect(day4.Part1(utils.ReadInput("sample_input.txt"))).To(Equal(4512))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 1924", func() {
			Expect(day4.Part2(utils.ReadInput("sample_input.txt"))).To(Equal(1924))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 64084", func() {
				Expect(day4.Part1(utils.ReadInput(inputFile))).To(Equal(64084))
			})
		})
	})
})
