package iam

import (
    "context"
    "net/http"

    "github.com/wxc/cmdb/ac"
    "github.com/wxc/cmdb/ac/meta"
    "github.com/wxc/cmdb/apimachinery"
    "github.com/wxc/cmdb/apimachinery/authserver"

    "github.com/wxc/cmdb/internel/scene_server/auth_server/sdk/types"
)

type authorizer struct {
    authClientSet authserver.AuthServerClientInterface
}

func NewAuthorizer(clientSet apimachinery.ClientSetInterface) ac.AuthorizeInterface {
    return &authorizer{authClientSet: clientSet.AuthServer()}
}

func (a *authorizer) AuthorizeBatch(ctx context.Context, h http.Header, user meta.UserInfo,
    resources ...meta.ResourceAttribute) ([]types.Decision, error) {
    return a.authorizeBatch(ctx, h, true, user, resources...)
}

func (a *authorizer) authorizeBatch(ctx context.Context, h http.Header, exact bool, user meta.UserInfo,
    resources ...meta.ResourceAttribute) ([]types.Decision, error) {
    panic("todo")
    return []types.Decision{}, nil
}
