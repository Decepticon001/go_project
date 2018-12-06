package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type LogRead struct {
	rc chan []byte
	wc chan string
	reader Reader
	writer Writer
}

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type ReadFromFile struct {
	path string
}
type WriteToFile struct {
	txtpath string
}
func (w *WriteToFile) Write(wc chan string){
	for v := range wc{
		fmt.Println(v)
	}

}
func (r *ReadFromFile) Read(rc chan []byte){
	f,err := os.Open(r.path)
	if err != nil{
		panic(fmt.Sprintf("open file err %s",err.Error()))
	}
	//fmt.Println(f)
	f.Seek(0,2)
	rd := bufio.NewReader(f)
	fmt.Println(rd)
	for {
		line ,err:=rd.ReadBytes('\n')
		if err == io.EOF{
			time.Sleep(500*time.Microsecond)
		}else if err != nil {
			panic(fmt.Sprintf("err %s",err.Error()))
		}
		rc <- line
	}

}

func (l *LogRead) Process(){
	//
	l.wc <- strings.ToUpper(string(<- l.wc))
	//for v := range l.rc{
	//
	//}


}

func main(){
	r := &ReadFromFile{
		path : "/Users/pengzhishen/go/src/gin_blog/log/2.txt",
	}
	w := &WriteToFile{
		txtpath:"./2.txt",
	}
	lp := LogRead{
		rc : make(chan []byte),
		wc : make(chan string),
		reader:r,
		writer:w,

	}
	go lp.reader.Read(lp.rc)
	go lp.Process()
	go lp.writer.Write(lp.wc)
	time.Sleep(40*time.Second)
}