package main

import "GO_mysql/DB"

func main() {
	DB.InitDB()
	DB.Select()
}
