package msg

// A ConnBlock can be used to form a modular Conn.
// The read and write sides can be modified independently, while being passable as a single value in function calls.
type ConnBlock[R Reader, W Writer] struct {
	ReaderChainer[R]
	WriterGiver[W]
}

type Void[R Reader] struct{}

func (x Void[R]) ReaderTake(r R) error {
	return r.Close()
}
