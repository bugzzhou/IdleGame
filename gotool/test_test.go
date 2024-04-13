package gotool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestGjson(t *testing.T) {
	a := 1
	b := 1
	assert.Equal(t, a, b)
	TestGjson()
}
