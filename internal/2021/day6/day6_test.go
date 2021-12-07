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
			Expect(day6.Part1(utils.ReadInput("sample_input.txt"))).To(Equal(5934))
		})
	})

})
