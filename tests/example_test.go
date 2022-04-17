package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	var a string = "Dılo"
	var b string = "Dılo"

	assert.Equal(t, a, b, "The two words should be the same.")

}
