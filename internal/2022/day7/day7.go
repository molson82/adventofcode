package day7

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/molson82/adventofcode/pkg/utils"
)

type FileSystem struct {
	Root    *Directory
	Current *Directory
}

type Directory struct {
	Name        string
	Parent      *Directory
	Files       map[string]*File
	Directories map[string]*Directory
	Size        int64
}

type File struct {
	Name   string
	Size   int64
	Parent *Directory
}

func (fs *FileSystem) ChangeDirectory(name string) {
	if name == "/" {
		fs.Current = fs.Root
		return
	}
	if name == ".." {
		fs.Current = fs.Current.Parent
		return
	}
	if dir, ok := fs.Current.Directories[name]; ok {
		fs.Current = dir
		if dir.Directories == nil {
			dir.Directories = make(map[string]*Directory)
		}
		if dir.Files == nil {
			dir.Files = make(map[string]*File)
		}
	}
}

func (fs *FileSystem) List() {
	for name, file := range fs.Current.Files {
		fmt.Printf("%d %s\n", file.Size, name)
	}
	for name := range fs.Current.Directories {
		fmt.Printf("dir %s\n", name)
	}
}

func (fs *FileSystem) AddFile(name string, size int64) {
	fs.Current.Files[name] = &File{Name: name, Size: size, Parent: fs.Current}
}

func (fs *FileSystem) AddDirectory(name string) {
	fs.Current.Directories[name] = &Directory{Name: name, Parent: fs.Current}
}

func NewFileSystem() *FileSystem {
	root := &Directory{Name: "/", Parent: nil, Files: make(map[string]*File), Directories: make(map[string]*Directory)}
	return &FileSystem{Root: root, Current: root}
}

func (fs *FileSystem) BuildFileSystem(input []string) {
	for _, line := range input {
		if strings.Contains(line, "cd") && !regexp.MustCompile(`\d`).MatchString(line) {
			dir := strings.Split(line, "$ cd ")[1]
			fs.ChangeDirectory(dir)

		} else if strings.Contains(line, "dir") {
			dirName := strings.Split(line, "dir ")[1]
			fs.AddDirectory(dirName)

		} else if regexp.MustCompile(`\d`).MatchString(line) {
			size := strings.Split(line, " ")[0]
			sizeInt, err := strconv.ParseInt(size, 10, 64)
			utils.CheckErr(err)
			name := strings.Split(line, " ")[1]
			fs.AddFile(name, sizeInt)
			fs.Current.Size += sizeInt
			updateParentSize(fs.Current, sizeInt)
		}
	}
}

func updateParentSize(dir *Directory, size int64) {
	if dir.Parent == nil {
		return
	}
	dir.Parent.Size += size
	updateParentSize(dir.Parent, size)
}

func BFS(root *Directory) []int64 {
	if root == nil {
		return []int64{}
	}

	var ans []int64
	ans = append(ans, root.Size)
	for _, v := range root.Directories {
		ans = append(ans, BFS(v)...)
	}

	return ans
}

func Part1(input []string) int64 {
	fs := NewFileSystem()
	fs.BuildFileSystem(input)
	fs.ChangeDirectory("/")

	var final []int64
	for _, v := range fs.Root.Directories {
		res := BFS(v)
		final = append(final, res...)
	}

	var sum int64
	for _, v := range final {
		if v <= 100000 {
			sum += v
		}
	}

	return sum
}

func Part2(input []string) int64 {
	fs := NewFileSystem()
	fs.BuildFileSystem(input)
	fs.ChangeDirectory("/")

	var sizes []int64
	for _, v := range fs.Root.Directories {
		sizes = append(sizes, v.Size)
	}
	for _, v := range fs.Root.Files {
		sizes = append(sizes, v.Size)
	}

	var sum int64
	for _, v := range sizes {
		sum = sum + v
	}

	remove := 30000000 - (70000000 - sum)

	var allSizes []int64
	for _, v := range fs.Root.Directories {
		res := BFS(v)
		allSizes = append(allSizes, res...)
	}

	var final []int64
	for _, v := range allSizes {
		if v >= remove {
			final = append(final, v)
		}
	}
	sort.Slice(final, func(i, j int) bool {
		return final[i] < final[j]
	})

	return final[0]
}
