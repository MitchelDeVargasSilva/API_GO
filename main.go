package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	DB = iniciaDB("tarefas.db")
	defer DB.Close()

	Rotas(router)

	router.Run(":3000")

}
