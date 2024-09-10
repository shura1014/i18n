package i18n

import (
	"os"
	"testing"
)

func TestReadContent(t *testing.T) {

	file, err := os.ReadFile("message/en_us.json")
	if err != nil {
		return
	}

	t.Log(string(file))
}
