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

	sampleInput := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	Context("Part 1 Sample", func() {
		It("should equal 150", func() {
			Expect(Part1(sampleInput)).To(Equal(150))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 900", func() {
			Expect(Part2(sampleInput)).To(Equal(900))
		})
	})

	Describe("Part 1 Solution", func() {
		Context("Answer", func() {
			It("should equal ", func() {
				Expect(Part1(utils.ReadInput(inputFile))).To(Equal(2039912))
			})
		})
	})

	Describe("Part 2 Solution", func() {
		Context("Answer", func() {
			It("should equal ", func() {
				Expect(Part2(utils.ReadInput(inputFile))).To(Equal(1942068080))
			})
		})
	})
})
