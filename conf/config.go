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
	Name             string //当前服务器名称
	Version          string //当前Zinx版本号
	TcpPort          int    //当前服务器主机监听端口号
	Host             string //当前服务器主机IP
	MaxConn          int    //当前服务器主机允许的最大链接个数
	MaxPacketSize    uint32 //都需数据包的最大值
	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
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
		MaxWorkerTaskLen: 1024,
		ConfFilePath:     "../conf/zinx.json",
	}
	ConfigInstance.Reload()
}
