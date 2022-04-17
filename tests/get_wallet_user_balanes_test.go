package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAllBalances(t *testing.T) {

	url := GetUrl()
	getUrl := fmt.Sprintf("%s/balance", url)
	get, err := http.Get(getUrl)
	if err != nil {
		return
	}
	assert.True(t, get.StatusCode == 200, "get all balances of users  page is not accessible")

}
