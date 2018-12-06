package test

import (
	"fmt"
	"time"

	//"math"
	"math/big"
)

func main()  {
	t1 := time.Now().Unix()
	for i := 0;i<100;i++{
		im := big.NewInt(1)
		//fmt.Println(im)
		for r:=1;r<101;r++{
			im1 := big.NewInt(int64(r))
			im = im.Mul(im,im1)
		}
		fmt.Println(im)
	}
	t2 := time.Now().Unix()
	fmt.Println(t2-t1)
}
