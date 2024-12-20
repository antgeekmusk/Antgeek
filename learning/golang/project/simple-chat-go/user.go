package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建User的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动一个监听C消息的goroutine,相当于用户接受消息的一个goroutine
	go user.ListenCMsg()

	return user
}

// 上线
func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "已上线")
}

// 下线
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "已下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理消息
func (this *User) DoMessage(msg string) {
	if "who" == msg { // 用户查询在线用户功能
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + " : " + "在线\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // 改名功能
		// 消息格式 rename|张三
		newName := strings.Split(msg, "|")[1]

		// 判断用户是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("用户名已存在")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName

			this.SendMsg("已改名 : " + this.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" { // 私聊功能
		// 消息格式 to|张三|msg
		// 获取私聊对象名称
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确,格式 : to|张三|msg")
			return
		}

		// 获取user
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("用户 : " + remoteName + " 不在线")
			return
		}

		// 私聊发消息
		content := strings.Split(msg, "|")[2]
		remoteUser.SendMsg(this.Name + " 对您说 : " + content + "\n")

	} else {
		this.server.BroadCast(this, msg)
	}

}

// 监听C的消息
func (this *User) ListenCMsg() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
