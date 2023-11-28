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

	Context("Part 1 Sample", func() {
		It("should equal 5934", func() {
			Expect(Solution(utils.ReadInput("sample_input.txt"), 18)).To(Equal(26))
			Expect(Solution(utils.ReadInput("sample_input.txt"), 80)).To(Equal(5934))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 26984457539", func() {
			Expect(Solution(utils.ReadInput("sample_input.txt"), 256)).To(Equal(26984457539))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 374927", func() {
				Expect(Solution(utils.ReadInput(inputFile), 80)).To(Equal(374927))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal 1687617803407", func() {
				Expect(Solution(utils.ReadInput(inputFile), 256)).To(Equal(1687617803407))
			})
		})
	})
})
