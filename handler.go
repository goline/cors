package cors

import (
	"github.com/goline/lapi"
	"strconv"
)

type Handler struct {
	Policy Policy
}

func (h *Handler) Handle(connection lapi.Connection) (interface{}, error) {
	response := connection.Response()
	h.applyResponseHeader(response.Header())
	return nil, response.Send()
}

func (h *Handler) applyResponseHeader(header lapi.Header) {
	if h.Policy.AllowOrigin != "" {
		header.Set("Access-Control-Allow-Origin", h.Policy.AllowOrigin)
	}
	if h.Policy.AllowCredentials != "" {
		header.Set("Access-Control-Allow-Credentials", h.Policy.AllowCredentials)
	}
	if h.Policy.AllowHeaders != "" {
		header.Set("Access-Control-Allow-Headers", h.Policy.AllowHeaders)
	}
	if h.Policy.AllowMethods != "" {
		header.Set("Access-Control-Allow-Methods", h.Policy.AllowMethods)
	}
	if h.Policy.ExposeHeaders != "" {
		header.Set("Access-Control-Expose-Headers", h.Policy.ExposeHeaders)
	}
	if h.Policy.MaxAge != 0 {
		header.Set("Access-Control-Max-Age", strconv.Itoa(h.Policy.MaxAge))
	}
}
