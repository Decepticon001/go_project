package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main()  {
	f,err := os.Open("/Users/pengzhishen/go/src/gin_blog/log/1.log")
	if err != nil{
		panic(fmt.Sprintf("err %s",err.Error()))
	}
	//a, _ := ioutil.ReadAll(f)
	//fmt.Printf("%s",a)
	//b1 := make([]byte,10)
	//for{
	//	_ ,err := f.Read(b1)
	//	if err == io.EOF{
	//		break
	//	}
	//	fmt.Printf("%s",string(b1))
	//
	//}
	f.Seek(0,2)
	r1 := bufio.NewReader(f)
	for{
		n ,err := r1.ReadBytes('\n')
		if err == io.EOF{
			time.Sleep(500*time.Millisecond)
		}
		fmt.Printf("%s",string(n))
	}

	//for{
	//	b4 , err := r1.Peek(5)
	//	fmt.Printf("%s\n",string(b4))
	//	if err == io.EOF{
	//		break
	//	}
	//}

	f.Close()
}
