package main

import (
	"context"
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/protocol/dubbo"
	"github.com/apache/dubbo-go/protocol/invocation"
)

func invoke(url common.URL, p dubbo.DubboPackage) {
	proto := dubbo.GetProtocol()
	// clientConf = &ClientConfig{}
	invoker := proto.Refer(url)

	// make sure url equal
	eq := invoker.GetUrl().URLEqual(url)
	println("Url equal:", eq)

	req := p.Body.([]interface{})
	/*
		req[0] = dubboVersion
		req[1] = target
		req[2] = serviceVersion
		req[3] = method
		req[4] = argsTypes
		req[5] = args ([]interface{})
		req[6] = attachments (map[interface{}]interface{})
	*/
	inv := invocation.NewRPCInvocation(req[3].(string), req[5].([]interface{}), req[6].(map[string]string))

	res := invoker.Invoke(context.Background(), inv)
	println(res.Result().(string))

}
