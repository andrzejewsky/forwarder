package connection

import (
	"io"
	"net"
)

// Forwarder forwarding tcp connection
type Forwarder struct {
	listener        net.Listener
	destinationAddr string
}

// NewForwarder new instance
func NewForwarder(listener net.Listener, destinationAddr string) *Forwarder {
	return &Forwarder{listener, destinationAddr}
}

// StartForwarding start frorwarding
func (f *Forwarder) StartForwarding() {
	for {
		client, _ := f.listener.Accept()

		go f.forward(client)
	}
}

func (f *Forwarder) forward(client net.Conn) {

	destConn, _ := net.Dial("tcp", f.destinationAddr)

	go func() {
		defer client.Close()
		defer destConn.Close()
		io.Copy(destConn, client)
	}()
	go func() {
		defer client.Close()
		defer destConn.Close()
		io.Copy(client, destConn)
	}()
}
