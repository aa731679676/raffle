package utils

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

func GetUuid() (string, error) {
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	id := u1.String()
	return strings.Replace(id, "-", "", -1), nil
}
