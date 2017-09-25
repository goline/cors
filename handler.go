package cors

import (
	. "github.com/goline/lapi"
)

func NewCorsHandler(policy Policy) Handler {
	return &corsHandler{Policy: policy}
}

type corsHandler struct {
	Policy Policy
}

func (h *corsHandler) Handle(c Connection) (interface{}, error) {
	response := c.Response()
	h.Policy.Apply(response.Header())
	return nil, response.Send()
}
