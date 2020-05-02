package porsist

import (
	"fmt"
	"log"
)

type Profile struct {
	Id int
	Title string
	Dynasty string
	Author string
	Content string
}

func ItemSaver() chan Profile{
	out:=make(chan Profile)
	go func() {
		itemCount := 0
		for {
			item:=<-out
			//Insert(item)
			//byte, _ := json.Marshal(<-out)
			//fmt.Println("插入状态：",)
			//fmt.Println(string(byte))
			log.Printf("Item Saver: got item #%d: %v",itemCount,item)
			itemCount ++
		}
	}()
	return out
}

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func Insert(p Profile)bool  {
	stmeInsert, err := Db.Prepare("insert into poem (title,author,dynasty,content) values (?,?,?,?)")
	if checkError (err) {
		fmt.Println("Prepare",err)
		return false
	}
	_, err = stmeInsert.Exec(&p.Title, &p.Author, &p.Dynasty, &p.Content)
	if checkError (err) {
		fmt.Println("Exec",err)
		return false
	}
	return true
}
