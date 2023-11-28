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
		It("should equal 4512", func() {
			Expect(Part1(utils.ReadInput("sample_input.txt"))).To(Equal(4512))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 1924", func() {
			Expect(Part2(utils.ReadInput("sample_input.txt"))).To(Equal(1924))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 64084", func() {
				Expect(Part1(utils.ReadInput(inputFile))).To(Equal(64084))
			})
		})
	})

	Describe("Part 2", func() {
		Context("Answer", func() {
			It("should equal ...", func() {
				// too high
				Expect(Part2(utils.ReadInput(inputFile))).To(Not(Equal(15960)))
				// too high
				Expect(Part2(utils.ReadInput(inputFile))).To(Not(Equal(22448)))
				Expect(Part2(utils.ReadInput(inputFile))).To(Not(Equal(6912)))
				Expect(Part2(utils.ReadInput(inputFile))).To(Not(Equal(13039)))
			})
		})
	})
})
