//create: 2015/05/14 15:55:20 Change: 2019/05/05 18:06:58 author:lijiao
package version

import (
	"fmt"
)

var (
	VERSION string
	COMPILE string
)

func Show() {
	fmt.Printf("version: %s   compile at: %s\n", VERSION, COMPILE)
}

func Show2() {
	fmt.Printf("version: %s   compile at: %s\n ( just for test)", VERSION, COMPILE)
}
