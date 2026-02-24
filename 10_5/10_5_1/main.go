package main

type Sender interface {
	Send(msg string)
}

type Closer interface {
	Close()
}
