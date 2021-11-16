package grpc

import (
	"context"

	"gitlab.warungpintar.co/sales-platform/brook/proto/brook"
)

func (obj *testObject) TestPing() {
	obj.Run("Success", func() {
		resp, err := obj.module.Ping(context.Background(), &brook.PingRequest{})
		obj.NoError(err)
		obj.Equal(resp.Message, "pong")
	})
}
