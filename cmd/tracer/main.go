package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type LoginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			err, _ := json.Marshal(r)
			_ = os.WriteFile("trace.json", err, os.ModePerm)
		}
	}()

	jsonBody := []byte(`{"password": "42246ccaa7b8852d1478", "username": "admin"}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://backend-mirror:8080/api/v1/login", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("incorrect status " + resp.Status)
	}
	jsonBody, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var loginResp LoginResponse
	err = json.Unmarshal(jsonBody, &loginResp)
	if err != nil {
		panic(err)
	}
	expire, err := time.Parse(time.RFC3339, loginResp.Expire)
	if !expire.After(time.Now()) {
		panic("incorrect expire " + loginResp.Expire)
	}

	req, err = http.NewRequest(http.MethodDelete, "http://backend-mirror:8080/api/v1/logout", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+loginResp.Token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("incorrect status " + resp.Status)
	}

	req, err = http.NewRequest(http.MethodGet, "http://jaeger:16686/api/traces?service=backend-mirror", nil)
	if err != nil {
		panic(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("incorrect status " + resp.Status)
	}
	jsonBody, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("traces/trace.json", jsonBody, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
