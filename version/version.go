//create: 2015/05/14 15:55:20 Change: 2019/06/04 18:54:10 author:lijiao
package version

import (
	"fmt"
)

var (
	VERSION string
	COMPILE string
)

func Show() {
	fmt.Printf("version: %s   compile at: %s  golib v2\n", VERSION, COMPILE)
}
