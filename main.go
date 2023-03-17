package main

import (
	"github.com/gogovan/ggx-kr-payment-service/startup"
)

// @title           Common Service API
// @version         1.0
// @description     service for paymet api
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	startup.Execute()
}
