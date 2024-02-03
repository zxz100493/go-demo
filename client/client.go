package main

import (
	"bufio"
	"demo/common"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	Target    string
	ClientMsg string
	Key       string
)

func init() {
	flag.StringVar(&Target, "t", "", "target address")
	flag.StringVar(&ClientMsg, "c", "", "msg")
	flag.StringVar(&Key, "k", "", "key")
}

func main() {
	flag.Parse()

	if Target == "" {
		fmt.Println("缺少目标地址")
		os.Exit(0)
	}

	if ClientMsg == "" {
		fmt.Println("缺少信息")
		os.Exit(0)
	}

	if Key == "" {
		fmt.Println("缺少key信息")
		os.Exit(0)
	}

	conn, err := net.Dial("tcp", Target)
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	defer conn.Close() // 关闭连接

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")

		if strings.EqualFold(inputInfo, "Q") { // 如果输入q就退出
			return
		}

		newInfo, err := common.EnPwdCode([]byte(inputInfo), Key)

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write([]byte(newInfo)) // 发送数据
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
