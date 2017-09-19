package cors

import "github.com/goline/lapi"

type Loader struct {
	AccessControlAllowOrigin string
}

func (l *Loader) Load(app lapi.App) {
	if l.AccessControlAllowOrigin == "" {
		l.AccessControlAllowOrigin = "*"
	}
	app.Router().Options("/{uri:.*}", &CorsHandler{l.AccessControlAllowOrigin})
}
