package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

// 测试连接
func pwdConnect(sshHost, sshUser, sshPassword string, sshPort int) (*ssh.Client, error) {
	// 创建ssh登录配置
	config := &ssh.ClientConfig{
		Timeout:         5 * time.Second,                             // 超时时间
		User:            sshUser,                                     // 登录账号
		Auth:            []ssh.AuthMethod{ssh.Password(sshPassword)}, // 密码
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),                 // 这个不够安全，生产环境不建议使用
		//HostKeyCallback: ssh.FixedHostKey(), // 建议使用这种，目前还没研究出怎么使用[todo]
	}

	// dial连接服务器
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	Client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("连接到服务器失败", err)
		return nil, err
	}

	return Client, nil
}

func main() {

}

func shell() {
	// 连接到服务器
	conn, err := pwdConnect("10.10.10.10", "root", "root", 22)
	if err != nil {
		return
	}
	defer func(conn *ssh.Client) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	// 创建 ssh session 会话
	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}

	fmt.Println(session.Shell())
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(session)

	// 执行远程命令
	cmd := "cd /home;ls -la;"
	cmdInfo, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cmdInfo))
}
