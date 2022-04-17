package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWelcome(t *testing.T) {

	url := GetUrl()
	get, err := http.Get(url)
	if err != nil {
		return
	}
	assert.Truef(t, get.StatusCode == 200, "welcome page is not accessible")

}
