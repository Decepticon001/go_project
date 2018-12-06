package main

import(
	db "gin_blog/database"
)
func main(){
	defer db.SqlDB.Close()
	router := intiRouter()
	router.Run(":8000")
}