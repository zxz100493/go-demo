package main

import (
	"bufio"
	"demo/common"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	Listen string
	Key1   string
	Key2   string
	Target string
)

func init() {
	flag.StringVar(&Listen, "l", "", "listener address")
	flag.StringVar(&Key1, "k1", "", "key1")
	flag.StringVar(&Key2, "k2", "", "key2")
	flag.StringVar(&Target, "t", "", "target address")
}

func main() {
	flag.Parse()

	if Listen == "" {
		fmt.Println("缺少监听地址")
		os.Exit(0)
	}

	if Key1 == "" {
		fmt.Println("缺少key1值")
		os.Exit(0)
	}

	if Key2 == "" {
		fmt.Println("缺少key2值")
		os.Exit(0)
	}

	if Target == "" {
		fmt.Println("缺少目标地址")
		os.Exit(0)
	}

	listen, err := net.Listen("tcp", Listen)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go process(conn) // 启动一个goroutine处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte

		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)

		newStr, err := common.DePwdCode(recvStr, Key1)
		if err != nil {
			fmt.Println("DePwdCode:", err)
			break
		}

		_, err = conn.Write(newStr) // 发送数据
		if err != nil {
			fmt.Println("write to client failed, err:", err)
			break
		}

		sendToServer(string(newStr))
	}
}

func sendToServer(msg string) {
	conn, err := net.Dial("tcp", Target)
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	defer conn.Close() // 关闭连接

	newStr, err := common.EnPwdCode([]byte(msg), Key2)

	if err != nil {
		fmt.Println("EnPwdCode err:", err)
	}

	_, err = conn.Write([]byte(newStr)) // 发送数据
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
