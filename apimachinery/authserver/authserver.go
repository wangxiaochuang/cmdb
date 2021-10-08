package authserver

import (
    "context"
    "fmt"
    "net/http"

    "github.com/wxc/cmdb/apimachinery/rest"
    "github.com/wxc/cmdb/apimachinery/util"
    "github.com/wxc/cmdb/internel/scene_server/auth_server/sdk/types"
)

type AuthServerClientInterface interface {
    AuthorizeBatch(ctx context.Context, h http.Header, input *types.AuthBatchOptions) ([]types.Decision, error)
}

func NewAuthServerClientInterface(c *util.Capability, version string) AuthServerClientInterface {
    base := fmt.Sprintf("/ac/%s", version)
    return &authServer{
        client: rest.NewRESTClient(c, base),
    }
}

type authServer struct {
    client rest.ClientInterface
}
