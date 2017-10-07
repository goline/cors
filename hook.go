package cors

import (
	"github.com/goline/errors"
	. "github.com/goline/lapi"
)

type corsHook struct {
	Policy Policy
}

func (h *corsHook) SetUp(c Connection) errors.Error {
	h.Policy.Apply(c.Response().Header())

	return nil
}

func (h *corsHook) TearDown(_ Connection, _ interface{}, _ errors.Error) errors.Error {
	return nil
}
