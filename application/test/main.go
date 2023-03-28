package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Model struct {
	Q    string `json:"q"`
	Hash string `json:"hash"`
}

func main() {
	bdy, err := getHash("hoge")
	if err != nil {
		log.Println(err)
	}
	println(bdy.Hash)
}

func getHash(q string) (*Model, error) {
	url := fmt.Sprintf("http://challenge-server.code-check.io/api/hash?q=%s", q)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("Hello Test")
	var model Model
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return nil, err
	}
	fmt.Println("Hello Test2")
	return &model, nil
}
