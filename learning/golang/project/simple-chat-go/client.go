package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	Conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       -1,
	}

	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err :", err)
		return nil
	}

	client.Conn = conn

	return client
}

func (this *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>>> 请输入合法数字0-3")
		return false
	}

}

func (this *Client) UpdateName() bool {
	fmt.Println(">>>>>>>> 请输入用户名")
	fmt.Scanln(&this.Name)

	sendMsg := "rename|" + this.Name + "\n"
	_, err := this.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("this.Conn.Write err :", err)
		return false
	}
	return true
}

func (this *Client) PublicChat() {
	var sendMsg string
	fmt.Println(">>>>>>>> 请输入聊天内容,exit退出")
	fmt.Scanln(&sendMsg)
	for sendMsg != "exit" {
		if len(sendMsg) > 0 {
			sendMsg = sendMsg + "\n"
			_, err := this.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("PublicChat err :", err)
				break
			}
		}
		sendMsg = ""
		fmt.Scanln(&sendMsg)
	}
}

// 查询在线用户
func (this *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := this.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("SelectUsers err : ", err)
		return
	}
}

// 私聊
func (this *Client) PrivateChat() {
	var remoteName string
	var sendMsg string

	// 查询在线用户
	this.SelectUsers()

	// 选择要私聊的用户
	fmt.Println(">>>>>>>> 选择要私聊的用户,exit 退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>>> 请输入消息内容,exit 退出")
		fmt.Scanln(&sendMsg)

		for sendMsg != "exit" {
			sendMsg = "to|" + remoteName + "|" + sendMsg + "\n\n"
			_, err := this.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("PrivateChat err :", err)
				break
			}

			sendMsg = ""
			fmt.Scanln(&sendMsg)
		}

		fmt.Println(">>>>>>>> 选择要私聊的用户,exit 退出")
		this.SelectUsers()
		fmt.Scanln(&remoteName)

	}

}

func (this *Client) DealResponse() {
	// 将返回的消息回显到标准输出
	io.Copy(os.Stdout, this.Conn)
}

func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() != true {
		}
		switch this.flag {
		case 1:
			this.PublicChat()
			break
		case 2:
			this.PrivateChat()
			break
		case 3:
			this.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器ip,默认127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "服务器端口,默认8888")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>客户端链接服务端失败")
		return
	}
	// 单独开启一个goroutine回显服务器数据
	go client.DealResponse()
	fmt.Println(">>>>>>>>客户端链接服务端成功")

	// 客户端的一些业务
	client.Run()

}
