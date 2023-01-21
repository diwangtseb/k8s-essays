package main

import (
	"net/http"
)

func main() {
	u := &User{}
	http.HandleFunc("/user", u.Get)
	http.ListenAndServe(":1001", nil)
}

type User struct{}

func (uh *User) Get(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("Hello Every Body!"))
}
