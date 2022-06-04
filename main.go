package main

import (
	"SimpleDY/handler"
	"SimpleDY/initial"
)

func main() {
	initial.LoadConfig()
	initial.Mysql()
	handler.Handler()

}
