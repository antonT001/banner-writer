package main

import (
	"fmt"
	"log"
	"time"

	"github.com/antonT001/banner-writer/templates"
)

func main() {
	word := requestWord()
	result := templates.New()
	for _, r := range word {
		result.SetDelimeter()
		result.JoinObject(templateByRune(r))
	}
	result.Println()
}

func requestDate() time.Time {
	fmt.Println("insert the year and press return(enter)")
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		log.Fatalf("failed scan eyar: %v", err)
	}
	t, err := time.Parse(time.DateTime, value+"-01-01 15:04:05")
	if err != nil {
		log.Fatalf("failed parse eyar: %v", err)
	}
	return t
}

func requestPath() string {
	fmt.Println("paste the path to save the repository and press return(enter)")
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		log.Fatalf("failed scan path: %v", err)
	}
	return value
}

func requestWord() string {
	fmt.Println("insert the word and press return(enter)")
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		log.Fatalf("failed scan word: %v", err)
	}
	return value
}

func templateByRune(r rune) templates.Template {
	switch r {
	case 'a', 'A':
		return templates.TmplA
	case 'b', 'B':
		return templates.TmplB
	case 'c', 'C':
		return templates.TmplC
	case 'd', 'D':
		return templates.TmplD
	case 'e', 'E':
		return templates.TmplE
	case 'f', 'F':
		return templates.TmplF
	case 'g', 'G':
		return templates.TmplG
	case 'h', 'H':
		return templates.TmplH
	case 'i', 'I':
		return templates.TmplI
	case 'j', 'J':
		return templates.TmplJ
	case 'k', 'K':
		return templates.TmplK
	case 'l', 'L':
		return templates.TmplL
	case 'm', 'M':
		return templates.TmplM
	case 'n', 'N':
		return templates.TmplN
	case 'o', 'O':
		return templates.TmplO
	case 'p', 'P':
		return templates.TmplP
	case 'q', 'Q':
		return templates.TmplQ
	case 'r', 'R':
		return templates.TmplR
	case 's', 'S':
		return templates.TmplS
	case 't', 'T':
		return templates.TmplT
	case 'u', 'U':
		return templates.TmplU
	case 'v', 'V':
		return templates.TmplV
	case 'w', 'W':
		return templates.TmplW
	case 'x', 'X':
		return templates.TmplX
	case 'y', 'Y':
		return templates.TmplY
	case 'z', 'Z':
		return templates.TmplZ
	case '0':
		return templates.Tmpl0
	case '1':
		return templates.Tmpl1
	case '2':
		return templates.Tmpl2
	case '3':
		return templates.Tmpl3
	case '4':
		return templates.Tmpl4
	case '5':
		return templates.Tmpl5
	case '6':
		return templates.Tmpl6
	case '7':
		return templates.Tmpl7
	case '8':
		return templates.Tmpl8
	case '9':
		return templates.Tmpl9
	default:
		return templates.TmplUnknownSymbol
	}

}
