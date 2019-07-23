package tlv

import (
	"encoding/binary"
	"io"
	"net"
)

type Transport struct {
	conn net.Conn
}

func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn}
}

func (t *Transport) WriteRequest(req interface{}) error {
	b, err := encode(req)
	if err != nil {
		return err
	}

	buf := make([]byte, 4+len(b))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(b)))
	copy(buf[4:], b)
	_, err = t.conn.Write(buf)
	return err
}

func (t *Transport) ReadRequest() (req interface{}, err error) {
	header := make([]byte, 4)
	_, err = io.ReadFull(t.conn, header)
	if err != nil {
		return req, err
	}

	data := make([]byte,
		binary.BigEndian.Uint32(header))

	_, err = io.ReadFull(t.conn, data)
	if err != nil {
		return req, err
	}

	rsp, err := decode(data)
	return rsp, err
}
