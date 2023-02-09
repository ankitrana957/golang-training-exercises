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
	Student string
	RollNo  int
	Id      int
	Subject string
}
