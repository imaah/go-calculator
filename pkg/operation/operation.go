package operation

type Operation interface {
	Eval() float64
	String() string
}
