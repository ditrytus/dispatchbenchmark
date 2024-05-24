package dispatchbenchmark

type foo struct {
	i int
}

func (f *foo) increment() {
	f.i++
}

type incrementor interface {
	increment()
}
