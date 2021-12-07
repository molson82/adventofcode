package day6_test

import (
	"github.com/molson82/adventofcode/internal/2021/day6"
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day6_input.txt"
)

var _ = Describe("Day6", func() {

	Context("Part 1 Sample", func() {
		It("should equal 5934", func() {
			Expect(day6.Solution(utils.ReadInput("sample_input.txt"), 18)).To(Equal(26))
			Expect(day6.Solution(utils.ReadInput("sample_input.txt"), 80)).To(Equal(5934))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 26984457539", func() {
			Expect(day6.Solution(utils.ReadInput("sample_input.txt"), 256)).To(Equal(26984457539))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 374927", func() {
				Expect(day6.Solution(utils.ReadInput(inputFile), 80)).To(Equal(374927))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal 1687617803407", func() {
				Expect(day6.Solution(utils.ReadInput(inputFile), 256)).To(Equal(1687617803407))
			})
		})
	})
})
