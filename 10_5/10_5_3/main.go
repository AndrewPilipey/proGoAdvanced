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

type ConfigurableNotifier interface {
	Notifier
	SetLevel(level int)
}
