package main

import (
	"gin-my/model"
	"gin-my/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
