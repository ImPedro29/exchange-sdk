package utils

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ImPedro29/exchange-sdk/constraints"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// GetURL provides json result decode to struct
func GetURL(url string, target interface{}, headers [][2]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), constraints.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	for _, h := range headers {
		req.Header.Set(h[0], h[1])
	}

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}
	r, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

func PostURL(url string, body interface{}, target interface{}, headers [][2]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), constraints.Timeout)
	defer cancel()

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	reqBody := strings.NewReader(string(bodyBytes))
	req, errNewReq := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if errNewReq != nil {
		return errNewReq
	}
	req.Header.Add("Content-type", "application/json; charset=UTF-8")
	for _, h := range headers {
		req.Header.Set(h[0], h[1])
	}

	r, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	data, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		return errRead
	}

	if r.StatusCode >= 400 {
		var message interface{}
		// if has a error message add to error return
		err := json.Unmarshal(data, &message)
		if err != nil {
			message = ""
		}

		return fmt.Errorf("request failed status %d - message :%v", r.StatusCode, message)
	}

	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	return nil
}

func SignSha256(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))

	return hex.EncodeToString(mac.Sum(nil))
}
