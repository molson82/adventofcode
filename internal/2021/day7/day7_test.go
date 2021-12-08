package day7_test

import (
	"github.com/molson82/adventofcode/internal/2021/day7"
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile = "../../../inputs/day7_input.txt"
)

var _ = Describe("Day7", func() {

	Context("Part 1 Sample", func() {
		It("should equal 37", func() {
			Expect(day7.Part1(utils.ReadInput("sample_input.txt"))).To(Equal(37))
		})
	})

	Describe("Part 1", func() {
		Context("Answer", func() {
			It("should equal 331067", func() {
				Expect(day7.Part1(utils.ReadInput(inputFile))).To(Equal(331067))
			})
		})
	})

})
