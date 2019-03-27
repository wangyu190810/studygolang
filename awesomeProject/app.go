package main

import (
	db "awesomeProject/database"
	"awesomeProject/middleware"
)



func main() {
	middleware.Conn()
	defer db.SqlDB.Close()
	defer middleware.AmqpConn.Close()
	router := InitRouter()
	router.Run(":8000")
}