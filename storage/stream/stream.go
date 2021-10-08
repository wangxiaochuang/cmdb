package stream

import (
	"context"
	"fmt"
	"time"

	"github.com/wxc/cmdb/apimachinery/discovery"
	"github.com/wxc/cmdb/storage/dal/mongo/local"
	"github.com/wxc/cmdb/storage/stream/event"
	"github.com/wxc/cmdb/storage/stream/loop"
	"github.com/wxc/cmdb/storage/stream/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// Stream Interface defines all the functionality it have.
type Interface interface {
	List(ctx context.Context, opts *types.ListOptions) (ch chan *types.Event, err error)
	Watch(ctx context.Context, opts *types.WatchOptions) (*types.Watcher, error)
	ListWatch(ctx context.Context, opts *types.ListWatchOptions) (*types.Watcher, error)
}

// NewStream create a list watch event stream
func NewStream(conf local.MongoConf) (Interface, error) {
	connStr, err := connstring.Parse(conf.URI)
	if nil != err {
		return nil, err
	}
	if conf.RsName == "" {
		return nil, fmt.Errorf("rsName not set")
	}

	timeout := 15 * time.Second
	conOpt := options.ClientOptions{
		MaxPoolSize:    &conf.MaxOpenConns,
		MinPoolSize:    &conf.MaxIdleConns,
		ConnectTimeout: &timeout,
		ReplicaSet:     &conf.RsName,
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI), &conOpt)
	if nil != err {
		return nil, err
	}
	if err := client.Connect(context.TODO()); nil != err {
		return nil, err
	}

	event, err := event.NewEvent(client, connStr.Database)
	if err != nil {
		return nil, fmt.Errorf("new event failed, err: %v", err)
	}
	return event, nil
}

type LoopInterface interface {
	WithOne(opts *types.LoopOneOptions) error
	WithBatch(opts *types.LoopBatchOptions) error
}

// NewLoopStream create a new event loop stream.
func NewLoopStream(conf local.MongoConf, isMaster discovery.ServiceManageInterface) (LoopInterface, error) {
	connStr, err := connstring.Parse(conf.URI)
	if nil != err {
		return nil, err
	}
	if conf.RsName == "" {
		return nil, fmt.Errorf("rsName not set")
	}

	timeout := 15 * time.Second
	conOpt := options.ClientOptions{
		MaxPoolSize:    &conf.MaxOpenConns,
		MinPoolSize:    &conf.MaxIdleConns,
		ConnectTimeout: &timeout,
		ReplicaSet:     &conf.RsName,
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI), &conOpt)
	if nil != err {
		return nil, err
	}
	if err := client.Connect(context.TODO()); nil != err {
		return nil, err
	}

	event, err := event.NewEvent(client, connStr.Database)
	if err != nil {
		return nil, fmt.Errorf("new event failed, err: %v", err)
	}

	loop, err := loop.NewLoopWatch(event, isMaster)
	if err != nil {
		return nil, err
	}

	return loop, nil
}
