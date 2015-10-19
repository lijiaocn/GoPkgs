//create: 2015/09/18 11:25:45 change: 2015/09/18 13:50:32 author:lijiao
package virtio

import(
	"testing"
	"time"
)

func TestVirtRead(t *testing.T){
	var size int = 100
	vr := VirtReader{Index:0, Capaticy: size, Step: 7, Delay: 1 *time.Second}

	var count int = 0
	b := make([]byte, 0, 0)
	for true{
		n, err := vr.Read(b)
		println(n)
		count = count + n
		if err != nil{
			break
		}
	}

	if size != count{
		t.Errorf("size is %d, but read %d\n %v\n", size, count, vr)
	}
}


