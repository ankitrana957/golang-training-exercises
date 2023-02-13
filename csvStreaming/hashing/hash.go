package hashing

import (
	"crypto/sha256"
	"fmt"

	"github.com/personhashing/models"
)

func HashString(i models.Person) models.Person {
	newObj := i
	hashValue := sha256.Sum256([]byte(newObj.Phone))
	phonehash := fmt.Sprintf("%x", hashValue)
	newObj.Phone = phonehash
	return newObj
}
