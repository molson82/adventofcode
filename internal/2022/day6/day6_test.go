package day6

import (
	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile  = "./input.txt"
	sampleFile = "./sample.txt"
)

type input struct {
	Input  string
	Result int
}

var _ = Describe("Day6", func() {

	Context("Part 1 Sample", func() {
		It("should equal list of inputs", func() {
			inputs := []input{
				{Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Result: 5},
				{Input: "nppdvjthqldpwncqszvftbrmjlhg", Result: 6},
				{Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Result: 10},
				{Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Result: 11},
			}
			for _, v := range inputs {
				Expect(Part1(v.Input)).To(Equal(v.Result))
			}
		})
	})

	Context("Part 1 Input", func() {
		It("should equal 1766", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input[0])).To(Equal(1766))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 0", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(0))
		})
	})
})
