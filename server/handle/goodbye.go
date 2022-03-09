package handle

import (
	"context"
	"github.com/wu-xie-888/micro-wire/api"
)

func (h *Handle) Goodbye(ctx context.Context, req *api.GoodbyeReq, rsp *api.GoodbyeRsp) error {
	msg := h.model.Goodbye(req.Content)
	rsp.Content = msg
	return nil
}
