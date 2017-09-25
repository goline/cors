package cors

import "github.com/goline/lapi"

type Loader struct {
	Policy Policy
}

func (l *Loader) Load(app lapi.App) {
	app.Router().Options("/.*", NewCorsHandler(l.Policy))
}
