//create: 2015/09/17 16:25:33 change: 2015/09/17 16:26:52 author:lijiao
package generator

import(
	"testing"
)

func TestStrGenerator(t *testing.T){
	
	strs := make([]string, 0, 0)
	strs = append(strs, "a", "b", "c", "d", "e", "f")
	gen := StrGenerator(strs)

	for i:=0; i <15; i++{
		println(gen())
	}

	strs[3]="dddddd"
	for i:=0; i <15; i++{
		println(gen())
	}
}


