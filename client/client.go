package client

import (
	"strings"
	"github.com/gimke/riff/api"
	"errors"
)

type Client interface {
	Services(serviceName string,state api.StateType) api.Service
	Robin(url string) (string,error)
}

func NewClient(url string) (Client,error) {
	if strings.Index(url,"riff://") == 0 {
		rpc := strings.Replace(url,"riff://","",1)
		return &RpcClient{rpc},nil
	} else {
		return nil,errors.New("not support")
	}
}
