package authserver

import (
    "context"
    "net/http"

    "github.com/wxc/cmdb/common/errors"
    "github.com/wxc/cmdb/common/metadata"
    "github.com/wxc/cmdb/internel/scene_server/auth_server/sdk/types"
)

type authorizeBatchResp struct {
    metadata.BaseResp `json:",inline"`
    Data              []types.Decision `json:"data"`
}

func (a *authServer) AuthorizeBatch(ctx context.Context, h http.Header, input *types.AuthBatchOptions) ([]types.Decision, error) {
    subPath := "/authorize/batch"
    response := new(authorizeBatchResp)

    err := a.client.Post().
        WithContext(ctx).
        Body(input).
        SubResourcef(subPath).
        WithHeaders(h).
        Do().
        Into(response)

    if err != nil {
        return nil, errors.CCHttpError
    }
    if response.Code != 0 {
        return nil, response.CCError()
    }

    return response.Data, nil
}
