package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	Tokens []Token
	Cursor int // This is your incrementer
}

// Peak at the current character without moving forward
func (p *Parser) peek() Token {
	if p.Cursor >= len(p.Tokens) {
		return Token{"EOF", "EOF"} // End of input
	}
	return p.Tokens[p.Cursor]
}

// Consume the current character and move the incrementer forward
func (p *Parser) advance() Token {
	ch := p.peek()
	if ch.Kind != "EOF" {
		p.Cursor++
	}
	return ch
}

// TODO (scalar-values): implement per the lesson description.
func ParseToken(parser *Parser) (string, error) {
	// fmt.Println(tokens[0].Value)
	// fmt.Println(line)
	switch parser.peek().Value {
	case "true":
		return "True", nil
	case "false":
		return "False", nil
	case "null":
		return "None", nil
	case "[":
		parser.advance()
		result, err := ParseArray(parser)
		if err != nil {
			return "", err
		}
		return "[" + strings.Join(result, ", ") + "]", nil
	}
	// if b, err := strconv.ParseInt(line, 10, 64); err == nil {
	// 	if b == -0 {
	// 		return "0"
	// 	}
	// 	return line
	// } else if f, err := strconv.ParseFloat(line, 64); err == nil {

	// 	if !math.IsNaN(f) && !math.IsInf(f, 0) {
	// 		float_str := strconv.FormatFloat(f, 'f', -1, 64)
	// 		if !strings.Contains(float_str, ".") {
	// 			return float_str + ".0"
	// 		}
	// 		return float_str
	// 	}
	// } else
	if parser.peek().Value[0] == '"' && parser.peek().Value[len(parser.peek().Value)-1] == '"' {
		return "'" + parser.peek().Value[1:len(parser.peek().Value)-1] + "'", nil
	} else {
		pattern := regexp.MustCompile(`\A-?(?:0|[1-9][0-9]*)(?:\.[0-9]+)?(?:[eE][+-]?[0-9]+)?\z`)
		if f, err := strconv.ParseFloat(parser.peek().Value, 64); pattern.MatchString(parser.peek().Value) && err == nil {

			float_str := strconv.FormatFloat(f, 'f', -1, 64)

			return float_str, nil
		}
	}

	return "", fmt.Errorf("Invalid Input")
}

func ParseArray(parser *Parser) ([]string, error) {
	// var output []string
	// // var stack []string
	// if parser.peek().Kind == "PUNCT" && parser.peek().Value == "]" {
	// 	parser.advance()
	// 	return output
	// }
	// val, err := ParseToken(parser)
	// if err != nil {
	// 	return []string{"Error 2"}
	// }
	// output = append(output, val)
	// parser.advance()
	// // fmt.Println(tokens)
	// for parser.advance().Value == "," {

	// 	// print(parser.Tokens)
	// 	if parser.peek().Kind == "EOF" || parser.peek().Kind == "PUNCT" {
	// 		fmt.Println("ERR trailing comma")
	// 	}
	// 	// fmt.Println(parser.Cursor)
	// 	val, err = ParseToken(parser)
	// 	// fmt.Println(parser.Cursor)
	// 	if err != nil {
	// 		return []string{"Error 1"}
	// 	}
	// 	output = append(output, val)
	// 	parser.advance()

	// }
	// // if parser.peek().Value != "]" {
	// // 	fmt.Println("ERR no ending ]")
	// // }
	// parser.advance()
	// return output

	var output []string
	if parser.peek().Kind == "PUNCT" && parser.peek().Value == "]" {
		parser.advance()
		return output, nil
	}
	val, err := ParseToken(parser)
	if err != nil {
		return nil, fmt.Errorf("Error")
	}
	output = append(output, val)
	parser.advance()
	for parser.peek().Kind == "PUNCT" && parser.peek().Value == "," {
		parser.advance()
		val, err := ParseToken(parser)
		if err != nil {
			return nil, fmt.Errorf("ERR trailing comma")
		}
		output = append(output, val)
		parser.advance()
	}
	// fmt.Println(parser.peek().Value)
	if parser.peek().Value != "]" {
		return nil, fmt.Errorf("Error 2")
	}

	return output, nil

}
