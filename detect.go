package main

import (
	"bytes"
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/common/constant"
	"github.com/apache/dubbo-go/protocol/dubbo"
	"github.com/apache/dubbo-go/registry/zookeeper"
	"net"
	"strconv"
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

/*
type baseUrl struct {
	Protocol string
	Location string // ip+port
	Ip       string
	Port     string
	//url.Values is not safe map, add to avoid concurrent map read and map write error
	paramsLock   sync.RWMutex
	params       url.Values
	PrimitiveURL string
}

// URL ...
type URL struct {
	baseUrl
	Path     string // like  /com.ikurento.dubbo.UserProvider3
	Username string
	Password string
	Methods  []string
	//special for registry
	SubURL *URL
}

// URL sample:
zookeeper://127.0.0.1:2181/org.apache.dubbo.registry.RegistryService?
application=demo-consumer&dubbo=2.0.2&interface=org.apache.dubbo.registry.RegistryService
&pid=1214&qos.port=33333&timestamp=1545721981946
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

		// req := p.Body.([]interface{})

		// req := &dubbo.Request{}

		/*
			req[0] = dubboVersion
			req[1] = target
			req[2] = serviceVersion
			req[3] = method
			req[4] = argsTypes
			req[5] = args ([]interface{})
			req[6] = attachments (map[interface{}]interface{})
		*/

		url := subscribe(p)
		invoke(url, p)
	}
}

/*
func assemble(p dubbo.DubboPackage) common.URL{
	myUrl := "zookeeper://152.136.97.145:2181/"
	myUrl += p.Service.Path
	url,err := common.NewURL(myUrl)
	if err != nil {
		panic(err)
	}

	return url
}
*/

func subscribe(p dubbo.DubboPackage) common.URL {
	regUrl, err := common.NewURL("zookeeper://152.136.97.145:2181", common.WithParamsValue(constant.ROLE_KEY, strconv.Itoa(common.CONSUMER)))
	if err != nil {
		panic(err)
	}

	urlStr := "dubbo://127.0.0.1:20000/"
	urlStr += p.Service.Path
	url, err := common.NewURL(urlStr)

	reg, err := zookeeper.NewZkRegistry(&regUrl)

	listener, err := reg.DoSubscribe(&url)
	if err != nil {
		panic(err)
	}
	serviceEvent, err := listener.Next()
	if err != nil {
		panic(err)
	}
	println(serviceEvent)

	return url
}

func main() {
	listen()
}
