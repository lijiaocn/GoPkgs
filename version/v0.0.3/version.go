//create: 2015/05/14 15:55:20 Change: 2019/05/05 18:54:03 author:lijiao
package v2

import (
	"fmt"
)

var (
	VERSION string
	COMPILE string
)

func Show() {
	fmt.Printf("version: %s   compile at: %s( just for test 0.0.3)\n", VERSION, COMPILE)
}
