package day1_test

import (
	"github.com/molson82/adventofcode/internal/2021/day1"
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day1_input.txt"
)

var _ = Describe("Day1", func() {

	sampleInput := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}

	Context("Part 1 Sample", func() {
		It("should equal 7", func() {
			Expect(day1.Part1(sampleInput)).To(Equal(7))
		})
	})

	Context("Part 1 input", func() {
		It("should equal 2,000", func() {
			Expect(len(utils.ReadInput(inputFile))).To(Equal(2000))
		})
	})

	Context("Part 1 ans", func() {
		It("should equal 1616", func() {
			Expect(day1.Part1(utils.ReadInput(inputFile))).To(Equal(1616))
		})
	})
})
