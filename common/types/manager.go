package types

var (
    needDiscoveryServiceName map[string]struct{} = make(map[string]struct{}, 0)
)

func DiscoveryAllService() {
    for name := range AllModule {
        needDiscoveryServiceName[name] = struct{}{}
    }
}

func AddDiscoveryService(name ...string) {
    for _, name := range name {
        needDiscoveryServiceName[name] = struct{}{}
    }
}

func GetDiscoveryService() map[string]struct{} {
    if len(needDiscoveryServiceName) == 0 {
        DiscoveryAllService()
    }
    return needDiscoveryServiceName
}
