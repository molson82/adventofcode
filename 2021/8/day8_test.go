package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/molson82/adventofcode/pkg/utils"
)

const (
	inputFile = "../../../inputs/input.txt"
)

var _ = Describe(", func() {

	Context("Part 1 Sample", func() {
		It("should equal 26", func() {
			Expect(Part1(utils.ReadInput("sample_input.txt"))).To(Equal(26))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 440", func() {
				Expect(Part1(utils.ReadInput(inputFile))).To(Equal(440))
			})
		})
	})
})
