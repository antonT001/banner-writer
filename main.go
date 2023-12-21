package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	word := requestWord()
	result := NewTemplate()
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

func templateByRune(r rune) template {
	switch r {
	case 'a', 'A':
		return tmplA
	case 'b', 'B':
		return tmplB
	case 'c', 'C':
		return tmplC
	case 'd', 'D':
		return tmplD
	case 'e', 'E':
		return tmplE
	case 'f', 'F':
		return tmplF
	case 'g', 'G':
		return tmplG
	case 'h', 'H':
		return tmplH
	case 'i', 'I':
		return tmplI
	case 'j', 'J':
		return tmplJ
	case 'k', 'K':
		return tmplK
	case 'l', 'L':
		return tmplL
	case 'm', 'M':
		return tmplM
	case 'n', 'N':
		return tmplN
	case 'o', 'O':
		return tmplO
	case 'p', 'P':
		return tmplP
	case 'q', 'Q':
		return tmplQ
	case 'r', 'R':
		return tmplR
	case 's', 'S':
		return tmplS
	case 't', 'T':
		return tmplT
	case 'u', 'U':
		return tmplU
	case 'v', 'V':
		return tmplV
	case 'w', 'W':
		return tmplW
	case 'x', 'X':
		return tmplX
	case 'y', 'Y':
		return tmplY
	case 'z', 'Z':
		return tmplZ
	case '0':
		return tmpl0
	case '1':
		return tmpl1
	case '2':
		return tmpl2
	case '3':
		return tmpl3
	case '4':
		return tmpl4
	case '5':
		return tmpl5
	case '6':
		return tmpl6
	case '7':
		return tmpl7
	case '8':
		return tmpl8
	case '9':
		return tmpl9
	default:
		return tmplUnknownSymbol
	}

}
