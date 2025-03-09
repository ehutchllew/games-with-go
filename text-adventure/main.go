package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd        string
	desc       string
	nextChoice *choices
	nextNode   *storyNode
}

type storyNode struct {
	choices *choices
	text    string
}

var scanner *bufio.Scanner

func (n *storyNode) addChoice(cmd, desc string, nextNode *storyNode) {
	choice := &choices{cmd, desc, nil, nextNode}

	if n.choices == nil {
		n.choices = choice
	} else {
		curChoice := n.choices
		for curChoice.nextChoice != nil {
			curChoice = curChoice.nextChoice
		}

		curChoice.nextChoice = choice
	}
}

func (n *storyNode) exeCmd(cmd string) *storyNode {
	curChoice := n.choices
	for curChoice != nil {
		if strings.ToLower(curChoice.cmd) == strings.ToLower(cmd) {
			return curChoice.nextNode
		}
		curChoice = curChoice.nextChoice
	}
	fmt.Println("Sorry I didn't understand that.")
	return n
}

func (n *storyNode) play() {
	n.render()
	if n.choices != nil {
		scanner.Scan()
		n.exeCmd(scanner.Text()).play()
	}
}

func (n *storyNode) render() {
	fmt.Println(n.text)
	curChoice := n.choices
	for curChoice != nil {
		fmt.Println(curChoice.cmd, ":", curChoice.desc)
		curChoice = curChoice.nextChoice
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := &storyNode{
		text: `
		You are in a large chamber, deep underground.
		You see three passages leading out. A north passage leads into darkness.
		To the south, a passage appears to head upward. The eastern passage appears
		flat and well traveled.
		`,
	}

	darkRoom := &storyNode{
		text: "It is pitch black. You cannot see a thing.",
	}

	darkRoomLit := &storyNode{
		text: "The dark passage is now lit by your lantern. You can continue north or head back south.",
	}

	grue := &storyNode{
		text: "While stumbling around in the darkness, you are eaten by a grue.",
	}

	trap := &storyNode{
		text: "You head down the well traveled path when suddently a trap door opens and you fall into a pit.",
	}

	treasure := &storyNode{
		text: "You enter a small chamber filled with treasure!",
	}

	start.addChoice("N", "Go North.", darkRoom)
	start.addChoice("S", "Go South.", darkRoom)
	start.addChoice("E", "Go East.", trap)

	darkRoom.addChoice("S", "Try to go back south.", grue)
	darkRoom.addChoice("O", "Turn on lantern.", darkRoomLit)

	darkRoomLit.addChoice("N", "Go North.", treasure)
	darkRoomLit.addChoice("S", "Go South.", start)

	start.play()

	fmt.Println()
	fmt.Println("The end.")
}
