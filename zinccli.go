package zinccli

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	ZincCli struct {
		Host     string
		User     string
		Password string
	}
)

func NewZincCli(host string, user string, password string) ZincCli {
	return ZincCli{
		Host:     host,
		User:     user,
		Password: password,
	}
}

func (z *ZincCli) ZincRequest(method string, url string, data string) ([]byte, error) {
	// Create a new request
	url = z.Host + "/api" + url
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(z.User, z.Password)
	req.Header.Set("Content-Type", "application/json")

	// call the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("zinc server not found")
		} else if resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("zinc user/password not right")
		}
		return nil, err
	}

	// parse response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
