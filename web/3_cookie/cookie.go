package main

import (
	"time"
	"net/http"
)

func main() {
	expiration := time.Time{}
	expiration.Year() +=1
	cookie := http.Cookie{Name:"hello",Value:"world",Expires:expiration}
	http.SetCookie()
}
