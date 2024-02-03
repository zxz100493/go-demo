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
	Key    string
)

func init() {
	flag.StringVar(&Listen, "l", "", "listener address")
	flag.StringVar(&Key, "k", "", "key")
}

func main() {
	flag.Parse()

	if Listen == "" {
		fmt.Println("缺少监听地址")
		os.Exit(0)
	}

	if Key == "" {
		fmt.Println("缺少key值")
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

		newStr, err := common.DePwdCode(recvStr, Key)
		if err != nil {
			fmt.Println("DePwdCode:", err)
			break
		}

		fmt.Println("收到client端发来的数据：", string(newStr))

		_, err = conn.Write(newStr) // 发送数据
		if err != nil {
			fmt.Println("write to client failed, err:", err)
			break
		}
	}
}
