package handle

import (
	"context"

	"wireDemo/api"
)

func (h *Handle) Goodbye(ctx context.Context, req *api.GoodbyeReq, rsp *api.GoodbyeRsp) error {
	msg := h.model.Goodbye(req.Content)
	rsp.Content = msg
	return nil
}
