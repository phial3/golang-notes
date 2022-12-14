package socket_test

import (
	"testing"

	"github.com/phial3/go-notes/network/socket/client"
	"github.com/phial3/go-notes/network/socket/server"
)

func TestTCPServer_SimpleMessage(t *testing.T) {
	server.TCPServer_SimpleMessage()
}

func TestTCPClient_SimpleMessage(t *testing.T) {
	client.TCPClient_SimpleMessage()
}

func TestTCPServer_StructMessage(t *testing.T) {
	server.TCPServer_StructMessage()
}

func TestTCPServer_MoreStructMessage(t *testing.T) {
	server.TCPServer_MoreStructMessage()
}

func TestTCPClient_StructMessage(t *testing.T) {
	client.TCPClient_StructMessage()
}

func TestTCPServer_MoreLongStructMessage(t *testing.T) {
	server.TCPServer_MoreLongStructMessage()
}

func TestTCPClient_LongStructMessage(t *testing.T) {
	client.TCPClient_LongStructMessage()
}

func TestUDPServer_MoreLongStructMessage(t *testing.T) {
	server.UDPServer_MoreLongStructMessage()
}

func TestUDPClient_LongStructMessage(t *testing.T) {
	client.UDPClient_MoreLongStructMessage()
}

func TestWServer_MoreLongStructMessage(t *testing.T) {
	server.WsServer_MoreLongStructMessage()
}

func TestWSClient_LongStructMessage(t *testing.T) {
	client.WSClient_MoreLongStructMessage()
}
