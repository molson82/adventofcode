package day8_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/molson82/adventofcode/internal/2021/day8"
	"github.com/molson82/adventofcode/pkg/utils"
)

const (
	inputFile = "../../../inputs/day8_input.txt"
)

var _ = Describe("Day8", func() {

	Context("Part 1 Sample", func() {
		It("should equal ...", func() {
			Expect(day8.Part1(utils.ReadInput("sample_input.txt"))).To(Equal(99))
		})
	})

})
