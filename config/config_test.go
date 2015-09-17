//create: 2015/09/16 14:55:07 change: 2015/09/16 17:35:35 author:lijiao
package config

import(
	"testing"
	"os"
)

type testConfig struct{
	MaxConnection int `json:"maxconnection"`
	MinConnection int `json:"minconnection"`
	TargetUrl string `json:"targeturl"`
}

var config testConfig

func TestConfig(t *testing.T){
	_, err := GenConfig(config, "./config.json")
	if err != nil{
		t.Error(err)
	}

	err = LoadConfig("./config.json",&config)
	if err != nil{
		t.Error(err)
	}
}

func TestMain(m *testing.M){
	config = testConfig{MaxConnection: 10, MinConnection: 5, TargetUrl: "znr.io"}
	os.Exit(m.Run())
}
