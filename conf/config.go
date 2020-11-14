package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	ConfigInstance *Config
)

type Config struct {
	Name             string
	Version          string
	TcpPort          int
	Host             string
	MaxConn          int
	MaxPacketSize    uint32
	WorkerPoolSize   uint32
	MaxMsgChanLen    uint32
	MaxWorkerTaskLen uint32
	ConfFilePath     string
}

func (g *Config) Reload() {
	if confFileExists, _ := PathExists(g.ConfFilePath); confFileExists != true {
		return
	}

	data, err := ioutil.ReadFile(g.ConfFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &ConfigInstance)
	if err != nil {
		panic(err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	ConfigInstance = &Config{
		Name:             "ZinxServerApp",
		Version:          "V0.5",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxMsgChanLen:    100,
		MaxWorkerTaskLen: 1024,
		ConfFilePath:     "../conf/zinx.json",
	}
	ConfigInstance.Reload()
}
