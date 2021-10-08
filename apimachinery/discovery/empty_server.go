package discovery

var (
	emptyServerInst = &emptyServer{}
)

// emptyServer 适配服务不存在的情况， 当服务不存在的时候，返回空的服务
type emptyServer struct {
}

func (es *emptyServer) GetServers() ([]string, error) {
	return []string{}, nil
}

func (es *emptyServer) GetServersChan() chan []string {
	return make(chan []string, 20)
}
