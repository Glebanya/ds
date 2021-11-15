package interfaces

type Comparable interface {
	Compare(Comparable) int
}

type SortCallback func([]Comparable) []Comparable
