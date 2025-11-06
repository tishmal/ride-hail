package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

func GenerateUUID() (string, error) {
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	// Устанавливаем версию (4)
	b[6] = (b[6] & 0x0f) | 0x40
	// Устанавливаем вариант (RFC 4122)
	b[8] = (b[8] & 0x3f) | 0x80

	uuid := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4],
		b[4:6],
		b[6:8],
		b[8:10],
		b[10:16],
	)

	return uuid, nil
}
