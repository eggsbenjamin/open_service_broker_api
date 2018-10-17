package uuid

import (
	"crypto/rand"
	"fmt"

	"github.com/pkg/errors"
)

type UUID string

func New() (UUID, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.Wrap(err, "New: error reading random number")
	}

	return UUID(fmt.Sprintf(
		"%X-%X-%X-%X-%X",
		b[0:4],
		b[4:6],
		b[6:8],
		b[8:10],
		b[10:],
	)), nil
}
