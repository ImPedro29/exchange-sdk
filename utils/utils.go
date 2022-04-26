package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// GetURL provides json result decode to struct
func GetURL(url string, target interface{}, headers [][2]string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func SignSha256(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))

	return hex.EncodeToString(mac.Sum(nil))
}
