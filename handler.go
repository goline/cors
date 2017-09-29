package cors

import (
	. "github.com/goline/lapi"
)

type corsHandler struct {
	Policy Policy
}

func (h *corsHandler) Handle(c Connection) (interface{}, error) {
	response := c.Response()
	h.Policy.Apply(response.Header())
	return nil, response.Send()
}
