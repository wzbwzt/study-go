package handle

import (
	"context"
	"wireDemo/api"
)

func (h *Handle) Greeting(ctx context.Context, req *api.GreetingReq, rsp *api.GreetingRsp) error {
	msg := h.model.Greet(req.Content)
	rsp.Content = msg
	return nil
}
