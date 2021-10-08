package confregdiscover

import (
	"github.com/wxc/cmdb/common/backbone/service_mange/zk"
)

//DiscoverEvent if servers changed, will create a discover event
type DiscoverEvent struct { //
	Err  error
	Key  string
	Data []byte
}

// ConfRegDiscover is config register and discover
type ConfRegDiscover struct {
	confRD ConfRegDiscvIf
}

// NewConfRegDiscover used to create a object of ConfRegDiscover
func NewConfRegDiscover(client *zk.ZkClient) *ConfRegDiscover {
	confRD := &ConfRegDiscover{
		confRD: nil,
	}

	confRD.confRD = ConfRegDiscvIf(NewZkRegDiscover(client))

	return confRD
}

// NewConfRegDiscoverWithTimeOut used to create a object
func NewConfRegDiscoverWithTimeOut(client *zk.ZkClient) *ConfRegDiscover {
	confRD := &ConfRegDiscover{
		confRD: nil,
	}

	confRD.confRD = ConfRegDiscvIf(NewZkRegDiscover(client))

	return confRD
}

// Ping to ping server
func (crd *ConfRegDiscover) Ping() error {
	return crd.confRD.Ping()
}

//Write the configure data
func (crd *ConfRegDiscover) Write(key string, data []byte) error {
	return crd.confRD.Write(key, data)
}

// Read the configure data
func (crd *ConfRegDiscover) Read(path string) (string, error) {
	return crd.confRD.Read(path)
}

//DiscoverConfig discover the config wether is changed
func (crd *ConfRegDiscover) DiscoverConfig(key string) (<-chan *DiscoverEvent, error) {
	return crd.confRD.Discover(key)
}
