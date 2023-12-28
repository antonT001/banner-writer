package templates

import (
	"fmt"
)

const (
	bannerHeight = 7
	bannerWidth  = 52
)

type Template [][]string

func New() Template {
	result := make(Template, bannerHeight)
	for i := 0; i < 7; i++ {
		result[i] = make([]string, 0, bannerWidth)
	}
	return result

}

func (t Template) SetDelimeter() Template {
	for i := 0; i < bannerHeight; i++ {
		t[i] = append(t[i], ".")
	}
	return t
}

func (t Template) JoinObject(j Template) Template {
	for i := 0; i < bannerHeight; i++ {
		t[i] = append(t[i], j[i]...)
	}
	return t
}

func (t Template) Println() {
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
