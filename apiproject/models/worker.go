package models

import (
//	"errors"
//	"strconv"
//	"time"
)


func init() {
	//UserList = make(map[string]*User)
	//u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	//UserList["user_11111"] = &u
}

type Worker struct {
	WorkerId    string
	Token 		string
	Index 		string
}

type ReWorker struct {
	Code string
	Message string `json:null`
	//Worker Worker `json:result`
	Result interface{}
}






