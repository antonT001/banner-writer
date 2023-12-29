package template

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	BANNER_HEIGHT = 7
	BANNER_WIDTH  = 52

	EMPTY_CELL   = " "
	BORDER       = "#"
	COLORED_CELL = "X"
)

type Template [][]string

func New() Template {
	result := make(Template, BANNER_HEIGHT)
	for i := 0; i < BANNER_HEIGHT; i++ {
		result[i] = make([]string, 0, BANNER_WIDTH)
	}
	return result

}

func (tmpl Template) SetDelimeter() Template {
	for i := 0; i < BANNER_HEIGHT; i++ {
		tmpl[i] = append(tmpl[i], EMPTY_CELL)
	}
	return tmpl
}

func (tmpl Template) JoinObject(j Template) Template {
	for i := 0; i < BANNER_HEIGHT; i++ {
		tmpl[i] = append(tmpl[i], j[i]...)
	}
	return tmpl
}

func (tmpl Template) Println() {
	d := display()

	length := BANNER_WIDTH
	if length > tmpl.Width() {
		length = tmpl.Width()
	}

	for i := 0; i < BANNER_HEIGHT; i++ {
		for j := 0; j < length; j++ {
			d[i+1][j+1] = tmpl[i][j]
		}
	}
	for i := 0; i < len(d); i++ {
		for _, v := range d[i] {
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
}

func (tmpl Template) Approve() bool {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Do you want to continue [y/n]? ")
		scanner.Scan()
		if strings.ToLower(scanner.Text()) == "y" {
			return true
		}
		if strings.ToLower(scanner.Text()) == "n" {
			return false
		}
	}
}

func (tmpl Template) Height() int {
	return len(tmpl)
}

func (tmpl Template) Width() int {
	if len(tmpl) == 0 {
		return 0
	}
	return len(tmpl[0])
}

func display() [][]string {
	result := make([][]string, BANNER_HEIGHT+2)
	for i := 0; i < BANNER_HEIGHT+2; i++ {
		result[i] = make([]string, BANNER_WIDTH+2)
		for j := 0; j < BANNER_WIDTH+2; j++ {
			if i == 0 || i == BANNER_HEIGHT+1 {
				result[i][j] = BORDER
			} else if j == 0 || j == BANNER_WIDTH+1 {
				result[i][j] = BORDER
			} else {
				result[i][j] = EMPTY_CELL
			}
		}
	}
	return result
}
