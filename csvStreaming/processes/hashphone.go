package processes

import (
	"crypto/sha256"
	"fmt"

	"github.com/personhashing/models"
)

func ProcessHashing(c models.JSON) models.JSON {
	phoneHash := fmt.Sprintf("%x", sha256.Sum256([]byte(c["phone"].(string))))
	c["phone"] = phoneHash
	return c
}
