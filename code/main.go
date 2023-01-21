package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var sevice_name = ""

func init() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int31n(100)
	sevice_name = fmt.Sprintf("%d", r)
}

func main() {
	u := &User{name: sevice_name, logger: log.Default()}
	http.HandleFunc("/user", u.Get)
	err := http.ListenAndServe(":1002", nil)
	if err != nil {
		panic(err)
	}
}

type User struct {
	name   string
	logger *log.Logger
}

func (uh *User) Get(rsp http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("Hello Every Body! By %s", uh.name)
	uh.logger.Println(msg)
	rsp.Write([]byte(msg))
}
