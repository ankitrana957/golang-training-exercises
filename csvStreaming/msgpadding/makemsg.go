package msgpadding

import (
	"fmt"

	"github.com/personhashing/models"
)

func MakeMsg(i models.Person) string {
	msg := fmt.Sprintf("%d%s%d%v", i.Id, i.Name, i.Age, i.Phone)
	msgSignature := fmt.Sprintf("%-*s", 100, msg)
	return msgSignature
}
