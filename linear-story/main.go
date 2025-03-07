package main

import (
	"fmt"
)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) addToEnd(text string) *storyPage {
	for page.nextPage != nil {
		page = page.nextPage
	}

	page.nextPage = &storyPage{
		text:     text,
		nextPage: nil,
	}

	return page.nextPage
}

func (page *storyPage) insertAfter(text string) *storyPage {
	newPage := &storyPage{
		text:     text,
		nextPage: page.nextPage,
	}
	page.nextPage = newPage

	return newPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func main() {

	page1 := storyPage{
		text:     "It's a dark and stormy night.",
		nextPage: nil,
	}
	page2 := page1.addToEnd("You are alone outdoors, and you need to find some shelter.")
	page3 := page1.addToEnd("You see a light ahead.")

	page2Insert := page2.insertAfter("page2InsertAfter")
	page2Insert.insertAfter("page2InsertAfterInsertAfter lol")
	page3.insertAfter("page3InsertAfter")
	page1.playStory()
}
