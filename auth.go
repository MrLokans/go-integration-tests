package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type authService struct {
	Base string
}

type loginResponse struct {
	Token string `json: "token"`
}

func (a *authService) Login(username, password string) loginResponse {
	_, body, err := post(a.Base+"/login", map[string]string{"username": username, "password": password})
	lr := loginResponse{}

	if err != nil {
		return lr
	}
	json.Unmarshal(body, &lr)

	return lr
}

func post(postURL string, payload map[string]string) (int, []byte, error) {
	form := url.Values{}
	for k, v := range payload {
		form.Add(k, v)
	}

	req, _ := http.NewRequest("POST", postURL, bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, body, nil
}
