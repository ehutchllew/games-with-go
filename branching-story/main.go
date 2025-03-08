package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	noPath  *storyNode
	yesPath *storyNode
}

func (n *storyNode) play() {
	fmt.Println(n.text)

	if n.yesPath == nil && n.noPath == nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		ans := scanner.Text()

		if ans == "yes" {
			n.yesPath.play()
			break
		}

		if ans == "no" {
			n.noPath.play()
			break
		}

		fmt.Println("That answer was not an option. Please answer 'yes' or 'no'.")

	}
}

func (n *storyNode) printStory(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Print(n.text)
	fmt.Println()
	if n.yesPath != nil {
		n.yesPath.printStory(depth + 1)
	}

	if n.noPath != nil {
		n.noPath.printStory(depth + 1)
	}
}

func main() {
	root := storyNode{"You are inside of a cabin in the middle of the woods. It is night. Do you want to go outside?", nil, nil}
	losing := storyNode{"You have lost...", nil, nil}
	winning := storyNode{"You have won!", nil, nil}
	root.yesPath = &losing
	root.noPath = &winning

	root.printStory(0)
}
