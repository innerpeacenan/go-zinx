package conf

import (
	"encoding/json"
	"go-zinx/ziface"
	"io/ioutil"
	"os"
)

var (
	ConfigInstance *Config
)

type Config struct {
	TcpServer     ziface.IServer //当前Zinx的全局Server对象
	Host          string         //当前服务器主机IP
	TcpPort       int            //当前服务器主机监听端口号
	Name          string         //当前服务器名称
	Version       string         //当前Zinx版本号
	MaxPacketSize uint32         //都需数据包的最大值
	MaxConn       int            //当前服务器主机允许的最大链接个数
	ConfFilePath  string
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
	// default
	ConfigInstance = &Config{
		Name:          "ZinxServerApp",
		Version:       "V0.5",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
		ConfFilePath:  "../conf/zinx.json",
	}
	ConfigInstance.Reload()
}
