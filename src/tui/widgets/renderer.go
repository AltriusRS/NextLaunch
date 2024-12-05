package widgets

import (
	"crypto/rand"
	"encoding/hex"
)

type Renderer interface {
	Id() string
	Render(int, int) []string
	Clear()
	SetWidth(width int)
	SetHeight(height int)
	SetActive(active bool)
	Size() (int, int)
}

func NodeID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
