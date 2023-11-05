package commom

import (
	"net/rpc"
)

func NewRpcClient(port string) (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", "localhost:"+port)
	if err != nil {
		return nil, err
	}

	return client, nil
}
