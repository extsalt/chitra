package main

import (
	"fmt"
	"os"
)

const (
	TokenLeftBrace = iota
	TokenRightBrace
	TokenRightParenthesis
	TokenLeftParenthesis
	TokenPlus
	TokenEqual
	TokenMinus
	TokenSemiColon
)

type token struct {
	literal   string
	tokenType int
	line      int
}

func main() {
	lexemes := lexer(rawSource())
	fmt.Print(lexemes)
}

func rawSource() string {
	file := "./code/test.cht"

	data, err := os.ReadFile(file)

	if err != nil {
		panic("Error in reading file")
	}

	return string(data)
}

func lexer(src string) *[]token {
	line := 1

	var lexeme []token

	for _, c := range src {
		if c == '\n' {
			line++
			continue
		}

		if c == ' ' || c == '\t' {
			continue
		}

		if c == '(' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenLeftParenthesis,
				line:      line,
			})
		}

		if c == ')' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenRightParenthesis,
				line:      line,
			})
		}

		if c == '{' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenLeftBrace,
				line:      line,
			})
		}

		if c == '}' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenRightBrace,
				line:      line,
			})
		}

		if c == '+' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenPlus,
				line:      line,
			})
		}

		if c == '=' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenEqual,
				line:      line,
			})
		}

		if c == '-' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenMinus,
				line:      line,
			})
		}

		if c == ';' {
			lexeme = append(lexeme, token{
				literal:   string(c),
				tokenType: TokenSemiColon,
				line:      line,
			})
		}
	}

	return &lexeme
}
