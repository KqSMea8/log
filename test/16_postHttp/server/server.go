package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Test struct {
	A int
	B int
}


func ReloadServer(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		fmt.Println("r nil")
		return
	}
	fmt.Println("method", r.Method)
	result, _:= ioutil.ReadAll(r.Body)
	test := &Test{}
	json.Unmarshal(result, test)
	fmt.Println("test:", test)

}
//json.Unmarshal(this.Ctx.Input.RequestBody, &debug)
func StartDebug() {
	fmt.Println("start debug")

	http.HandleFunc("/reload", ReloadServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}


func main() {
	var sig = make(chan os.Signal, 1)
	go StartDebug()

	for {
		select {
		case <-sig:
			fmt.Println("Signal received.  Shutting down...\n")
			return
			//case <-ticker.C:
			//	fmt.Println("break heart")
		}
	}
}

