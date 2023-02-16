package models

type Person struct {
	Id    int
	Name  string
	Age   int
	Phone string
}

func (p Person) SetHashPhone(phonehash string) Person {
	p.Phone = phonehash
	return p
}
