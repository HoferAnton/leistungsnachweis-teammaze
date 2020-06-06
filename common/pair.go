package common

type Pair struct {
	first, second interface{}
}

func NewPair(first interface{}, second interface{}) Pair {
	return Pair{first, second}
}

func (p Pair) GetFirst() interface{} {
	return p.first
}

func (p Pair) GetSecond() interface{} {
	return p.second
}

func (p Pair) Compare(that Pair) bool {

	return p.GetFirst() == that.GetFirst() &&
		p.GetSecond() == that.GetSecond()
}
