package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/zyedidia/generic/list"
)

type NamedObject interface {
	name() string
}

type File struct {
	fileName string
	size     uint64
}

func (file *File) name() string {
	return file.fileName
}

type Directory struct {
	parent  *Directory
	dirName string

	childDirectories *list.List[*Directory]
	childFiles       *list.List[*File]
}

func (dir *Directory) name() string {
	return dir.dirName
}

func (dir *Directory) totalSize() uint64 {
	var total uint64 = 0

	childDirNode := dir.childDirectories.Front
	for childDirNode != nil {
		total += childDirNode.Value.totalSize()
		childDirNode = childDirNode.Next
	}

	childFileNode := dir.childFiles.Front
	for childFileNode != nil {
		total += childFileNode.Value.size
		childFileNode = childFileNode.Next
	}

	return total
}

func createDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		parent:           parent,
		dirName:          name,
		childDirectories: list.New[*Directory](),
		childFiles:       list.New[*File](),
	}
}

func findNamedObject[T NamedObject](name string, items *list.List[T]) T {
	node := items.Front
	for node != nil {
		if node.Value.name() == name {
			return node.Value
		}
		node = node.Next
	}

	panic(fmt.Sprintf("Could not find named object: %s", name))
}

func computePart1(lines []string) uint64 {
	root := readDirectories(lines)

	total := uint64(0)
	totalSizeBelow100000(root, &total)
	return total
}

func computePart2(lines []string) uint64 {
	root := readDirectories(lines)

	diskSize := uint64(70000000)
	neededSize := uint64(30000000)

	availableSpace := diskSize - root.totalSize()
	deletionNeeded := neededSize - availableSpace

	var smallestDirectorySize uint64 = 0
	smallestDirectoryAbove(root, deletionNeeded, &smallestDirectorySize)

	return smallestDirectorySize
}

func readDirectories(lines []string) *Directory {
	/*
		$ cd /
		$ ls
		dir dscbfp
		283653 fsdfddfv
		dir mjzqq
	*/
	cdCommandRegex := regexp.MustCompile(`^\$ cd (.*)`)
	lsCommandRegex := regexp.MustCompile(`^\$ ls`)
	directoryResultRegex := regexp.MustCompile(`^dir (.*)`)
	fileResultRegex := regexp.MustCompile(`^(\d+) (.*)`)

	var root *Directory = nil
	var currentDir *Directory = nil

	for _, line := range lines {
		cdCommandMatch := cdCommandRegex.FindAllStringSubmatch(line, -1)
		if cdCommandMatch != nil {
			targetDirName := cdCommandMatch[0][1]
			if root == nil {
				root = createDirectory("/", nil)
				currentDir = root
			} else {
				if targetDirName == ".." {
					currentDir = currentDir.parent
				} else {
					currentDir = findNamedObject(targetDirName, currentDir.childDirectories)
				}
			}
			continue
		}
		if lsCommandRegex.MatchString(line) {
			continue
		}
		directoryResultMatch := directoryResultRegex.FindAllStringSubmatch(line, -1)
		if directoryResultMatch != nil {
			name := directoryResultMatch[0][1]
			newDir := createDirectory(name, currentDir)
			currentDir.childDirectories.PushBack(newDir)
		}

		fileResultMatch := fileResultRegex.FindAllStringSubmatch(line, -1)
		if fileResultMatch != nil {
			size, _ := strconv.ParseUint(fileResultMatch[0][1], 10, 0)
			name := fileResultMatch[0][2]
			newFile := File{
				fileName: name,
				size:     size,
			}
			currentDir.childFiles.PushBack(&newFile)
		}
	}

	return root
}

func totalSizeBelow100000(directory *Directory, total *uint64) {
	childDirNode := directory.childDirectories.Front
	for childDirNode != nil {
		size := childDirNode.Value.totalSize()
		if size < 100000 {
			*total += size
		}
		totalSizeBelow100000(childDirNode.Value, total)
		childDirNode = childDirNode.Next
	}
}

func smallestDirectoryAbove(directory *Directory, minimumSize uint64, smallestSize *uint64) {
	childDirNode := directory.childDirectories.Front
	for childDirNode != nil {
		dir := childDirNode.Value
		size := dir.totalSize()
		if size > minimumSize {
			if *smallestSize == 0 || size < *smallestSize {
				*smallestSize = size
			}
		}
		smallestDirectoryAbove(dir, minimumSize, smallestSize)
		childDirNode = childDirNode.Next
	}
}

func dump(directory *Directory, depth int) {
	fmt.Printf("%*s- %s (dir) (total: %d)\n", depth*2, "", directory.name(), directory.totalSize())

	childDirNode := directory.childDirectories.Front
	for childDirNode != nil {
		dump(childDirNode.Value, depth+1)
		childDirNode = childDirNode.Next
	}

	childFileNode := directory.childFiles.Front
	for childFileNode != nil {
		fmt.Printf("%*s- %s (file, size=%d)\n", depth*2+2, "", childFileNode.Value.fileName, childFileNode.Value.size)
		childFileNode = childFileNode.Next
	}
}
