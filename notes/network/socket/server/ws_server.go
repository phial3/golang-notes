package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/phial3/go-notes/network/socket"
	"github.com/gorilla/websocket"
)

func NewWsServer(port int) *WsServer {
	return &WsServer{
		addr: socket.ListenAddr + ":" + strconv.Itoa(port),
		upgrade: &websocket.Upgrader{
			HandshakeTimeout: 5 * time.Second, // 握手超时时间
			ReadBufferSize:   2048,            // 读缓冲大小
			WriteBufferSize:  1024,            // 写缓冲大小
			// 请求检查函数，用于统一的链接检查，以防止跨站点请求伪造；
			// 如果Origin请求头存在且原始主机不等于请求主机头，则返回false
			CheckOrigin: func(r *http.Request) bool {
				fmt.Printf("Request url %s\n", r.URL)
				fmt.Println("Handshake request header")
				for key, values := range r.Header {
					fmt.Printf("%s:%s\n", key, values[0])
				}
				return true
			},
			// http错误响应函数
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
				httpCode := http.StatusInternalServerError
				rePhrase := http.StatusText(httpCode)
				if status == httpCode {
					fmt.Print("Path error ", rePhrase)
					http.Error(w, rePhrase, httpCode)
				}
			},
		},
	}
}

type WsServer struct {
	addr     string
	listener net.Listener
	upgrade  *websocket.Upgrader
}

func (ws *WsServer) Start() (err error) {
	ws.listener, err = net.Listen("tcp", ws.addr) // 建立tcp监听，http和websocket都是建立在tcp协议之上的
	if err != nil {
		fmt.Printf("Listen error:%s\n", err)
		return
	}

	err = http.Serve(ws.listener, ws) //开始对外提供http服务，可以接收多个连接请求，即使一个连接处理出错，也不会影响其他连接
	if err != nil {
		fmt.Printf("Http server error: %v\n", err)
		return
	}

	// Listen和Serve两步合成一步
	// if err:=http.ListenAndServe(ws.addr, ws);err!=nil{
	// 	fmt.Printf("Http server error: %v\n", err)
	// 	return
	// }

	return nil
}

// http.Handler 必须实现ServeHTTP接口
func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		httpCode := http.StatusInternalServerError // 返回一个内部错误的信息
		rePhrase := http.StatusText(httpCode)      // 组织错误信息
		fmt.Println("Path error ", rePhrase)       // 把错误信息写到ResponseWriter里
		http.Error(w, rePhrase, httpCode)
		return
	}

	// 升级http为websocket
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Upgrade http to websocket error: %v\n", err)
		return
	}
	fmt.Printf("Establish conection to client %s\n", conn.RemoteAddr().String())
	go ws.handleConnection(conn)
}

// handleConnection 处理websocket连接里发来的请求数据
func (ws *WsServer) handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	// 长连接，即连接建立后进行多轮的读写交互
	for {
		conn.SetReadDeadline(time.Now().Add(20 * time.Second)) // 设置读超时时间
		var request socket.Request
		if err := conn.ReadJSON(&request); err != nil {
			// 判断是否超时
			if netError, ok := err.(net.Error); ok {
				if netError.Timeout() {
					fmt.Printf("Read message timeout, remote %s\n", conn.RemoteAddr().String())
					return
				}
			}
			// 忽略websocket.CloseGoingAway/websocket.CloseNormalClosure这2种closeErr，如果是其他closeErr就打一条错误日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("Read message from %s error %v\n", conn.RemoteAddr().String(), err)
			}
			return // 只要ReadMessage发生错误，就关闭这条连接
		} else {
			response := socket.Response{Sum: request.A + request.B}
			if err = conn.WriteJSON(&response); err != nil {
				fmt.Printf("Write response failed: %v", err)
			} else {
				fmt.Printf("Write response %d\n", response.Sum)
			}
		}
	}
}

func WsServer_MoreLongStructMessage() {
	ws := NewWsServer(socket.Port)
	ws.Start()
}
