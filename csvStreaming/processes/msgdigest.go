package processes

import (
	"fmt"

	"github.com/personhashing/models"
)

func ProcessMsg(m models.JSON) models.JSON {
	msg := fmt.Sprintf("%d%s%d%v", m["id"], m["name"], m["age"], m["phone"])
	msgSignature := fmt.Sprintf("%-*s", 100, msg)
	k := models.JSON{"value": msgSignature}
	return k
}
