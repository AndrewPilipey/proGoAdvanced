package main

type Student interface {
	Name() string
	ListeningLection()
	DoHomework(exercise string)
	PrepareForExam()
}

//просто интерфейс
