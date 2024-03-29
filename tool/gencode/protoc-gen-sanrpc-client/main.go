package main

import (
	"fmt"

	"github.com/hillguo/sanrpc/tool/gencode"
)

var formatPart1 = `package %s

import (
	"context"
	"github.com/hillguo/sanrpc/client"
)

type %sClient struct {
	client.XClient
}

func New%sClient() *%sClient {
	client := %sClient{}
	return &client
}
`

var formatFunc = `
func (c *%sClient) %s(ctx *context.Context, req *%s, resp *%s) error {
	return c.Call(ctx, "%s", "%s", req, resp)
}
`

func genClient(protoInfo *gencode.ProtoFileInfo) (string, string) {
	data := fmt.Sprintf(formatPart1, protoInfo.PackageName, protoInfo.ServiceName, protoInfo.ServiceName,
		protoInfo.ServiceName, protoInfo.ServiceName)

	for _, methodInfo := range protoInfo.Methods {
		data += fmt.Sprintf(formatFunc, protoInfo.ServiceName, methodInfo.MethodName,
			methodInfo.InputType, methodInfo.OutputType, protoInfo.ServiceName, methodInfo.MethodName)
	}
	return protoInfo.ModuleName + "client.go", data
}

func main() {
	gencode.Main(genClient)
}
