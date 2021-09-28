module github.com/wxc/cmdb

go 1.16

require (
	github.com/emicklei/go-restful v2.15.0+incompatible
	github.com/go-redis/redis/v7 v7.4.1
	github.com/juju/ratelimit v1.0.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/xid v1.3.0
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
)

replace github.com/wxc/cmdb => ./
