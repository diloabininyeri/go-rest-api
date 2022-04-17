package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAllBalances(t *testing.T) {

	url := GetUrl()
	get, err := http.Get(url)
	if err != nil {
		return
	}
	assert.True(t, get.StatusCode == 200, "get all balances of users  page is not accessible")

}
