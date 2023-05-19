package plugin

import (
	"encoding/gob"
	"fmt"
	"github.com/BoyChai/CoralBot/task"
	"net"
)

// WriteInfo 传输
func WriteInfo(data task.Info, network net.Conn) {
	encoder := gob.NewEncoder(network)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("插件gob编码错误,err:", err)
	}
}

// ReadInfo 接收
func ReadInfo(network net.Conn) task.Info {
	var data task.Info
	decoder := gob.NewDecoder(network)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("插件读取错误,gob解码错误,err:", err)
	}
	return data
}

// WriteData 传输
func WriteData(data []byte, network net.Conn) {
	encoder := gob.NewEncoder(network)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("插件gob编码错误,err:", err)
	}
}

// ReadData 接收
func ReadData(network net.Conn) []byte {
	var data []byte
	decoder := gob.NewDecoder(network)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("插件读取错误,gob解码错误,err:", err)
	}
	return data
}
