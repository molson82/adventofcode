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
			Expect(Part1(input)).To(Equal(95437))
		})
	})

	Context("Part 1 Sample Temp", func() {
		It("should equal 95437", func() {
			ans := Part1Temp(sampleFile)
			Expect(ans).To(Equal(95437))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 912377: too low", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(912377)))
		})
	})

	Context("Part 1 Input", func() {
		It("should not equal 1118530", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Not(Equal(1118530)))
		})
	})

	Context("Part 1 Input", func() {
		It("should equal 1086293", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal(1086293))
		})
	})

	//Context("Part 2 Sample", func() {
	//It("should equal 0", func() {
	//input := utils.ReadInput(sampleFile)
	//Expect(Part2(input)).To(Equal(0))
	//})
	//})
})
