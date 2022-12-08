package day7

import (
	"fmt"
	"regexp"
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
	Size        int
}

type File struct {
	Name   string
	Size   int
	Parent *Directory
}

func (d *Directory) GetSize(currentTotal int) int {
	newTotal := currentTotal
	if d.Files != nil {
		for _, v := range d.Files {
			newTotal += v.Size
		}
		if d.Directories == nil || len(d.Directories) == 0 {
			d.Size = newTotal
			return newTotal
		} else {
			for _, v := range d.Directories {
				return v.GetSize(newTotal)
			}
		}
	}
	return 0
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

func (fs *FileSystem) AddFile(name string, size int) {
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
			fs.AddFile(name, int(sizeInt))
		}
	}
}

func BFS(root *Directory) []int {
	if root == nil {
		return []int{}
	}

	var ans []int
	fmt.Println("root: ", root.Name)
	root.GetSize(0)
	ans = append(ans, root.Size)
	fmt.Println("ans: ", ans)
	for _, v := range root.Directories {
		ans = append(ans, BFS(v)...)
	}

	return ans
}

func Part1(input []string) int {
	fs := NewFileSystem()
	fs.BuildFileSystem(input)
	fs.ChangeDirectory("/")

	var final []int
	for _, v := range fs.Root.Directories {
		res := BFS(v)
		final = append(final, res...)
	}
	fmt.Println(final)

	return 0
}

func Part2(input []string) int {
	return 0
}
