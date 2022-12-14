package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/phial3/go-notes/historys/week11/socket"
)

// 序列化请求和响应的结构体
func main() {
	// 设置TCP端点的地址
	ip := "127.0.0.1" // ip也可设置成 0.0.0.0 和 空字符串
	port := 5656      // 改成 1023，会报错 bind: permission denied
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(port))
	socket.CheckError(err)

	// 设置监听地址
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	socket.CheckError(err)
	fmt.Println("Waiting for client connection ...")

	// 等待TCP连接
	conn, err := listener.Accept()
	socket.CheckError(err)
	fmt.Printf("Establish connection to client %s\n", conn.RemoteAddr().String()) // 操作系统会随机给客户端分配一个 49152 z~ 65535 上的端口号

	// 获取客户端Request
	requestBytes := make([]byte, 256) // 设定一个最大长度，防止 flood attack；初始化后byte数组每个元素都是0
	read_len, err := conn.Read(requestBytes)
	socket.CheckError(err)
	fmt.Printf("Receive request %s\n", string(requestBytes)) // []byte转string时，0后面的会自动被截掉

	// 返回Response
	var request socket.Request
	json.Unmarshal(requestBytes[:read_len], &request) // json反序列化时会把0都考虑在内，需要指定只读前read_len个字节
	response := socket.Response{Sum: request.A + request.B}
	responseBytes, _ := json.Marshal(response)
	_, err = conn.Write(responseBytes)
	socket.CheckError(err)
	fmt.Printf("Write response %s\n", string(responseBytes))

	// 中断TCP连接
	conn.Close()
}
