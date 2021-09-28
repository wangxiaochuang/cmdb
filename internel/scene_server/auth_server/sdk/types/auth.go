package types

import (
    "errors"
    "github.com/prometheus/client_golang/prometheus"
)

type Options struct {
    Metric prometheus.Registerer
}

type ResourceType string

// Decision describes the authorize decision, have already been authorized(true) or not(false)
type Decision struct {
    Authorized bool `json:"authorized"`
}

// AuthOptions describes a item to be authorized
type AuthOptions struct {
    System    string     `json:"system"`
    Subject   Subject    `json:"subject"`
    Action    Action     `json:"action"`
    Resources []Resource `json:"resources"`
}

func (a AuthOptions) Validate() error {
    if len(a.System) == 0 {
        return errors.New("system is empty")
    }

    if len(a.Subject.Type) == 0 {
        return errors.New("subject.type is empty")
    }

    if len(a.Subject.ID) == 0 {
        return errors.New("subject.id is empty")
    }

    if len(a.Action.ID) == 0 {
        return errors.New("action.id is empty")
    }

    return nil
}

type AuthBatchOptions struct {
    System  string       `json:"system"`
    Subject Subject      `json:"subject"`
    Batch   []*AuthBatch `json:"batch"`
}

func (a AuthBatchOptions) Validate() error {
    if len(a.System) == 0 {
        return errors.New("system is empty")
    }

    if len(a.Subject.Type) == 0 {
        return errors.New("subject.type is empty")
    }

    if len(a.Subject.ID) == 0 {
        return errors.New("subject.id is empty")
    }

    if len(a.Batch) == 0 {
        return nil
    }

    for _, b := range a.Batch {
        if len(b.Action.ID) == 0 {
            return errors.New("empty action id")
        }
    }
    return nil
}

type AuthBatch struct {
    Action    Action     `json:"action"`
    Resources []Resource `json:"resources"`
}

type Subject struct {
    Type ResourceType `json:"type"`
    ID   string       `json:"id"`
}

// Action define's the use's action, which is must correspond to the registered action ids in iam.
type Action struct {
    ID string `json:"id"`
}

// Resource defines all the information used to authorize a resource.
type Resource struct {
    System    string             `json:"system"`
    Type      ResourceType       `json:"type"`
    ID        string             `json:"id"`
    Attribute ResourceAttributes `json:"attribute"`
}

type ResourceAttributes map[string]interface{}
