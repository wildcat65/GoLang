package main

import (
	"fmt"
	"time"
)

func main() {

	tomlFrontMatter()

}

func tomlFrontMatter() { // Create the front matter for post

	tm := time.Now()
	postDate := tm.Format(time.ANSIC)
	fmt.Println("+++")
	fmt.Println("title = 'New Blog Post'")
	fmt.Println("date = \"", postDate, "\"")
	fmt.Println("draft = false")
	fmt.Println("tags = ['Birthday','Grief','Sorrow']")
	fmt.Println("categories = [\"Grief\"]")
	fmt.Println("author = \"Alan Greenwell\"")
	fmt.Println("[[images]]")
	fmt.Println("\tsrc = \"img/h-photo.jpg\"")
	fmt.Println("\talt = \"Photo\"")
	fmt.Println("\tstretch = \"\"")
	fmt.Println("+++")

}
