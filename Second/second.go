package main

import (
	"os"
	"time"
)

func main() {

	tomlFrontMatter()
	markDownBlog()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func tomlFrontMatter() { // Create the front matter for post
	f, error := os.Create("test.md")
	check(error)
	defer f.Close()

	tm := time.Now()
	postDate := tm.Format(time.ANSIC)

	f.WriteString("+++\n")
	f.WriteString("title = 'New Blog Post'\n")
	f.WriteString("date = \"" + postDate + "\"\n")
	f.WriteString("draft = false\n")
	f.WriteString("tags = ['Birthday','Grief','Sorrow']\n")
	f.WriteString("categories = [\"Grief\"]\n")
	f.WriteString("author = \"Alan Greenwell\"\n")
	f.WriteString("[[images]]\n")
	f.WriteString("\tsrc = \"img/h-photo.jpg\"\n")
	f.WriteString("\talt = \"Photo\"\n")
	f.WriteString("\tstretch = \"\"\n")
	f.WriteString("+++\n")
	f.Sync()
	f.Close()

}

func markDownBlog() {
	f, error := os.OpenFile("test.md", os.O_APPEND|os.O_WRONLY, 0600)
	check(error)
	defer f.Close()
	f.WriteString("\n\n\n")
	f.WriteString("BLOG Story\n")
	f.Sync()
	f.Close()
}
