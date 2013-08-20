package generic

type T interface {
	Cols() int
	Rows() int
	Size() int
	Slice() []float32
	Get(row, col int) float32
	IsZero() bool
}
