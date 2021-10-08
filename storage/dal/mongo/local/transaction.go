package local

import (
	"context"
	"fmt"

	"github.com/wxc/cmdb/common"
	"github.com/wxc/cmdb/common/blog"
	"github.com/wxc/cmdb/common/metadata"
	"go.mongodb.org/mongo-driver/mongo"
)

// CommitTransaction 提交事务
func (c *Mongo) CommitTransaction(ctx context.Context, cap *metadata.TxnCapable) error {
	rid := ctx.Value(common.ContextRequestIDField)
	reloadSession, err := c.tm.PrepareTransaction(cap, c.dbc)
	if err != nil {
		blog.Errorf("commit transaction, but prepare transaction failed, err: %v, rid: %v", err, rid)
		return err
	}
	// reset the transaction state, so that we can commit the transaction after start the
	// transaction immediately.
	mongo.CmdbPrepareCommitOrAbort(reloadSession)

	// we commit the transaction with the session id
	err = reloadSession.CommitTransaction(ctx)
	if err != nil {
		return fmt.Errorf("commit transaction: %s failed, err: %v, rid: %v", cap.SessionID, err, rid)
	}

	err = c.tm.RemoveSessionKey(cap.SessionID)
	if err != nil {
		// this key has ttl, it's ok if we not delete it, cause this key has a ttl.
		blog.Errorf("commit transaction, but delete txn session: %s key failed, err: %v, rid: %v", cap.SessionID, err, rid)
		// do not return.
	}

	return nil
}

// AbortTransaction 取消事务
func (c *Mongo) AbortTransaction(ctx context.Context, cap *metadata.TxnCapable) error {
	rid := ctx.Value(common.ContextRequestIDField)
	reloadSession, err := c.tm.PrepareTransaction(cap, c.dbc)
	if err != nil {
		blog.Errorf("abort transaction, but prepare transaction failed, err: %v, rid: %v", err, rid)
		return err
	}
	// reset the transaction state, so that we can abort the transaction after start the
	// transaction immediately.
	mongo.CmdbPrepareCommitOrAbort(reloadSession)

	// we abort the transaction with the session id
	err = reloadSession.AbortTransaction(ctx)
	if err != nil {
		return fmt.Errorf("abort transaction: %s failed, err: %v, rid: %v", cap.SessionID, err, rid)
	}

	err = c.tm.RemoveSessionKey(cap.SessionID)
	if err != nil {
		// this key has ttl, it's ok if we not delete it, cause this key has a ttl.
		blog.Errorf("abort transaction, but delete txn session: %s key failed, err: %v, rid: %v", cap.SessionID, err, rid)
		// do not return.
	}

	return nil
}
