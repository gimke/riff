package riff

import (
	"github.com/gimke/riff/api"
	"gopkg.in/yaml.v2"
	"net"
	"strconv"
)

type Resolver struct {
}

//Query
func (_ *Resolver) Riff() *RiffResolver {
	return &RiffResolver{}
}

func (_ *Resolver) Nodes() *[]*NodeResolver {
	var l []*NodeResolver
	for _, node := range server.api.Nodes() {
		l = append(l, &NodeResolver{node})
	}
	return &l
}

func (_ *Resolver) Node(args struct{ Name string }) *NodeResolver {
	node := server.api.Node(args.Name)
	if node == nil {
		return nil
	} else {
		return &NodeResolver{node}
	}
}

func (_ *Resolver) Server() *NodeResolver {
	return &NodeResolver{server.api.Node(server.Self.Name)}
}

func (_ *Resolver) Services() *[]*ServiceResolver {
	var l []*ServiceResolver
	for _, service := range server.api.Services() {
		l = append(l, &ServiceResolver{service})
	}
	return &l
}

func (_ *Resolver) Service(args struct {
	Name  string
	State *string
}) *ServiceResolver {
	state := api.StateAll
	if args.State != nil {
		state = api.StateType_FromName(*args.State)
	}
	service := server.api.Service(args.Name, state)
	if service == nil {
		return nil
	} else {
		return &ServiceResolver{service}
	}
}

func (_ *Resolver) MutationService(args struct {
	Services *[]*MutationServiceInput
}) *[]*MutationService {
	var l []*MutationService
	for _, service := range *args.Services {
		result := &MutationService{
			cmd:  service.Cmd,
			ip:   service.Ip,
			port: int(service.Port),
		}
		if err := mutationService(service.Name, net.JoinHostPort(service.Ip, strconv.Itoa(int(service.Port))), api.CmdType_FromName(service.Cmd)); err != nil {
			result.error = err.Error()
			result.success = false
		} else {
			result.error = ""
			result.success = true
		}
		l = append(l, result)
	}
	return &l
}

func (_ *Resolver) RegisteService(args struct {
	Service struct {
		Name string
		Ip   string
		Port int32
	}
}) *bool {
	service := &Service{
		ServiceConfig: &ServiceConfig{
			Name: args.Service.Name,
			Ip:   args.Service.Ip,
			Port: int(args.Service.Port),
		},
	}
	config, _ := yaml.Marshal(service.ServiceConfig)
	service.Config = string(config)
	server.AddService(service)
	result := true
	return &result
}

func (_ *Resolver) UnregisteService(args struct {
	Name string
}) *bool {
	delete(server.Self.Services, args.Name)
	result := true
	return &result
}
