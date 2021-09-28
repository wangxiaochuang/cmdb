package configcenter

type ProcessConfig struct {
    ConfigData []byte
}

func ParseConfigWithData(data []byte) *ProcessConfig {
    return &ProcessConfig{ConfigData: data}
}
