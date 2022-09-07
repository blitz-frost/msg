// Package msg defines high level interfaces for communication through discrete messages, rather than continous streams.
package msg

// A Canceler provides a mechanism to cancel the current operation, typically message writing, by not actually sending any data, or at least informing the peer.
// The Cancel method should generally be treated as an alternative to the Close method.
// In particular, implementations should perform any relevant cleanup without relying on a Close call.
type Canceler interface {
	Cancel() error
}

type Closer interface {
	Close() error
}

type Conn[R Reader, W Writer] interface {
	ReaderChainer[R]
	WriterGiver[W]
}

type ExchangeReader[W Writer] interface {
	Reader         // read message; close when done
	WriterGiver[W] // get response writer
}

type ExchangeWriter[R Reader] interface {
	ReaderGiver[R] // finalize and send message and return answer
	Writer         // write message; close to abort
}

// A Reader must also define a method for reading message data.
type Reader interface {
	Closer
}

type ReaderChainer[R Reader] interface {
	ReaderChain(ReaderTaker[R]) error
}

type ReaderGiver[R Reader] interface {
	Reader() (R, error)
}

type ReaderTaker[R Reader] interface {
	ReaderTake(R) error
}

// A Writer must also define a method for writing message data.
type Writer interface {
	Closer
}

type WriteCanceler interface {
	Writer
	Canceler
}

type WriterGiver[W Writer] interface {
	Writer() (W, error)
}
