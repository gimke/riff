package client

import (
	"github.com/gimke/riff/api"
	"fmt"
	"qiniupkg.com/x/errors.v7"
	"math/rand"
	"strconv"
	"strings"
)

type RpcClient struct {
	rpc string
}

func (this *RpcClient) Services(name string,state api.StateType) (service api.Service) {
	client, err := api.NewClient(this.rpc)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	err = client.Call("Query.Service", api.ParamService{Name: name, State: state}, &service)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

/*
robin
url: http://serviceName or rpc://serviceName
http url return http://ip:port
rpc url only return ip:port
 */
func (this *RpcClient) Robin(url string) (string,error) {
	prefix := "http://"
	serviceName := ""
	urls := strings.SplitN(url,"//",2)
	prefix = urls[0]+"//"
	serviceName = urls[1]
	if prefix == "rpc://" {
		prefix = ""
	}
	service := this.Services(serviceName,api.StateAlive)
	count:=len(service.NestNodes)
	if count > 0 {
		r := GenerateNumber(0,count-1)
		return prefix+service.NestNodes[r].IP+":"+strconv.Itoa(service.NestNodes[r].Port),nil
	}
	return "",errors.New("404")
}


func GenerateNumber(min, max int) int {
	i := rand.Intn(max-min) + min
	return i
}