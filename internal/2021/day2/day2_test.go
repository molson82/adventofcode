package day2_test

import (
	"github.com/molson82/adventofcode/internal/2021/day2"
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day2_input.txt"
)

var _ = Describe("Day2", func() {

	sampleInput := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	Context("Part 1 Sample", func() {
		It("should equal 150", func() {
			Expect(day2.Part1(sampleInput)).To(Equal(150))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 900", func() {
			Expect(day2.Part2(sampleInput)).To(Equal(900))
		})
	})

	Describe("Part 1 Solution", func() {
		Context("Answer", func() {
			It("should equal ", func() {
				Expect(day2.Part1(utils.ReadInput(inputFile))).To(Equal(2039912))
			})
		})
	})

	Describe("Part 2 Solution", func() {
		Context("Answer", func() {
			It("should equal ", func() {
				Expect(day2.Part2(utils.ReadInput(inputFile))).To(Equal(1942068080))
			})
		})
	})
})
