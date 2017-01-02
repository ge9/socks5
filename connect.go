package socks5

import (
	"log"
	"net"
)

// Connect remote conn which u want to connect.
// You may should write your method instead of use this method.
func (r *Request) Connect(c net.Conn) (net.Conn, error) {
	if Debug {
		log.Println("Call:", r.Address())
	}
	rc, err := net.Dial("tcp", r.Address())
	if err != nil {
		p := NewReply(RepHostUnreachable, ATYPIPv4, []byte{0x00, 0x00, 0x00, 0x00}, []byte{0x00, 0x00})
		if err := p.WriteTo(c); err != nil {
			return nil, err
		}
		return nil, err
	}

	a, addr, port := ParseAddress(rc.LocalAddr())
	p := NewReply(RepSuccess, a, addr, port)
	if err := p.WriteTo(c); err != nil {
		return nil, err
	}

	return rc, nil
}
