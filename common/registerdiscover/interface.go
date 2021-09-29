package registerdiscover

type RegDiscvServer interface {
    Ping() error
    RegisterAndWatch(key string, data []byte) error
    GetServNodes(key string) ([]string, error)
    Discover(key string) (<-chan *DiscoverEvent, error)
    Cancel()
    ClearRegisterPath() error
}
