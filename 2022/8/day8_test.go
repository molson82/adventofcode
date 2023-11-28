package day8

import (
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	inputFile  = "./input.txt"
	sampleFile = "./sample.txt"
)

var _ = Describe("Day8", func() {

	Context("Part 1 Sample", func() {
		It("should equal 21", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part1(input)).To(Equal(21))
		})
	})

	Context("Part 1 Input", func() {
		It("should equal 1684", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part1(input)).To(Equal(1684))
		})
	})

	Context("Part 2 Sample", func() {
		It("should equal 0", func() {
			input := utils.ReadInput(sampleFile)
			Expect(Part2(input)).To(Equal(8))
		})
	})

	Context("Part 2 Input", func() {
		It("should equal 486540", func() {
			input := utils.ReadInput(inputFile)
			Expect(Part2(input)).To(Equal(486540))
		})
	})

	Describe("Individual Tree tests", func() {
		var treeGrid [][]int
		var width int
		var height int
		BeforeEach(func() {
			input := utils.ReadInput(sampleFile)
			for _, v := range input {
				row := strings.Split(v, "")
				rowInt := utils.StrList_to_IntList(row)
				treeGrid = append(treeGrid, rowInt)
			}
			width = len(treeGrid[0])
			height = len(treeGrid)
		})

		AfterEach(func() {
			treeGrid = [][]int{}
			width = 0
			height = 0
		})

		Context("Part 2 Sample", func() {
			It("row 2 col 3 should equal 4", func() {
				tree := treeGrid[1][2]
				up := countVerticleUp(tree, 1, 2, treeGrid)
				left := countHorizLeft(tree, 1, 2, treeGrid)
				right := countHorizRight(tree, 1, 2, treeGrid, width)
				down := countVerticleDown(tree, 1, 2, treeGrid, height)
				//fmt.Println("tree: ", tree)
				//fmt.Printf("up: %v | left: %v | right: %v | down: %v\n", up, left, right, down)
				ans := up * left * right * down
				Expect(ans).To(Equal(4))
			})
		})

		Context("Part 2 Sample", func() {
			It("row 4 col 3 should equal 8", func() {
				input := utils.ReadInput(sampleFile)
				var treeGrid [][]int
				for _, v := range input {
					row := strings.Split(v, "")
					rowInt := utils.StrList_to_IntList(row)
					treeGrid = append(treeGrid, rowInt)
				}
				width = len(treeGrid[0])
				height = len(treeGrid)
				tree := treeGrid[3][2]
				up := countVerticleUp(tree, 3, 2, treeGrid)
				left := countHorizLeft(tree, 3, 2, treeGrid)
				right := countHorizRight(tree, 3, 2, treeGrid, width)
				down := countVerticleDown(tree, 3, 2, treeGrid, height)
				//fmt.Println("tree: ", tree)
				//fmt.Printf("up: %v | left: %v | right: %v | down: %v\n", up, left, right, down)
				ans := up * left * right * down
				Expect(ans).To(Equal(8))
			})
		})
	})
})
