package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Test struct {
	A int
	B int
}

func main() {
	//postValue := url.Values{
	//	"email": {"xx@xx.com"},
	//	"password": {"123456"},
	//}

	//postString := postValue.Encode()

	test := Test{1, 2}

 	testBy, _ := json.Marshal(test)

 	bytes.NewReader(testBy)

	req, err := http.NewRequest("POST","http://127.0.0.1:12345/reload", bytes.NewReader(testBy))
	if err != nil {
		// handle error
	}

	// 表单方式(必须)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//AJAX 方式请求
	//req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}