package solution

import (
	_ "embed"
	"errors"
	"github.com/rbutcher/aoc2024/internal/helpers"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

//go:embed input/day3.txt
var day3Input string

type day3 struct {
	operators *helpers.Stack[string]
	operands  *helpers.Stack[int]
}

func NewDay3() Solution {
	result := &day3{
		operands:  helpers.NewStack[int](),
		operators: helpers.NewStack[string](),
	}

	for i := 0; i < len(day3Input); {
		startIdx := strings.Index(day3Input[i:], "mul(")
		if startIdx == -1 {
			break
		}
		startIdx += 4 + i

		endIdx := strings.Index(day3Input[startIdx:], ")")
		if endIdx == -1 {
			break
		}
		endIdx += startIdx

		i = startIdx + 1
		rawOperands := day3Input[startIdx:endIdx]
		operandSplit := strings.Split(rawOperands, ",")
		if len(operandSplit) != 2 {
			ins := day3Input[startIdx:endIdx]
			log.Debug().Str("ins", ins).Msg("invalid instruction")
			continue
		}

		l, err := strconv.Atoi(operandSplit[0])
		if err != nil {
			ins := day3Input[startIdx:endIdx]
			log.Debug().Str("ins", ins).Msg("invalid instruction")
			continue
		}

		r, err := strconv.Atoi(operandSplit[1])
		if err != nil {
			ins := day3Input[startIdx:endIdx]
			log.Debug().Str("ins", ins).Msg("invalid instruction")
			continue
		}

		result.operands.Push(r)
		result.operands.Push(l)
		result.operators.Push("mul")
	}

	return result
}

func (d *day3) Part1() (string, error) {
	sum := 0
	for !d.operators.IsEmpty() && !d.operands.IsEmpty() {
		d.operators.Pop()
		l, _ := d.operands.Pop()
		r, _ := d.operands.Pop()

		sum += l * r
	}

	return strconv.Itoa(sum), nil
}

func (d *day3) Part2() (string, error) {
	return "", errors.New("not implemented")
}
