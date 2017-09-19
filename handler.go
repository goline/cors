package cors

import "github.com/goline/lapi"

type CorsHandler struct {
	AccessControlAllowOrigin string
}

func (h *CorsHandler) Handle(connection lapi.Connection) (interface{}, error) {
	response := connection.Response()
	response.Header().Set("Access-Control-Allow-Origin", h.AccessControlAllowOrigin)
	return nil, response.Send()
}
