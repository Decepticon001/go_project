package models

import (
	db "gin_blog/database"
)

type Arts struct {
	Id int
	Yb string
	Zh string
}

func (a *Arts) AddCha() (id int64,err error){
	rs,err := db.SqlDB.Exec("INSERT  INTO hanzi(yuanzi,zhuanhuanhou) VALUES(?,?)",a.Yb,a.Zh)
	if err != nil{
		return
	}
	id,err = rs.LastInsertId()
	return
}

func (a *Arts) GetCha()(art []Arts,err error){
	art = make([]Arts,0)
	rows,err := db.SqlDB.Query("SELECT id,yuanzi,zhuanhuanhou FROM hanzi")
	defer rows.Close()
	if err != nil{
		return
	}
	for rows.Next(){
		var ar Arts
		rows.Scan(&ar.Id,&ar.Yb,&ar.Zh)
		art = append(art,ar)
	}
	//fmt.Print(art)
	if err = rows.Err(); err != nil {
		return
	}
	return
}
