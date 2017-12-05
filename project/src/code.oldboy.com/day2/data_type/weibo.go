package main

/*
daren	hongming
0  			0  			0  0  0  0  0  0 

1. 设置其中一位为1
	用1左移n与目标数做或操作;
2. 
*/

import(
	"fmt"
)

const (
	HongMing = 1
	DaRen = 1 << 1
	Vip = 1 << 2
)
type User struct {
	name string
	//表示0000 0000,用每一位表示不同的属性
	flag uint8
}


func set_flag(user User, isSet bool, flag int) User {
	if isSet == true {
		user.flag = user.flag | flag
	}
}


func set_hongming(user User, isSet bool) User {
	if isSet == true {
		user.flag = user.flag | 1
	} else {
		user.flag = user.flag & 0
	}
	return user
}

func is_hongming(user User) bool {
	result := user.flag & 1
	return result == 1
}

func set_daren(user User, isSet bool) User {
	if isSet == true {
		user.flag = user.flag | DaRen
	}else {
		user.flag = user.flag & DaRen
	}
}

func weibo_test() {
	var user User
	user.name = "test01"
	user.flag = 0

	result := is_hongming(user)
	fmt.Printf("user is hongming:%t\n", result)

	user = set_hongming(user, true)
	result = is_hongming(user)
	fmt.Printf("user is hongming:%t\n", result)

}