package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"miaosha-jjl/common/global"
)

type Consul struct {
	Host string
	Port int
}

func NewConsul(host string, port int) *Consul {
	return &Consul{
		Host: host,
		Port: port,
	}
}

// 注册consul

func (c *Consul) RegisterConsul(name string, address string, port int, tags []string) error {
	if tags == nil {
		tags = []string{}
	}
	registration := new(api.AgentServiceRegistration)
	registration.ID = uuid.NewString()
	registration.Name = name
	registration.Address = address
	registration.Port = port
	registration.Tags = tags
	err := global.ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Error("error register service")
	}
	zap.S().Infof("service %s register successful", name)
	return nil
}

// 服务过滤 Filtration

func (c *Consul) FilterConsul(name string) (map[string]*api.AgentService, error) {

	sprintf := fmt.Sprintf(`Service == "%s"`, name)
	zap.S().Infof("filter %s", sprintf)
	filter, err := global.ConsulClient.Agent().ServicesWithFilter(sprintf)
	if err != nil {
		zap.S().Error("error filter consul service")
		return nil, err
	}
	return filter, err
}

// 获取consul服务列表

func (c *Consul) GetConsulServices() (map[string]*api.AgentService, error) {
	services, err := global.ConsulClient.Agent().Services()
	if err != nil {
		return nil, err
	}
	return services, nil
}

// consul注销

func (c *Consul) ServiceDeregister(serviceID string) error {
	err := global.ConsulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		zap.S().Error("error ServiceDeregister service")
		return err
	}
	zap.S().Infof("service %s ServiceDeregister successfully", serviceID)
	return nil
}
