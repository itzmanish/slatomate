package utils

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	out := RandomString(28)
	fmt.Println(out, len(out))
}
