//create: 2015/07/16 10:29:14 change: 2015/09/17 15:52:16 author:lijiao
package container

import(
	"testing"
	"log"
	"os"
	"time"
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
)

var (
	err    error
	client *docker.Client
)
func TestDefaultConConfig(t *testing.T){
	handler := DefaultHandler()
	handler.SetClient(client)
	handler.SetImage("127.0.0.1:5000/redis-port", "127.0.0.1:5000", "latest")
	handler.SetName("redis-port")
	handler.SetAuth("", "", "", "")
	handler.SetMemLimit(0,-1)
	handler.SetNetworkMode("host")
	handler.SetCmds("--port", "7000")

	err = handler.Create()
	if err != nil{
		t.Fatal(err)
	}

	t.Log(handler.Container)
	
	err = handler.Start()
	if err != nil{
		t.Error(err)
	}

	conlist,err := handler.ListAll()
	if err != nil{
		t.Error(err)
	}
	for _,v := range conlist{
		log.Println(v)
	}

	names,err := handler.ListAllName()
	if err != nil{
		t.Error(err)
	}
	log.Println(names)

	fmt.Printf("sleep 10 scond\n")
	time.Sleep(10 *time.Second)
	err = handler.Remove()
	if err != nil{
		t.Log(err)
	}
}

func TestMain(m *testing.M){
	client,err = docker.NewClient("tcp://127.0.0.1:2375")
	if err != nil{
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}
