package solution

import (
	_ "embed"
	"github.com/rbutcher/aoc2024/internal/helpers"
	"github.com/rs/zerolog/log"
	"math"
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
		doIdx := strings.Index(day3Input[i:], "do")
		if doIdx == -1 {
			doIdx = math.MaxInt
		} else {
			doIdx += i
		}

		dontIdx := strings.Index(day3Input[i:], "don't")
		if dontIdx == -1 {
			dontIdx = math.MaxInt
		} else {
			dontIdx += i
		}

		mulIdx := strings.Index(day3Input[i:], "mul(")
		if mulIdx == -1 {
			mulIdx = math.MaxInt
		} else {
			mulIdx += 4 + i
		}

		startIdx := helpers.Min(helpers.Min(dontIdx, doIdx), mulIdx)
		if startIdx > len(day3Input) {
			break
		}

		i = startIdx + 1
		if startIdx == dontIdx {
			result.operators.Push("don't")
			continue
		}

		if startIdx == doIdx {
			result.operators.Push("do")
			continue
		}

		endIdx := strings.Index(day3Input[startIdx:], ")")
		if endIdx == -1 {
			break
		}
		endIdx += startIdx

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
		var op string
		for ; op != "mul"; op, _ = d.operators.Pop() {
		}

		l, _ := d.operands.Pop()
		r, _ := d.operands.Pop()

		sum += l * r
	}

	return strconv.Itoa(sum), nil
}

func (d *day3) Part2() (string, error) {
	sum := 0
	for !d.operators.IsEmpty() && !d.operands.IsEmpty() {
		op, _ := d.operators.Pop()
		isum := 0
		for ; op == "mul"; op, _ = d.operators.Pop() {
			l, _ := d.operands.Pop()
			r, _ := d.operands.Pop()
			isum += l * r
		}

		if op != "don't" {
			sum += isum
		}
	}

	return strconv.Itoa(sum), nil
}
