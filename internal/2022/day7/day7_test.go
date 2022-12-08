package day7

import (
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile  = "./input.txt"
	sampleFile = "./sample.txt"
)

var _ = Describe("Day0", func() {

	Context("Part 1 Sample", func() {
		It("should equal 95437", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(int64(95437)))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 912377: too low", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(int64(912377))))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 1118530", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(int64(1118530))))
		})
	})

	Context("Part 1 Input", func() {
		It("should equal 1086293", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal(int64(1086293)))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 24933642", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(int64(24933642)))
		})
	})

	Context("Part 2 Input", func() {
		It("should not equal 355", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Not(Equal(int64(355))))
		})
	})

	Context("Part 2 Input", func() {
		It("should equal 366028", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Equal(int64(366028)))
		})
	})

})
