package healthz

type HealthzInterface interface {
    HealthCheck(moduleName string) (healthy bool, err error)
}
