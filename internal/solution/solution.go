package solution

type Solution interface {
	Part1() (string, error)
	Part2() (string, error)
}

type Constructor func() Solution
