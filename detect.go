package main

import (
	"bytes"
	"github.com/apache/dubbo-go/protocol/dubbo"
	"net"
)

/*
// Request ...
type Request struct {
	addr   string
	svcUrl common.URL
	method string
	args   interface{}
	atta   map[string]string
}

// NewRequest ...
func NewRequest(addr string, svcUrl common.URL, method string, args interface{}, atta map[string]string) *Request {
	return &Request{
		addr:   addr,
		svcUrl: svcUrl,
		method: method,
		args:   args,
		atta:   atta,
	}
}

// Response ...
type Response struct {
	reply interface{}
	atta  map[string]string
}

// NewResponse ...
func NewResponse(reply interface{}, atta map[string]string) *Response {
	return &Response{
		reply: reply,
		atta:  atta,
	}
}

// DubboPackage ...
type DubboPackage struct {
	Header  hessian.DubboHeader
	Service hessian.Service
	Body    interface{}
	Err     error
}
*/

func listen() {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	for {

		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		var buf []byte
		_, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}
		bytesBuffer := bytes.NewBuffer(buf)

		var p dubbo.DubboPackage
		err = p.Unmarshal(bytesBuffer)
		if err != nil {
			panic(err)
		}
		println(p.String())

		req := p.Body.([]interface{})
		// req := &dubbo.Request{}

		/*
			req[0] = dubboVersion
			req[1] = target
			req[2] = serviceVersion
			req[3] = method
			req[4] = argsTypes
			req[5] = args
		*/
		println(req[3], req[5])

	}
}

func main() {
	listen()
}
