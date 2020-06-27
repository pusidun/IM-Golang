package main

import (
	"im-golang/dao"
	"im-golang/server"
)

func main() {
	dao.InitDB()
	r := server.NewRouter()
	r.Run(":8080")
}
