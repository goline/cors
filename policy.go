package cors

import (
	"strconv"

	. "github.com/goline/lapi"
)

type Policy struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowMethods     string
	AllowCredentials string
	Origin           string
	RequestMethod    string
	RequestHeaders   string
	ExposeHeaders    string
	MaxAge           int
}

func (p *Policy) Apply(h Header) {
	if p.AllowOrigin != "" {
		h.Set("Access-Control-Allow-Origin", p.AllowOrigin)
	}
	if p.AllowCredentials != "" {
		h.Set("Access-Control-Allow-Credentials", p.AllowCredentials)
	}
	if p.AllowHeaders != "" {
		h.Set("Access-Control-Allow-Headers", p.AllowHeaders)
	}
	if p.AllowMethods != "" {
		h.Set("Access-Control-Allow-Methods", p.AllowMethods)
	}
	if p.ExposeHeaders != "" {
		h.Set("Access-Control-Expose-Headers", p.ExposeHeaders)
	}
	if p.MaxAge != 0 {
		h.Set("Access-Control-Max-Age", strconv.Itoa(p.MaxAge))
	}
}
