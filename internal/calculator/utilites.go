package calculator

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	pb "GoComputeFlow/internal/worker/proto"
)

// ParseExpression разбивает выражение на токены в польской нотации
func ParseExpression(expr string) ([]*pb.Token, error) {
	var tokens []*pb.Token
	var ops []rune
	var buffer string

	for _, ch := range expr {
		if unicode.IsDigit(ch) {
			buffer += string(ch)
		} else if strings.ContainsRune("+-*/", ch) {
			if buffer != "" {
				tokens = append(tokens, &pb.Token{Value: buffer, IsOp: false})
				buffer = ""
			}
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				tokens = append(tokens, &pb.Token{Value: string(ops[len(ops)-1]), IsOp: true})
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, ch)
		} else if ch != ' ' {
			return nil, fmt.Errorf("invalid character: %c", ch)
		}
	}

	if buffer != "" {
		tokens = append(tokens, &pb.Token{Value: buffer, IsOp: false})
	}

	for len(ops) > 0 {
		tokens = append(tokens, &pb.Token{Value: string(ops[len(ops)-1]), IsOp: true})
		ops = ops[:len(ops)-1]
	}

	return tokens, nil
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func parseAndSetTimeout(timeout string, timer *time.Duration) {
	duration, err := time.ParseDuration(timeout + "s")
	if err == nil {
		*timer = duration
	}
}
