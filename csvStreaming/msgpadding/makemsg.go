package msgpadding

import (
	"fmt"

	"github.com/personhashing/models"
)

func MakeMsg(c chan models.Person, m chan string) {
	for i := range c {
		msg := fmt.Sprintf("%d%s%d%v", i.Id, i.Name, i.Age, i.Phone)
		msgSignature := fmt.Sprintf("%-*s", 100, msg)
		m <- msgSignature
	}
	close(m)
}
