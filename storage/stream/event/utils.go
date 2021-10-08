package event

import (
	"reflect"

	"github.com/wxc/cmdb/storage/stream/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	esType = reflect.TypeOf(types.EventStream{})
)

// newEventStruct construct a change stream event data structure
// which can help us to adjust different kind of collection structure.
func newEventStruct(typ reflect.Type) reflect.Value {
	f := reflect.StructOf([]reflect.StructField{
		{
			Name:      "EventStream",
			Type:      esType,
			Anonymous: true,
			Tag:       `bson:",inline"`,
		},
		{
			Name:      "FullDocument",
			Type:      typ,
			Anonymous: false,
			Tag:       `bson:"fullDocument"`,
		},
	})
	return reflect.New(f).Elem()
}

const fullDocPrefix = "fullDocument."

var eventFields = []string{"_id", "operationType", "clusterTime", "ns", "documentKey", "updateDescription"}

func generateOptions(opts *types.Options) (mongo.Pipeline, *options.ChangeStreamOptions) {

	fields := make([]bson.E, 0)
	if opts.OperationType != nil {
		fields = append(fields, bson.E{Key: "operationType", Value: *opts.OperationType})
	}

	if opts.Collection == "" {
		fields = append(fields, bson.E{Key: "ns.coll", Value: opts.CollectionFilter})
	}

	if opts.Filter != nil {
		for k, v := range opts.Filter {
			fields = append(fields, bson.E{Key: fullDocPrefix + k, Value: v})
		}
	}

	var pipeline mongo.Pipeline
	if len(fields) != 0 {
		pipeline = []bson.D{{{Key: "$match", Value: fields}}}
	}

	if len(opts.Fields) != 0 {
		project := make(map[string]int)
		for _, f := range opts.Fields {
			project[fullDocPrefix+f] = 1
		}

		// add default event fields, otherwise, these fields will not be returned.
		for _, f := range eventFields {
			project[f] = 1
		}

		pipeline = append(pipeline, bson.D{{Key: "$project", Value: project}})
	}

	streamOptions := new(options.ChangeStreamOptions)
	if opts.MajorityCommitted != nil {
		if *opts.MajorityCommitted {
			major := options.UpdateLookup
			streamOptions.FullDocument = &major
		} else {
			def := options.Default
			streamOptions.FullDocument = &def
		}
	}

	if opts.MaxAwaitTime != nil {
		streamOptions.MaxAwaitTime = opts.MaxAwaitTime
	}

	if opts.StartAfterToken != nil && opts.StartAtTime == nil {
		streamOptions.SetStartAfter(opts.StartAfterToken)
	}

	if opts.StartAfterToken == nil && opts.StartAtTime != nil {
		streamOptions.SetStartAtOperationTime(&primitive.Timestamp{
			T: opts.StartAtTime.Sec,
			I: opts.StartAtTime.Nano,
		})
	}

	// if all set, then use token to resume after, this is accurate.
	if opts.StartAfterToken != nil && opts.StartAtTime != nil {
		streamOptions.SetStartAfter(opts.StartAfterToken)
	}

	// set batch size, otherwise,
	// it will take as much as about 16MB data one cycle with unlimited batch size as default.
	var batchSize int32 = 2000
	streamOptions.BatchSize = &batchSize

	return pipeline, streamOptions
}
