package cors

import (
	. "github.com/goline/lapi"
)

type corsHook struct {
	Policy Policy
}

func (h *corsHook) SetUp(c Connection) error {
	h.Policy.Apply(c.Response().Header())

	return nil
}

func (h *corsHook) TearDown(_ Connection, _ interface{}, _ error) error {
	return nil
}
