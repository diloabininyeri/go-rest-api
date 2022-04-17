package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindUserWallet(t *testing.T) {

	var userId int
	userId = 2
	urlGet := fmt.Sprintf("%s/user/balance?user_id=%d", GetUrl(), userId)
	get, err := http.Get(urlGet)
	if err != nil {
		return
	}
	assert.True(t, get.StatusCode == 200, "get all balances of users  page is not accessible")

}
