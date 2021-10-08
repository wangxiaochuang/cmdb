module github.com/wxc/cmdb

go 1.16

require (
	github.com/coccyx/timeparser v0.0.0-20161029180942-5644122b3667
	github.com/emicklei/go-restful v2.15.0+incompatible
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis/v7 v7.4.1
	github.com/joyt/godate v0.0.0-20150226210126-7151572574a7 // indirect
	github.com/json-iterator/go v1.1.11
	github.com/juju/ratelimit v1.0.1
	github.com/mitchellh/mapstructure v1.4.2
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/mssola/user_agent v0.5.3
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/xid v1.3.0
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	github.com/tidwall/gjson v1.9.1
	github.com/xdg/scram v1.0.3 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	go.mongodb.org/mongo-driver v1.7.2
)

replace github.com/wxc/cmdb => ./

replace go.mongodb.org/mongo-driver => ./other/mongo-driver
