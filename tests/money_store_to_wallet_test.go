package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMoneyStore(t *testing.T) {

	values := map[string]int{"user_id": 3, "amount": 450}
	payload, _ := json.Marshal(values)

	url := fmt.Sprintf("%s/balance", GetUrl())
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	assert.Truef(t, response.StatusCode == 200, "welcome page is not accessible")

}
