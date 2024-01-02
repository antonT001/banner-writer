package template

import "log"

func ByRune(r rune) Template {
	switch r {
	case 'a', 'A':
		return TmplA
	case 'b', 'B':
		return TmplB
	case 'c', 'C':
		return TmplC
	case 'd', 'D':
		return TmplD
	case 'e', 'E':
		return TmplE
	case 'f', 'F':
		return TmplF
	case 'g', 'G':
		return TmplG
	case 'h', 'H':
		return TmplH
	case 'i', 'I':
		return TmplI
	case 'j', 'J':
		return TmplJ
	case 'k', 'K':
		return TmplK
	case 'l', 'L':
		return TmplL
	case 'm', 'M':
		return TmplM
	case 'n', 'N':
		return TmplN
	case 'o', 'O':
		return TmplO
	case 'p', 'P':
		return TmplP
	case 'q', 'Q':
		return TmplQ
	case 'r', 'R':
		return TmplR
	case 's', 'S':
		return TmplS
	case 't', 'T':
		return TmplT
	case 'u', 'U':
		return TmplU
	case 'v', 'V':
		return TmplV
	case 'w', 'W':
		return TmplW
	case 'x', 'X':
		return TmplX
	case 'y', 'Y':
		return TmplY
	case 'z', 'Z':
		return TmplZ
	case '0':
		return Tmpl0
	case '1':
		return Tmpl1
	case '2':
		return Tmpl2
	case '3':
		return Tmpl3
	case '4':
		return Tmpl4
	case '5':
		return Tmpl5
	case '6':
		return Tmpl6
	case '7':
		return Tmpl7
	case '8':
		return Tmpl8
	case '9':
		return Tmpl9
	case ' ':
		return TmplSpace
	default:
		log.Fatalf("unsupported character: %s", string(r))
	}
	return nil
}
