package main

import (
	"fmt"
	"goTutorial4/employee"
)

type author struct {
	firstName string
	lastName string
	bio string
}

func (a author) getFullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title string
	content string
	author
}

func (b blogPost) getDetails() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.author)
	fmt.Println("Bio: ", b.bio)
}

type website struct {
	blogPosts []blogPost
}

func (w website) getWebsiteContent() {
	fmt.Println("Contents of the website")
	for _, v := range w.blogPosts {
		v.getDetails()
		fmt.Println()
	}
}

func main() {
	e := employee.New("Rachel", "Green", 30, 20)
	e.LeavesRemaining()

	fmt.Println()

	//blogpost
	author1 := author{
		firstName: "Anne",
		lastName:  "Marie",
		bio:       "Travelling is life",
	}

	author2 := author{
		firstName: "Peter",
		lastName:  "Brook",
		bio:       "Music make me happy",
	}

	blogPost1 := blogPost{
		title:   "Travelling",
		content: "The best way to live your life is by travelling",
		author:  author1,
	}

	blogPost2 := blogPost{
		title:   "Music and Mind",
		content: "Music relaxes your mind",
		author:  author2,
	}

	blogPost3 := blogPost{
		title:   "Mysterious places on earth",
		content: "There are lot of hidden and mysterious places",
		author:  author1,
	}

	website := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}

	website.getWebsiteContent()

}




