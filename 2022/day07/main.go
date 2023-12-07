package day07

import (
	"advent/util"
	"bytes"
	"errors"
	"fmt"
	"log"
)

type file struct {
	name string
	size int
}

type directory struct {
	name     string
	parent   *directory
	children map[string]*directory
	files    map[string]*file
	size     int
}

func (d *directory) getChild(name string) *directory {
	c, ok := d.children[name]
	if ok {
		return c
	} else {
		c = &directory{
			name:     name,
			parent:   d,
			children: map[string]*directory{},
			files:    map[string]*file{},
		}
		d.children[name] = c
		return c
	}
}

func (d *directory) getFile(name string, size int) *file {
	f, ok := d.files[name]
	if ok {
		return f
	} else {
		f = &file{
			name: name,
			size: size,
		}
		d.files[name] = f
		return f
	}
}

func (d *directory) updateSize() int {
	size := 0
	for _, c := range d.children {
		size += c.updateSize()
	}
	for _, f := range d.files {
		size += f.size
	}
	d.size = size
	return size
}

func (d *directory) sumSizeOfSmallDirectories() int {
	answer := 0
	if d.size <= 100_000 {
		answer += d.size
	}
	for _, c := range d.children {
		answer += c.sumSizeOfSmallDirectories()
	}
	return answer
}

func (d *directory) getSmallestDirToDelete(minSize, haveSize int) int {
	if d.size >= minSize && d.size < haveSize {
		haveSize = d.size
	}
	for _, c := range d.children {
		haveSize = c.getSmallestDirToDelete(minSize, haveSize)
	}
	return haveSize
}

func Run(input []byte) error {
	root := &directory{
		name:     "/",
		children: map[string]*directory{},
		files:    map[string]*file{},
	}
	cd := root
	lsing := false

	lines := util.ParseInputLines(input)
	for _, line := range lines {
		if line[0] == '$' {
			lsing = false
			command := line[2:4]
			if bytes.Equal(command, []byte("cd")) {
				dirName := line[5:]
				if bytes.Equal(dirName, []byte("..")) {
					cd = cd.parent
				} else if bytes.Equal(dirName, []byte("/")) {
					cd = root
				} else {
					cd = cd.getChild(string(dirName))
				}
			} else if bytes.Equal(command, []byte("ls")) {
				lsing = true
			} else {
				return errors.New(fmt.Sprintf("unrecognized command: %s", line[2:4]))
			}
		} else if lsing {
			if bytes.Equal(line[0:3], []byte("dir")) {
				dirName := string(line[4:])
				cd.getChild(dirName)
			} else {
				parts := bytes.Split(line, []byte(" "))
				size := util.Btoi(parts[0])
				fileName := string(parts[1])
				cd.getFile(fileName, size)
			}
		} else {

			return errors.New(fmt.Sprintf("unrecognized line: %s", line))
		}
	}

	root.updateSize()
	part1 := root.sumSizeOfSmallDirectories()
	totalSize := 70_000_000
	needSize := 30_000_000
	emptySize := totalSize - root.size
	minSize := needSize - emptySize
	part2 := root.getSmallestDirToDelete(minSize, root.size)

	log.Printf("The sum of the sizes of directories at most 100,000 size is %d", part1)
	log.Printf("The size of the smallest directory that could be deleted to create space is %d", part2)

	return nil
}
