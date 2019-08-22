package main

import (
	"filestore-server/config"
	"filestore-server/router"
)

func main() {
	r := router.Router()
	r.Run(config.ServerPort)
}
