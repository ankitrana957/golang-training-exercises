package msgdigest

import (
	"fmt"

	"github.com/personhashing/models"
)

func CreateMsgSignature(p models.Person) string {
	msg := fmt.Sprintf("%d%s%d%v", p.Id, p.Name, p.Age, p.Phone)
	msgSignature := fmt.Sprintf("%-*s", 100, msg)
	return msgSignature
}
