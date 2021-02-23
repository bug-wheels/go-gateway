package proxy

type ServiceInstance interface {

	// return The unique instance ID as registered.
	GetInstanceId() string

	// return The service ID as registered.
	GetServiceId() string

	// return The hostname of the registered service instance.
	GetHost() string

	// return The port of the registered service instance.
	GetPort() int

	// return Whether the port of the registered service instance uses HTTPS.
	IsSecure() bool

	// return The key / value pair metadata associated with the service instance.
	GetMetadata() map[string]string
}

type DefaultServiceInstance struct {
	InstanceId string
	ServiceId  string
	Host       string
	Port       int
	Secure     bool
	Metadata   map[string]string
}

func (serviceInstance DefaultServiceInstance) GetInstanceId() string {
	return serviceInstance.InstanceId
}

func (serviceInstance DefaultServiceInstance) GetServiceId() string {
	return serviceInstance.ServiceId
}

func (serviceInstance DefaultServiceInstance) GetHost() string {
	return serviceInstance.Host
}

func (serviceInstance DefaultServiceInstance) GetPort() int {
	return serviceInstance.Port
}

func (serviceInstance DefaultServiceInstance) IsSecure() bool {
	return serviceInstance.Secure
}

func (serviceInstance DefaultServiceInstance) GetMetadata() map[string]string {
	return serviceInstance.Metadata
}

type DiscoveryClient interface {

	GetInstances(serviceId string) ([]ServiceInstance, error)

	GetServices() ([]string, error)
}
