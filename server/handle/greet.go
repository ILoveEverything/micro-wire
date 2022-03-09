package handle

import (
	"context"
	"github.com/wu-xie-888/micro-wire/api"
)

func (h *Handle) Greeting(ctx context.Context, req *api.GreetingReq, rsp *api.GreetingRsp) error {
	msg := h.model.Greet(req.Content)
	rsp.Content = msg
	return nil
}
