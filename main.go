package main

import (
	"github.com/Joao-Carlos2002/Api-User/config"
	"github.com/Joao-Carlos2002/Api-User/router"
)

func main() {
	config.Init()
	router.Router()
}
