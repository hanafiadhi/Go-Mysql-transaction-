package main

import (
	"Go-Mysql/database"
	"Go-Mysql/routers"
)

func main() {
	database.StartDB()
	routers.RootHandler().Run(":8000")
}
