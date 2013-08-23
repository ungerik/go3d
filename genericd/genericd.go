package generic

type T interface {
	Cols() int
	Rows() int
	Size() int
	Slice() []float64
	Get(row, col int) float64
	IsZero() bool
}
