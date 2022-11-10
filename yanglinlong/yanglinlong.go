package yanglinlong

import (
	"fmt"
	"moduledemo/mypackage"
)

func YangHello(a int) int {
	mypackage.New()
	fmt.Println("杨林龙")
	return a
}
