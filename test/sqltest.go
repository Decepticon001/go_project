package test

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

func main(){
	db,err := sql.Open("mysql","root:spz1234@tcp(127.0.0.1:3306)/info?parseTime=true")
	if err != nil{
		log.Fatal(err.Error())
	}
	//var yuan,zhh string
	rows,err := db.Query("SELECT id,yuanzi,zhuanhuanhou FROM hanzi")
	for rows.Next(){
		var id,yz,zh string
		if err != nil{
			fmt.Println(err)
		}
		err := rows.Scan(&id,&yz,&zh);
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(id)
		fmt.Println(yz)
		fmt.Println(zh)
	}
	var arr = [4]int{2,1}
	fmt.Println(arr)
}
