package random

import (
	"github.com/google/uuid"
	"strings"
)

func GetSessionID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func GetToken() string {
	u1 := uuid.New()
	u2 := uuid.New()
	return strings.ReplaceAll(u1.String()+u2.String(), "-", "")
}
