package cors

import (
	"github.com/goline/errors"
	. "github.com/goline/lapi"
)

type corsHandler struct {
	Policy Policy
}

func (h *corsHandler) Handle(c Connection) (interface{}, errors.Error) {
	response := c.Response()
	h.Policy.Apply(response.Header())
	return nil, response.Send()
}
