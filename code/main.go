package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var sevice_name = ""

func init() {
	rand.Seed(time.Second.Nanoseconds())
	r := rand.Int31n(100)
	sevice_name = fmt.Sprintf("%d", r)
}

func main() {
	u := &User{name: sevice_name}
	http.HandleFunc("/user", u.Get)
	http.ListenAndServe(":1001", nil)
}

type User struct {
	name string
}

func (uh *User) Get(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte(fmt.Sprintf("Hello Every Body! By %s", uh.name)))
}
