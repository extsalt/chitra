package main

import (
	"fmt"
	"os"
	"regexp"
	"unicode"
)

// Const for representation
const (
	TokenLeftBrace        = "Left_Brace"
	TokenRightBrace       = "Right_Brace"
	TokenRightParenthesis = "Right_Parenthesis"
	TokenLeftParenthesis  = "Left_Parenthesis"
	TokenPlus             = "Plus"
	TokenEqual            = "Equal"
	TokenEqualEqual       = "Equal_Equal"
	TokenMinus            = "Minus"
	TokenSemiColon        = "Semicolon"
	TokenLessThan         = "Less_Than"
	TokenGreaterThan      = "Greater_Than"
	TokenGreaterOrEqual   = "Greater_Or_Equal"
	TokenLessOrEqual      = "Less_Or_Equal"
	TokenBang             = "Bang"
	TokenBangEqual        = "Bang_Equal"
	TokenInteger          = "Integer"
	TokenFloat            = "Float"
	TokenString           = "String"
	TokenChar             = "Char"
	TokenDigit            = "Digit"
)

// check if next string is reserved identifier
func hasKey(k string) bool {
	reservedKeywords := map[string]string{
		"int":    TokenInteger,
		"float":  TokenFloat,
		"string": TokenString,
		"char":   TokenChar,
	}

	_, ok := reservedKeywords[k]

	return ok
}

type token struct {
	literal   string
	tokenType string
	line      int
}

func main() {
	fmt.Println(lexer(rawSource()))
}

// read raw source code
func rawSource() string {
	file := "./code/test.cht"

	data, err := os.ReadFile(file)

	if err != nil {
		panic("Error in reading file")
	}

	return string(data)
}

func lexer(src string) *[]token {
	var tokens []token
	line := 1
	whitespace := 0

	for i := 0; i < len(src); i++ {
		// Ignore comment for now.
		switch string(src[i]) {

		// update new line
		case "\n":
			line++
			break

			// ignore whitespaces
		case " ":
			whitespace++
			for string(src[i+1]) == " " {
				i++
				whitespace++
			}
			break

		case "(":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenLeftParenthesis,
				line:      line,
			})
			break

		case "!":
			if string(src[i+1]) == "=" {
				tokens = append(tokens, token{
					literal:   string(src[i]) + string(src[i+1]),
					tokenType: TokenBangEqual,
					line:      line,
				})
				i++
			} else {
				tokens = append(tokens, token{
					literal:   string(src[i]),
					tokenType: TokenBang,
					line:      line,
				})
			}
			break

		case ")":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenRightParenthesis,
				line:      line,
			})
			break

		case "{":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenLeftBrace,
				line:      line,
			})
			break

		case "}":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenRightBrace,
				line:      line,
			})
			break

		case "=":
			if string(src[i+1]) == "=" {
				tokens = append(tokens, token{
					literal:   string(src[i]) + string(src[i+1]),
					tokenType: TokenEqualEqual,
					line:      line,
				})
				i++
			} else {
				tokens = append(tokens, token{
					literal:   string(src[i]),
					tokenType: TokenEqual,
					line:      line,
				})
			}
			break

		case ";":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenSemiColon,
				line:      line,
			})
			break

		case "+":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenPlus,
				line:      line,
			})
			break

		case "-":
			tokens = append(tokens, token{
				literal:   string(src[i]),
				tokenType: TokenMinus,
				line:      line,
			})
			break

		case ">":
			if string(src[i+1]) == "=" {
				tokens = append(tokens, token{
					literal:   string(src[i]) + string(src[i+1]),
					tokenType: TokenGreaterOrEqual,
					line:      line,
				})
				i++
			} else {
				tokens = append(tokens, token{
					literal:   string(src[i]),
					tokenType: TokenGreaterThan,
					line:      line,
				})
			}
			break

		case "<":
			if string(src[i+1]) == "=" {
				tokens = append(tokens, token{
					literal:   string(src[i]) + string(src[i+1]),
					tokenType: TokenLessOrEqual,
					line:      line,
				})
				i++
			} else {
				tokens = append(tokens, token{
					literal:   string(src[i]),
					tokenType: TokenLessThan,
					line:      line,
				})
			}
			break

		default:
			if unicode.IsDigit(rune(src[i])) {
				d := digit(&i, src)
				tokens = append(tokens, token{
					literal:   d,
					tokenType: TokenDigit,
					line:      line,
				})
			}

			result, err := regexp.Match("[a-zA-Z_]", []byte(string(src[i])))

			if err != nil {
				panic("something happened while executing regex")
			}

			if result {

			}
		}
	}

	return &tokens
}

// Capture digit from src
func digit(i *int, src string) string {
	start := *i
	end := *i

	for unicode.IsDigit(rune(src[*i])) {
		*i++
	}

	end = *i

	return src[start:end]
}
