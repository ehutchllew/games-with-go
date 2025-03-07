package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	fmt.Println(page.text)
	if page.nextPage == nil {
		return
	}
	playStory(page.nextPage)
}

func main() {

	page1 := storyPage{
		text:     "It's a dark and stormy night.",
		nextPage: nil,
	}
	page2 := storyPage{
		text:     "You are alone outdoors, and you need to find some shelter.",
		nextPage: nil,
	}
	page3 := storyPage{
		text:     "You see a light ahead.",
		nextPage: nil,
	}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)
}
