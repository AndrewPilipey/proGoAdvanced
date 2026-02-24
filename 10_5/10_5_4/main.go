package main

import "fmt"

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

type FullNotifier struct {
	Name  string
	Level int
}

func (f *FullNotifier) Send(msg string) {
	fmt.Printf("%s (уровень %d) отправка: %s\n", f.Name, f.Level, msg)
}

func (f *FullNotifier) Close() {
	fmt.Printf("%s закрыт\n", f.Name)
}

func (f *FullNotifier) SetLevel(level int) {
	f.Level = level
	fmt.Printf("%s уровень установлен на %d\n", f.Name, level)
}

func main() {
	fn := &FullNotifier{Name: "SuperNotifier"}

	fn.SetLevel(2)
	fn.Send("Сообщение с уровнем важности")
	fn.Close()
}
