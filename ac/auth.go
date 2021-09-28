package ac

import (
    "context"
    "errors"
    "net/http"

    "github.com/wxc/cmdb/ac/meta"
    //"github.com/wxc/cmdb/common/metadata"
    "github.com/wxc/cmdb/internel/scene_server/auth_server/sdk/types"
)

var NoAuthorizeError = errors.New("no authorize")

type AuthInterface interface {
    RegisterSystem(ctx context.Context, host string) error
}

type AuthorizeInterface interface {
    AuthorizeBatch(ctx context.Context, h http.Header, user meta.UserInfo, resources ...meta.ResourceAttribute) (
        []types.Decision, error)

    // AuthorizeAnyBatch(ctx context.Context, h http.Header, user meta.UserInfo, resources ...meta.ResourceAttribute) (
    //     []types.Decision, error)

    // ListAuthorizedResources(ctx context.Context, h http.Header, input meta.ListAuthorizedResourcesParam) ([]string, error)
    // GetNoAuthSkipUrl(ctx context.Context, h http.Header, input *metadata.IamPermission) (string, error)
    // GetPermissionToApply(ctx context.Context, h http.Header, input []meta.ResourceAttribute) (*metadata.IamPermission, error)
    // RegisterResourceCreatorAction(ctx context.Context, h http.Header, input metadata.IamInstanceWithCreator) (
    //     []metadata.IamCreatorActionPolicy, error)

    // BatchRegisterResourceCreatorAction(ctx context.Context, h http.Header, input metadata.IamInstancesWithCreator) (
    //     []metadata.IamCreatorActionPolicy, error)
}
