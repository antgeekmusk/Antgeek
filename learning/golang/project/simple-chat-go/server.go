package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 服务端结构体
type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播channel
	Message chan string
}

// 创建服务端
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + " : " + msg
	this.Message <- sendMsg
}

// 监听Message
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		// 将消息发送所有User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 处理链接
func (this *Server) Handler(conn net.Conn) {
	//fmt.Println("链接建立成功")

	user := NewUser(conn, this)

	// 用户上线
	user.Online()

	isLive := make(chan bool)

	// 处理用户发送的消息,每来一个链接(用户)就单独启动一个goroutine来处理用户消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn err :", err)
				return
			}

			// 提取用户消息(去掉回车)
			msg := string(buf[:n-1])

			// 将消息广播
			user.DoMessage(msg)

			// 发消息就代表活跃
			isLive <- true
		}
	}()

	// 超时强踢功能
	go func() {
		for {
			select {
			case <-isLive:
			case <-time.After(time.Second * 300):
				// 超时
				user.SendMsg("已超时.....")
				// 释放资源
				close(user.C)
				conn.Close()

				return
			}
		}
	}()
}

// 启动服务器
func (this *Server) Start() {
	// Socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err :", err)
	}

	// close listen
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessage()

	// accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err", err)
			continue
		}

		this.Handler(conn)
	}
}
