package main

type Pay interface {
	Pay(user_id int64, money float32) error
}