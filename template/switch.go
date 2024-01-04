package template

import (
	"log"
)

var data = [128]Template{}

func init() {
	data['A'] = TmplA
	data['B'] = TmplB
	data['C'] = TmplC
	data['D'] = TmplD
	data['E'] = TmplE
	data['F'] = TmplF
	data['G'] = TmplG
	data['H'] = TmplH
	data['I'] = TmplI
	data['J'] = TmplJ
	data['K'] = TmplK
	data['L'] = TmplL
	data['M'] = TmplM
	data['N'] = TmplN
	data['O'] = TmplO
	data['P'] = TmplP
	data['Q'] = TmplQ
	data['R'] = TmplR
	data['S'] = TmplS
	data['T'] = TmplT
	data['U'] = TmplU
	data['V'] = TmplV
	data['W'] = TmplW
	data['X'] = TmplX
	data['Y'] = TmplY
	data['Z'] = TmplZ
	data['0'] = Tmpl0
	data['1'] = Tmpl1
	data['2'] = Tmpl2
	data['3'] = Tmpl3
	data['4'] = Tmpl4
	data['5'] = Tmpl5
	data['6'] = Tmpl6
	data['7'] = Tmpl7
	data['8'] = Tmpl8
	data['9'] = Tmpl9
	data[' '] = TmplSpace
	data['!'] = TmplExclamationMark
	data['"'] = TmplDoubleQuote
	data['$'] = TmplDollarSign
	data['%'] = TmplPercentSign
	data['('] = TmplRightParenthesis
	data['\''] = TmplSingleQuote
	data[')'] = TmplLeftParenthesis
	data['*'] = TmplAsterisk
	data['+'] = TmplPlusSign
	data[','] = TmplComma
	data['-'] = TmplHyphenMinus
	data['.'] = TmplPeriod
	data['/'] = TmplSlash
	data[':'] = TmplColon
	data[';'] = TmplSemicolon
	data['<'] = TmplLessThanSign
	data['='] = TmplEqualSign
	data['>'] = TmplGreaterThanSign
	data['?'] = TmplQuestionMark
	data['['] = TmplLeftSquareBracket
	data['\\'] = TmplBackslash
	data[']'] = TmplRightSquareBracket
	data['^'] = TmplCaret
	data['_'] = TmplUnderscore
	data['|'] = TmplVerticalBar
	data['~'] = TmplTilde
}

func ByRune(r rune) Template {
	// to upper
	if r >= 'a' && r <= 'z' {
		r -= 'a' - 'A'
	}

	if r >= 128 || data[r] == nil {
		log.Fatalf("unsupported character: %s", string(r))
	}

	return data[r]
}
