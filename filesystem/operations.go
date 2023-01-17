package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Student struct {
	Name   string
	Age    int
	RollNo int
	PhnNo  []int
}

func CreateJsonFile(name string) (myfile *os.File) {
	myfile, e := os.Create(name)
	if e != nil {
		panic("Error in creation of file")
	}
	return myfile
}

func ReadFromJson() []byte {
	myfile, err := os.Open("data.json")
	if err != nil {
		panic("Error in reading from the data file")
	}
	byteValue, _ := ioutil.ReadAll(myfile)
	defer myfile.Close()
	return byteValue
}

func Analyse(byteValue []byte) ([]Student, []Student) {
	var s []Student
	var primary []Student
	var secondary []Student
	json.Unmarshal(byteValue, &s)
	for i := 0; i < len(s); i++ {
		if s[i].Age < 10 {
			primary = append(primary, s[i])
		} else {
			secondary = append(secondary, s[i])
		}
	}
	return primary, secondary
}

func WriteFile(primary, secondary []Student) {
	a, err := json.Marshal(primary)
	b, err1 := json.Marshal(secondary)
	if err != nil && err1 != nil {
		panic("Error in marshaling the data")
	}
	primaryFile := CreateJsonFile("primaryData.json")
	primaryFile.WriteString(string(a))
	primaryFile.Close()

	secondaryFile := CreateJsonFile("secondaryData.json")
	secondaryFile.WriteString(string(b))
	secondaryFile.Close()
}
