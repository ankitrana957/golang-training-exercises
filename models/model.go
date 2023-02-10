package models

type Student struct {
	Name   string
	Age    int
	RollNo int
}

type Subject struct {
	Id   int
	Name string
}

type Record struct {
	RollNo int
	Id     int
}

type SubName struct {
	Subject string
}
