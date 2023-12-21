package main

import (
	"fmt"
)

const (
	bannerHeight = 7
	bannerWidth  = 52
)

type template [][]string

func NewTemplate() template {
	result := make(template, bannerHeight)
	for i := 0; i < 7; i++ {
		result[i] = make([]string, 0, bannerWidth)
	}
	return result

}

func (t template) SetDelimeter() template {
	for i := 0; i < bannerHeight; i++ {
		t[i] = append(t[i], ".")
	}
	return t
}

func (t template) JoinObject(j template) template {
	for i := 0; i < bannerHeight; i++ {
		t[i] = append(t[i], j[i]...)
	}
	return t
}

func (t template) Println() {
	for i := 0; i < bannerHeight; i++ {
		for _, v := range t[i] {
			if v == "." {
				v = " "
			}
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
}
