package hashing

import (
	"crypto/sha256"
	"fmt"

	"github.com/personhashing/models"
)

func HashString(c chan models.Person, m chan models.Person) {
	for i := range c {
		newObj := i
		hashValue := sha256.Sum256([]byte(newObj.Phone))
		phonehash := fmt.Sprintf("%x", hashValue)
		newObj.Phone = phonehash
		m <- newObj
	}

	close(m)
}
