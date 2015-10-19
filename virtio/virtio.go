//create: 2015/09/18 11:12:13 change: 2015/09/18 13:51:03 author:lijiao
package virtio

import(
	"io"
	"sync"
	"time"
)

type VirtReader struct{
	mu       sync.Mutex
	Index    int
	Capaticy int
	Step     int
	Delay    time.Duration
}

func (r *VirtReader) Read(p []byte)(n int, err error){
	time.Sleep(r.Delay)
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.Index < r. Capaticy - r.Step {
		r.Index = r.Index + r.Step
		return r.Step, nil
	}else if r.Index >= r.Capaticy{
		return 0, io.EOF
	}
	n = r.Capaticy - r.Index
	r.Index = r.Capaticy
	return  n, io.EOF
}

