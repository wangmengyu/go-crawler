package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
  封装RPC服务端
  注册服务，不断的监听TCP端口，如果有数据就执行运算
*/
func ServeRpc(host string, service interface{}) error {
	//注册rpc服务
	_ = rpc.Register(service)

	//监听tcp 1234端口
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	log.Printf("listen on port:%s", host)

	//一直从tcp接收信息，
	for {
		conn, err := listener.Accept()
		if err != nil {
			//发生错误，接收下一个请求
			log.Printf("accept err:%v", err)
			continue
		}
		//接收成功，异步的完成服务
		go jsonrpc.ServeConn(conn)
	}
}

/**
  创建客户端进行连接
*/
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(conn)
	return client, nil
}
