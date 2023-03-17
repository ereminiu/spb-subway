package models

type Station struct {
	X     int
	Y     int
	Name  string
	Line  string
	Adj   []string
	Dist  []int
	Shift string
	Dx    int
	Dy    int
}

type Line struct {
	X1    int
	X2    int
	Y1    int
	Y2    int
	Color string
}

type Lable struct {
	X    int
	Y    int
	Name string
}

func NewLable() Lable {
	return Lable{0, 0, ""}
}

type Edge struct {
	To   int
	Cost int
}

func NewEdge(To, Cost int) *Edge {
	return &Edge{To, Cost}
}
