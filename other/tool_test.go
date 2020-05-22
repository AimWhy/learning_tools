package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

//获取本地的IP
func Test_GetIp(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(addrs)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

//设置配置文件路径
func Test_Path(t *testing.T) {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	fmt.Println(appConfigPath)
}

func Test_Defer(t *testing.T) {
	defer fmt.Println(1)
	A()
	defer fmt.Println(3)
	return
}
func A() {
	defer func() {
		fmt.Println(2)
	}()
}

func Test_lottery(t *testing.T) {
	var lottery = make(map[string]int)
	lottery["特等奖"] = 5
	lottery["一等奖"] = 10
	lottery["二等奖"] = 35
	lottery["三等奖"] = 50
	//计算概率
	rand.Seed(time.Now().Unix())
	var (
		randNum int
	)
	for _, v := range lottery {
		randNum += v
	}
	fmt.Println("从 ", randNum, "中产生随机数")
	for j := 0; j < 20; j++ {
		i := rand.Intn(randNum)
		var (
			start int
			end   int
		)
		for k, v := range lottery {
			end += v
			if start <= i && i < end {
				fmt.Println("恭喜你中了 ", k)
			}
			start = end
		}
	}
}
