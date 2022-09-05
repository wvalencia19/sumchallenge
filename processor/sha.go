package processor

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func IntToHexSHA(value int) string {
	bytes := []byte(strconv.Itoa(value))
	h := sha256.New()
	h.Write(bytes)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
