package main

import "fmt"

type Course struct {
	name   string
	class  int
	skills []string
}

func (c Course) subscribe(name string) {
	fmt.Printf("The user %s, is subscribed in %s\n", name, c.name)
}

func main() {
	html := new(Course)
	html.name = "Go lang"
	html.class = 40
	html.skills = []string{"Frontend"}
	html.subscribe("Alberto")
}
