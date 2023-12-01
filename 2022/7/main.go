package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	Commands []string
	Results  [][]string
}

func parse(fileName string) (*Data, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	d := &Data{}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0] == '$' {
			d.Commands = append(d.Commands, strings.TrimPrefix(line, "$ "))
			d.Results = append(d.Results, nil)
		} else {
			d.Results[len(d.Results)-1] = append(d.Results[len(d.Results)-1], line)
		}
	}
	return d, nil
}

type Node struct {
	Name  string
	IsDir bool
	Size  int

	Parent   *Node
	Children []*Node
}

func run(_ context.Context) error {
	d, err := parse("input")
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	// Build tree.
	root := &Node{Name: "/", IsDir: true}
	curNode := root
	for i, cmdLine := range d.Commands {
		parts := strings.Split(cmdLine, " ")
		cmd, args := parts[0], parts[1:]
		switch cmd {
		case "cd":
			if len(args) == 0 {
				args = append(args, "/")
			}
			switch args[0] {
			case "/":
				curNode = root
			case "..":
				if curNode.Parent != nil {
					curNode = curNode.Parent
				}
			default:
				for _, child := range curNode.Children {
					if child.Name == args[0] {
						curNode = child
						break
					}
				}
			}
		case "ls":
			for _, line := range d.Results[i] {
				parts := strings.Split(line, " ")
				if len(parts) != 2 {
					return fmt.Errorf("invalid ls output %q", line)
				}
				if parts[0] == "dir" {
					child := &Node{Parent: curNode, IsDir: true, Name: parts[1]}
					curNode.Children = append(curNode.Children, child)
				} else {
					n, err := strconv.Atoi(parts[0])
					if err != nil {
						return fmt.Errorf("invalid size %q: %w", parts[1], err)
					}
					child := &Node{Parent: curNode, IsDir: false, Name: parts[0], Size: n}
					curNode.Children = append(curNode.Children, child)
				}
			}
		}
	}

	// Process dir sizes.
	processDirSizes(root)

	var flatDirs []*Node
	flattenDirs(root, &flatDirs)

	// Part 1:
	// n := 0
	// for _, elem := range flatDirs {
	// 	if elem.Size < 100000 {
	// 		n += elem.Size
	// 	}
	// }
	// fmt.Printf("%d\n", n)

	// Part 2:
	var candidates []int
	for _, elem := range flatDirs {
		if elem.Size > 30000000-(70000000-root.Size) {
			candidates = append(candidates, elem.Size)
		}
	}
	sort.Sort(sort.IntSlice(candidates))
	fmt.Printf("%d\n", candidates[0])
	return nil
}

func processDirSizes(node *Node) int {
	if !node.IsDir {
		return node.Size
	}
	for _, child := range node.Children {
		node.Size += processDirSizes(child)
	}
	return node.Size
}

func flattenDirs(node *Node, out *[]*Node) {
	if !node.IsDir {
		return
	}
	*out = append(*out, node)
	for _, child := range node.Children {
		flattenDirs(child, out)
	}
}

func main() {
	if err := run(context.Background()); err != nil {
		println("Fail:", err.Error())
		return
	}
	println("success")
}
