package main

type Sender interface {
	Send(msg string)
}

type Closer interface {
	Close()
}

type Notifier interface {
	Sender
	Closer
}
