package proxy

import (
	"errors"
	"github.com/hashicorp/consul/api"
	"strconv"
	"unsafe"
)

type consulServiceRegistry struct {
	serviceInstances     map[string]map[string]ServiceInstance
	client               api.Client
	localServiceInstance ServiceInstance
}

func (c consulServiceRegistry) GetInstances(serviceId string) ([]ServiceInstance, error) {
	catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	if len(catalogService) > 0 {
		result := make([]ServiceInstance, len(catalogService))
		for index, sever := range catalogService {
			s := DefaultServiceInstance{
				InstanceId: sever.ServiceID,
				ServiceId:  sever.ServiceName,
				Host:       sever.Address,
				Port:       sever.ServicePort,
				Metadata:   sever.ServiceMeta,
			}
			result[index] = s
		}
		return result, nil
	}
	return nil, nil
}

func (c consulServiceRegistry) GetServices() ([]string, error) {
	services, _, _ := c.client.Catalog().Services(nil)
	result := make([]string, unsafe.Sizeof(services))
	index := 0
	for serviceName, _ := range services {
		result[index] = serviceName
		index++
	}
	return result, nil
}

func NewConsulServiceRegistry(host string, port int, token string) (*consulServiceRegistry, error) {
	if len(host) < 3 {
		return nil, errors.New("check host")
	}

	if port <= 0 || port > 65535 {
		return nil, errors.New("check port, port should between 1 and 65535")
	}

	config := api.DefaultConfig()
	config.Address = host + ":" + strconv.Itoa(port)
	config.Token = token
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &consulServiceRegistry{client: *client}, nil
}
